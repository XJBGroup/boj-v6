package types

type LanguageCode = uint8

const (
	LanguageGCCCpp LanguageCode = iota
	LanguageGNUCpp
	LanguageCLANGCpp
	LanguageJava
	LanguageJava8
	LanguagePython2
	LanguagePython3
	LanguageRust
	LanguageGCCCpp11
	LanguageGNUCpp11
	LanguageCLANGCpp11
	LanguageGCCCpp14
	LanguageGNUCpp14
	LanguageCLANGCpp14
	LanguageGCCCpp17
	LanguageGNUCpp17
	LanguageCLANGCpp17
	LanguageGCCC
	LanguageGNUC
)

var LanguageTypeMapping = map[string]LanguageCode{
	"gcc-c": LanguageGCCC,
	"gnu-c": LanguageGNUC,

	"gcc-c++":   LanguageGCCCpp,
	"gcc-c++11": LanguageGCCCpp11,
	"gcc-c++14": LanguageGCCCpp14,
	"gcc-c++17": LanguageGCCCpp17,

	"gnu-c++":   LanguageGNUCpp,
	"gnu-c++11": LanguageGNUCpp11,
	"gnu-c++14": LanguageGNUCpp14,
	"gnu-c++17": LanguageGNUCpp17,

	"clang-c++":   LanguageCLANGCpp,
	"clang-c++11": LanguageCLANGCpp11,
	"clang-c++14": LanguageCLANGCpp14,
	"clang-c++17": LanguageCLANGCpp17,

	"rust":    LanguageRust,
	"java":    LanguageJava,
	"java8":   LanguageJava8,
	"python2": LanguagePython2,
	"python3": LanguagePython3,
}

var LanguageSuffixMapping = map[LanguageCode]string{
	LanguageGCCC: ".c",
	LanguageGNUC: ".c",

	LanguageGCCCpp:   ".cpp",
	LanguageGCCCpp11: ".cpp",
	LanguageGCCCpp14: ".cpp",
	LanguageGCCCpp17: ".cpp",

	LanguageGNUCpp:   ".cpp",
	LanguageGNUCpp11: ".cpp",
	LanguageGNUCpp14: ".cpp",
	LanguageGNUCpp17: ".cpp",

	LanguageCLANGCpp:   ".cpp",
	LanguageCLANGCpp11: ".cpp",
	LanguageCLANGCpp14: ".cpp",
	LanguageCLANGCpp17: ".cpp",

	LanguageRust:    ".rs",
	LanguageJava:    ".java",
	LanguageJava8:   ".java",
	LanguagePython2: ".py",
	LanguagePython3: ".py",
}
