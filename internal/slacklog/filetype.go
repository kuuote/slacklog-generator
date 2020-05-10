package slacklog

// https://api.slack.com/types/file
var filetypeToExtension = map[string]string{
	"auto":         "",             // Auto Detect Type,
	"text":         ".txt",         // Plain Text,
	"ai":           ".ai",          // Illustrator File,
	"apk":          ".apk",         // APK,
	"applescript":  ".applescript", // AppleScript,
	"binary":       "",             // Binary,
	"bmp":          ".bmp",         // Bitmap,
	"boxnote":      ".boxnote",     // BoxNote,
	"c":            ".c",           // C,
	"csharp":       ".cs",          // C#,
	"cpp":          ".cpp",         // C++,
	"css":          ".css",         // CSS,
	"csv":          ".csv",         // CSV,
	"clojure":      ".clj",         // Clojure,
	"coffeescript": ".coffee",      // CoffeeScript,
	"cfm":          ".cfm",         // ColdFusion,
	"d":            ".d",           // D,
	"dart":         ".dart",        // Dart,
	"diff":         ".diff",        // Diff,
	"doc":          ".doc",         // Word Document,
	"docx":         ".docx",        // Word document,
	"dockerfile":   ".dockerfile",  // Docker,
	"dotx":         ".dotx",        // Word template,
	"email":        ".eml",         // Email,
	"eps":          ".eps",         // EPS,
	"epub":         ".epub",        // EPUB,
	"erlang":       ".erl",         // Erlang,
	"fla":          ".fla",         // Flash FLA,
	"flv":          ".flv",         // Flash video,
	"fsharp":       ".fs",          // F#,
	"fortran":      ".f90",         // Fortran,
	"gdoc":         ".gdoc",        // GDocs Document,
	"gdraw":        ".gdraw",       // GDocs Drawing,
	"gif":          ".gif",         // GIF,
	"go":           ".go",          // Go,
	"gpres":        ".gpres",       // GDocs Presentation,
	"groovy":       ".groovy",      // Groovy,
	"gsheet":       ".gsheet",      // GDocs Spreadsheet,
	"gzip":         ".gz",          // GZip,
	"html":         ".html",        // HTML,
	"handlebars":   ".handlebars",  // Handlebars,
	"haskell":      ".hs",          // Haskell,
	"haxe":         ".hx",          // Haxe,
	"indd":         ".indd",        // InDesign Document,
	"java":         ".java",        // Java,
	"javascript":   ".js",          // JavaScript/JSON,
	"jpg":          ".jpeg",        // JPEG,
	"keynote":      ".keynote",     // Keynote Document,
	"kotlin":       ".kt",          // Kotlin,
	"latex":        ".tex",         // LaTeX/sTeX,
	"lisp":         ".lisp",        // Lisp,
	"lua":          ".lua",         // Lua,
	"m4a":          ".m4a",         // MPEG 4 audio,
	"markdown":     ".md",          // Markdown (raw),
	"matlab":       ".m",           // MATLAB,
	"mhtml":        ".mhtml",       // MHTML,
	"mkv":          ".mkv",         // Matroska video,
	"mov":          ".mov",         // QuickTime video,
	"mp3":          ".mp3",         // mp4,
	"mp4":          ".mp4",         // MPEG 4 video,
	"mpg":          ".mpeg",        // MPEG video,
	"mumps":        ".m",           // MUMPS,
	"numbers":      ".numbers",     // Numbers Document,
	"nzb":          ".nzb",         // NZB,
	"objc":         ".objc",        // Objective-C,
	"ocaml":        ".ml",          // OCaml,
	"odg":          ".odg",         // OpenDocument Drawing,
	"odi":          ".odi",         // OpenDocument Image,
	"odp":          ".odp",         // OpenDocument Presentation,
	"ods":          ".ods",         // OpenDocument Spreadsheet,
	"odt":          ".odt",         // OpenDocument Text,
	"ogg":          ".ogg",         // Ogg Vorbis,
	"ogv":          ".ogv",         // Ogg video,
	"pages":        ".pages",       // Pages Document,
	"pascal":       ".pp",          // Pascal,
	"pdf":          ".pdf",         // PDF,
	"perl":         ".pl",          // Perl,
	"php":          ".php",         // PHP,
	"pig":          ".pig",         // Pig,
	"png":          ".png",         // PNG,
	"post":         ".post",        // Slack Post,
	"powershell":   ".ps1",         // PowerShell,
	"ppt":          ".ppt",         // PowerPoint presentation,
	"pptx":         ".pptx",        // PowerPoint presentation,
	"psd":          ".psd",         // Photoshop Document,
	"puppet":       ".pp",          // Puppet,
	"python":       ".py",          // Python,
	"qtz":          ".qtz",         // Quartz Composer Composition,
	"r":            ".r",           // R,
	"rtf":          ".rtf",         // Rich Text File,
	"ruby":         ".rb",          // Ruby,
	"rust":         ".rs",          // Rust,
	"sql":          ".sql",         // SQL,
	"sass":         ".sass",        // Sass,
	"scala":        ".scala",       // Scala,
	"scheme":       ".scm",         // Scheme,
	"sketch":       ".sketch",      // Sketch File,
	"shell":        ".sh",          // Shell,
	"smalltalk":    ".st",          // Smalltalk,
	"svg":          ".svg",         // SVG,
	"swf":          ".swf",         // Flash SWF,
	"swift":        ".swift",       // Swift,
	"tar":          ".tar",         // Tarball,
	"tiff":         ".tiff",        // TIFF,
	"tsv":          ".tsv",         // TSV,
	"vb":           ".vb",          // VB.NET,
	"vbscript":     ".vbs",         // VBScript,
	"vcard":        ".vcf",         // vCard,
	"velocity":     ".vm",          // Velocity,
	"verilog":      ".v",           // Verilog,
	"wav":          ".wav",         // Waveform audio,
	"webm":         ".webm",        // WebM,
	"wmv":          ".wmv",         // Windows Media Video,
	"xls":          ".xls",         // Excel spreadsheet,
	"xlsx":         ".xlsx",        // Excel spreadsheet,
	"xlsb":         ".xlsb",        // Excel Spreadsheet (Binary, Macro Enabled),
	"xlsm":         ".xlsm",        // Excel Spreadsheet (Macro Enabled),
	"xltx":         ".xltx",        // Excel template,
	"xml":          ".xml",         // XML,
	"yaml":         ".yaml",        // YAML,
	"zip":          ".zip",         // Zip,
}