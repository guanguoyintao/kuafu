package evalidator

import (
	"regexp"
	"testing"
)

func TestValidator(t *testing.T) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// --- 定义测试用例 ---
	testCases := []struct {
		name            string
		field           string
		value           any
		handlers        []ValidationHandler[any]
		expectError     bool
		expectedErrType ValidationErrorType
	}{
		{
			name:  "Valid Email",
			field: "email",
			value: "test@example.com",
			handlers: []ValidationHandler[any]{
				Required[any](),
				MinLength[any](5),
				MaxLength[any](50),
				Regex[any](emailRegex),
			},
			expectError: false,
		},
		{
			name:  "Empty Value - Fails Required",
			field: "email",
			value: "",
			handlers: []ValidationHandler[any]{
				Required[any](),
				MinLength[any](5),
			},
			expectError:     true,
			expectedErrType: ErrRequired,
		},
		{
			name:            "Too Short - Fails MinLength",
			field:           "username",
			value:           "usr",
			handlers:        []ValidationHandler[any]{Required[any](), MinLength[any](4)},
			expectError:     true,
			expectedErrType: ErrMinLength,
		},
		{
			name:            "Too Long - Fails MaxLength",
			field:           "username",
			value:           "this_is_a_very_long_username",
			handlers:        []ValidationHandler[any]{MaxLength[any](20)},
			expectError:     true,
			expectedErrType: ErrMaxLength,
		},
		{
			name:            "Invalid Format - Fails Regex",
			field:           "email",
			value:           "not-a-valid-email",
			handlers:        []ValidationHandler[any]{Required[any](), Regex[any](emailRegex)},
			expectError:     true,
			expectedErrType: ErrInvalidFormat,
		},
		{
			name:  "Custom Rule - Fails",
			field: "role",
			value: "guest",
			handlers: []ValidationHandler[any]{
				Rule[any](func(field string, value any) *ValidationError {
					if v, ok := value.(string); ok && v != "admin" {
						return &ValidationError{Field: field, Type: ErrCustom, Msg: "must be 'admin'"}
					}
					return nil
				}),
			},
			expectError:     true,
			expectedErrType: ErrCustom,
		},
		{
			name:        "Valid Slice Length",
			field:       "tags",
			value:       []any{"go", "generic", "validator"},
			handlers:    []ValidationHandler[any]{MinLength[any](2), MaxLength[any](5)},
			expectError: false,
		},
		{
			name:            "Invalid Slice Length - Fails MinLength",
			field:           "tags",
			value:           []any{"go"},
			handlers:        []ValidationHandler[any]{MinLength[any](2)},
			expectError:     true,
			expectedErrType: ErrMinLength,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 注意：这里我们使用 any 作为泛型类型，因为测试用例的值类型不同。
			// 在实际应用中，您会使用具体的类型，例如 string 或 int。
			err := Validate(tc.field, tc.value, tc.handlers...)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error for field '%s' with value '%v', but got none", tc.field, tc.value)
					return
				}
				if err.Type != tc.expectedErrType {
					t.Errorf("Expected error type %v, but got %v. Error message: %s", tc.expectedErrType, err.Type, err.Error())
				}
				t.Logf("Successfully caught expected error: %v", err)
			} else if err != nil {
				t.Errorf("Expected no error for field '%s' with value '%v', but got: %s", tc.field, tc.value, err.Error())
			}
		})
	}
}

func TestPasswordValidation_UsageExample(t *testing.T) {
	// --- 1. 定义校验规则 ---
	// 规则1: 长度范围控制在 10-20
	minLengthRule := MinLength[string](10)
	maxLengthRule := MaxLength[string](20)

	// 规则2: 必须包含数字和字母
	containsDigitRule := Regex[string](regexp.MustCompile(`[0-9]`))
	containsLetterRule := Regex[string](regexp.MustCompile(`[a-zA-Z]`))

	// 规则3: 不能有特殊字符
	noSpecialCharsRule := Regex[string](regexp.MustCompile(`^[a-zA-Z0-9]+$`))

	// --- 2. 定义测试用例表 ---
	testCases := []struct {
		name              string              // 测试用例的描述
		password          string              // 待测试的密码
		expectError       bool                // 是否期望出现错误
		expectedFailureOn ValidationErrorType // 如果期望出错，期望是哪种类型的错误
	}{
		{
			name:        "Valid password",
			password:    "password123",
			expectError: false,
		},
		{
			name:              "Fails on MinLength",
			password:          "short",
			expectError:       true,
			expectedFailureOn: ErrMinLength,
		},
		{
			name:              "Fails on MaxLength",
			password:          "thispasswordiswaytoolong",
			expectError:       true,
			expectedFailureOn: ErrMaxLength,
		},
		{
			name:              "Fails on missing digit",
			password:          "longenoughtext",
			expectError:       true,
			expectedFailureOn: ErrInvalidFormat, // 由 containsDigitRule 触发
		},
		{
			name:              "Fails on missing letter",
			password:          "12345678901",
			expectError:       true,
			expectedFailureOn: ErrInvalidFormat, // 由 containsLetterRule 触发
		},
		{
			name:              "Fails on special character",
			password:          "password123!",
			expectError:       true,
			expectedFailureOn: ErrInvalidFormat, // 由 noSpecialCharsRule 触发
		},
	}

	// --- 3. 执行测试 ---
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用 Validate 函数，传入所有要应用的规则
			err := Validate("password", tc.password,
				minLengthRule,
				maxLengthRule,
				containsDigitRule,
				containsLetterRule,
				noSpecialCharsRule,
			)

			// --- 4.断言测试结果 ---
			if tc.expectError {
				if err == nil {
					t.Fatalf("Expected a validation error for password '%s', but got none", tc.password)
				}
				if err.Type != tc.expectedFailureOn {
					t.Errorf("Expected error type %v, but got %v. Full error: %s", tc.expectedFailureOn, err.Type, err.Error())
				}
				// 打印捕获到的预期错误，便于调试
				t.Logf("Successfully caught expected error: %s", err.Error())
			} else {
				if err != nil {
					t.Fatalf("Expected no error for password '%s', but got: %s", tc.password, err.Error())
				}
				t.Log("Password validated successfully as expected.")
			}
		})
	}
}
