package emime

import (
	"strings"
)

// MIMEType 表示MIME类型的枚举
type MIMEType string

// 定义所有支持的MIME类型常量
const (
	// 视频相关的 MIME 类型
	MIMEVIDEOM2TS  MIMEType = "video/mp2t"            // M2TS视频
	MIMEVIDEOMP4   MIMEType = "video/mp4"             // M4V视频
	MIMEVIDEOMOV   MIMEType = "video/quicktime"       // MOV视频
	MIMEVIDEOMPEG  MIMEType = "video/mpeg"            // MPEG视频
	MIMEVIDEOMKV   MIMEType = "video/x-matroska"      // MKV视频
	MIMEVIDEOWMV   MIMEType = "video/x-ms-wmv"        // WMV视频
	MIMEVIDEOM3U8  MIMEType = "application/x-mpegURL" // M3U8视频
	MIMEVIDEORTSP  MIMEType = "application/x-rtsp"    // RTSP流
	MIMEVIDEORTMP  MIMEType = "application/x-rtmp"    // RTMP流
	MIMEVIDEO3G2   MIMEType = "video/3gpp2"
	MIMEVIDEO3GPP  MIMEType = "video/3gpp"
	MIMEVIDEODVB   MIMEType = "video/vnd.dvb.file" // DVB视频
	MIMEVIDEOFLI   MIMEType = "video/x-fli"        // FLI视频
	MIMEVIDEOFVT   MIMEType = "video/vnd.fvt"      // FVT视频
	MIMEVIDEOH261  MIMEType = "video/h261"         // H.261视频
	MIMEVIDEOH263  MIMEType = "video/h263"         // H.263视频
	MIMEVIDEOH264  MIMEType = "video/h264"         // H.264视频
	MIMEVIDEOJPGM  MIMEType = "video/jpm"          // JPGM视频
	MIMEVIDEOJPGV  MIMEType = "video/jpeg"         // JPGV视频
	MIMEVIDEOM4U   MIMEType = "video/vnd.mpegurl"  // M4U视频
	MIMEVIDEOFLV   MIMEType = "video/x-flv"        // FLV视频
	MIMEVIDEOASF   MIMEType = "video/x-ms-asf"
	MIMEVIDEOASX   MIMEType = "video/x-ms-asf"
	MIMEVIDEOAVI   MIMEType = "video/x-msvideo"
	MIMEVIDEOMK3D  MIMEType = "video/x-matroska-3d" // MK3D视频
	MIMEVIDEOMOVIE MIMEType = "video/x-sgi-movie"
	MIMEVIDEOF4V   MIMEType = "video/x-f4v"
	MIMEVIDEOMJ2   MIMEType = "video/mj2"
	MIMEVIDEOMNG   MIMEType = "video/x-mng"
	MIMEVIDEOMXU   MIMEType = "video/vnd.mpegurl"
	MIMEVIDEOOGV   MIMEType = "video/ogg"
	MIMEVIDEOPYV   MIMEType = "video/vnd.ms-playready.media.pyv"
	MIMEVIDEOSMV   MIMEType = "video/x-smv"
	MIMEVIDEOUVH   MIMEType = "video/vnd.dece.hd"
	MIMEVIDEOUVM   MIMEType = "video/vnd.dece.mobile"
	MIMEVIDEOUVP   MIMEType = "video/vnd.dece.pd"
	MIMEVIDEOUVS   MIMEType = "video/vnd.dece.sd"
	MIMEVIDEOUVU   MIMEType = "video/vnd.uvvu.mp4"
	MIMEVIDEOUVV   MIMEType = "video/vnd.dece.video"
	MIMEVIDEOUVVH  MIMEType = "video/vnd.dece.hd"
	MIMEVIDEOUVVM  MIMEType = "video/vnd.dece.mobile"
	MIMEVIDEOUVVP  MIMEType = "video/vnd.dece.pd"
	MIMEVIDEOUVVS  MIMEType = "video/vnd.dece.sd"
	MIMEVIDEOUVVU  MIMEType = "video/vnd.uvvu.mp4"
	MIMEVIDEOUVVV  MIMEType = "video/vnd.dece.video"
	MIMEVIDEOVIV   MIMEType = "video/vnd.vivo"
	MIMEVIDEOVOB   MIMEType = "video/x-ms-vob"
	MIMEVIDEOWEBM  MIMEType = "video/webm"
	MIMEVIDEOWM    MIMEType = "video/x-ms-wm"
	MIMEVIDEOWVX   MIMEType = "video/x-ms-wvx"
	MIMEVIDEOWMX   MIMEType = "video/x-ms-wmx"

	// 音频相关的 MIME 类型
	MIMEAUDIOADPCM      MIMEType = "audio/adpcm"
	MIMEAUDIOWAV        MIMEType = "audio/wav"                           // WAV音频
	MIMEAUDIOAAC        MIMEType = "audio/aac"                           // AAC音频
	MIMEAUDIOFLAC       MIMEType = "audio/flac"                          // FLAC音频
	MIMEAUDIOOGG        MIMEType = "audio/ogg"                           // OGG音频
	MIMEAUDIOAMR        MIMEType = "audio/amr"                           // AMR音频
	MIMEAUDIOAC3        MIMEType = "audio/ac3"                           // AC3音频
	MIMEAUDIOMPEG       MIMEType = "audio/mpeg"                          // MPGA音频
	MIMEAUDIOMKA        MIMEType = "audio/x-matroska"                    // MKA音频
	MIMEAUDIOMP4A       MIMEType = "audio/mp4"                           // MP4音频
	MIMEAPPLICATIONMP4S MIMEType = "application/mp4"                     // MP4文件
	MIMEAPPLICATIONMPKG MIMEType = "application/vnd.apple.installer+xml" // MPKG安装包
	MIMEAUDIORAP        MIMEType = "audio/x-pn-realaudio-plugin"
	MIMEAUDIOAIFF       MIMEType = "audio/x-aiff"
	MIMEAUDIOBASIC      MIMEType = "audio/basic"
	MIMEAUDIOCAF        MIMEType = "audio/x-caf"
	MIMEAUDIODTS        MIMEType = "audio/vnd.dts"
	MIMEAUDIODTSHD      MIMEType = "audio/vnd.dts.hd"
	MIMEAUDIODRA        MIMEType = "audio/vnd.dra"
	MIMEAUDIOECELP4800  MIMEType = "audio/vnd.nuera.ecelp4800"
	MIMEAUDIOECELP7470  MIMEType = "audio/vnd.nuera.ecelp7470"
	MIMEAUDIOECELP9600  MIMEType = "audio/vnd.nuera.ecelp9600"
	MIMEAUDIOEOL        MIMEType = "audio/vnd.digital-winds"
	MIMEAUDIOKAR        MIMEType = "audio/midi"
	MIMEAUDIOLVP        MIMEType = "audio/vnd.lucent.voice"
	MIMEAUDIOM3U        MIMEType = "audio/x-mpegurl"
	MIMEAUDIOM4P        MIMEType = "audio/mp4a-latm"
	MIMEAUDIOMIDI       MIMEType = "audio/midi"
	MIMEAUDIOMP4        MIMEType = "audio/mp4"
	MIMEAUDIOPLS        MIMEType = "audio/x-scpls"
	MIMEAUDIOPYA        MIMEType = "audio/vnd.ms-playready.media.pya"
	MIMEAUDIORAM        MIMEType = "audio/x-pn-realaudio"
	MIMEAUDIORIP        MIMEType = "audio/vnd.rip"
	MIMEAUDIORMI        MIMEType = "audio/midi"
	MIMEAUDIORMP        MIMEType = "audio/x-pn-realaudio-plugin"
	MIMEAUDIOSIL        MIMEType = "audio/silk"
	MIMEAUDIOSND        MIMEType = "audio/basic"
	MIMEAUDIOSTM        MIMEType = "audio/x-stm"
	MIMEAUDIOUVA        MIMEType = "audio/vnd.dece.audio"
	MIMEAUDIOWAX        MIMEType = "audio/x-ms-wax"
	MIMEAUDIOWEBA       MIMEType = "audio/webm"
	MIMEAUDIOWMA        MIMEType = "audio/x-ms-wma"
	MIMEAUDIOXM         MIMEType = "audio/xm"

	// 图片类型
	MIMEIMAGEUVI                     MIMEType = "image/vnd.dece.graphic"
	MIMEIMAGEX3DS                    MIMEType = "image/x-3ds"                      // 3DS 图像文件
	MIMEIMAGEAPNG                    MIMEType = "image/apng"                       // APNG 图像文件
	MIMEIMAGEXSONYARW                MIMEType = "image/x-sony-arw"                 // Sony ARW 图像文件
	MIMEIMAGEXPRSBTIF                MIMEType = "image/prs.btif"                   // BTIF 图像文件
	MIMEIMAGEXMSBMP                  MIMEType = "image/bmp"                        // BMP 图像文件
	MIMEIMAGEVNDAIRZIPACCELERATORAZV MIMEType = "image/vnd.airzip.accelerator.azv" // AirZip 加速器 AZV 图像文件
	MIMEIMAGECGM                     MIMEType = "image/cgm"                        // CGM 图像文件
	MIMEIMAGEXCMX                    MIMEType = "image/x-cmx"                      // CMX 图像文件
	MIMEIMAGEXCANONCR2               MIMEType = "image/x-canon-cr2"                // Canon CR2 图像文件
	MIMEIMAGEXCANONCRW               MIMEType = "image/x-canon-crw"                // Canon CRW 图像文件
	MIMEIMAGEJVU                     MIMEType = "image/vnd.djvu"                   // DjVu图像文件
	MIMEIMAGEDNG                     MIMEType = "image/x-adobe-dng"                // Adobe数码负片格式
	MIMEIMAGEDICOMRLE                MIMEType = "image/dicom-rle"                  // DICOM RLE压缩图像
	MIMEIMAGEDWG                     MIMEType = "image/vnd.dwg"                    // DWG工程图
	MIMEIMAGEDXF                     MIMEType = "image/vnd.dxf"                    // DXF工程图
	MIMEIMAGEMF                      MIMEType = "image/emf"                        // 增强图元文件
	MIMEIMAGEERF                     MIMEType = "image/x-epson-erf"                // Epson RAW格式
	MIMEIMAGEEXR                     MIMEType = "image/x-exr"                      // EXR图像文件
	MIMEIMAGEFLIF                    MIMEType = "image/x-flif"                     // FLIF图像文件
	MIMEIMAGEGIF                     MIMEType = "image/gif"                        // GIF图像文件
	MIMEIMAGEXACES                   MIMEType = "image/aces"                       // ACES图像
	MIMEIMAGEFREEHAND                MIMEType = "image/x-freehand"                 // FreeHand图像
	MIMEIMAGEFITS                    MIMEType = "image/fits"                       // FITS图像格式
	MIMEIMAGEFASTBIDSHEET            MIMEType = "image/vnd.fastbidsheet"           // FastBid图像
	MIMEIMAGEFLO                     MIMEType = "image/vnd.micrografx.flo"         // Micrografx Flo图像
	MIMEIMAGEFPX                     MIMEType = "image/vnd.fpx"                    // FlashPix图像
	MIMEIMAGEFST                     MIMEType = "image/vnd.fst"                    // FST图像格式
	MIMEIMAGEG3FAX                   MIMEType = "image/g3fax"                      // G3传真图像
	MIMEIMAGEEIF                     MIMEType = "image/ief"                        // IEF图像
	MIMEIMAGEXICON                   MIMEType = "image/x-icon"                     // ICO图标
	MIMEIMAGEHEIFC                   MIMEType = "image/heic"                       // HEIC图像
	MIMEIMAGEHEIFSERIES              MIMEType = "image/heic-sequence"              // HEIC图像序列
	MIMEIMAGEJLS                     MIMEType = "image/jls"                        // JPEG-LS图像
	MIMEIMAGEJNG                     MIMEType = "image/x-jng"                      // JNG图像
	MIMEIMAGEJP2                     MIMEType = "image/jp2"                        // JPEG 2000(jp2, jpg2)
	MIMEIMAGEJPEG                    MIMEType = "image/jpeg"                       // JPEG图像
	MIMEIMAGEJPX                     MIMEType = "image/jpx"                        // JPEG XR(jpf, jpx)
	MIMEIMAGEKODAKK25                MIMEType = "image/x-kodak-k25"                // Kodak K25图像
	MIMEIMAGEKODAKKDC                MIMEType = "image/x-kodak-kdc"                // Kodak KDC图像
	MIMEIMAGEKTX                     MIMEType = "image/ktx"                        // KTX纹理
	MIMEIMAGEMSMODI                  MIMEType = "image/vnd.ms-modi"                // MS MODI图像
	MIMEIMAGEMMR                     MIMEType = "image/vnd.fujixerox.edmics-mmr"   // Fujixerox MMR图像
	MIMEIMAGENETFPX                  MIMEType = "image/vnd.net-fpx"                // Net FPX图像
	MIMEIMAGEPBM                     MIMEType = "image/x-portable-bitmap"          // PBM图像
	MIMEIMAGEPHOTOCD                 MIMEType = "image/x-photo-cd"                 // Photo CD图像
	MIMEIMAGEPICT                    MIMEType = "image/x-pict"                     // PICT图像
	MIMEIMAGEPICT2                   MIMEType = "image/pict"                       // PICT图像
	MIMEIMAGEPICTURE                 MIMEType = "image/x-pcx"                      // PCX图像
	MIMEIMAGEPGM                     MIMEType = "image/x-portable-graymap"         // PGM图像
	MIMEIMAGEPNG                     MIMEType = "image/png"                        // PNG图像
	MIMEIMAGEPNM                     MIMEType = "image/x-portable-anymap"          // PNM图像
	MIMEIMAGEMACPAINT                MIMEType = "image/x-macpaint"                 // MacPaint图像
	MIMEIMAGEPPM                     MIMEType = "image/x-portable-pixmap"          // PPM图像
	MIMEIMAGEPHOTOSHOP               MIMEType = "image/vnd.adobe.photoshop"        // Photoshop图像
	MIMEIMAGECMURASTER               MIMEType = "image/x-cmu-raster"               // CMU光栅图
	MIMEIMAGERGB                     MIMEType = "image/x-rgb"                      // RGB图像
	MIMEIMAGEMRSID                   MIMEType = "image/x-mrsid-image"              // MrSID图像
	MIMEIMAGEWBMP                    MIMEType = "image/vnd.wap.wbmp"               // WBMP
	MIMEIMAGEMSPHOTOS                MIMEType = "image/vnd.ms-photo"               // MS照片
	MIMEIMAGEWEPB                    MIMEType = "image/webp"                       // WebP
	MIMEIMAGEXBITMAP                 MIMEType = "image/x-xbitmap"                  // X位图
	MIMEIMAGEXIFF                    MIMEType = "image/vnd.xiff"                   // XIFF
	MIMEIMAGEXPIXMAP                 MIMEType = "image/x-xpixmap"                  // X像素图
	MIMEIMAGEXWINDOWDUMP             MIMEType = "image/x-xwindowdump"              // X窗口转储
	MIMEIMAGEMINOLTA                 MIMEType = "image/x-minolta-mrw"              // Minolta RAW
	MIMEIMAGENIKON                   MIMEType = "image/x-nikon-nef"                // Nikon NEF

	// text
	MIMETEXTIN3D3DML      MIMEType = "text/vnd.in3d.3dml"               // 3DML
	MIMETEXTXASM          MIMEType = "text/x-asm"                       // 汇编语言文件
	MIMETEXTCACHEMANIFEST MIMEType = "text/cache-manifest"              // Cache Manifest 文件
	MIMETEXTCSS           MIMEType = "text/css"                         // CSS 文件
	MIMETEXTCSV           MIMEType = "text/csv"                         // CSV 文件
	MIMETEXTCOFFEESCRIPT  MIMEType = "text/coffeescript"                // CoffeeScript 文件
	MIMETEXTCURL          MIMEType = "text/vnd.curl"                    // Curl文本文件
	MIMETEXTDCURL         MIMEType = "text/vnd.curl.dcurl"              // Curl数据文件
	MIMETEXTPLAIN         MIMEType = "text/plain"                       // 纯文本文件
	MIMETEXTC             MIMEType = "text/x-c"                         // C语言源代码文件
	MIMETEXTLINESTAG      MIMEType = "text/prs.lines.tag"               // 行标记文本
	MIMETEXTSETEXT        MIMEType = "text/x-setext"                    // SEText格式
	MIMETEXTFORTRAN       MIMEType = "text/x-fortran"                   // Fortran源代码
	MIMETEXTFLEXSTOR      MIMEType = "text/vnd.fmi.flexstor"            // FlexStor文本
	MIMETEXTFLY           MIMEType = "text/vnd.fly"                     // Fly文本格式
	MIMETEXTGRAPHVIZ      MIMEType = "text/vnd.graphviz"                // Graphviz图形描述
	MIMETEXTHANDLEBARS    MIMEType = "text/x-handlebars-template"       // Handlebars模板
	MIMETEXTCOMPONENT     MIMEType = "text/x-component"                 // 组件文本
	MIMETEXTHTML          MIMEType = "text/html"                        // HTML文件
	MIMETEXTCALENDAR      MIMEType = "text/calendar"                    // iCalendar日历
	MIMETEXTCALENDAR3     MIMEType = "text/calendar"                    // iCalendar日历
	MIMETEXTJADE          MIMEType = "text/jade"                        // Jade模板
	MIMETEXTJAVASOURCE    MIMEType = "text/x-java-source"               // Java源代码
	MIMETEXTJSX           MIMEType = "text/jsx"                         // JSX代码
	MIMETEXTLESS          MIMEType = "text/less"                        // Less样式表
	MIMETEXTTROFF         MIMEType = "text/troff"                       // Troff文档
	MIMETEXTCURLMCURL     MIMEType = "text/vnd.curl.mcurl"              // Curl mcurl文件
	MIMETEXTJ2MEAPP       MIMEType = "text/vnd.sun.j2me.app-descriptor" // J2ME应用描述符
	MIMETEXTML            MIMEType = "text/ml"                          // ML文本
	MIMETEXTPASCAL        MIMEType = "text/x-pascal"                    // Pascal源代码
	MIMETEXTORG           MIMEType = "text/x-org"                       // Org文本
	MIMETEXTRICHTEXT      MIMEType = "text/richtext"                    // 富文本
	MIMETEXTASM           MIMEType = "text/x-asm"                       // 汇编代码
	MIMETEXTCURLSCURL     MIMEType = "text/vnd.curl.scurl"              // CURL脚本
	MIMETEXTSGML          MIMEType = "text/sgml"                        // SGML文档
	MIMETEXTSPOT          MIMEType = "text/vnd.in3d.spot"               // 3D点
	MIMETEXTSUBTITLE      MIMEType = "text/vnd.dvb.subtitle"            // DVB字幕
	MIMETEXTTABSEPARATED  MIMEType = "text/tab-separated-values"        // TSV
	MIMETEXTTURTLE        MIMEType = "text/turtle"                      // Turtle
	MIMETEXTURILIST       MIMEType = "text/uri-list"                    // URI列表
	MIMETEXTUUENCODE      MIMEType = "text/x-uuencode"                  // UUEncode编码
	MIMETEXTVCARD         MIMEType = "text/vcard"                       // vCard
	MIMETEXTVCARD2        MIMEType = "text/x-vcard"                     // vCard
	MIMETEXTVCALENDAR     MIMEType = "text/x-vcalendar"                 // vCalendar
	MIMETEXTWML           MIMEType = "text/vnd.wap.wml"                 // WML
	MIMETEXTWMLSCRIPT     MIMEType = "text/vnd.wap.wmlscript"           // WMLScript
	MIMETEXTYAML          MIMEType = "text/yaml"                        // YAML(yaml,yml)
	MIMETEXTYMP           MIMEType = "text/x-suse-ymp"                  // SUSE YMP
	MIMETEXTLUA           MIMEType = "text/x-lua"                       // Lua
	MIMETEXTMARKDOWN      MIMEType = "text/markdown"                    // Markdown(md,mkd)
	MIMETEXTMATHML        MIMEType = "text/mathml"                      // MathML
	MIMETEXTNFO           MIMEType = "text/x-nfo"                       // NFO
	MIMETEXTOPML          MIMEType = "text/x-opml"                      // OPML

	// chemical类型
	CHEMICALXCSML     MIMEType = "chemical/x-csml" // CSML 文件
	MIMECHEMICALXCMDF MIMEType = "chemical/x-cmdf" // CMDF 文件
	MIMECHEMICALXCML  MIMEType = "chemical/x-cml"  // CML 文件
	MIMECHEMICALXCIF  MIMEType = "chemical/x-cif"  // CIF 文件
	MIMECHEMICALXCDX  MIMEType = "chemical/x-cdx"  // CDX 文件

	// Model类型
	MIMEMODELCOLLADA   MIMEType = "model/vnd.collada+xml" // COLLADA 3D建模文件
	MIMEMODELDWF       MIMEType = "model/vnd.dwf"         // DWF 3D模型
	MIMEMODELGTW       MIMEType = "model/vnd.gtw"         // GTW模型文件
	MIMEMODELIGES      MIMEType = "model/iges"            // IGES工程图(iges, igs)
	MIMEMODELMESH      MIMEType = "model/mesh"            // 网格模型
	MIMEMODELMTS       MIMEType = "model/vnd.mts"         // MTS模型
	MIMEMODELVRML      MIMEType = "model/vrml"            // VRML
	MIMEMODELVTU       MIMEType = "model/vnd.vtu"         // VTU
	MIMEMODELX3DXML    MIMEType = "model/x3d+xml"         // X3D XML
	MIMEMODELX3DBINARY MIMEType = "model/x3d+binary"      // X3D二进制
	MIMEMODELX3DVRML   MIMEType = "model/x3d+vrml"        // X3D VRML
	MIMECHEMICALXYZ    MIMEType = "chemical/x-xyz"        // XYZ化学文件

	// Message类型
	MIMEMESSAGEDISPOSITION       MIMEType = "message/disposition-notification"        // 邮件回执通知
	MIMEMESSAGERFC822            MIMEType = "message/rfc822"                          // RFC822邮件
	MIMEMESSAGEGLOBALDELIVERY    MIMEType = "message/global-delivery-status"          // 全局投递状态
	MIMEMESSAGEGLOBALHEADERS     MIMEType = "message/global-headers"                  // 全局头部
	MIMEMESSAGEGLOBALDISPOSITION MIMEType = "message/global-disposition-notification" // 全局处理通知
	MIMEMESSAGEGLOBAL            MIMEType = "message/global"                          // 全局消息
	MIMEMESSAGEWSC               MIMEType = "message/vnd.wfa.wsc"                     // WSC

	// Font类型
	MIMEFONTCOLLECTION MIMEType = "font/collection" // 字体集合
	MIMEFONTTTF        MIMEType = "font/ttf"        // TrueType字体
	MIMEFONTWOFF       MIMEType = "font/woff"       // WOFF字体
	MIMEFONTWOFF2      MIMEType = "font/woff2"      // WOFF2字体
	MIMEFONTOTF        MIMEType = "font/otf"        // OpenType字体

	// Conference类型
	MIMECONFERENCECOOLTALK MIMEType = "x-conference/x-cooltalk" // CoolTalk会议

	// application
	MIMEAPPLICATIONM3U8                                      MIMEType = "application/vnd.apple.mpegurl" // 音视频播放列表
	MIMEAPPLICATIONMP21                                      MIMEType = "application/mp21"              // MP21文件
	MIMEAPPLICATIONSMZIP                                     MIMEType = "application/vnd.stepmania.package"
	MIMEAPPLICATIONUVVT                                      MIMEType = "application/vnd.dece.ttml+xml"
	MIMEAPPLICATIONVNDLOTUS123                               MIMEType = "application/vnd.lotus-1-2-3"                                             // Lotus 1-2-3 文件
	MIMEAPPLICATIONX7ZCOMPRESSED                             MIMEType = "application/x-7z-compressed"                                             // 7z 压缩文件
	MIMEAPPLICATIONXAUTHORWAREBIN                            MIMEType = "application/x-authorware-bin"                                            // Authorware 二进制文件
	MIMEAPPLICATIONXAUTHORWAREMAP                            MIMEType = "application/x-authorware-map"                                            // Authorware 地图文件
	MIMEAPPLICATIONXAUTHORWARESEG                            MIMEType = "application/x-authorware-seg"                                            // Authorware 分段文件
	MIMEAPPLICATIONXABIWORD                                  MIMEType = "application/x-abiword"                                                   // AbiWord 文档
	MIMEAPPLICATIONPKIXATTRCERT                              MIMEType = "application/pkix-attr-cert"                                              // PKIX 属性证书
	MIMEAPPLICATIONVNDAMERICANDYNAMICSACC                    MIMEType = "application/vnd.americandynamics.acc"                                    // American Dynamics ACC 文件
	MIMEAPPLICATIONXACECOMPRESSED                            MIMEType = "application/x-ace-compressed"                                            // ACE 压缩文件
	MIMEAPPLICATIONVNDACUCOBOL                               MIMEType = "application/vnd.acucobol"                                                // Acucobol 文件
	MIMEAPPLICATIONVNDACUCORP                                MIMEType = "application/vnd.acucorp"                                                 // Acucorp 文件
	MIMEAPPLICATIONVNDAUDIOGRAPH                             MIMEType = "application/vnd.audiograph"                                              // Audiograph 文件
	MIMEAPPLICATIONXFONTTYPE1                                MIMEType = "application/x-font-type1"                                                // Type 1 字体文件
	MIMEAPPLICATIONVNDIBMMODCAP                              MIMEType = "application/vnd.ibm.modcap"                                              // IBM ModCap 文件
	MIMEAPPLICATIONVNDAHEADSPACE                             MIMEType = "application/vnd.ahead.space"                                             // Ahead Space 文件
	MIMEAPPLICATIONPOSTSCRIPT                                MIMEType = "application/postscript"                                                  // PostScript 文件
	MIMEAPPLICATIONVNDADOBEAIRAPPLICATIONINSTALLERPACKAGEZIP MIMEType = "application/vnd.adobe.air-application-installer-package+zip"             // Adobe AIR 安装包
	MIMEAPPLICATIONVNDDVBAIT                                 MIMEType = "application/vnd.dvb.ait"                                                 // DVB AIT 文件
	MIMEAPPLICATIONVNDAMIGAAMI                               MIMEType = "application/vnd.amiga.ami"                                               // Amiga 文件
	MIMEAPPLICATIONVNDANDROIDPACKAGEARCHIVE                  MIMEType = "application/vnd.android.package-archive"                                 // Android 包文件
	MIMEAPPLICATIONXMSAPPLICATION                            MIMEType = "application/x-ms-application"                                            // Microsoft 应用程序文件
	MIMEAPPLICATIONVNDLOTUSAPPROACH                          MIMEType = "application/vnd.lotus-approach"                                          // Lotus Approach 文件
	MIMEAPPLICATIONXFREEARC                                  MIMEType = "application/x-freearc"                                                   // FreeArc 压缩文件
	MIMEAPPLICATIONXARJ                                      MIMEType = "application/x-arj"                                                       // ARJ 压缩文件
	MIMEAPPLICATIONPGPSIGNATURE                              MIMEType = "application/pgp-signature"                                               // PGP 签名文件
	MIMEAPPLICATIONVNDACCPACSIMPLYASO                        MIMEType = "application/vnd.accpac.simply.aso"                                       // Accpac Simply ASO 文件
	MIMEAPPLICATIONATOMXML                                   MIMEType = "application/atom+xml"                                                    // Atom XML 文件
	MIMEAPPLICATIONATOMCATXML                                MIMEType = "application/atomcat+xml"                                                 // AtomCat XML 文件
	MIMEAPPLICATIONATOMSVCXML                                MIMEType = "application/atomsvc+xml"                                                 // AtomSvc XML 文件
	MIMEAPPLICATIONVNDANTIXGAMECOMPONENT                     MIMEType = "application/vnd.antix.game-component"                                    // Antix 游戏组件文件
	MIMEAPPLICATIONAPPLIXWARE                                MIMEType = "application/applixware"                                                  // Applixware 文件
	MIMEAPPLICATIONVNDAIRZIPFILESECUREAZF                    MIMEType = "application/vnd.airzip.filesecure.azf"                                   // AirZip 文件安全 AZF 文件
	MIMEAPPLICATIONVNDAIRZIPFILESECUREAZS                    MIMEType = "application/vnd.airzip.filesecure.azs"                                   // AirZip 文件安全 AZS 文件
	MIMEAPPLICATIONVNDAMAZONBOOK                             MIMEType = "application/vnd.amazon.ebook"                                            // 亚马逊电子书
	MIMEAPPLICATIONXBCPIO                                    MIMEType = "application/x-bcpio"                                                     // BCPIO 文件
	MIMEAPPLICATIONXFONTBDF                                  MIMEType = "application/x-font-bdf"                                                  // BDF 字体文件
	MIMEAPPLICATIONVNDSYNCMLDMWBXML                          MIMEType = "application/vnd.syncml.dm+wbxml"                                         // SyncML WBXML 文件
	MIMEAPPLICATIONXBDOC                                     MIMEType = "application/x-bdoc"                                                      // BDOC 文件
	MIMEAPPLICATIONVNDRREALVNCBED                            MIMEType = "application/vnd.realvnc.bed"                                             // RealVNC BED 文件
	MIMEAPPLICATIONVNDFUJITSUOASYSPRS                        MIMEType = "application/vnd.fujitsu.oasysprs"                                        // Fujitsu OASYSPRS 文件
	MIMEAPPLICATIONOCTETSTREAM                               MIMEType = "application/octet-stream"                                                // 二进制流文件
	MIMEAPPLICATIONXBLORB                                    MIMEType = "application/x-blorb"                                                     // Blorb 文件
	MIMEAPPLICATIONVNDBMI                                    MIMEType = "application/vnd.bmi"                                                     // BMI 文件
	MIMEAPPLICATIONVNDFRAMEMAKER                             MIMEType = "application/vnd.framemaker"                                              // Framemaker 文件
	MIMEAPPLICATIONVDNPREVIEWSYSTEMSBOX                      MIMEType = "application/vnd.previewsystems.box"                                      // Preview Systems BOX 文件
	MIMEAPPLICATIONXBZ2                                      MIMEType = "application/x-bzip2"                                                     // BZIP2 压缩文件
	MIMEAPPLICATIONXBZIP                                     MIMEType = "application/x-bzip"                                                      // BZIP 压缩文件
	MIMEAPPLICATIONXBCBZ2                                    MIMEType = "application/x-bzip2"                                                     // BZIP2 压缩文件
	MIMEAPPLICATIONVNDCLUETRUSTCARTOMOBILECONFIG             MIMEType = "application/vnd.cluetrust.cartomobile-config"                            // Cluetrust Cartomobile 配置文件
	MIMEAPPLICATIONVNDCLUETRUSTCARTOMOBILECONFIGPKG          MIMEType = "application/vnd.cluetrust.cartomobile-config-pkg"                        // Cluetrust Cartomobile 配置包
	MIMEAPPLICATIONVNDCLONKC4GROUP                           MIMEType = "application/vnd.clonk.c4group"                                           // Clonk C4Group 文件
	MIMEAPPLICATIONVNDMSCABCOMPRESSED                        MIMEType = "application/vnd.ms-cab-compressed"                                       // Microsoft CAB 压缩文件
	MIMEAPPLICATIONVNDTCPDUMPCAP                             MIMEType = "application/vnd.tcpdump.pcap"                                            // TCPDump pcap 文件
	MIMEAPPLICATIONVNDCURLCAR                                MIMEType = "application/vnd.curl.car"                                                // Curl CAR 文件
	MIMEAPPLICATIONVNDMSPKISECCAT                            MIMEType = "application/vnd.ms-pki.seccat"                                           // Microsoft PKI 证书文件
	MIMEAPPLICATIONXCBR                                      MIMEType = "application/x-cbr"                                                       // CBA 压缩文件
	MIMEAPPLICATIONXCOCOA                                    MIMEType = "application/x-cocoa"                                                     // Cocoa 文件
	MIMEAPPLICATIONCCXML                                     MIMEType = "application/ccxml+xml"                                                   // CCXML 文件
	MIMEAPPLICATIONVNDCONTACTCMSG                            MIMEType = "application/vnd.contact.cmsg"                                            // Contact CMSG 文件
	MIMEAPPLICATIONXNETCDF                                   MIMEType = "application/x-netcdf"                                                    // NetCDF 文件
	MIMEAPPLICATIONVNDMEDIASTATIONCDKEY                      MIMEType = "application/vnd.mediastation.cdkey"                                      // MediaStation CDKey 文件
	MIMEAPPLICATIONCDMICAPABILITY                            MIMEType = "application/cdmi-capability"                                             // CDMI 能力文件
	MIMEAPPLICATIONCDMICONTAINER                             MIMEType = "application/cdmi-container"                                              // CDMI 容器文件
	MIMEAPPLICATIONCDMIDOMAIN                                MIMEType = "application/cdmi-domain"                                                 // CDMI 域文件
	MIMEAPPLICATIONCDMIOBJECT                                MIMEType = "application/cdmi-object"                                                 // CDMI 对象文件
	MIMEAPPLICATIONCDMIQUEUE                                 MIMEType = "application/cdmi-queue"                                                  // CDMI 队列文件
	MIMEAPPLICATIONVNDCHEMDRAWXML                            MIMEType = "application/vnd.chemdraw+xml"                                            // ChemDraw XML 文件
	MIMEAPPLICATIONVNDCINDERELLA                             MIMEType = "application/vnd.cinderella"                                              // Cinderella 文件
	MIMEAPPLICATIONPKIXCERT                                  MIMEType = "application/pkix-cert"                                                   // PKIX 证书文件
	MIMEAPPLICATIONXCFSCOMPRESSED                            MIMEType = "application/x-cfs-compressed"                                            // CFS 压缩文件
	MIMEAPPLICATIONXCHAT                                     MIMEType = "application/x-chat"                                                      // Chat 文件
	MIMEAPPLICATIONVNDMSHTMLHELP                             MIMEType = "application/vnd.ms-htmlhelp"                                             // MS HTMLHelp 文件
	MIMEAPPLICATIONVNDKDEKCHART                              MIMEType = "application/vnd.kde.kchart"                                              // KDE KChart 文件
	MIMEAPPLICATIONVNDANSERWEBCERTIFICATEISSUEINITIATION     MIMEType = "application/vnd.anser-web-certificate-issue-initiation"                  // Anser Web 证书请求
	MIMEAPPLICATIONVNDMSARTGALRY                             MIMEType = "application/vnd.ms-artgalry"                                             // MS ArtGalry 文件
	MIMEAPPLICATIONVNDCLAYMORE                               MIMEType = "application/vnd.claymore"                                                // Claymore 文件
	MIMEAPPLICATIONJAVAVM                                    MIMEType = "application/java-vm"                                                     // Java 虚拟机文件
	MIMEAPPLICATIONVNDCRICKCLICKERKEYBOARD                   MIMEType = "application/vnd.crick.clicker.keyboard"                                  // Crick Clicker 键盘文件
	MIMEAPPLICATIONVNDCRICKCLICKERPALETTE                    MIMEType = "application/vnd.crick.clicker.palette"                                   // Crick Clicker 调色板文件
	MIMEAPPLICATIONVNDCRICKCLICKERTEMPLATE                   MIMEType = "application/vnd.crick.clicker.template"                                  // Crick Clicker 模板文件
	MIMEAPPLICATIONVNDCRICKCLICKERWORDBANK                   MIMEType = "application/vnd.crick.clicker.wordbank"                                  // Crick Clicker 单词库文件
	MIMEAPPLICATIONVNDCRICKCLICKER                           MIMEType = "application/vnd.crick.clicker"                                           // Crick Clicker 文件
	MIMEAPPLICATIONXMSCLIP                                   MIMEType = "application/x-msclip"                                                    // MS Clip 文件
	MIMEAPPLICATIONVNDCOSMOCALLER                            MIMEType = "application/vnd.cosmocaller"                                             // Cosmocaller 文件
	MIMEAPPLICATIONVNDYELLOWRIVERCUSTOMMENU                  MIMEType = "application/vnd.yellowriver-custom-menu"                                 // YellowRiver 自定义菜单文件
	MIMEAPPLICATIONVNDRIMCOD                                 MIMEType = "application/vnd.rim.cod"                                                 // RIM COD 文件
	MIMEAPPLICATIONXMSDOWNLOAD                               MIMEType = "application/x-msdownload"                                                // MS 下载文件
	MIMEAPPLICATIONXCPIO                                     MIMEType = "application/x-cpio"                                                      // CPIO 文件
	MIMEAPPLICATIONCUSEEME                                   MIMEType = "application/cu-seeme"                                                    // CU-SeeMe视频会议文件
	MIMEAPPLICATIONCWW                                       MIMEType = "application/prs.cww"                                                     // CommonWorks工作文件
	MIMEAPPLICATIONDIRECTOR                                  MIMEType = "application/x-director"                                                  // Macromedia Director文件
	MIMEAPPLICATIONMOBIUSDAF                                 MIMEType = "application/vnd.mobius.daf"                                              // Mobius DAF文件
	MIMEAPPLICATIONDART                                      MIMEType = "application/vnd.dart"                                                    // Dart编程语言文件
	MIMEAPPLICATIONFDSNDATA                                  MIMEType = "application/vnd.fdsn.seed"                                               // FDSN地震数据文件
	MIMEAPPLICATIONDAVMOUNT                                  MIMEType = "application/davmount+xml"                                                // WebDAV挂载配置文件
	MIMEAPPLICATIONDOCBOOK                                   MIMEType = "application/docbook+xml"                                                 // DocBook文档文件
	MIMEAPPLICATIONDD2                                       MIMEType = "application/vnd.oma.dd2+xml"                                             // OMA DD2配置文件
	MIMEAPPLICATIONFUJIXEROXDDD                              MIMEType = "application/vnd.fujixerox.ddd"                                           // 富士施乐DDD文件
	MIMEAPPLICATIONDEB                                       MIMEType = "application/x-debian-package"                                            // Debian软件包
	MIMEAPPLICATIONX509CERT                                  MIMEType = "application/x-x509-ca-cert"                                              // X.509证书文件
	MIMEAPPLICATIONDREAMFACTORY                              MIMEType = "application/vnd.dreamfactory"                                            // DreamFactory应用文件
	MIMEAPPLICATIONDGC                                       MIMEType = "application/x-dgc-compressed"                                            // DGC压缩文件
	MIMEAPPLICATIONMOBIUSDIS                                 MIMEType = "application/vnd.mobius.dis"                                              // Mobius DIS文件
	MIMEAPPLICATIONMSEXE                                     MIMEType = "application/x-msdownload"                                                // Windows可执行文件
	MIMEAPPLICATIONDMG                                       MIMEType = "application/x-apple-diskimage"                                           // Apple磁盘镜像
	MIMEAPPLICATIONPCAP                                      MIMEType = "application/vnd.tcpdump.pcap"                                            // 网络抓包文件
	MIMEAPPLICATIONDNA                                       MIMEType = "application/vnd.dna"                                                     // DNA基因数据文件
	MIMEAPPLICATIONXFIG                                      MIMEType = "application/x-xfig"                                                      // XFig图形
	MIMEAPPLICATIONFORMSCENTRAL                              MIMEType = "application/vnd.adobe.formscentral.fcdt"                                 // Adobe FormsCentral
	MIMEAPPLICATIONISAC                                      MIMEType = "application/vnd.isac.fcs"                                                // ISAC数据
	MIMEAPPLICATIONFDF                                       MIMEType = "application/vnd.fdf"                                                     // PDF表单数据
	MIMEAPPLICATIONFCSELAYOUT                                MIMEType = "application/vnd.denovo.fcselayout-link"                                  // Layout链接
	MIMEAPPLICATIONOASYSGP                                   MIMEType = "application/vnd.fujitsu.oasysgp"                                         // Fujitsu文件
	MIMEAPPLICATIONNOVADIGMEXT                               MIMEType = "application/vnd.novadigm.ext"                                            // Novadigm EXT文件
	MIMEAPPLICATIONANDREWINSET                               MIMEType = "application/andrew-inset"                                                // Andrew Inset文件
	MIMEAPPLICATIONEZPIXALBUM                                MIMEType = "application/vnd.ezpix-album"                                             // Ezpix相册
	MIMEAPPLICATIONEZPIXPACKAGE                              MIMEType = "application/vnd.ezpix-package"                                           // Ezpix包
	MIMEAPPLICATIONEVA                                       MIMEType = "application/x-eva"                                                       // EVA文件
	MIMEAPPLICATIONENVOY                                     MIMEType = "application/x-envoy"                                                     // Envoy文件
	MIMEAPPLICATIONEXI                                       MIMEType = "application/exi"                                                         // EXI文件
	MIMEAPPLICATIONECMASCRIPT2                               MIMEType = "application/ecmascript"                                                  // ECMAScript
	MIMEAPPLICATIONESZIGNO                                   MIMEType = "application/vnd.eszigno3+xml"                                            // eSign XML文件
	MIMEAPPLICATIONOSGISUBSYSTEM                             MIMEType = "application/vnd.osgi.subsystem"                                          // OSGi子系统
	MIMEAPPLICATIONEPSONSFM                                  MIMEType = "application/vnd.epson.esf"                                               // Epson格式文件
	MIMEAPPLICATIONESZIGNO3                                  MIMEType = "application/vnd.eszigno3+xml"                                            // eSign XML
	MIMEAPPLICATIONEMMA                                      MIMEType = "application/emma+xml"                                                    // EMMA标记语言
	MIMEAPPLICATIONMSMETAFILE                                MIMEType = "application/x-msmetafile"                                                // 微软图元文件
	MIMEAPPLICATIONMSFONT                                    MIMEType = "application/vnd.ms-fontobject"                                           // 微软字体对象
	MIMEAPPLICATIONEPUB                                      MIMEType = "application/epub+zip"                                                    // EPUB电子书
	MIMEAPPLICATIONSPOTFIRE                                  MIMEType = "application/vnd.spotfire.dxp"                                            // Spotfire分析文件
	MIMEAPPLICATIONJAVAARCHIVE                               MIMEType = "application/java-archive"                                                // Java归档文件
	MIMEAPPLICATIONECMASCRIPT                                MIMEType = "application/ecmascript"                                                  // ECMAScript脚本
	MIMEAPPLICATIONNOVADIGM                                  MIMEType = "application/vnd.novadigm.edm"                                            // Novadigm EDM文件
	MIMEAPPLICATIONNOVADIGMEDX                               MIMEType = "application/vnd.novadigm.edx"                                            // Novadigm EDX文件
	MIMEAPPLICATIONPICSEL                                    MIMEType = "application/vnd.picsel"                                                  // Picsel文件
	MIMEAPPLICATIONOSASLI                                    MIMEType = "application/vnd.pg.osasli"                                               // OSASLI文件
	MIMEAPPLICATIONDSSCDER                                   MIMEType = "application/dssc+der"                                                    // DSSC DER编码文件
	MIMEAPPLICATIONDTBOOK                                    MIMEType = "application/x-dtbook+xml"                                                // DTBook电子书
	MIMEAPPLICATIONXMLDTD                                    MIMEType = "application/xml-dtd"                                                     // XML DTD文件
	MIMEAPPLICATIONDVI                                       MIMEType = "application/x-dvi"                                                       // DVI文档文件
	MIMEAPPLICATIONMSWORD                                    MIMEType = "application/msword"                                                      // Microsoft Word文档
	MIMEAPPLICATIONWORDMACRO                                 MIMEType = "application/vnd.ms-word.document.macroenabled.12"                        // Word启用宏的文档
	MIMEAPPLICATIONWORDOOXML                                 MIMEType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document" // Word OOXML文档
	MIMEAPPLICATIONWORDTEMPLATE                              MIMEType = "application/vnd.ms-word.template.macroenabled.12"                        // Word启用宏的模板
	MIMEAPPLICATIONWORDOOXMLTEMPLATE                         MIMEType = "application/vnd.openxmlformats-officedocument.wordprocessingml.template" // Word OOXML模板
	MIMEAPPLICATIONOSGIDP                                    MIMEType = "application/vnd.osgi.dp"                                                 // OSGi部署包
	MIMEAPPLICATIONDPGRAPH                                   MIMEType = "application/vnd.dpgraph"                                                 // DPGraph图表文件
	MIMEAPPLICATIONGROOVEIDENTITY                            MIMEType = "application/vnd.groove-identity-message"                                 // Groove身份消息
	MIMEAPPLICATIONGLTFBINARY                                MIMEType = "model/gltf-binary"                                                       // glTF二进制3D模型
	MIMEAPPLICATIONGLTFJSON                                  MIMEType = "model/gltf+json"                                                         // glTF JSON 3D模型
	MIMEAPPLICATIONGML                                       MIMEType = "application/gml+xml"                                                     // 地理标记语言
	MIMEAPPLICATIONGMX                                       MIMEType = "application/vnd.gmx"                                                     // GameMaker扩展
	MIMEAPPLICATIONGNUMERIC                                  MIMEType = "application/x-gnumeric"                                                  // Gnumeric电子表格
	MIMEAPPLICATIONFLOGRAPHIT                                MIMEType = "application/vnd.flographit"                                              // FloGraphit图表
	MIMEAPPLICATIONGPX                                       MIMEType = "application/gpx+xml"                                                     // GPS交换格式
	MIMEAPPLICATIONGRAFEQ                                    MIMEType = "application/vnd.grafeq"                                                  // GrafEq方程
	MIMEAPPLICATIONSRGS                                      MIMEType = "application/srgs"                                                        // 语音识别语法
	MIMEAPPLICATIONGRAMPSXML                                 MIMEType = "application/x-gramps-xml"                                                // Gramps家谱数据
	MIMEAPPLICATIONGEOMETRYEXPLORER                          MIMEType = "application/vnd.geometry-explorer"                                       // 几何探索器
	MIMEAPPLICATIONGROOVEINJECT                              MIMEType = "application/vnd.groove-injector"                                         // Groove注入器
	MIMEAPPLICATIONSRGSXML                                   MIMEType = "application/srgs+xml"                                                    // 语音识别语法XML
	MIMEAPPLICATIONGHOSTSCRIPT                               MIMEType = "application/x-font-ghostscript"                                          // Ghostscript字体
	MIMEAPPLICATIONGSHEET                                    MIMEType = "application/vnd.google-apps.spreadsheet"                                 // Google表格
	MIMEAPPLICATIONGSLIDES                                   MIMEType = "application/vnd.google-apps.presentation"                                // Google演示文稿
	MIMEAPPLICATIONGEOSPACE                                  MIMEType = "application/vnd.geospace"                                                // GeoSpace文件
	MIMEAPPLICATIONGROOVEACCOUNT                             MIMEType = "application/vnd.groove-account"                                          // Groove账户
	MIMEAPPLICATIONTADS                                      MIMEType = "application/x-tads"                                                      // TADS游戏文件
	MIMEAPPLICATIONRPKI                                      MIMEType = "application/rpki-ghostbusters"                                           // RPKI Ghostbusters
	MIMEAPPLICATIONGCA                                       MIMEType = "application/x-gca-compressed"                                            // GCA压缩文件
	MIMEAPPLICATIONGDL                                       MIMEType = "model/vnd.gdl"                                                           // GDL模型
	MIMEAPPLICATIONGOOGLEDOC                                 MIMEType = "application/vnd.google-apps.document"                                    // Google文档
	MIMEAPPLICATIONDYNAGEO                                   MIMEType = "application/vnd.dynageo"                                                 // DynaGeo文件
	MIMEAPPLICATIONGEOJSON                                   MIMEType = "application/geo+json"                                                    // GeoJSON地理数据
	MIMEAPPLICATIONGEOMETRY                                  MIMEType = "application/vnd.geometry-explorer"                                       // 几何探索器
	MIMEAPPLICATIONGEOGEBRA                                  MIMEType = "application/vnd.geogebra.file"                                           // GeoGebra文件
	MIMEAPPLICATIONGEOGEBRATOOL                              MIMEType = "application/vnd.geogebra.tool"                                           // GeoGebra工具
	MIMEAPPLICATIONGROOVEHELP                                MIMEType = "application/vnd.groove-help"                                             // Groove帮助
	MIMEAPPLICATIONFLUXTIME                                  MIMEType = "application/vnd.fluxtime.clip"                                           // FluxTime剪辑
	MIMEAPPLICATIONWEBFUNDS                                  MIMEType = "application/vnd.anser-web-funds-transfer-initiation"                     // Web资金转账
	MIMEAPPLICATIONFXP                                       MIMEType = "application/vnd.adobe.fxp"                                               // Adobe FXP文件
	MIMEAPPLICATIONFXPL                                      MIMEType = "application/vnd.adobe.fxp"                                               // Adobe FXP文件
	MIMEAPPLICATIONFUZZYSHEET                                MIMEType = "application/vnd.fuzzysheet"                                              // FuzzySheet文件
	MIMEAPPLICATIONGEOPLAN                                   MIMEType = "application/vnd.geoplan"                                                 // GeoPlan文件
	MIMEAPPLICATIONFRAMEMAKER                                MIMEType = "application/vnd.framemaker"                                              // FrameMaker文档
	MIMEAPPLICATIONFROGANS                                   MIMEType = "application/vnd.frogans.fnc"                                             // Frogans文件
	MIMEAPPLICATIONKIVIO                                     MIMEType = "application/vnd.kde.kivio"                                               // KDE Kivio文件
	MIMEAPPLICATIONGTAR                                      MIMEType = "application/x-gtar"                                                      // GNU Tar归档文件
	MIMEAPPLICATIONGROOVETOOL                                MIMEType = "application/vnd.groove-tool-message"                                     // Groove工具消息
	MIMEAPPLICATIONGXF                                       MIMEType = "application/gxf"                                                         // GXF文件
	MIMEAPPLICATIONGEONEXT                                   MIMEType = "application/vnd.geonext"                                                 // GEONExT几何文件
	MIMEAPPLICATIONGZIP                                      MIMEType = "application/gzip"                                                        // GZIP压缩文件
	MIMEAPPLICATIONHAL                                       MIMEType = "application/vnd.hal+xml"                                                 // HAL超媒体文件
	MIMEAPPLICATIONHBCI                                      MIMEType = "application/vnd.hbci"                                                    // HBCI银行接口文件
	MIMEAPPLICATIONVBOXHDD                                   MIMEType = "application/x-virtualbox-hdd"                                            // VirtualBox硬盘镜像
	MIMEAPPLICATIONHDF                                       MIMEType = "application/x-hdf"                                                       // HDF数据文件
	MIMEAPPLICATIONHJSON                                     MIMEType = "application/hjson"                                                       // HJSON数据格式
	MIMEAPPLICATIONWINHELP                                   MIMEType = "application/winhlp"                                                      // Windows帮助文件
	MIMEAPPLICATIONHPGL                                      MIMEType = "application/vnd.hp-hpgl"                                                 // HP图形语言
	MIMEAPPLICATIONHPID                                      MIMEType = "application/vnd.hp-hpid"                                                 // HP标识符
	MIMEAPPLICATIONHPS                                       MIMEType = "application/vnd.hp-hps"                                                  // HP服务
	MIMEAPPLICATIONBINHEX                                    MIMEType = "application/mac-binhex40"                                                // BinHex编码文件
	MIMEAPPLICATIONKENAMEAAPP                                MIMEType = "application/vnd.kenameaapp"                                              // Kenameaapp文件
	MIMEAPPLICATIONYAMAHADIC                                 MIMEType = "application/vnd.yamaha.hv-dic"                                           // Yamaha电子词典
	MIMEAPPLICATIONYAMAHAVOICE                               MIMEType = "application/vnd.yamaha.hv-voice"                                         // Yamaha语音
	MIMEAPPLICATIONYAMAHASCRIPT                              MIMEType = "application/vnd.yamaha.hv-script"                                        // Yamaha脚本
	MIMEAPPLICATIONINTERGEO                                  MIMEType = "application/vnd.intergeo"                                                // Intergeo文件
	MIMEAPPLICATIONICCPROFILE                                MIMEType = "application/vnd.iccprofile"                                              // ICC配置文件
	MIMEAPPLICATIONINFORMEDFORM                              MIMEType = "application/vnd.shana.informed.formdata"                                 // Informed表单数据
	MIMEAPPLICATIONIGLOADER                                  MIMEType = "application/vnd.igloader"                                                // IGLoader文件
	MIMEAPPLICATIONINSORSIGM                                 MIMEType = "application/vnd.insors.igm"                                              // Insors IGM文件
	MIMEAPPLICATIONMICROGRAFX                                MIMEType = "application/vnd.micrografx.igx"                                          // Micrografx IGX文件
	MIMEAPPLICATIONINFORMEDINTERCHANGE                       MIMEType = "application/vnd.shana.informed.interchange"                              // Informed数据交换
	MIMEAPPLICATIONACCPACIMP                                 MIMEType = "application/vnd.accpac.simply.imp"                                       // Accpac导入
	MIMEAPPLICATIONMSIMS                                     MIMEType = "application/vnd.ms-ims"                                                  // MS IMS
	MIMEAPPLICATIONINKML                                     MIMEType = "application/inkml+xml"                                                   // InkML标记
	MIMEAPPLICATIONINSTALL                                   MIMEType = "application/x-install-instructions"                                      // 安装说明
	MIMEAPPLICATIONASTRAEA                                   MIMEType = "application/vnd.astraea-software.iota"                                   // Astraea软件
	MIMEAPPLICATIONIPFIX                                     MIMEType = "application/ipfix"                                                       // IPFIX协议
	MIMEAPPLICATIONINFORMEDPACKAGE                           MIMEType = "application/vnd.shana.informed.package"                                  // Informed包
	MIMEAPPLICATIONIBMRIGHTS                                 MIMEType = "application/vnd.ibm.rights-management"                                   // IBM权限管理
	MIMEAPPLICATIONIREPOSITORY                               MIMEType = "application/vnd.irepository.package+xml"                                 // iRepository包
	MIMEAPPLICATIONISO                                       MIMEType = "application/x-iso9660-image"                                             // ISO镜像
	MIMEAPPLICATIONINFORMEDTEMPLATE                          MIMEType = "application/vnd.shana.informed.formtemplate"                             // Informed表单模板
	MIMEAPPLICATIONIMMERVISIONIVP                            MIMEType = "application/vnd.immervision-ivp"                                         // Immervision IVP
	MIMEAPPLICATIONIMMERVISIONIVU                            MIMEType = "application/vnd.immervision-ivu"                                         // Immervision IVU
	MIMEAPPLICATIONJAM                                       MIMEType = "application/vnd.jam"                                                     // JAM文件
	MIMEAPPLICATIONJAVADIFF                                  MIMEType = "application/x-java-archive-diff"                                         // Java归档差异
	MIMEAPPLICATIONJISP                                      MIMEType = "application/vnd.jisp"                                                    // JISP文件
	MIMEAPPLICATIONHPJLYT                                    MIMEType = "application/vnd.hp-jlyt"                                                 // HP JLYT
	MIMEAPPLICATIONJNLP                                      MIMEType = "application/x-java-jnlp-file"                                            // Java网络启动
	MIMEAPPLICATIONJODAARCHIVE                               MIMEType = "application/vnd.joost.joda-archive"                                      // Joda存档
	MIMEAPPLICATIONJAVASCRIPT                                MIMEType = "application/javascript"                                                  // JavaScript
	MIMEAPPLICATIONJSON                                      MIMEType = "application/json"                                                        // JSON数据
	MIMEAPPLICATIONJSON5                                     MIMEType = "application/json5"                                                       // JSON5数据
	MIMEAPPLICATIONJSONLD                                    MIMEType = "application/ld+json"                                                     // JSON-LD数据
	MIMEAPPLICATIONJSONML                                    MIMEType = "application/jsonml+json"                                                 // JsonML数据
	MIMEAPPLICATIONMETALINK4                                 MIMEType = "application/metalink4+xml"                                               // Metalink v4文件
	MIMEAPPLICATIONMETALINK                                  MIMEType = "application/metalink+xml"                                                // Metalink文件
	MIMEAPPLICATIONMETS                                      MIMEType = "application/mets+xml"                                                    // METS文档
	MIMEAPPLICATIONMFMP                                      MIMEType = "application/vnd.mfmp"                                                    // MFMP文件
	MIMEAPPLICATIONMAPGUIDE                                  MIMEType = "application/vnd.osgeo.mapguide.package"                                  // MapGuide包
	MIMEAPPLICATIONPROTEUS                                   MIMEType = "application/vnd.proteus.magazine"                                        // Proteus杂志
	MIMEAPPLICATIONMIE                                       MIMEType = "application/x-mie"                                                       // MIE文件
	MIMEAPPLICATIONMIF                                       MIMEType = "application/vnd.mif"                                                     // MIF文件
	MIMEAPPLICATIONKARAOKE                                   MIMEType = "application/vnd.chipnuts.karaoke-mmd"                                    // 卡拉OK MMD
	MIMEAPPLICATIONSMAF                                      MIMEType = "application/vnd.smaf"                                                    // SMAF文件
	MIMEAPPLICATIONMOBIEBOOK                                 MIMEType = "application/x-mobipocket-ebook"                                          // Mobi电子书
	MIMEAPPLICATIONMODS                                      MIMEType = "application/mods+xml"                                                    // MODS文档
	MIMEAPPLICATIONMOPHUNCERT                                MIMEType = "application/vnd.mophun.certificate"                                      // Mophun证书
	MIMEAPPLICATIONBLUEICE                                   MIMEType = "application/vnd.blueice.multipass"                                       // BlueIce多通道
	MIMEAPPLICATIONMOPHUNAPP                                 MIMEType = "application/vnd.mophun.application"                                      // Mophun应用
	MIMEAPPLICATIONMSPROJECT                                 MIMEType = "application/vnd.ms-project"                                              // MS Project
	MIMEAPPLICATIONMINIPAY                                   MIMEType = "application/vnd.ibm.minipay"                                             // IBM Mini Pay
	MIMEAPPLICATIONMOBIUSMQY                                 MIMEType = "application/vnd.mobius.mqy"                                              // Mobius MQY
	MIMEAPPLICATIONMARC                                      MIMEType = "application/marc"                                                        // MARC文件
	MIMEAPPLICATIONMARCXML                                   MIMEType = "application/marcxml+xml"                                                 // MARC XML
	MIMEAPPLICATIONMEDIASERVER                               MIMEType = "application/mediaservercontrol+xml"                                      // 媒体服务器控制
	MIMEAPPLICATIONMSEED                                     MIMEType = "application/vnd.fdsn.mseed"                                              // Mini-SEED数据
	MIMEAPPLICATIONMSEQ                                      MIMEType = "application/vnd.mseq"                                                    // MSEQ序列
	MIMEAPPLICATIONEPSONMSF                                  MIMEType = "application/vnd.epson.msf"                                               // Epson MSF
	MIMEAPPLICATIONMSDOWNLOAD                                MIMEType = "application/x-msdownload"                                                // MS下载文件
	MIMEAPPLICATIONMOBIUSMSL                                 MIMEType = "application/vnd.mobius.msl"                                              // Mobius MSL
	MIMEAPPLICATIONMUVEE                                     MIMEType = "application/vnd.muvee.style"                                             // Muvee样式
	MIMEAPPLICATIONMUSICIAN                                  MIMEType = "application/vnd.musician"                                                // Musician文件
	MIMEAPPLICATIONMUSICXML                                  MIMEType = "application/vnd.recordare.musicxml+xml"                                  // MusicXML
	MIMEAPPLICATIONMSMEDIAVIEW                               MIMEType = "application/x-msmediaview"                                               // MS媒体查看器
	MIMEAPPLICATIONMFER                                      MIMEType = "application/vnd.mfer"                                                    // MFER文件
	MIMEAPPLICATIONMXF                                       MIMEType = "application/mxf"                                                         // MXF文件
	MIMEAPPLICATIONMUSICXML2                                 MIMEType = "application/vnd.recordare.musicxml"                                      // MusicXML文件
	MIMEAPPLICATIONXVXML                                     MIMEType = "application/xv+xml"                                                      // XV XML文件
	MIMEAPPLICATIONTRISCAPE                                  MIMEType = "application/vnd.triscape.mxs"                                            // Triscape MXS
	MIMEAPPLICATIONLINUXPSF                                  MIMEType = "application/x-font-linux-psf"                                            // Linux PSF字体
	MIMEAPPLICATIONPSKC                                      MIMEType = "application/pskc+xml"                                                    // PSKC文档
	MIMEAPPLICATIONPTID                                      MIMEType = "application/vnd.pvi.ptid1"                                               // PTID文件
	MIMEAPPLICATIONMSPUB                                     MIMEType = "application/x-mspublisher"                                               // MS Publisher
	MIMEAPPLICATION3GPPBWVAR                                 MIMEType = "application/vnd.3gpp.pic-bw-var"                                         // 3GPP图片
	MIMEAPPLICATIONPOSTIT                                    MIMEType = "application/vnd.3m.post-it-notes"                                        // 3M便签
	MIMEAPPLICATIONQUICKANIME                                MIMEType = "application/vnd.epson.quickanime"                                        // Epson动画
	MIMEAPPLICATIONQBO                                       MIMEType = "application/vnd.intu.qbo"                                                // QuickBooks
	MIMEAPPLICATIONQFX                                       MIMEType = "application/vnd.intu.qfx"                                                // Quicken
	MIMEAPPLICATIONDELTATREE                                 MIMEType = "application/vnd.publishare-delta-tree"                                   // Delta树
	MIMEAPPLICATIONQUICKTIME                                 MIMEType = "application/x-quicktimeplayer"                                           // QuickTime
	MIMEAPPLICATIONQUARKXPRESS                               MIMEType = "application/vnd.quark.quarkxpress"                                       // QuarkXPress
	MIMEAPPLICATIONRAR                                       MIMEType = "application/x-rar-compressed"                                            // RAR压缩
	MIMEAPPLICATIONRAWDISK                                   MIMEType = "application/x-raw-disk-image"                                            // 原始磁盘镜像
	MIMEAPPLICATIONRDF                                       MIMEType = "application/rdf+xml"                                                     // RDF数据
	MIMEAPPLICATIONRDZ                                       MIMEType = "application/vnd.data-vision.rdz"                                         // RDZ数据
	MIMEAPPLICATIONBUSINESSOBJ                               MIMEType = "application/vnd.businessobjects"                                         // 商业对象
	MIMEAPPLICATIONDTBRESOURCE                               MIMEType = "application/x-dtbresource+xml"                                           // DTB资源
	MIMEAPPLICATIONREGINFO                                   MIMEType = "application/reginfo+xml"                                                 // 注册信息
	MIMEAPPLICATIONRIS                                       MIMEType = "application/x-research-info-systems"                                     // 研究信息系统
	MIMEAPPLICATIONRESOURCELIST                              MIMEType = "application/resource-lists+xml"                                          // 资源列表
	MIMEAPPLICATIONRESOURCELISTDIFF                          MIMEType = "application/resource-lists-diff+xml"                                     // 资源列表差异
	MIMEAPPLICATIONREALMEDIA                                 MIMEType = "application/vnd.rn-realmedia"                                            // RealMedia
	MIMEAPPLICATIONJAVAMEMIDLET                              MIMEType = "application/vnd.jcp.javame.midlet-rms"                                   // Java ME Midlet
	MIMEAPPLICATIONTROFF                                     MIMEType = "application/x-troff"                                                     // Troff
	MIMEAPPLICATIONRP9                                       MIMEType = "application/vnd.cloanto.rp9"                                             // RP9
	MIMEAPPLICATIONRPM                                       MIMEType = "application/x-rpm"                                                       // RPM包
	MIMEAPPLICATIONRSS                                       MIMEType = "application/rss+xml"                                                     // RSS源
	MIMEAPPLICATIONRTF                                       MIMEType = "application/rtf"                                                         // RTF文档
	MIMEAPPLICATIONSMAFAUDIO                                 MIMEType = "application/vnd.yamaha.smaf-audio"                                       // Yamaha SMAF音频
	MIMEAPPLICATIONSMAFPHRASE                                MIMEType = "application/vnd.yamaha.smaf-phrase"                                      // Yamaha SMAF短语
	MIMEAPPLICATIONSBML                                      MIMEType = "application/sbml+xml"                                                    // SBML
	MIMEAPPLICATIONSECURECONTAINER                           MIMEType = "application/vnd.ibm.secure-container"                                    // IBM安全容器
	MIMEAPPLICATIONSCHEDULE                                  MIMEType = "application/x-msschedule"                                                // MS日程表
	MIMEAPPLICATIONSCREENCAM                                 MIMEType = "application/vnd.lotus-screencam"                                         // Lotus屏幕录像
	MIMEAPPLICATIONSCVPREQUEST                               MIMEType = "application/scvp-cv-request"                                             // SCVP请求
	MIMEAPPLICATIONSCVPRESPONSE                              MIMEType = "application/scvp-cv-response"                                            // SCVP响应
	MIMEAPPLICATIONUSTAR                                     MIMEType = "application/x-ustar"                                                     // USTAR归档
	MIMEAPPLICATIONUIQTHEME                                  MIMEType = "application/vnd.uiq.theme"                                               // UIQ主题
	MIMEAPPLICATIONDECE                                      MIMEType = "application/vnd.dece.data"                                               // DECE数据
	MIMEAPPLICATIONDECETTML                                  MIMEType = "application/vnd.dece.ttml+xml"                                           // DECE TTML
	MIMEAPPLICATIONDECEUNSPEC                                MIMEType = "application/vnd.dece.unspecified"                                        // DECE未指定
	MIMEAPPLICATIONDECEZIP                                   MIMEType = "application/vnd.dece.zip"                                                // DECE压缩
	MIMEAPPLICATIONVBOX                                      MIMEType = "application/x-virtualbox-vbox"                                           // VirtualBox
	MIMEAPPLICATIONVBOXEXT                                   MIMEType = "application/x-virtualbox-vbox-extpack"                                   // VirtualBox扩展
	MIMEAPPLICATIONCDLINK                                    MIMEType = "application/x-cdlink"                                                    // CD链接
	MIMEAPPLICATIONGROOVECARD                                MIMEType = "application/vnd.groove-vcard"                                            // Groove名片
	MIMEAPPLICATIONVCX                                       MIMEType = "application/vnd.vcx"                                                     // VCX
	MIMEAPPLICATIONVDI                                       MIMEType = "application/x-virtualbox-vdi"                                            // VirtualBox磁盘
	MIMEAPPLICATIONVHD                                       MIMEType = "application/x-virtualbox-vhd"                                            // VirtualBox硬盘
	MIMEAPPLICATIONVISIONARY                                 MIMEType = "application/vnd.visionary"                                               // Visionary
	MIMEAPPLICATIONVMDK                                      MIMEType = "application/x-virtualbox-vmdk"                                           // VMDK
	MIMEAPPLICATIONSTARWRITER                                MIMEType = "application/vnd.stardivision.writer"                                     // StarWriter
	MIMEAPPLICATIONAUTHORWARE                                MIMEType = "application/x-authorware-bin"                                            // Authorware
	MIMEAPPLICATIONVISIO                                     MIMEType = "application/vnd.visio"                                                   // Visio
	MIMEAPPLICATIONVSF                                       MIMEType = "application/vnd.vsf"                                                     // VSF
	MIMEAPPLICATIONVOICEXML                                  MIMEType = "application/voicexml+xml"                                                // VoiceXML
	MIMEAPPLICATIONDOOM                                      MIMEType = "application/x-doom"                                                      // Doom
	MIMEAPPLICATIONWADL                                      MIMEType = "application/vnd.sun.wadl+xml"                                            // WADL
	MIMEAPPLICATIONWASM                                      MIMEType = "application/wasm"                                                        // WebAssembly
	MIMEAPPLICATIONWBS                                       MIMEType = "application/vnd.criticaltools.wbs+xml"                                   // WBS
	MIMEAPPLICATIONWBXML                                     MIMEType = "application/vnd.wap.wbxml"                                               // WBXML
	MIMEAPPLICATIONMSWORKS                                   MIMEType = "application/vnd.ms-works"                                                // MS Works
	MIMEAPPLICATIONWEBMANIFEST                               MIMEType = "application/manifest+json"                                               // Web Manifest
	MIMEAPPLICATIONWEBAPP                                    MIMEType = "application/x-web-app-manifest+json"                                     // Web App Manifest
	MIMEAPPLICATIONWIDGET                                    MIMEType = "application/widget"                                                      // Widget
	MIMEAPPLICATIONPMIWIDGET                                 MIMEType = "application/vnd.pmi.widget"                                              // PMI Widget
	MIMEAPPLICATIONWMD                                       MIMEType = "application/x-ms-wmd"                                                    // WMD
	MIMEAPPLICATIONMETAFILE                                  MIMEType = "application/x-msmetafile"                                                // MS Metafile
	MIMEAPPLICATIONWMLC                                      MIMEType = "application/vnd.wap.wmlc"                                                // WMLC
	MIMEAPPLICATIONWMLSCRIPTC                                MIMEType = "application/vnd.wap.wmlscriptc"                                          // WMLScript编译
	MIMEAPPLICATIONWMZ                                       MIMEType = "application/x-ms-wmz"                                                    // WMZ
	MIMEAPPLICATIONWORDPERFECT                               MIMEType = "application/vnd.wordperfect"                                             // WordPerfect
	MIMEAPPLICATIONWPL                                       MIMEType = "application/vnd.ms-wpl"                                                  // Windows播放列表
	MIMEAPPLICATIONWQD                                       MIMEType = "application/vnd.wqd"                                                     // WQD
	MIMEAPPLICATIONMSWRITE                                   MIMEType = "application/x-mswrite"                                                   // MS Write
	MIMEAPPLICATIONWSDL                                      MIMEType = "application/wsdl+xml"                                                    // WSDL
	MIMEAPPLICATIONWSPOLICY                                  MIMEType = "application/wspolicy+xml"                                                // WS-Policy
	MIMEAPPLICATIONWEBTURBO                                  MIMEType = "application/vnd.webturbo"                                                // WebTurbo
	MIMEAPPLICATIONXPR                                       MIMEType = "application/vnd.is-xpr"                                                  // XPR文件
	MIMEAPPLICATIONXPS                                       MIMEType = "application/vnd.ms-xpsdocument"                                          // XPS文档
	MIMEAPPLICATIONFORMNET                                   MIMEType = "application/vnd.intercon.formnet"                                        // FormNet
	MIMEAPPLICATIONXML                                       MIMEType = "application/xml"                                                         // XML(xsd,xsl)
	MIMEAPPLICATIONXSLT                                      MIMEType = "application/xslt+xml"                                                    // XSLT
	MIMEAPPLICATIONSYNCML                                    MIMEType = "application/vnd.syncml+xml"                                              // SyncML
	MIMEAPPLICATIONXSPF                                      MIMEType = "application/xspf+xml"                                                    // XSPF
	MIMEAPPLICATIONXUL                                       MIMEType = "application/vnd.mozilla.xul+xml"                                         // XUL
	MIMEAPPLICATIONXV                                        MIMEType = "application/xv+xml"                                                      // XV(xvm,xvml)
	MIMEAPPLICATIONXZ                                        MIMEType = "application/x-xz"                                                        // XZ压缩
	MIMEAPPLICATIONYANG                                      MIMEType = "application/yang"                                                        // YANG
	MIMEAPPLICATIONYIN                                       MIMEType = "application/yin+xml"                                                     // YIN
	MIMEAPPLICATIONZMACHINE                                  MIMEType = "application/x-zmachine"                                                  // Z-machine(z1-z8)
	MIMEAPPLICATIONZZAZZ                                     MIMEType = "application/vnd.zzazz.deck+xml"                                          // Zzazz
	MIMEAPPLICATIONZIP                                       MIMEType = "application/zip"                                                         // ZIP
	MIMEAPPLICATIONZUL                                       MIMEType = "application/vnd.zul"                                                     // ZUL(zir,zirz)
	MIMEAPPLICATIONHANDHELD                                  MIMEType = "application/vnd.handheld-entertainment+xml"                              // 手持设备
	MIMEAPPLICATIONSHORTCUT                                  MIMEType = "application/x-ms-shortcut"                                               // MS快捷方式
	MIMEAPPLICATIONLUABYTECODE                               MIMEType = "application/x-lua-bytecode"                                              // Lua字节码
	MIMEAPPLICATIONMOBIUS                                    MIMEType = "application/vnd.mobius.mbk"                                              // Mobius
	MIMEAPPLICATIONMEDCALC                                   MIMEType = "application/vnd.medcalcdata"                                             // MedCalc
	MIMEAPPLICATIONMCD                                       MIMEType = "application/vnd.mcd"                                                     // MCD
	MIMEAPPLICATIONDOLBY                                     MIMEType = "application/vnd.dolby.mlp"                                               // Dolby
	MIMEAPPLICATIONMSMONEY                                   MIMEType = "application/x-msmoney"                                                   // MS Money
	MIMEAPPLICATIONDASH                                      MIMEType = "application/dash+xml"                                                    // DASH
	MIMEAPPLICATIONOUTLOOK                                   MIMEType = "application/vnd.ms-outlook"                                              // MS Outlook
	MIMEAPPLICATIONMATHEMATICA                               MIMEType = "application/mathematica"                                                 // Mathematica
	MIMEAPPLICATIONWOLFRAM                                   MIMEType = "application/vnd.wolfram.player"                                          // Wolfram
	MIMEAPPLICATIONNITF                                      MIMEType = "application/vnd.nitf"                                                    // NITF
	MIMEAPPLICATIONNUMBERS                                   MIMEType = "application/vnd.apple.numbers"                                           // Numbers
	MIMEAPPLICATIONNZB                                       MIMEType = "application/x-nzb"                                                       // NZB
	MIMEAPPLICATIONOASYS2                                    MIMEType = "application/vnd.fujitsu.oasys2"                                          // Oasys 2
	MIMEAPPLICATIONOASYS3                                    MIMEType = "application/vnd.fujitsu.oasys3"                                          // Oasys 3
	MIMEAPPLICATIONMSBINDER                                  MIMEType = "application/x-msbinder"                                                  // MS Binder
	MIMEAPPLICATIONTGIF                                      MIMEType = "application/x-tgif"                                                      // TGIF
	MIMEAPPLICATIONPALM                                      MIMEType = "application/vnd.palm"                                                    // Palm
	MIMEAPPLICATIONRTSP                                      MIMEType = "application/x-rtsp"                                                      // RTSP流
	MIMEAPPLICATIONRTMP                                      MIMEType = "application/x-rtmp"
)

var mimeTypes = map[string]MIMEType{
	"123":                      "application/vnd.lotus-1-2-3",
	"3dml":                     "text/vnd.in3d.3dml",
	"3ds":                      "image/x-3ds",
	"3g2":                      "video/3gpp2",
	"3gp":                      "video/3gpp",
	"3gpp":                     "video/3gpp",
	"7z":                       "application/x-7z-compressed",
	"aab":                      "application/x-authorware-bin",
	"aac":                      "audio/x-aac",
	"aam":                      "application/x-authorware-map",
	"aas":                      "application/x-authorware-seg",
	"abw":                      "application/x-abiword",
	"ac":                       "application/pkix-attr-cert",
	"acc":                      "application/vnd.americandynamics.acc",
	"ace":                      "application/x-ace-compressed",
	"acu":                      "application/vnd.acucobol",
	"acutc":                    "application/vnd.acucorp",
	"adp":                      "audio/adpcm",
	"aep":                      "application/vnd.audiograph",
	"afm":                      "application/x-font-type1",
	"afp":                      "application/vnd.ibm.modcap",
	"ahead":                    "application/vnd.ahead.space",
	"ai":                       "application/postscript",
	"aif":                      "audio/x-aiff",
	"aifc":                     "audio/x-aiff",
	"aiff":                     "audio/x-aiff",
	"air":                      "application/vnd.adobe.air-application-installer-package+zip",
	"ait":                      "application/vnd.dvb.ait",
	"ami":                      "application/vnd.amiga.ami",
	"apk":                      "application/vnd.android.package-archive",
	"apng":                     "image/apng",
	"appcache":                 "text/cache-manifest",
	"application":              "application/x-ms-application",
	"apr":                      "application/vnd.lotus-approach",
	"arc":                      "application/x-freearc",
	"arj":                      "application/x-arj",
	"arw":                      "image/x-sony-arw",
	"asc":                      "application/pgp-signature",
	"asf":                      "video/x-ms-asf",
	"asm":                      "text/x-asm",
	"aso":                      "application/vnd.accpac.simply.aso",
	"asx":                      "video/x-ms-asf",
	"atc":                      "application/vnd.acucorp",
	"atom":                     "application/atom+xml",
	"atomcat":                  "application/atomcat+xml",
	"atomsvc":                  "application/atomsvc+xml",
	"atx":                      "application/vnd.antix.game-component",
	"au":                       "audio/basic",
	"avi":                      "video/x-msvideo",
	"aw":                       "application/applixware",
	"azf":                      "application/vnd.airzip.filesecure.azf",
	"azs":                      "application/vnd.airzip.filesecure.azs",
	"azv":                      "image/vnd.airzip.accelerator.azv",
	"azw":                      "application/vnd.amazon.ebook",
	"bat":                      "application/x-msdownload",
	"bcpio":                    "application/x-bcpio",
	"bdf":                      "application/x-font-bdf",
	"bdm":                      "application/vnd.syncml.dm+wbxml",
	"bdoc":                     "application/x-bdoc",
	"bed":                      "application/vnd.realvnc.bed",
	"bh2":                      "application/vnd.fujitsu.oasysprs",
	"bin":                      "application/octet-stream",
	"blb":                      "application/x-blorb",
	"blorb":                    "application/x-blorb",
	"bmi":                      "application/vnd.bmi",
	"bmp":                      "image/bmp",
	"book":                     "application/vnd.framemaker",
	"box":                      "application/vnd.previewsystems.box",
	"boz":                      "application/x-bzip2",
	"bpk":                      "application/octet-stream",
	"btif":                     "image/prs.btif",
	"buffer":                   "application/octet-stream",
	"bz":                       "application/x-bzip",
	"bz2":                      "application/x-bzip2",
	"c":                        "text/x-c",
	"c11amc":                   "application/vnd.cluetrust.cartomobile-config",
	"c11amz":                   "application/vnd.cluetrust.cartomobile-config-pkg",
	"c4d":                      "application/vnd.clonk.c4group",
	"c4f":                      "application/vnd.clonk.c4group",
	"c4g":                      "application/vnd.clonk.c4group",
	"c4p":                      "application/vnd.clonk.c4group",
	"c4u":                      "application/vnd.clonk.c4group",
	"cab":                      "application/vnd.ms-cab-compressed",
	"caf":                      "audio/x-caf",
	"cap":                      "application/vnd.tcpdump.pcap",
	"car":                      "application/vnd.curl.car",
	"cat":                      "application/vnd.ms-pki.seccat",
	"cb7":                      "application/x-cbr",
	"cba":                      "application/x-cbr",
	"cbr":                      "application/x-cbr",
	"cbt":                      "application/x-cbr",
	"cbz":                      "application/x-cbr",
	"cc":                       "text/x-c",
	"cco":                      "application/x-cocoa",
	"cct":                      "application/x-director",
	"ccxml":                    "application/ccxml+xml",
	"cdbcmsg":                  "application/vnd.contact.cmsg",
	"cdf":                      "application/x-netcdf",
	"cdkey":                    "application/vnd.mediastation.cdkey",
	"cdmia":                    "application/cdmi-capability",
	"cdmic":                    "application/cdmi-container",
	"cdmid":                    "application/cdmi-domain",
	"cdmio":                    "application/cdmi-object",
	"cdmiq":                    "application/cdmi-queue",
	"cdx":                      "chemical/x-cdx",
	"cdxml":                    "application/vnd.chemdraw+xml",
	"cdy":                      "application/vnd.cinderella",
	"cer":                      "application/pkix-cert",
	"cfs":                      "application/x-cfs-compressed",
	"cgm":                      "image/cgm",
	"chat":                     "application/x-chat",
	"chm":                      "application/vnd.ms-htmlhelp",
	"chrt":                     "application/vnd.kde.kchart",
	"cif":                      "chemical/x-cif",
	"cii":                      "application/vnd.anser-web-certificate-issue-initiation",
	"cil":                      "application/vnd.ms-artgalry",
	"cla":                      "application/vnd.claymore",
	"class":                    "application/java-vm",
	"clkk":                     "application/vnd.crick.clicker.keyboard",
	"clkp":                     "application/vnd.crick.clicker.palette",
	"clkt":                     "application/vnd.crick.clicker.template",
	"clkw":                     "application/vnd.crick.clicker.wordbank",
	"clkx":                     "application/vnd.crick.clicker",
	"clp":                      "application/x-msclip",
	"cmc":                      "application/vnd.cosmocaller",
	"cmdf":                     "chemical/x-cmdf",
	"cml":                      "chemical/x-cml",
	"cmp":                      "application/vnd.yellowriver-custom-menu",
	"cmx":                      "image/x-cmx",
	"cod":                      "application/vnd.rim.cod",
	"coffee":                   "text/coffeescript",
	"com":                      "application/x-msdownload",
	"conf":                     "text/plain",
	"cpio":                     "application/x-cpio",
	"cpp":                      "text/x-c",
	"cpt":                      "application/mac-compactpro",
	"cr2":                      "image/x-canon-cr2",
	"crd":                      "application/x-mscardfile",
	"crl":                      "application/pkix-crl",
	"crt":                      "application/x-x509-ca-cert",
	"crw":                      "image/x-canon-crw",
	"crx":                      "application/x-chrome-extension",
	"cryptonote":               "application/vnd.rig.cryptonote",
	"csh":                      "application/x-csh",
	"csl":                      "application/vnd.citationstyles.style+xml",
	"csml":                     "chemical/x-csml",
	"csp":                      "application/vnd.commonspace",
	"css":                      "text/css",
	"cst":                      "application/x-director",
	"csv":                      "text/csv",
	"cu":                       "application/cu-seeme",
	"curl":                     "text/vnd.curl",
	"cww":                      "application/prs.cww",
	"cxt":                      "application/x-director",
	"cxx":                      "text/x-c",
	"dae":                      "model/vnd.collada+xml",
	"daf":                      "application/vnd.mobius.daf",
	"dart":                     "application/vnd.dart",
	"dataless":                 "application/vnd.fdsn.seed",
	"davmount":                 "application/davmount+xml",
	"dbk":                      "application/docbook+xml",
	"dcr":                      "application/x-director",
	"dcurl":                    "text/vnd.curl.dcurl",
	"dd2":                      "application/vnd.oma.dd2+xml",
	"ddd":                      "application/vnd.fujixerox.ddd",
	"deb":                      "application/x-debian-package",
	"def":                      "text/plain",
	"deploy":                   "application/octet-stream",
	"der":                      "application/x-x509-ca-cert",
	"dfac":                     "application/vnd.dreamfactory",
	"dgc":                      "application/x-dgc-compressed",
	"dic":                      "text/x-c",
	"dir":                      "application/x-director",
	"dis":                      "application/vnd.mobius.dis",
	"disposition-notification": "message/disposition-notification",
	"dist":                     "application/octet-stream",
	"distz":                    "application/octet-stream",
	"djv":                      "image/vnd.djvu",
	"djvu":                     "image/vnd.djvu",
	"dll":                      "application/x-msdownload",
	"dmg":                      "application/x-apple-diskimage",
	"dmp":                      "application/vnd.tcpdump.pcap",
	"dms":                      "application/octet-stream",
	"dna":                      "application/vnd.dna",
	"dng":                      "image/x-adobe-dng",
	"doc":                      "application/msword",
	"docm":                     "application/vnd.ms-word.document.macroenabled.12",
	"docx":                     "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"dot":                      "application/msword",
	"dotm":                     "application/vnd.ms-word.template.macroenabled.12",
	"dotx":                     "application/vnd.openxmlformats-officedocument.wordprocessingml.template",
	"dp":                       "application/vnd.osgi.dp",
	"dpg":                      "application/vnd.dpgraph",
	"dra":                      "audio/vnd.dra",
	"drle":                     "image/dicom-rle",
	"dsc":                      "text/prs.lines.tag",
	"dssc":                     "application/dssc+der",
	"dtb":                      "application/x-dtbook+xml",
	"dtd":                      "application/xml-dtd",
	"dts":                      "audio/vnd.dts",
	"dtshd":                    "audio/vnd.dts.hd",
	"dump":                     "application/octet-stream",
	"dvb":                      "video/vnd.dvb.file",
	"dvi":                      "application/x-dvi",
	"dwf":                      "model/vnd.dwf",
	"dwg":                      "image/vnd.dwg",
	"dxf":                      "image/vnd.dxf",
	"dxp":                      "application/vnd.spotfire.dxp",
	"dxr":                      "application/x-director",
	"ear":                      "application/java-archive",
	"ecelp4800":                "audio/vnd.nuera.ecelp4800",
	"ecelp7470":                "audio/vnd.nuera.ecelp7470",
	"ecelp9600":                "audio/vnd.nuera.ecelp9600",
	"ecma":                     "application/ecmascript",
	"edm":                      "application/vnd.novadigm.edm",
	"edx":                      "application/vnd.novadigm.edx",
	"efif":                     "application/vnd.picsel",
	"ei6":                      "application/vnd.pg.osasli",
	"elc":                      "application/octet-stream",
	"emf":                      "image/emf",
	"eml":                      "message/rfc822",
	"emma":                     "application/emma+xml",
	"emz":                      "application/x-msmetafile",
	"eol":                      "audio/vnd.digital-winds",
	"eot":                      "application/vnd.ms-fontobject",
	"eps":                      "application/postscript",
	"epub":                     "application/epub+zip",
	"erf":                      "image/x-epson-erf",
	"es":                       "application/ecmascript",
	"es3":                      "application/vnd.eszigno3+xml",
	"esa":                      "application/vnd.osgi.subsystem",
	"esf":                      "application/vnd.epson.esf",
	"et3":                      "application/vnd.eszigno3+xml",
	"etx":                      "text/x-setext",
	"eva":                      "application/x-eva",
	"evy":                      "application/x-envoy",
	"exe":                      "application/x-msdownload",
	"exi":                      "application/exi",
	"exr":                      "image/aces",
	"ext":                      "application/vnd.novadigm.ext",
	"ez":                       "application/andrew-inset",
	"ez2":                      "application/vnd.ezpix-album",
	"ez3":                      "application/vnd.ezpix-package",
	"f":                        "text/x-fortran",
	"f4v":                      "video/x-f4v",
	"f77":                      "text/x-fortran",
	"f90":                      "text/x-fortran",
	"fbs":                      "image/vnd.fastbidsheet",
	"fcdt":                     "application/vnd.adobe.formscentral.fcdt",
	"fcs":                      "application/vnd.isac.fcs",
	"fdf":                      "application/vnd.fdf",
	"fe_launch":                "application/vnd.denovo.fcselayout-link",
	"fg5":                      "application/vnd.fujitsu.oasysgp",
	"fgd":                      "application/x-director",
	"fh":                       "image/x-freehand",
	"fh4":                      "image/x-freehand",
	"fh5":                      "image/x-freehand",
	"fh7":                      "image/x-freehand",
	"fhc":                      "image/x-freehand",
	"fig":                      "application/x-xfig",
	"fits":                     "image/fits",
	"flac":                     "audio/x-flac",
	"fli":                      "video/x-fli",
	"flo":                      "application/vnd.micrografx.flo",
	"flv":                      "video/x-flv",
	"flw":                      "application/vnd.kde.kivio",
	"flx":                      "text/vnd.fmi.flexstor",
	"fly":                      "text/vnd.fly",
	"fm":                       "application/vnd.framemaker",
	"fnc":                      "application/vnd.frogans.fnc",
	"for":                      "text/x-fortran",
	"fpx":                      "image/vnd.fpx",
	"frame":                    "application/vnd.framemaker",
	"fsc":                      "application/vnd.fsc.weblaunch",
	"fst":                      "image/vnd.fst",
	"ftc":                      "application/vnd.fluxtime.clip",
	"fti":                      "application/vnd.anser-web-funds-transfer-initiation",
	"fvt":                      "video/vnd.fvt",
	"fxp":                      "application/vnd.adobe.fxp",
	"fxpl":                     "application/vnd.adobe.fxp",
	"fzs":                      "application/vnd.fuzzysheet",
	"g2w":                      "application/vnd.geoplan",
	"g3":                       "image/g3fax",
	"g3w":                      "application/vnd.geospace",
	"gac":                      "application/vnd.groove-account",
	"gam":                      "application/x-tads",
	"gbr":                      "application/rpki-ghostbusters",
	"gca":                      "application/x-gca-compressed",
	"gdl":                      "model/vnd.gdl",
	"gdoc":                     "application/vnd.google-apps.document",
	"geo":                      "application/vnd.dynageo",
	"geojson":                  "application/geo+json",
	"gex":                      "application/vnd.geometry-explorer",
	"ggb":                      "application/vnd.geogebra.file",
	"ggt":                      "application/vnd.geogebra.tool",
	"ghf":                      "application/vnd.groove-help",
	"gif":                      "image/gif",
	"gim":                      "application/vnd.groove-identity-message",
	"glb":                      "model/gltf-binary",
	"gltf":                     "model/gltf+json",
	"gml":                      "application/gml+xml",
	"gmx":                      "application/vnd.gmx",
	"gnumeric":                 "application/x-gnumeric",
	"gph":                      "application/vnd.flographit",
	"gpx":                      "application/gpx+xml",
	"gqf":                      "application/vnd.grafeq",
	"gqs":                      "application/vnd.grafeq",
	"gram":                     "application/srgs",
	"gramps":                   "application/x-gramps-xml",
	"gre":                      "application/vnd.geometry-explorer",
	"grv":                      "application/vnd.groove-injector",
	"grxml":                    "application/srgs+xml",
	"gsf":                      "application/x-font-ghostscript",
	"gsheet":                   "application/vnd.google-apps.spreadsheet",
	"gslides":                  "application/vnd.google-apps.presentation",
	"gtar":                     "application/x-gtar",
	"gtm":                      "application/vnd.groove-tool-message",
	"gtw":                      "model/vnd.gtw",
	"gv":                       "text/vnd.graphviz",
	"gxf":                      "application/gxf",
	"gxt":                      "application/vnd.geonext",
	"gz":                       "application/gzip",
	"h":                        "text/x-c",
	"h261":                     "video/h261",
	"h263":                     "video/h263",
	"h264":                     "video/h264",
	"hal":                      "application/vnd.hal+xml",
	"hbci":                     "application/vnd.hbci",
	"hbs":                      "text/x-handlebars-template",
	"hdd":                      "application/x-virtualbox-hdd",
	"hdf":                      "application/x-hdf",
	"heic":                     "image/heic",
	"heics":                    "image/heic-sequence",
	"heif":                     "image/heif",
	"heifs":                    "image/heif-sequence",
	"hh":                       "text/x-c",
	"hjson":                    "application/hjson",
	"hlp":                      "application/winhlp",
	"hpgl":                     "application/vnd.hp-hpgl",
	"hpid":                     "application/vnd.hp-hpid",
	"hps":                      "application/vnd.hp-hps",
	"hqx":                      "application/mac-binhex40",
	"htc":                      "text/x-component",
	"htke":                     "application/vnd.kenameaapp",
	"htm":                      "text/html",
	"html":                     "text/html",
	"hvd":                      "application/vnd.yamaha.hv-dic",
	"hvp":                      "application/vnd.yamaha.hv-voice",
	"hvs":                      "application/vnd.yamaha.hv-script",
	"i2g":                      "application/vnd.intergeo",
	"icc":                      "application/vnd.iccprofile",
	"ice":                      "x-conference/x-cooltalk",
	"icm":                      "application/vnd.iccprofile",
	"ico":                      "image/x-icon",
	"ics":                      "text/calendar",
	"ief":                      "image/ief",
	"ifb":                      "text/calendar",
	"ifm":                      "application/vnd.shana.informed.formdata",
	"iges":                     "model/iges",
	"igl":                      "application/vnd.igloader",
	"igm":                      "application/vnd.insors.igm",
	"igs":                      "model/iges",
	"igx":                      "application/vnd.micrografx.igx",
	"iif":                      "application/vnd.shana.informed.interchange",
	"img":                      "application/octet-stream",
	"imp":                      "application/vnd.accpac.simply.imp",
	"ims":                      "application/vnd.ms-ims",
	"in":                       "text/plain",
	"ini":                      "text/plain",
	"ink":                      "application/inkml+xml",
	"inkml":                    "application/inkml+xml",
	"install":                  "application/x-install-instructions",
	"iota":                     "application/vnd.astraea-software.iota",
	"ipfix":                    "application/ipfix",
	"ipk":                      "application/vnd.shana.informed.package",
	"irm":                      "application/vnd.ibm.rights-management",
	"irp":                      "application/vnd.irepository.package+xml",
	"iso":                      "application/x-iso9660-image",
	"itp":                      "application/vnd.shana.informed.formtemplate",
	"ivp":                      "application/vnd.immervision-ivp",
	"ivu":                      "application/vnd.immervision-ivu",
	"jad":                      "text/vnd.sun.j2me.app-descriptor",
	"jade":                     "text/jade",
	"jam":                      "application/vnd.jam",
	"jar":                      "application/java-archive",
	"jardiff":                  "application/x-java-archive-diff",
	"java":                     "text/x-java-source",
	"jisp":                     "application/vnd.jisp",
	"jls":                      "image/jls",
	"jlt":                      "application/vnd.hp-jlyt",
	"jng":                      "image/x-jng",
	"jnlp":                     "application/x-java-jnlp-file",
	"joda":                     "application/vnd.joost.joda-archive",
	"jp2":                      "image/jp2",
	"jpe":                      "image/jpeg",
	"jpeg":                     "image/jpeg",
	"jpf":                      "image/jpx",
	"jpg":                      "image/jpeg",
	"jpg2":                     "image/jp2",
	"jpgm":                     "video/jpm",
	"jpgv":                     "video/jpeg",
	"jpm":                      "video/jpm",
	"jpx":                      "image/jpx",
	"js":                       "application/javascript",
	"json":                     "application/json",
	"json5":                    "application/json5",
	"jsonld":                   "application/ld+json",
	"jsonml":                   "application/jsonml+json",
	"jsx":                      "text/jsx",
	"k25":                      "image/x-kodak-k25",
	"kar":                      "audio/midi",
	"karbon":                   "application/vnd.kde.karbon",
	"kdc":                      "image/x-kodak-kdc",
	"keynote":                  "application/vnd.apple.keynote",
	"kfo":                      "application/vnd.kde.kformula",
	"kia":                      "application/vnd.kidspiration",
	"kml":                      "application/vnd.google-earth.kml+xml",
	"kmz":                      "application/vnd.google-earth.kmz",
	"kne":                      "application/vnd.kinar",
	"knp":                      "application/vnd.kinar",
	"kon":                      "application/vnd.kde.kontour",
	"kpr":                      "application/vnd.kde.kpresenter",
	"kpt":                      "application/vnd.kde.kpresenter",
	"kpxx":                     "application/vnd.ds-keypoint",
	"ksp":                      "application/vnd.kde.kspread",
	"ktr":                      "application/vnd.kahootz",
	"ktx":                      "image/ktx",
	"ktz":                      "application/vnd.kahootz",
	"kwd":                      "application/vnd.kde.kword",
	"kwt":                      "application/vnd.kde.kword",
	"lasxml":                   "application/vnd.las.las+xml",
	"latex":                    "application/x-latex",
	"lbd":                      "application/vnd.llamagraphics.life-balance.desktop",
	"lbe":                      "application/vnd.llamagraphics.life-balance.exchange+xml",
	"les":                      "application/vnd.hhe.lesson-player",
	"less":                     "text/less",
	"lha":                      "application/x-lzh-compressed",
	"link66":                   "application/vnd.route66.link66+xml",
	"list":                     "text/plain",
	"list3820":                 "application/vnd.ibm.modcap",
	"listafp":                  "application/vnd.ibm.modcap",
	"litcoffee":                "text/coffeescript",
	"lnk":                      "application/x-ms-shortcut",
	"log":                      "text/plain",
	"lostxml":                  "application/lost+xml",
	"lrf":                      "application/octet-stream",
	"lrm":                      "application/vnd.ms-lrm",
	"ltf":                      "application/vnd.frogans.ltf",
	"lua":                      "text/x-lua",
	"luac":                     "application/x-lua-bytecode",
	"lvp":                      "audio/vnd.lucent.voice",
	"lwp":                      "application/vnd.lotus-wordpro",
	"lzh":                      "application/x-lzh-compressed",
	"m13":                      "application/x-msmediaview",
	"m14":                      "application/x-msmediaview",
	"m1v":                      "video/mpeg",
	"m21":                      "application/mp21",
	"m2a":                      "audio/mpeg",
	"m2v":                      "video/mpeg",
	"m3a":                      "audio/mpeg",
	"m3u":                      "audio/x-mpegurl",
	"m3u8":                     "application/vnd.apple.mpegurl",
	"m4a":                      "audio/mp4",
	"m4p":                      "application/mp4",
	"m4u":                      "video/vnd.mpegurl",
	"m4v":                      "video/x-m4v",
	"ma":                       "application/mathematica",
	"mads":                     "application/mads+xml",
	"mag":                      "application/vnd.ecowin.chart",
	"maker":                    "application/vnd.framemaker",
	"man":                      "text/troff",
	"manifest":                 "text/cache-manifest",
	"map":                      "application/json",
	"mar":                      "application/octet-stream",
	"markdown":                 "text/markdown",
	"mathml":                   "application/mathml+xml",
	"mb":                       "application/mathematica",
	"mbk":                      "application/vnd.mobius.mbk",
	"mbox":                     "application/mbox",
	"mc1":                      "application/vnd.medcalcdata",
	"mcd":                      "application/vnd.mcd",
	"mcurl":                    "text/vnd.curl.mcurl",
	"md":                       "text/markdown",
	"mdb":                      "application/x-msaccess",
	"mdi":                      "image/vnd.ms-modi",
	"me":                       "text/troff",
	"mesh":                     "model/mesh",
	"meta4":                    "application/metalink4+xml",
	"metalink":                 "application/metalink+xml",
	"mets":                     "application/mets+xml",
	"mfm":                      "application/vnd.mfmp",
	"mft":                      "application/rpki-manifest",
	"mgp":                      "application/vnd.osgeo.mapguide.package",
	"mgz":                      "application/vnd.proteus.magazine",
	"mid":                      "audio/midi",
	"midi":                     "audio/midi",
	"mie":                      "application/x-mie",
	"mif":                      "application/vnd.mif",
	"mime":                     "message/rfc822",
	"mj2":                      "video/mj2",
	"mjp2":                     "video/mj2",
	"mjs":                      "application/javascript",
	"mk3d":                     "video/x-matroska",
	"mka":                      "audio/x-matroska",
	"mkd":                      "text/x-markdown",
	"mks":                      "video/x-matroska",
	"mkv":                      "video/x-matroska",
	"mlp":                      "application/vnd.dolby.mlp",
	"mmd":                      "application/vnd.chipnuts.karaoke-mmd",
	"mmf":                      "application/vnd.smaf",
	"mml":                      "text/mathml",
	"mmr":                      "image/vnd.fujixerox.edmics-mmr",
	"mng":                      "video/x-mng",
	"mny":                      "application/x-msmoney",
	"mobi":                     "application/x-mobipocket-ebook",
	"mods":                     "application/mods+xml",
	"mov":                      "video/quicktime",
	"movie":                    "video/x-sgi-movie",
	"mp2":                      "audio/mpeg",
	"mp21":                     "application/mp21",
	"mp2a":                     "audio/mpeg",
	"mp3":                      "audio/mpeg",
	"mp4":                      "video/mp4",
	"mp4a":                     "audio/mp4",
	"mp4s":                     "application/mp4",
	"mp4v":                     "video/mp4",
	"mpc":                      "application/vnd.mophun.certificate",
	"mpd":                      "application/dash+xml",
	"mpe":                      "video/mpeg",
	"mpeg":                     "video/mpeg",
	"mpg":                      "video/mpeg",
	"mpg4":                     "video/mp4",
	"mpga":                     "audio/mpeg",
	"mpkg":                     "application/vnd.apple.installer+xml",
	"mpm":                      "application/vnd.blueice.multipass",
	"mpn":                      "application/vnd.mophun.application",
	"mpp":                      "application/vnd.ms-project",
	"mpt":                      "application/vnd.ms-project",
	"mpy":                      "application/vnd.ibm.minipay",
	"mqy":                      "application/vnd.mobius.mqy",
	"mrc":                      "application/marc",
	"mrcx":                     "application/marcxml+xml",
	"mrw":                      "image/x-minolta-mrw",
	"ms":                       "text/troff",
	"mscml":                    "application/mediaservercontrol+xml",
	"mseed":                    "application/vnd.fdsn.mseed",
	"mseq":                     "application/vnd.mseq",
	"msf":                      "application/vnd.epson.msf",
	"msg":                      "application/vnd.ms-outlook",
	"msh":                      "model/mesh",
	"msi":                      "application/x-msdownload",
	"msl":                      "application/vnd.mobius.msl",
	"msm":                      "application/octet-stream",
	"msp":                      "application/octet-stream",
	"msty":                     "application/vnd.muvee.style",
	"mts":                      "model/vnd.mts",
	"mus":                      "application/vnd.musician",
	"musicxml":                 "application/vnd.recordare.musicxml+xml",
	"mvb":                      "application/x-msmediaview",
	"mwf":                      "application/vnd.mfer",
	"mxf":                      "application/mxf",
	"mxl":                      "application/vnd.recordare.musicxml",
	"mxml":                     "application/xv+xml",
	"mxs":                      "application/vnd.triscape.mxs",
	"mxu":                      "video/vnd.mpegurl",
	"n-gage":                   "application/vnd.nokia.n-gage.symbian.install",
	"n3":                       "text/n3",
	"nb":                       "application/mathematica",
	"nbp":                      "application/vnd.wolfram.player",
	"nc":                       "application/x-netcdf",
	"ncx":                      "application/x-dtbncx+xml",
	"nef":                      "image/x-nikon-nef",
	"nfo":                      "text/x-nfo",
	"ngdat":                    "application/vnd.nokia.n-gage.data",
	"nitf":                     "application/vnd.nitf",
	"nlu":                      "application/vnd.neurolanguage.nlu",
	"nml":                      "application/vnd.enliven",
	"nnd":                      "application/vnd.noblenet-directory",
	"nns":                      "application/vnd.noblenet-sealer",
	"nnw":                      "application/vnd.noblenet-web",
	"npx":                      "image/vnd.net-fpx",
	"nsc":                      "application/x-conference",
	"nsf":                      "application/vnd.lotus-notes",
	"ntf":                      "application/vnd.nitf",
	"numbers":                  "application/vnd.apple.numbers",
	"nzb":                      "application/x-nzb",
	"oa2":                      "application/vnd.fujitsu.oasys2",
	"oa3":                      "application/vnd.fujitsu.oasys3",
	"oas":                      "application/vnd.fujitsu.oasys",
	"obd":                      "application/x-msbinder",
	"obj":                      "application/x-tgif",
	"oda":                      "application/oda",
	"odb":                      "application/vnd.oasis.opendocument.database",
	"odc":                      "application/vnd.oasis.opendocument.chart",
	"odf":                      "application/vnd.oasis.opendocument.formula",
	"odft":                     "application/vnd.oasis.opendocument.formula-template",
	"odg":                      "application/vnd.oasis.opendocument.graphics",
	"odi":                      "application/vnd.oasis.opendocument.image",
	"odm":                      "application/vnd.oasis.opendocument.text-master",
	"odp":                      "application/vnd.oasis.opendocument.presentation",
	"ods":                      "application/vnd.oasis.opendocument.spreadsheet",
	"odt":                      "application/vnd.oasis.opendocument.text",
	"oga":                      "audio/ogg",
	"ogg":                      "audio/ogg",
	"ogv":                      "video/ogg",
	"ogx":                      "application/ogg",
	"omdoc":                    "application/omdoc+xml",
	"onepkg":                   "application/onenote",
	"onetmp":                   "application/onenote",
	"onetoc":                   "application/onenote",
	"onetoc2":                  "application/onenote",
	"opf":                      "application/oebps-package+xml",
	"opml":                     "text/x-opml",
	"oprc":                     "application/vnd.palm",
	"orf":                      "image/x-olympus-orf",
	"org":                      "text/x-org",
	"osf":                      "application/vnd.yamaha.openscoreformat",
	"osfpvg":                   "application/vnd.yamaha.openscoreformat.osfpvg+xml",
	"otc":                      "application/vnd.oasis.opendocument.chart-template",
	"otf":                      "font/otf",
	"otg":                      "application/vnd.oasis.opendocument.graphics-template",
	"oth":                      "application/vnd.oasis.opendocument.text-web",
	"oti":                      "application/vnd.oasis.opendocument.image-template",
	"otp":                      "application/vnd.oasis.opendocument.presentation-template",
	"ots":                      "application/vnd.oasis.opendocument.spreadsheet-template",
	"ott":                      "application/vnd.oasis.opendocument.text-template",
	"ova":                      "application/x-virtualbox-ova",
	"ovf":                      "application/x-virtualbox-ovf",
	"owl":                      "application/rdf+xml",
	"oxps":                     "application/oxps",
	"oxt":                      "application/vnd.openofficeorg.extension",
	"p":                        "text/x-pascal",
	"p10":                      "application/pkcs10",
	"p12":                      "application/x-pkcs12",
	"p7b":                      "application/x-pkcs7-certificates",
	"p7c":                      "application/pkcs7-mime",
	"p7m":                      "application/pkcs7-mime",
	"p7r":                      "application/x-pkcs7-certreqresp",
	"p7s":                      "application/pkcs7-signature",
	"p8":                       "application/pkcs8",
	"pac":                      "application/x-ns-proxy-autoconfig",
	"pages":                    "application/vnd.apple.pages",
	"pas":                      "text/x-pascal",
	"paw":                      "application/vnd.pawaafile",
	"pbd":                      "application/vnd.powerbuilder6",
	"pbm":                      "image/x-portable-bitmap",
	"pcap":                     "application/vnd.tcpdump.pcap",
	"pcf":                      "application/x-font-pcf",
	"pcl":                      "application/vnd.hp-pcl",
	"pclxl":                    "application/vnd.hp-pclxl",
	"pct":                      "image/x-pict",
	"pcurl":                    "application/vnd.curl.pcurl",
	"pcx":                      "image/x-pcx",
	"pdb":                      "application/x-pilot",
	"pde":                      "text/x-processing",
	"pdf":                      "application/pdf",
	"pef":                      "image/x-pentax-pef",
	"pem":                      "application/x-x509-ca-cert",
	"pfa":                      "application/x-font-type1",
	"pfb":                      "application/x-font-type1",
	"pfm":                      "application/x-font-type1",
	"pfr":                      "application/font-tdpfr",
	"pfx":                      "application/x-pkcs12",
	"pgm":                      "image/x-portable-graymap",
	"pgn":                      "application/x-chess-pgn",
	"pgp":                      "application/pgp-encrypted",
	"php":                      "application/x-httpd-php",
	"pic":                      "image/x-pict",
	"pkg":                      "application/octet-stream",
	"pki":                      "application/pkixcmp",
	"pkipath":                  "application/pkix-pkipath",
	"pkpass":                   "application/vnd.apple.pkpass",
	"pl":                       "application/x-perl",
	"plb":                      "application/vnd.3gpp.pic-bw-large",
	"plc":                      "application/vnd.mobius.plc",
	"plf":                      "application/vnd.pocketlearn",
	"pls":                      "application/pls+xml",
	"pm":                       "application/x-perl",
	"pml":                      "application/vnd.ctc-posml",
	"png":                      "image/png",
	"pnm":                      "image/x-portable-anymap",
	"portpkg":                  "application/vnd.macports.portpkg",
	"pot":                      "application/vnd.ms-powerpoint",
	"potm":                     "application/vnd.ms-powerpoint.template.macroenabled.12",
	"potx":                     "application/vnd.openxmlformats-officedocument.presentationml.template",
	"ppam":                     "application/vnd.ms-powerpoint.addin.macroenabled.12",
	"ppd":                      "application/vnd.cups-ppd",
	"ppm":                      "image/x-portable-pixmap",
	"pps":                      "application/vnd.ms-powerpoint",
	"ppsm":                     "application/vnd.ms-powerpoint.slideshow.macroenabled.12",
	"ppsx":                     "application/vnd.openxmlformats-officedocument.presentationml.slideshow",
	"ppt":                      "application/vnd.ms-powerpoint",
	"pptm":                     "application/vnd.ms-powerpoint.presentation.macroenabled.12",
	"pptx":                     "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	"pqa":                      "application/vnd.palm",
	"prc":                      "application/x-pilot",
	"pre":                      "application/vnd.lotus-freelance",
	"prf":                      "application/pics-rules",
	"ps":                       "application/postscript",
	"psb":                      "application/vnd.3gpp.pic-bw-small",
	"psd":                      "image/vnd.adobe.photoshop",
	"psf":                      "application/x-font-linux-psf",
	"pskcxml":                  "application/pskc+xml",
	"pti":                      "image/prs.pti",
	"ptid":                     "application/vnd.pvi.ptid1",
	"pub":                      "application/x-mspublisher",
	"pvb":                      "application/vnd.3gpp.pic-bw-var",
	"pwn":                      "application/vnd.3m.post-it-notes",
	"pya":                      "audio/vnd.ms-playready.media.pya",
	"pyv":                      "video/vnd.ms-playready.media.pyv",
	"qam":                      "application/vnd.epson.quickanime",
	"qbo":                      "application/vnd.intu.qbo",
	"qfx":                      "application/vnd.intu.qfx",
	"qps":                      "application/vnd.publishare-delta-tree",
	"qt":                       "video/quicktime",
	"qwd":                      "application/vnd.quark.quarkxpress",
	"qwt":                      "application/vnd.quark.quarkxpress",
	"qxb":                      "application/vnd.quark.quarkxpress",
	"qxd":                      "application/vnd.quark.quarkxpress",
	"qxl":                      "application/vnd.quark.quarkxpress",
	"qxt":                      "application/vnd.quark.quarkxpress",
	"ra":                       "audio/x-realaudio",
	"raf":                      "image/x-fuji-raf",
	"ram":                      "audio/x-pn-realaudio",
	"raml":                     "application/raml+yaml",
	"rar":                      "application/x-rar-compressed",
	"ras":                      "image/x-cmu-raster",
	"raw":                      "image/x-panasonic-raw",
	"rcprofile":                "application/vnd.ipunplugged.rcprofile",
	"rdf":                      "application/rdf+xml",
	"rdz":                      "application/vnd.data-vision.rdz",
	"rep":                      "application/vnd.businessobjects",
	"res":                      "application/x-dtbresource+xml",
	"rgb":                      "image/x-rgb",
	"rif":                      "application/reginfo+xml",
	"rip":                      "audio/vnd.rip",
	"ris":                      "application/x-research-info-systems",
	"rl":                       "application/resource-lists+xml",
	"rlc":                      "image/vnd.fujixerox.edmics-rlc",
	"rld":                      "application/resource-lists-diff+xml",
	"rm":                       "application/vnd.rn-realmedia",
	"rmi":                      "audio/midi",
	"rmp":                      "audio/x-pn-realaudio-plugin",
	"rms":                      "application/vnd.jcp.javame.midlet-rms",
	"rmvb":                     "application/vnd.rn-realmedia-vbr",
	"rnc":                      "application/relax-ng-compact-syntax",
	"rng":                      "application/xml",
	"roa":                      "application/rpki-roa",
	"roff":                     "text/troff",
	"rp9":                      "application/vnd.cloanto.rp9",
	"rpm":                      "application/x-redhat-package-manager",
	"rpss":                     "application/vnd.nokia.radio-presets",
	"rpst":                     "application/vnd.nokia.radio-preset",
	"rq":                       "application/sparql-query",
	"rs":                       "application/rls-services+xml",
	"rsd":                      "application/rsd+xml",
	"rss":                      "application/rss+xml",
	"rtf":                      "text/rtf",
	"rtx":                      "text/richtext",
	"run":                      "application/x-makeself",
	"s":                        "text/x-asm",
	"s3m":                      "audio/s3m",
	"saf":                      "application/vnd.yamaha.smaf-audio",
	"sass":                     "text/x-sass",
	"sbml":                     "application/sbml+xml",
	"sc":                       "application/vnd.ibm.secure-container",
	"scd":                      "application/x-msschedule",
	"scm":                      "application/vnd.lotus-screencam",
	"scq":                      "application/scvp-cv-request",
	"scs":                      "application/scvp-cv-response",
	"scss":                     "text/x-scss",
	"scurl":                    "text/vnd.curl.scurl",
	"sda":                      "application/vnd.stardivision.draw",
	"sdc":                      "application/vnd.stardivision.calc",
	"sdd":                      "application/vnd.stardivision.impress",
	"sdkd":                     "application/vnd.solent.sdkm+xml",
	"sdkm":                     "application/vnd.solent.sdkm+xml",
	"sdp":                      "application/sdp",
	"sdw":                      "application/vnd.stardivision.writer",
	"sea":                      "application/x-sea",
	"see":                      "application/vnd.seemail",
	"seed":                     "application/vnd.fdsn.seed",
	"sema":                     "application/vnd.sema",
	"semd":                     "application/vnd.semd",
	"semf":                     "application/vnd.semf",
	"ser":                      "application/java-serialized-object",
	"setpay":                   "application/set-payment-initiation",
	"setreg":                   "application/set-registration-initiation",
	"sfd-hdstx":                "application/vnd.hydrostatix.sof-data",
	"sfs":                      "application/vnd.spotfire.sfs",
	"sfv":                      "text/x-sfv",
	"sgi":                      "image/sgi",
	"sgl":                      "application/vnd.stardivision.writer-global",
	"sgm":                      "text/sgml",
	"sgml":                     "text/sgml",
	"sh":                       "application/x-sh",
	"shar":                     "application/x-shar",
	"shex":                     "text/shex",
	"shf":                      "application/shf+xml",
	"shtml":                    "text/html",
	"sid":                      "image/x-mrsid-image",
	"sig":                      "application/pgp-signature",
	"sil":                      "audio/silk",
	"silo":                     "model/mesh",
	"sis":                      "application/vnd.symbian.install",
	"sisx":                     "application/vnd.symbian.install",
	"sit":                      "application/x-stuffit",
	"sitx":                     "application/x-stuffitx",
	"skd":                      "application/vnd.koan",
	"skm":                      "application/vnd.koan",
	"skp":                      "application/vnd.koan",
	"skt":                      "application/vnd.koan",
	"sldm":                     "application/vnd.ms-powerpoint.slide.macroenabled.12",
	"sldx":                     "application/vnd.openxmlformats-officedocument.presentationml.slide",
	"slim":                     "text/slim",
	"slm":                      "text/slim",
	"slt":                      "application/vnd.epson.salt",
	"sm":                       "application/vnd.stepmania.stepchart",
	"smf":                      "application/vnd.stardivision.math",
	"smi":                      "application/smil+xml",
	"smil":                     "application/smil+xml",
	"smv":                      "video/x-smv",
	"smzip":                    "application/vnd.stepmania.package",
	"snd":                      "audio/basic",
	"snf":                      "application/x-font-snf",
	"so":                       "application/octet-stream",
	"spc":                      "application/x-pkcs7-certificates",
	"spf":                      "application/vnd.yamaha.smaf-phrase",
	"spl":                      "application/x-futuresplash",
	"spot":                     "text/vnd.in3d.spot",
	"spp":                      "application/scvp-vp-response",
	"spq":                      "application/scvp-vp-request",
	"spx":                      "audio/ogg",
	"sql":                      "application/x-sql",
	"sr2":                      "image/x-sony-sr2",
	"src":                      "application/x-wais-source",
	"srf":                      "image/x-sony-srf",
	"srt":                      "application/x-subrip",
	"sru":                      "application/sru+xml",
	"srx":                      "application/sparql-results+xml",
	"ssdl":                     "application/ssdl+xml",
	"sse":                      "application/vnd.kodak-descriptor",
	"ssf":                      "application/vnd.epson.ssf",
	"ssml":                     "application/ssml+xml",
	"st":                       "application/vnd.sailingtracker.track",
	"stc":                      "application/vnd.sun.xml.calc.template",
	"std":                      "application/vnd.sun.xml.draw.template",
	"stf":                      "application/vnd.wt.stf",
	"sti":                      "application/vnd.sun.xml.impress.template",
	"stk":                      "application/hyperstudio",
	"stl":                      "application/vnd.ms-pki.stl",
	"str":                      "application/vnd.pg.format",
	"stw":                      "application/vnd.sun.xml.writer.template",
	"styl":                     "text/stylus",
	"stylus":                   "text/stylus",
	"sub":                      "text/vnd.dvb.subtitle",
	"sus":                      "application/vnd.sus-calendar",
	"susp":                     "application/vnd.sus-calendar",
	"sv4cpio":                  "application/x-sv4cpio",
	"sv4crc":                   "application/x-sv4crc",
	"svc":                      "application/vnd.dvb.service",
	"svd":                      "application/vnd.svd",
	"svg":                      "image/svg+xml",
	"svgz":                     "image/svg+xml",
	"swa":                      "application/x-director",
	"swf":                      "application/x-shockwave-flash",
	"swi":                      "application/vnd.aristanetworks.swi",
	"sxc":                      "application/vnd.sun.xml.calc",
	"sxd":                      "application/vnd.sun.xml.draw",
	"sxg":                      "application/vnd.sun.xml.writer.global",
	"sxi":                      "application/vnd.sun.xml.impress",
	"sxm":                      "application/vnd.sun.xml.math",
	"sxw":                      "application/vnd.sun.xml.writer",
	"t":                        "text/troff",
	"t3":                       "application/x-t3vm-image",
	"t38":                      "image/t38",
	"taglet":                   "application/vnd.mynfc",
	"tao":                      "application/vnd.tao.intent-module-archive",
	"tap":                      "image/vnd.tencent.tap",
	"tar":                      "application/x-tar",
	"tcap":                     "application/vnd.3gpp2.tcap",
	"tcl":                      "application/x-tcl",
	"teacher":                  "application/vnd.smart.teacher",
	"tei":                      "application/tei+xml",
	"teicorpus":                "application/tei+xml",
	"tex":                      "application/x-tex",
	"texi":                     "application/x-texinfo",
	"texinfo":                  "application/x-texinfo",
	"text":                     "text/plain",
	"tfi":                      "application/thraud+xml",
	"tfm":                      "application/x-tex-tfm",
	"tfx":                      "image/tiff-fx",
	"tga":                      "image/x-tga",
	"thmx":                     "application/vnd.ms-officetheme",
	"tif":                      "image/tiff",
	"tiff":                     "image/tiff",
	"tk":                       "application/x-tcl",
	"tmo":                      "application/vnd.tmobile-livetv",
	"torrent":                  "application/x-bittorrent",
	"tpl":                      "application/vnd.groove-tool-template",
	"tpt":                      "application/vnd.trid.tpt",
	"tr":                       "text/troff",
	"tra":                      "application/vnd.trueapp",
	"trm":                      "application/x-msterminal",
	"ts":                       "video/mp2t",
	"tsd":                      "application/timestamped-data",
	"tsv":                      "text/tab-separated-values",
	"ttc":                      "font/collection",
	"ttf":                      "font/ttf",
	"ttl":                      "text/turtle",
	"twd":                      "application/vnd.simtech-mindmapper",
	"twds":                     "application/vnd.simtech-mindmapper",
	"txd":                      "application/vnd.genomatix.tuxedo",
	"txf":                      "application/vnd.mobius.txf",
	"txt":                      "text/plain",
	"u32":                      "application/x-authorware-bin",
	"u8dsn":                    "message/global-delivery-status",
	"u8hdr":                    "message/global-headers",
	"u8mdn":                    "message/global-disposition-notification",
	"u8msg":                    "message/global",
	"udeb":                     "application/x-debian-package",
	"ufd":                      "application/vnd.ufdl",
	"ufdl":                     "application/vnd.ufdl",
	"ulx":                      "application/x-glulx",
	"umj":                      "application/vnd.umajin",
	"unityweb":                 "application/vnd.unity",
	"uoml":                     "application/vnd.uoml+xml",
	"uri":                      "text/uri-list",
	"uris":                     "text/uri-list",
	"urls":                     "text/uri-list",
	"ustar":                    "application/x-ustar",
	"utz":                      "application/vnd.uiq.theme",
	"uu":                       "text/x-uuencode",
	"uva":                      "audio/vnd.dece.audio",
	"uvd":                      "application/vnd.dece.data",
	"uvf":                      "application/vnd.dece.data",
	"uvg":                      "image/vnd.dece.graphic",
	"uvh":                      "video/vnd.dece.hd",
	"uvi":                      "image/vnd.dece.graphic",
	"uvm":                      "video/vnd.dece.mobile",
	"uvp":                      "video/vnd.dece.pd",
	"uvs":                      "video/vnd.dece.sd",
	"uvt":                      "application/vnd.dece.ttml+xml",
	"uvu":                      "video/vnd.uvvu.mp4",
	"uvv":                      "video/vnd.dece.video",
	"uvva":                     "audio/vnd.dece.audio",
	"uvvd":                     "application/vnd.dece.data",
	"uvvf":                     "application/vnd.dece.data",
	"uvvg":                     "image/vnd.dece.graphic",
	"uvvh":                     "video/vnd.dece.hd",
	"uvvi":                     "image/vnd.dece.graphic",
	"uvvm":                     "video/vnd.dece.mobile",
	"uvvp":                     "video/vnd.dece.pd",
	"uvvs":                     "video/vnd.dece.sd",
	"uvvt":                     "application/vnd.dece.ttml+xml",
	"uvvu":                     "video/vnd.uvvu.mp4",
	"uvvv":                     "video/vnd.dece.video",
	"uvvx":                     "application/vnd.dece.unspecified",
	"uvvz":                     "application/vnd.dece.zip",
	"uvx":                      "application/vnd.dece.unspecified",
	"uvz":                      "application/vnd.dece.zip",
	"vbox":                     "application/x-virtualbox-vbox",
	"vbox-extpack":             "application/x-virtualbox-vbox-extpack",
	"vcard":                    "text/vcard",
	"vcd":                      "application/x-cdlink",
	"vcf":                      "text/x-vcard",
	"vcg":                      "application/vnd.groove-vcard",
	"vcs":                      "text/x-vcalendar",
	"vcx":                      "application/vnd.vcx",
	"vdi":                      "application/x-virtualbox-vdi",
	"vhd":                      "application/x-virtualbox-vhd",
	"vis":                      "application/vnd.visionary",
	"viv":                      "video/vnd.vivo",
	"vmdk":                     "application/x-virtualbox-vmdk",
	"vob":                      "video/x-ms-vob",
	"vor":                      "application/vnd.stardivision.writer",
	"vox":                      "application/x-authorware-bin",
	"vrml":                     "model/vrml",
	"vsd":                      "application/vnd.visio",
	"vsf":                      "application/vnd.vsf",
	"vss":                      "application/vnd.visio",
	"vst":                      "application/vnd.visio",
	"vsw":                      "application/vnd.visio",
	"vtf":                      "image/vnd.valve.source.texture",
	"vtt":                      "text/vtt",
	"vtu":                      "model/vnd.vtu",
	"vxml":                     "application/voicexml+xml",
	"w3d":                      "application/x-director",
	"wad":                      "application/x-doom",
	"wadl":                     "application/vnd.sun.wadl+xml",
	"war":                      "application/java-archive",
	"wasm":                     "application/wasm",
	"wav":                      "audio/x-wav",
	"wax":                      "audio/x-ms-wax",
	"wbmp":                     "image/vnd.wap.wbmp",
	"wbs":                      "application/vnd.criticaltools.wbs+xml",
	"wbxml":                    "application/vnd.wap.wbxml",
	"wcm":                      "application/vnd.ms-works",
	"wdb":                      "application/vnd.ms-works",
	"wdp":                      "image/vnd.ms-photo",
	"weba":                     "audio/webm",
	"webapp":                   "application/x-web-app-manifest+json",
	"webm":                     "video/webm",
	"webmanifest":              "application/manifest+json",
	"webp":                     "image/webp",
	"wg":                       "application/vnd.pmi.widget",
	"wgt":                      "application/widget",
	"wks":                      "application/vnd.ms-works",
	"wm":                       "video/x-ms-wm",
	"wma":                      "audio/x-ms-wma",
	"wmd":                      "application/x-ms-wmd",
	"wmf":                      "image/wmf",
	"wml":                      "text/vnd.wap.wml",
	"wmlc":                     "application/vnd.wap.wmlc",
	"wmls":                     "text/vnd.wap.wmlscript",
	"wmlsc":                    "application/vnd.wap.wmlscriptc",
	"wmv":                      "video/x-ms-wmv",
	"wmx":                      "video/x-ms-wmx",
	"wmz":                      "application/x-msmetafile",
	"woff":                     "font/woff",
	"woff2":                    "font/woff2",
	"wpd":                      "application/vnd.wordperfect",
	"wpl":                      "application/vnd.ms-wpl",
	"wps":                      "application/vnd.ms-works",
	"wqd":                      "application/vnd.wqd",
	"wri":                      "application/x-mswrite",
	"wrl":                      "model/vrml",
	"wsc":                      "message/vnd.wfa.wsc",
	"wsdl":                     "application/wsdl+xml",
	"wspolicy":                 "application/wspolicy+xml",
	"wtb":                      "application/vnd.webturbo",
	"wvx":                      "video/x-ms-wvx",
	"x32":                      "application/x-authorware-bin",
	"x3d":                      "model/x3d+xml",
	"x3db":                     "model/x3d+binary",
	"x3dbz":                    "model/x3d+binary",
	"x3dv":                     "model/x3d+vrml",
	"x3dvz":                    "model/x3d+vrml",
	"x3dz":                     "model/x3d+xml",
	"x3f":                      "image/x-sigma-x3f",
	"xaml":                     "application/xaml+xml",
	"xap":                      "application/x-silverlight-app",
	"xar":                      "application/vnd.xara",
	"xbap":                     "application/x-ms-xbap",
	"xbd":                      "application/vnd.fujixerox.docuworks.binder",
	"xbm":                      "image/x-xbitmap",
	"xdf":                      "application/xcap-diff+xml",
	"xdm":                      "application/vnd.syncml.dm+xml",
	"xdp":                      "application/vnd.adobe.xdp+xml",
	"xdssc":                    "application/dssc+xml",
	"xdw":                      "application/vnd.fujixerox.docuworks",
	"xenc":                     "application/xenc+xml",
	"xer":                      "application/patch-ops-error+xml",
	"xfdf":                     "application/vnd.adobe.xfdf",
	"xfdl":                     "application/vnd.xfdl",
	"xht":                      "application/xhtml+xml",
	"xhtml":                    "application/xhtml+xml",
	"xhvml":                    "application/xv+xml",
	"xif":                      "image/vnd.xiff",
	"xla":                      "application/vnd.ms-excel",
	"xlam":                     "application/vnd.ms-excel.addin.macroenabled.12",
	"xlc":                      "application/vnd.ms-excel",
	"xlf":                      "application/x-xliff+xml",
	"xlm":                      "application/vnd.ms-excel",
	"xls":                      "application/vnd.ms-excel",
	"xlsb":                     "application/vnd.ms-excel.sheet.binary.macroenabled.12",
	"xlsm":                     "application/vnd.ms-excel.sheet.macroenabled.12",
	"xlsx":                     "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"xlt":                      "application/vnd.ms-excel",
	"xltm":                     "application/vnd.ms-excel.template.macroenabled.12",
	"xltx":                     "application/vnd.openxmlformats-officedocument.spreadsheetml.template",
	"xlw":                      "application/vnd.ms-excel",
	"xm":                       "audio/xm",
	"xml":                      "text/xml",
	"xo":                       "application/vnd.olpc-sugar",
	"xop":                      "application/xop+xml",
	"xpi":                      "application/x-xpinstall",
	"xpl":                      "application/xproc+xml",
	"xpm":                      "image/x-xpixmap",
	"xpr":                      "application/vnd.is-xpr",
	"xps":                      "application/vnd.ms-xpsdocument",
	"xpw":                      "application/vnd.intercon.formnet",
	"xpx":                      "application/vnd.intercon.formnet",
	"xsd":                      "application/xml",
	"xsl":                      "application/xml",
	"xslt":                     "application/xslt+xml",
	"xsm":                      "application/vnd.syncml+xml",
	"xspf":                     "application/xspf+xml",
	"xul":                      "application/vnd.mozilla.xul+xml",
	"xvm":                      "application/xv+xml",
	"xvml":                     "application/xv+xml",
	"xwd":                      "image/x-xwindowdump",
	"xyz":                      "chemical/x-xyz",
	"xz":                       "application/x-xz",
	"yaml":                     "text/yaml",
	"yang":                     "application/yang",
	"yin":                      "application/yin+xml",
	"yml":                      "text/yaml",
	"ymp":                      "text/x-suse-ymp",
	"z1":                       "application/x-zmachine",
	"z2":                       "application/x-zmachine",
	"z3":                       "application/x-zmachine",
	"z4":                       "application/x-zmachine",
	"z5":                       "application/x-zmachine",
	"z6":                       "application/x-zmachine",
	"z7":                       "application/x-zmachine",
	"z8":                       "application/x-zmachine",
	"zaz":                      "application/vnd.zzazz.deck+xml",
	"zip":                      "application/zip",
	"zir":                      "application/vnd.zul",
	"zirz":                     "application/vnd.zul",
	"zmm":                      "application/vnd.handheld-entertainment+xml",
	"m4s":                      "video/iso.segment",
	"dash":                     "video/dash+xml",
	"hls":                      "video/hls",
	"fxm":                      "video/x-fxm",
	"divx":                     "video/divx",
	"ogm":                      "video/ogg",
	"rv":                       "video/vnd.rn-realvideo",
	"opus":                     "audio/opus",
	"m4b":                      "audio/mp4",
	"m4r":                      "audio/x-m4r",
	"gsm":                      "audio/x-gsm",
	"wv":                       "audio/wavpack",
	"ape":                      "audio/x-ape",
	"amr":                      "audio/amr",
	"avif":                     "image/avif",
	"jxr":                      "image/vnd.ms-photo",
	"jxl":                      "image/jxl",
	"xcf":                      "image/x-xcf",
	"cr3":                      "image/x-canon-cr3",
	"dib":                      "image/bmp",
}

// 存储MIME类型到扩展名的映射关系
var mimeToExtension = make(map[MIMEType]string)

func init() {
	// 初始化时建立MIME类型到扩展名的映射
	for ext, mime := range mimeTypes {
		// 如果已存在映射,优先使用更常用的扩展名
		if existing, ok := mimeToExtension[mime]; ok {
			if isPrimaryExtension(ext, existing) {
				mimeToExtension[mime] = ext
			}
		} else {
			mimeToExtension[mime] = ext
		}
	}
}

// isPrimaryExtension 判断是否为主要扩展名
func isPrimaryExtension(newExt, existingExt string) bool {
	// 定义一些扩展名的优先级规则
	primaryExts := map[string]bool{
		"jpeg": true, // 优先于jpg, jpe
		"tiff": true, // 优先于tif
		"html": true, // 优先于htm, xhtml
		"mp3":  true, // 优先于mpga
		"mp4":  true, // 优先于m4v, m4p
		"avi":  true, // 优先于divx, xvid
		"mkv":  true, // 优先于mk3d, mka
		"webm": true, // 优先于weba
		"flv":  true, // 优先于f4v
		"m2ts": true, // 优先于ts, mts
		"mov":  true, // 优先于qt
		"mpg":  true, // 优先于mpeg, m1v, m2v
		"mpeg": true, // 优先于mpg, mpe
		"mp4a": true, // 优先于m4a
		"xls":  true, // 优先于xlt
		"doc":  true, // 优先于dot
		"bmp":  true, // 优先于dib
	}

	return primaryExts[newExt]
}

// GetMIMETypeFromFileExtension returns the content type based on the extension of the file
func GetMIMETypeFromFileExtension(extension string) (MIMEType, bool) {
	ext := strings.ToLower(extension)
	if strings.HasPrefix(ext, ".") {
		// Remove the period
		ext = ext[1:]
	}
	mimeType, ok := mimeTypes[ext]
	if ok {
		return mimeType, true
	}
	return "application/unknown", ok
}

// GetFileExtensionFromMIMEType 根据MIME类型获取文件扩展名
func GetFileExtensionFromMIMEType(mimeType MIMEType) string {
	if ext, ok := mimeToExtension[mimeType]; ok {
		return ext
	}
	return ""
}

// IsVideoMIMEType 判断mime type是否为视频文件
func IsVideoMIMEType(mimeType MIMEType) bool {
	// 检查MIME类型字符串是否以"video/"开头
	return strings.HasPrefix(string(mimeType), "video/")
}

// IsAudioMIMEType 判断mime type是否为音频文件
func IsAudioMIMEType(mimeType MIMEType) bool {
	// 检查MIME类型字符串是否以"audio/"开头
	return strings.HasPrefix(string(mimeType), "audio/")
}

// IsImageMIMEType 判断mime type是否为图片文件
func IsImageMIMEType(mimeType MIMEType) bool {
	// 检查MIME类型字符串是否以"image/"开头
	return strings.HasPrefix(string(mimeType), "image/")
}

// IsTextMIMEType 判断mime type是否为文本文件
func IsTextMIMEType(mimeType MIMEType) bool {
	// 检查MIME类型字符串是否以"image/"开头
	return strings.HasPrefix(string(mimeType), "text/")
}

func IsArchiveMIMEType(mimeType MIMEType) bool {
	// 检查MIME类型字符串是否以"image/"开头
	mimeMap := map[MIMEType]struct{}{
		"application/zip":                   {},
		"application/gzip":                  {},
		"application/x-tar":                 {},
		"application/x-rar-compressed":      {},
		"application/x-7z-compressed":       {},
		"application/x-bzip2":               {},
		"application/x-xz":                  {},
		"application/x-lzma":                {},
		"application/x-arj-compressed":      {},
		"application/vnd.ms-cab-compressed": {},
		"application/x-compress":            {},
		"application/x-lzh-compressed":      {},
		"application/x-tar.xz":              {},
		"application/x-wim":                 {},
		"application/x-iso9660-image":       {},
		"application/x-lz4":                 {},
		"application/x-zpaq":                {},
		"application/x-ace-compressed":      {},
		"application/x-squashfs":            {},
		"application/x-apple-diskimage":     {},
	}

	_, ok := mimeMap[mimeType]
	return ok
}

func IsDocumentMIMEType(mimeType MIMEType) bool {
	mimeMap := map[MIMEType]struct{}{
		"doc":  {},
		"docm": {},
		"docx": {},
		"dotx": {},
		"gdoc": {},
		"xps":  {},
	}

	_, ok := mimeMap[mimeType]
	return ok
}

func IsPPTMIMEType(mimeType MIMEType) bool {
	mimeMap := map[MIMEType]struct{}{
		"ppt":  {},
		"pptm": {},
		"pptx": {},
	}

	_, ok := mimeMap[mimeType]
	return ok
}

func IsPdfMIMEType(mimeType MIMEType) bool {
	return mimeType == "application/pdf"
}
