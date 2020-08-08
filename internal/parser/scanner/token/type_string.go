// Code generated by "stringer -type=Type"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Error-1]
	_ = x[EOF-2]
	_ = x[StatementSeparator-3]
	_ = x[KeywordAbort-4]
	_ = x[KeywordAction-5]
	_ = x[KeywordAdd-6]
	_ = x[KeywordAfter-7]
	_ = x[KeywordAll-8]
	_ = x[KeywordAlter-9]
	_ = x[KeywordAlways-10]
	_ = x[KeywordAnalyze-11]
	_ = x[KeywordAnd-12]
	_ = x[KeywordAs-13]
	_ = x[KeywordAsc-14]
	_ = x[KeywordAttach-15]
	_ = x[KeywordAutoincrement-16]
	_ = x[KeywordBefore-17]
	_ = x[KeywordBegin-18]
	_ = x[KeywordBetween-19]
	_ = x[KeywordBy-20]
	_ = x[KeywordCascade-21]
	_ = x[KeywordCase-22]
	_ = x[KeywordCast-23]
	_ = x[KeywordCheck-24]
	_ = x[KeywordCollate-25]
	_ = x[KeywordColumn-26]
	_ = x[KeywordCommit-27]
	_ = x[KeywordConflict-28]
	_ = x[KeywordConstraint-29]
	_ = x[KeywordCreate-30]
	_ = x[KeywordCross-31]
	_ = x[KeywordCurrent-32]
	_ = x[KeywordCurrentDate-33]
	_ = x[KeywordCurrentTime-34]
	_ = x[KeywordCurrentTimestamp-35]
	_ = x[KeywordDatabase-36]
	_ = x[KeywordDefault-37]
	_ = x[KeywordDeferrable-38]
	_ = x[KeywordDeferred-39]
	_ = x[KeywordDelete-40]
	_ = x[KeywordDesc-41]
	_ = x[KeywordDetach-42]
	_ = x[KeywordDistinct-43]
	_ = x[KeywordDo-44]
	_ = x[KeywordDrop-45]
	_ = x[KeywordEach-46]
	_ = x[KeywordElse-47]
	_ = x[KeywordEnd-48]
	_ = x[KeywordEscape-49]
	_ = x[KeywordExcept-50]
	_ = x[KeywordExclude-51]
	_ = x[KeywordExclusive-52]
	_ = x[KeywordExists-53]
	_ = x[KeywordExplain-54]
	_ = x[KeywordFail-55]
	_ = x[KeywordFilter-56]
	_ = x[KeywordFirst-57]
	_ = x[KeywordFollowing-58]
	_ = x[KeywordFor-59]
	_ = x[KeywordForeign-60]
	_ = x[KeywordFrom-61]
	_ = x[KeywordFull-62]
	_ = x[KeywordGenerated-63]
	_ = x[KeywordGlob-64]
	_ = x[KeywordGroup-65]
	_ = x[KeywordGroups-66]
	_ = x[KeywordHaving-67]
	_ = x[KeywordIf-68]
	_ = x[KeywordIgnore-69]
	_ = x[KeywordImmediate-70]
	_ = x[KeywordIn-71]
	_ = x[KeywordIndex-72]
	_ = x[KeywordIndexed-73]
	_ = x[KeywordInitially-74]
	_ = x[KeywordInner-75]
	_ = x[KeywordInsert-76]
	_ = x[KeywordInstead-77]
	_ = x[KeywordIntersect-78]
	_ = x[KeywordInto-79]
	_ = x[KeywordIs-80]
	_ = x[KeywordIsnull-81]
	_ = x[KeywordJoin-82]
	_ = x[KeywordKey-83]
	_ = x[KeywordLast-84]
	_ = x[KeywordLeft-85]
	_ = x[KeywordLike-86]
	_ = x[KeywordLimit-87]
	_ = x[KeywordMatch-88]
	_ = x[KeywordNatural-89]
	_ = x[KeywordNo-90]
	_ = x[KeywordNot-91]
	_ = x[KeywordNothing-92]
	_ = x[KeywordNotnull-93]
	_ = x[KeywordNull-94]
	_ = x[KeywordNulls-95]
	_ = x[KeywordOf-96]
	_ = x[KeywordOffset-97]
	_ = x[KeywordOn-98]
	_ = x[KeywordOr-99]
	_ = x[KeywordOrder-100]
	_ = x[KeywordOthers-101]
	_ = x[KeywordOuter-102]
	_ = x[KeywordOver-103]
	_ = x[KeywordPartition-104]
	_ = x[KeywordPlan-105]
	_ = x[KeywordPragma-106]
	_ = x[KeywordPreceding-107]
	_ = x[KeywordPrimary-108]
	_ = x[KeywordQuery-109]
	_ = x[KeywordRaise-110]
	_ = x[KeywordRange-111]
	_ = x[KeywordRecursive-112]
	_ = x[KeywordReferences-113]
	_ = x[KeywordRegexp-114]
	_ = x[KeywordReIndex-115]
	_ = x[KeywordRelease-116]
	_ = x[KeywordRename-117]
	_ = x[KeywordReplace-118]
	_ = x[KeywordRestrict-119]
	_ = x[KeywordRight-120]
	_ = x[KeywordRollback-121]
	_ = x[KeywordRow-122]
	_ = x[KeywordRows-123]
	_ = x[KeywordSavepoint-124]
	_ = x[KeywordSelect-125]
	_ = x[KeywordSet-126]
	_ = x[KeywordStored-127]
	_ = x[KeywordTable-128]
	_ = x[KeywordTemp-129]
	_ = x[KeywordTemporary-130]
	_ = x[KeywordThen-131]
	_ = x[KeywordTies-132]
	_ = x[KeywordTo-133]
	_ = x[KeywordTransaction-134]
	_ = x[KeywordTrigger-135]
	_ = x[KeywordUnbounded-136]
	_ = x[KeywordUnion-137]
	_ = x[KeywordUnique-138]
	_ = x[KeywordUpdate-139]
	_ = x[KeywordUsing-140]
	_ = x[KeywordVacuum-141]
	_ = x[KeywordValues-142]
	_ = x[KeywordView-143]
	_ = x[KeywordVirtual-144]
	_ = x[KeywordWhen-145]
	_ = x[KeywordWhere-146]
	_ = x[KeywordWindow-147]
	_ = x[KeywordWith-148]
	_ = x[KeywordWithout-149]
	_ = x[Literal-150]
	_ = x[LiteralNumeric-151]
	_ = x[UnaryOperator-152]
	_ = x[BinaryOperator-153]
	_ = x[Delimiter-154]
}

const _Type_name = "UnknownErrorEOFStatementSeparatorKeywordAbortKeywordActionKeywordAddKeywordAfterKeywordAllKeywordAlterKeywordAlwaysKeywordAnalyzeKeywordAndKeywordAsKeywordAscKeywordAttachKeywordAutoincrementKeywordBeforeKeywordBeginKeywordBetweenKeywordByKeywordCascadeKeywordCaseKeywordCastKeywordCheckKeywordCollateKeywordColumnKeywordCommitKeywordConflictKeywordConstraintKeywordCreateKeywordCrossKeywordCurrentKeywordCurrentDateKeywordCurrentTimeKeywordCurrentTimestampKeywordDatabaseKeywordDefaultKeywordDeferrableKeywordDeferredKeywordDeleteKeywordDescKeywordDetachKeywordDistinctKeywordDoKeywordDropKeywordEachKeywordElseKeywordEndKeywordEscapeKeywordExceptKeywordExcludeKeywordExclusiveKeywordExistsKeywordExplainKeywordFailKeywordFilterKeywordFirstKeywordFollowingKeywordForKeywordForeignKeywordFromKeywordFullKeywordGeneratedKeywordGlobKeywordGroupKeywordGroupsKeywordHavingKeywordIfKeywordIgnoreKeywordImmediateKeywordInKeywordIndexKeywordIndexedKeywordInitiallyKeywordInnerKeywordInsertKeywordInsteadKeywordIntersectKeywordIntoKeywordIsKeywordIsnullKeywordJoinKeywordKeyKeywordLastKeywordLeftKeywordLikeKeywordLimitKeywordMatchKeywordNaturalKeywordNoKeywordNotKeywordNothingKeywordNotnullKeywordNullKeywordNullsKeywordOfKeywordOffsetKeywordOnKeywordOrKeywordOrderKeywordOthersKeywordOuterKeywordOverKeywordPartitionKeywordPlanKeywordPragmaKeywordPrecedingKeywordPrimaryKeywordQueryKeywordRaiseKeywordRangeKeywordRecursiveKeywordReferencesKeywordRegexpKeywordReindexKeywordReleaseKeywordRenameKeywordReplaceKeywordRestrictKeywordRightKeywordRollbackKeywordRowKeywordRowsKeywordSavepointKeywordSelectKeywordSetKeywordStoredKeywordTableKeywordTempKeywordTemporaryKeywordThenKeywordTiesKeywordToKeywordTransactionKeywordTriggerKeywordUnboundedKeywordUnionKeywordUniqueKeywordUpdateKeywordUsingKeywordVacuumKeywordValuesKeywordViewKeywordVirtualKeywordWhenKeywordWhereKeywordWindowKeywordWithKeywordWithoutLiteralLiteralNumericUnaryOperatorBinaryOperatorDelimiter"

var _Type_index = [...]uint16{0, 7, 12, 15, 33, 45, 58, 68, 80, 90, 102, 115, 129, 139, 148, 158, 171, 191, 204, 216, 230, 239, 253, 264, 275, 287, 301, 314, 327, 342, 359, 372, 384, 398, 416, 434, 457, 472, 486, 503, 518, 531, 542, 555, 570, 579, 590, 601, 612, 622, 635, 648, 662, 678, 691, 705, 716, 729, 741, 757, 767, 781, 792, 803, 819, 830, 842, 855, 868, 877, 890, 906, 915, 927, 941, 957, 969, 982, 996, 1012, 1023, 1032, 1045, 1056, 1066, 1077, 1088, 1099, 1111, 1123, 1137, 1146, 1156, 1170, 1184, 1195, 1207, 1216, 1229, 1238, 1247, 1259, 1272, 1284, 1295, 1311, 1322, 1335, 1351, 1365, 1377, 1389, 1401, 1417, 1434, 1447, 1461, 1475, 1488, 1502, 1517, 1529, 1544, 1554, 1565, 1581, 1594, 1604, 1617, 1629, 1640, 1656, 1667, 1678, 1687, 1705, 1719, 1735, 1747, 1760, 1773, 1785, 1798, 1811, 1822, 1836, 1847, 1859, 1872, 1883, 1897, 1904, 1918, 1931, 1945, 1954}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}