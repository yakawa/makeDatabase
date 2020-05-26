// Code generated by "stringer -type=Type"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[INVALID - -1]
	_ = x[EOS-1]
	_ = x[IDENT-2]
	_ = x[KEYWORD-3]
	_ = x[SYMBOL-4]
	_ = x[NUMBER-5]
	_ = x[STRING-6]
	_ = x[DOUBLEQUOTE-10]
	_ = x[PERCENT-11]
	_ = x[AMPERSAND-12]
	_ = x[QUOTE-13]
	_ = x[LEFTPAREN-14]
	_ = x[RIGHTPAREN-15]
	_ = x[ASTERISK-16]
	_ = x[PLUSSIGN-17]
	_ = x[COMMA-18]
	_ = x[MINUSSIGN-19]
	_ = x[PERIOD-20]
	_ = x[SOLIDAS-21]
	_ = x[COLON-22]
	_ = x[SEMICOLON-23]
	_ = x[LESSTHAN-24]
	_ = x[EQUALS-25]
	_ = x[GREATERTHAN-26]
	_ = x[QUESTION-27]
	_ = x[LEFTBRACKET-28]
	_ = x[RIGHTBRACKET-29]
	_ = x[CIRCUMFLEX-30]
	_ = x[UNDERSCORE-31]
	_ = x[VERTICALBAR-32]
	_ = x[LEFTBRACE-33]
	_ = x[RIGHTBRACE-34]
	_ = x[NOTEQUALS-35]
	_ = x[GREATERTHANEQUALS-36]
	_ = x[LESSTHANEQUALS-37]
	_ = x[CONCAT-38]
	_ = x[K_ALL-100]
	_ = x[K_AND-101]
	_ = x[K_AS-102]
	_ = x[K_ASC-103]
	_ = x[K_BETWEEN-104]
	_ = x[K_BY-105]
	_ = x[K_CASE-106]
	_ = x[K_CAST-107]
	_ = x[K_COLLATE-108]
	_ = x[K_CROSS-109]
	_ = x[K_CURRENT-110]
	_ = x[K_CURRENT_DATE-111]
	_ = x[K_CURRENT_TIME-112]
	_ = x[K_CURRENT_TIMESTAMP-113]
	_ = x[K_DESC-114]
	_ = x[K_DISTINCT-115]
	_ = x[K_ELSE-116]
	_ = x[K_END-117]
	_ = x[K_EXCLUDE-118]
	_ = x[K_ESCAPE-119]
	_ = x[K_EXCEPT-120]
	_ = x[K_EXISTS-121]
	_ = x[K_FALSE-122]
	_ = x[K_FIRST-123]
	_ = x[K_FOLLOWING-124]
	_ = x[K_FROM-125]
	_ = x[K_GLOB-126]
	_ = x[K_GROUP-127]
	_ = x[K_GROUPS-128]
	_ = x[K_HAVING-129]
	_ = x[K_INTERSECT-130]
	_ = x[K_IN-131]
	_ = x[K_INDEXED-132]
	_ = x[K_INNER-133]
	_ = x[K_IS-134]
	_ = x[K_ISNULL-135]
	_ = x[K_JOIN-136]
	_ = x[K_LAST-137]
	_ = x[K_LEFT-138]
	_ = x[K_LIKE-139]
	_ = x[K_LIMIT-140]
	_ = x[K_MATCH-141]
	_ = x[K_NATURAL-142]
	_ = x[K_NO-143]
	_ = x[K_NOT-144]
	_ = x[K_NOTNULL-145]
	_ = x[K_NULL-146]
	_ = x[K_NULLS-147]
	_ = x[K_OFFSET-148]
	_ = x[K_ON-149]
	_ = x[K_ORDER-150]
	_ = x[K_OTHERS-151]
	_ = x[K_OUTER-152]
	_ = x[K_OVER-153]
	_ = x[K_PARTITION-154]
	_ = x[K_PRECEDING-155]
	_ = x[K_RANGE-156]
	_ = x[K_RECURSIVE-157]
	_ = x[K_REGEXP-158]
	_ = x[K_RIGHT-159]
	_ = x[K_ROW-160]
	_ = x[K_ROWS-161]
	_ = x[K_SELECT-162]
	_ = x[K_THEN-163]
	_ = x[K_TIES-164]
	_ = x[K_TRUE-165]
	_ = x[K_UNBOUNDED-166]
	_ = x[K_UNION-167]
	_ = x[K_USING-168]
	_ = x[K_VALUES-169]
	_ = x[K_WHEN-170]
	_ = x[K_WHERE-171]
	_ = x[K_WINDOW-172]
	_ = x[K_WITH-173]
}

const (
	_Type_name_0 = "INVALID"
	_Type_name_1 = "EOSIDENTKEYWORDSYMBOLNUMBERSTRING"
	_Type_name_2 = "DOUBLEQUOTEPERCENTAMPERSANDQUOTELEFTPARENRIGHTPARENASTERISKPLUSSIGNCOMMAMINUSSIGNPERIODSOLIDASCOLONSEMICOLONLESSTHANEQUALSGREATERTHANQUESTIONLEFTBRACKETRIGHTBRACKETCIRCUMFLEXUNDERSCOREVERTICALBARLEFTBRACERIGHTBRACENOTEQUALSGREATERTHANEQUALSLESSTHANEQUALSCONCAT"
	_Type_name_3 = "K_ALLK_ANDK_ASK_ASCK_BETWEENK_BYK_CASEK_CASTK_COLLATEK_CROSSK_CURRENTK_CURRENT_DATEK_CURRENT_TIMEK_CURRENT_TIMESTAMPK_DESCK_DISTINCTK_ELSEK_ENDK_EXCLUDEK_ESCAPEK_EXCEPTK_EXISTSK_FALSEK_FIRSTK_FOLLOWINGK_FROMK_GLOBK_GROUPK_GROUPSK_HAVINGK_INTERSECTK_INK_INDEXEDK_INNERK_ISK_ISNULLK_JOINK_LASTK_LEFTK_LIKEK_LIMITK_MATCHK_NATURALK_NOK_NOTK_NOTNULLK_NULLK_NULLSK_OFFSETK_ONK_ORDERK_OTHERSK_OUTERK_OVERK_PARTITIONK_PRECEDINGK_RANGEK_RECURSIVEK_REGEXPK_RIGHTK_ROWK_ROWSK_SELECTK_THENK_TIESK_TRUEK_UNBOUNDEDK_UNIONK_USINGK_VALUESK_WHENK_WHEREK_WINDOWK_WITH"
)

var (
	_Type_index_1 = [...]uint8{0, 3, 8, 15, 21, 27, 33}
	_Type_index_2 = [...]uint16{0, 11, 18, 27, 32, 41, 51, 59, 67, 72, 81, 87, 94, 99, 108, 116, 122, 133, 141, 152, 164, 174, 184, 195, 204, 214, 223, 240, 254, 260}
	_Type_index_3 = [...]uint16{0, 5, 10, 14, 19, 28, 32, 38, 44, 53, 60, 69, 83, 97, 116, 122, 132, 138, 143, 152, 160, 168, 176, 183, 190, 201, 207, 213, 220, 228, 236, 247, 251, 260, 267, 271, 279, 285, 291, 297, 303, 310, 317, 326, 330, 335, 344, 350, 357, 365, 369, 376, 384, 391, 397, 408, 419, 426, 437, 445, 452, 457, 463, 471, 477, 483, 489, 500, 507, 514, 522, 528, 535, 543, 549}
)

func (i Type) String() string {
	switch {
	case i == -1:
		return _Type_name_0
	case 1 <= i && i <= 6:
		i -= 1
		return _Type_name_1[_Type_index_1[i]:_Type_index_1[i+1]]
	case 10 <= i && i <= 38:
		i -= 10
		return _Type_name_2[_Type_index_2[i]:_Type_index_2[i+1]]
	case 100 <= i && i <= 173:
		i -= 100
		return _Type_name_3[_Type_index_3[i]:_Type_index_3[i+1]]
	default:
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
