package econcurrent

import (
	"context"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	econtext "github.com/guanguoyintao/kuafu/context"
)

// DAG 结构，使用 sync.Map 来存储任务和依赖关系
type DAG struct {
	ctx     context.Context
	tasks   map[string]*Runnable // 存储任务，键: 任务名称，值: 任务函数
	taskSub map[string][]string
	states  *sync.Map // 存储任务子任务
	lock    *sync.Mutex
}

// NewDAG 创建一个新的 DAG
func NewDAG(ctx context.Context) *DAG {
	return &DAG{
		ctx:     ctx,
		tasks:   make(map[string]*Runnable),
		taskSub: make(map[string][]string),
		lock:    &sync.Mutex{},
		states:  &sync.Map{},
	}
}

// SetState sets the state for a task
func (d *DAG) SetState(taskName string, state any) {
	d.states.Store(taskName, state)
}

// GetState gets the state for a task
func (d *DAG) GetState(taskName string) (any, bool) {
	return d.states.Load(taskName)
}

// AddTask 添加任务到 DAG
func (d *DAG) AddTask(runnables ...*Runnable) {
	for _, runnable := range runnables {
		d.tasks[runnable.TaskName] = runnable
	}
}

// AddDependency 设置任务依赖关系
func (d *DAG) AddDependency(taskName, dependencyTaskName string) {
	if _, ok := d.tasks[taskName]; !ok {
		return
	}
	if _, ok := d.tasks[dependencyTaskName]; !ok {
		return
	}
	_, ok := d.taskSub[dependencyTaskName]
	if !ok {
		d.taskSub[dependencyTaskName] = make([]string, 0)
	}
	d.taskSub[dependencyTaskName] = append(d.taskSub[dependencyTaskName], taskName)
}

// 递归计算任务的入度和所有子任务
// 递归计算任务的入度和所有子任务
func (d *DAG) calculateDependencies() (map[string]int, map[string][]string) {
	inDegree := make(map[string]int)         // 任务的入度
	allSubTasks := make(map[string][]string) // 所有子任务的完整列表
	// 初始化所有任务的入度和子任务列表
	for task := range d.tasks {
		inDegree[task] = 0
		allSubTasks[task] = make([]string, 0)
	} // 计算所有子任务并更新入度
	visited := make(map[string]bool)
	var getAllSubTasks func(task string) []string
	getAllSubTasks = func(task string) []string {
		if visited[task] {
			return allSubTasks[task]
		}
		visited[task] = true
		result := make([]string, 0)
		directSubs, exists := d.taskSub[task]
		if !exists {
			return result
		}
		// 添加直接子任务并更新入度
		for _, sub := range directSubs {
			inDegree[sub]++ // 直接依赖的入度加1
			result = append(result, sub)
			// 递归获取子任务的所有后代
			grandChildren := getAllSubTasks(sub)
			result = append(result, grandChildren...)
		}
		// 去除重复项
		unique := make(map[string]bool)
		final := make([]string, 0)
		for _, t := range result {
			if !unique[t] {
				unique[t] = true
				final = append(final, t)
			}
		}
		allSubTasks[task] = final
		return final
	}
	// 对每个任务计算完整的子任务列表和入度
	for task := range d.tasks {
		if !visited[task] {
			getAllSubTasks(task)
		}
	}
	return inDegree, allSubTasks
}

// Execute 执行 DAG 中的所有任务
func (d *DAG) Execute(ctx context.Context, l *log.Helper, timeout time.Duration) error {
	ctx, cancel := econtext.NewTimeoutContext(ctx, timeout)
	go func() {
		defer cancel()
		inDegree, _ := d.calculateDependencies()
		// 使用 channel 来控制并发任务
		taskChan := make(chan string, len(inDegree)) // 用于存放可以执行的任务
		doneChan := make(chan string, len(inDegree)) // 用于通知任务执行完毕
		defer func() {
			// 关闭 channel
			close(taskChan)
			close(doneChan)
		}()
		// 启动并发 goroutine 来处理任务
		wg := &sync.WaitGroup{}
		go func() {
			// 从 channel 获取任务
			for taskName := range taskChan {
				go func() {
					defer wg.Done()
					defer func() {
						doneChan <- taskName
					}()
					d.lock.Lock()
					// 获取任务对应的执行方法
					task, _ := d.tasks[taskName]
					d.lock.Unlock()
					// 执行任务
					defer GoroutineRecover(ctx, l)
					_ = goExecTracerRunnableHandler(ctx, l, task)
					return
				}()
			}
		}()
		// 将所有入度为 0 的任务放入任务 channel 中
		for taskName, degree := range inDegree {
			if degree == 0 {
				wg.Add(1)
				taskChan <- taskName
				delete(inDegree, taskName)
			}
		}
		for doneTaskName := range doneChan {
			// 更新入度
			_, ok := d.taskSub[doneTaskName]
			if ok {
				for _, taskName := range d.taskSub[doneTaskName] {
					inDegree[taskName]--
				}
			}
			// 将所有入度为 0 的任务放入任务 channel 中
			for taskName, degree := range inDegree {
				if degree == 0 {
					wg.Add(1)
					taskChan <- taskName
					delete(inDegree, taskName)
				}
			}
			if len(inDegree) == 0 {
				break
			}
		}
		wg.Wait()
	}()
	return nil
}
