package mime_utils

var (
	MimeTypes map[string]string = map[string]string{
		"txt":        "text/plain",
		"3fr":        "image/x-hasselblad-3fr",
		"3gp":        "video/3gpp",
		"3gpp":       "video/3gpp",
		"7z":         "application/x-7z-compressed",
		"ai":         "application/postscript",
		"aif":        "audio/x-aiff",
		"aiff":       "audio/x-aiff",
		"apk":        "application/vnd.android.package-archive",
		"arw":        "image/x-sony-arw",
		"asf":        "video/x-ms-asf",
		"asx":        "video/x-ms-asf",
		"atom":       "application/atom+xml",
		"avi":        "video/x-msvideo",
		"bin":        "application/octet-stream",
		"bmp":        "image/x-ms-bmp",
		"bz2":        "application/x-bz2",
		"cab":        "application/vnd.ms-cab-compressed",
		"cco":        "application/x-cocoa",
		"cr2":        "image/x-canon-cr2",
		"crt":        "application/x-x509-ca-cert",
		"crw":        "image/x-canon-crw",
		"css":        "text/css",
		"csv":        "text/csv",
		"dcr":        "image/x-kodak-dcr",
		"deb":        "application/octet-stream",
		"der":        "application/x-x509-ca-cert",
		"dll":        "application/octet-stream",
		"dmg":        "application/octet-stream",
		"dng":        "image/x-adobe-dng",
		"doc":        "application/msword",
		"docx":       "application/word",
		"dpkg":       "application/dpkg-www-installer",
		"ds_store":   "application/octet-stream",
		"ear":        "application/java-archive",
		"eot":        "application/vnd.ms-fontobject",
		"eps":        "application/postscript",
		"erf":        "image/x-epson-erf",
		"exe":        "application/octet-stream",
		"flac":       "audio/flac",
		"flv":        "video/x-flv",
		"form":       "application/x-form",
		"gif":        "image/gif",
		"gz":         "application/x-gzip",
		"hqx":        "application/mac-binhex40",
		"htc":        "text/x-component",
		"htm":        "text/html",
		"html":       "text/html",
		"ico":        "image/x-icon",
		"ics":        "text/calendar",
		"img":        "application/octet-stream",
		"ini":        "text/x-ini",
		"iso":        "application/octet-stream",
		"jad":        "text/vnd.sun.j2me.app-descriptor",
		"jar":        "application/java-archive",
		"jardiff":    "application/x-java-archive-diff",
		"jng":        "image/x-jng",
		"jnlp":       "application/x-java-jnlp-file",
		"jpeg":       "image/jpeg",
		"jpg":        "image/jpeg",
		"js":         "application/javascript",
		"json":       "application/json",
		"kar":        "audio/midi",
		"kdc":        "image/x-kodak-kdc",
		"kml":        "application/vnd.google-earth.kml+xml",
		"kmz":        "application/vnd.google-earth.kmz",
		"m3u8":       "application/vnd.apple.mpegurl",
		"m4a":        "audio/x-m4a",
		"m4v":        "video/x-m4v",
		"md":         "text/markdown",
		"mdc":        "image/x-minolta-mdc",
		"mef":        "image/x-mamiya-mef",
		"mid":        "audio/midi",
		"midi":       "application/x-midi",
		"mkv":        "video/x-matroska",
		"mml":        "text/mathml",
		"mng":        "video/x-mng",
		"mos":        "image/x-aptus-mos",
		"mov":        "video/quicktime",
		"mp3":        "audio/mp3",
		"mp4":        "video/mp4",
		"mpeg":       "audio/mpeg",
		"mpg":        "video/mpeg",
		"mrw":        "image/x-minolta-mrw",
		"msi":        "application/octet-stream",
		"msm":        "application/octet-stream",
		"msp":        "application/octet-stream",
		"nef":        "image/x-nikon-nef",
		"nrw":        "image/x-nikon-nrw",
		"odg":        "application/vnd.oasis.opendocument.graphics",
		"odp":        "application/vnd.oasis.opendocument.presentation",
		"ods":        "application/vnd.oasis.opendocument.spreadsheet",
		"odt":        "application/vnd.oasis.opendocument.text",
		"ogg":        "audio/ogg",
		"ogv":        "application/ogg",
		"orf":        "image/x-olympus-orf",
		"org":        "text/org",
		"pdb":        "application/x-pilot",
		"pdf":        "application/pdf",
		"pef":        "image/x-pentax-pef",
		"pem":        "application/x-x509-ca-cert",
		"pkg":        "application/x-newton-compatible-pkg",
		"pl":         "application/x-perl",
		"pm":         "application/x-perl",
		"png":        "image/png",
		"pps":        "application/vnd.ms-powerpoint",
		"ppt":        "application/vnd.ms-powerpoint",
		"pptx":       "application/powerpoint",
		"prc":        "application/x-pilot",
		"properties": "text/x-ini",
		"ps":         "application/postscript",
		"ra":         "audio/x-realaudio",
		"raf":        "image/x-fuji-raf",
		"ram":        "audio/x-pn-realaudio",
		"rar":        "application/x-rar-compressed",
		"raw":        "image/x-raw",
		"rpm":        "application/x-redhat-package-manager",
		"rss":        "application/rss+xml",
		"rtf":        "application/rtf",
		"run":        "application/x-makeself",
		"rw2":        "image/x-panasonic-rw2",
		"sea":        "application/x-sea",
		"shtml":      "text/html",
		"sit":        "application/x-stuffit",
		"so":         "application/octet-stream",
		"sr2":        "image/x-sony-sr2",
		"srw":        "image/x-samsung-srw",
		"svg":        "image/svg+xml",
		"svgz":       "image/svg+xml",
		"swf":        "application/x-shockwave-flash",
		"tar":        "application/x-tar",
		"tcl":        "application/x-tcl",
		"tif":        "image/x-tif",
		"tiff":       "image/tiff",
		"tk":         "application/x-tcl",
		"ts":         "video/mp2t",
		"vcf":        "text/vcard",
		"vrml":       "application/x-vrml",
		"war":        "application/java-archive",
		"wav":        "audio/wave",
		"wbmp":       "image/vnd.wap.wbmp",
		"webm":       "video/webm",
		"webp":       "image/webp",
		"wma":        "audio/x-ms-wma",
		"wml":        "text/vnd.wap.wml",
		"wmlc":       "application/vnd.wap.wmlc",
		"wmv":        "video/x-ms-wmv",
		"woff":       "application/font-woff",
		"wrl":        "x-world/x-vrml",
		"x3f":        "image/x-x3f",
		"xhtml":      "application/xhtml+xml",
		"xls":        "application/vnd.ms-excel",
		"xlsx":       "application/excel",
		"xml":        "application/xml",
		"xpi":        "application/x-xpinstall",
		"xspf":       "application/xspf+xml",
		"zip":        "application/zip",
	}
)
