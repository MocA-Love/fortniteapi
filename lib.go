package fortniteapi

import (
	"time"
)

type Language string

const (
	LanguageAR     Language = "ar"
	LanguageDE     Language = "de"
	LanguageEN     Language = "en"
	LanguageES     Language = "es"
	LanguageES419  Language = "es-419"
	LanguageFR     Language = "fr"
	LanguageIT     Language = "it"
	LanguageJA     Language = "ja"
	LanguageKO     Language = "ko"
	LanguagePL     Language = "pl"
	LanguagePTBR   Language = "pt-BR"
	LanguageRU     Language = "ru"
	LanguageTR     Language = "tr"
	LanguageZH     Language = "zh-CN"
	LanguageZHHant Language = "zh-Hant"
)

type AccountType string

const (
	AccountTypeEpic     AccountType = "epic"
	AccountTypePSN      AccountType = "psn"
	AccountTypeXboxLive AccountType = "xbl"
)

type TimeWindow string

const (
	TimeWindowSeason   TimeWindow = "season"
	TimeWindowLifetime TimeWindow = "lifetime"
)

type ImageType string

const (
	ImageAll           ImageType = "all"
	ImageKeyboardMouse ImageType = "keyboardMouse"
	ImageGamepad       ImageType = "gamepad"
	ImageTouch         ImageType = "touch "
	ImageNone          ImageType = "none"
)

type KeyFormat string

const (
	KeyFormatHEX    KeyFormat = "hex"
	KeyFormatBase64 KeyFormat = "base64"
)

type MatchMethod string

const (
	MatchMethodFull     MatchMethod = "full"
	MatchMethodContains MatchMethod = "contains"
	MatchMethodStarts   MatchMethod = "starts"
	MatchMethodEnds     MatchMethod = "ends"
)

