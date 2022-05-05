// Code generated from asn1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package asn1_parser // asn1
import "github.com/antlr/antlr4/runtime/Go/antlr"

// asn1Listener is a complete listener for a parse tree produced by asn1Parser.
type asn1Listener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterRuleModules is called when entering the ruleModules production.
	EnterRuleModules(c *RuleModulesContext)

	// EnterRuleModuleDefinition is called when entering the ruleModuleDefinition production.
	EnterRuleModuleDefinition(c *RuleModuleDefinitionContext)

	// EnterRuleModuleReference is called when entering the ruleModuleReference production.
	EnterRuleModuleReference(c *RuleModuleReferenceContext)

	// EnterRuleModuleIdentifier is called when entering the ruleModuleIdentifier production.
	EnterRuleModuleIdentifier(c *RuleModuleIdentifierContext)

	// EnterRuleDefinitiveIdentification is called when entering the ruleDefinitiveIdentification production.
	EnterRuleDefinitiveIdentification(c *RuleDefinitiveIdentificationContext)

	// EnterRuleDefinitiveOID is called when entering the ruleDefinitiveOID production.
	EnterRuleDefinitiveOID(c *RuleDefinitiveOIDContext)

	// EnterRuleDefinitiveObjIdComponentList is called when entering the ruleDefinitiveObjIdComponentList production.
	EnterRuleDefinitiveObjIdComponentList(c *RuleDefinitiveObjIdComponentListContext)

	// EnterRuleDefinitiveObjIdComponent is called when entering the ruleDefinitiveObjIdComponent production.
	EnterRuleDefinitiveObjIdComponent(c *RuleDefinitiveObjIdComponentContext)

	// EnterRuleNameForm is called when entering the ruleNameForm production.
	EnterRuleNameForm(c *RuleNameFormContext)

	// EnterRuleDefinitiveNumberForm is called when entering the ruleDefinitiveNumberForm production.
	EnterRuleDefinitiveNumberForm(c *RuleDefinitiveNumberFormContext)

	// EnterRuleDefinitiveNameAndNumberForm is called when entering the ruleDefinitiveNameAndNumberForm production.
	EnterRuleDefinitiveNameAndNumberForm(c *RuleDefinitiveNameAndNumberFormContext)

	// EnterRuleDefinitiveOIDAndIRI is called when entering the ruleDefinitiveOIDAndIRI production.
	EnterRuleDefinitiveOIDAndIRI(c *RuleDefinitiveOIDAndIRIContext)

	// EnterRuleFirstArcIdentifier is called when entering the ruleFirstArcIdentifier production.
	EnterRuleFirstArcIdentifier(c *RuleFirstArcIdentifierContext)

	// EnterRuleSubsequentArcIdentifier is called when entering the ruleSubsequentArcIdentifier production.
	EnterRuleSubsequentArcIdentifier(c *RuleSubsequentArcIdentifierContext)

	// EnterRuleIRIValue is called when entering the ruleIRIValue production.
	EnterRuleIRIValue(c *RuleIRIValueContext)

	// EnterRuleEncodingReferenceDefault is called when entering the ruleEncodingReferenceDefault production.
	EnterRuleEncodingReferenceDefault(c *RuleEncodingReferenceDefaultContext)

	// EnterRuleTagDefault is called when entering the ruleTagDefault production.
	EnterRuleTagDefault(c *RuleTagDefaultContext)

	// EnterRuleExtensionDefault is called when entering the ruleExtensionDefault production.
	EnterRuleExtensionDefault(c *RuleExtensionDefaultContext)

	// EnterRuleModuleBody is called when entering the ruleModuleBody production.
	EnterRuleModuleBody(c *RuleModuleBodyContext)

	// EnterRuleExports is called when entering the ruleExports production.
	EnterRuleExports(c *RuleExportsContext)

	// EnterRuleSymbolsExported is called when entering the ruleSymbolsExported production.
	EnterRuleSymbolsExported(c *RuleSymbolsExportedContext)

	// EnterRuleSymbolList is called when entering the ruleSymbolList production.
	EnterRuleSymbolList(c *RuleSymbolListContext)

	// EnterRuleSymbol is called when entering the ruleSymbol production.
	EnterRuleSymbol(c *RuleSymbolContext)

	// EnterRuleReference is called when entering the ruleReference production.
	EnterRuleReference(c *RuleReferenceContext)

	// EnterRuleIdentifier is called when entering the ruleIdentifier production.
	EnterRuleIdentifier(c *RuleIdentifierContext)

	// EnterRuleTypeReference is called when entering the ruleTypeReference production.
	EnterRuleTypeReference(c *RuleTypeReferenceContext)

	// EnterRuleValueReference is called when entering the ruleValueReference production.
	EnterRuleValueReference(c *RuleValueReferenceContext)

	// EnterRuleObjectReference is called when entering the ruleObjectReference production.
	EnterRuleObjectReference(c *RuleObjectReferenceContext)

	// EnterRuleObjectClassReference is called when entering the ruleObjectClassReference production.
	EnterRuleObjectClassReference(c *RuleObjectClassReferenceContext)

	// EnterRuleObjectSetReference is called when entering the ruleObjectSetReference production.
	EnterRuleObjectSetReference(c *RuleObjectSetReferenceContext)

	// EnterRuleParameterizedReference is called when entering the ruleParameterizedReference production.
	EnterRuleParameterizedReference(c *RuleParameterizedReferenceContext)

	// EnterRuleExternalTypeReference is called when entering the ruleExternalTypeReference production.
	EnterRuleExternalTypeReference(c *RuleExternalTypeReferenceContext)

	// EnterRuleExternalValueReference is called when entering the ruleExternalValueReference production.
	EnterRuleExternalValueReference(c *RuleExternalValueReferenceContext)

	// EnterRuleExternalObjectClassReference is called when entering the ruleExternalObjectClassReference production.
	EnterRuleExternalObjectClassReference(c *RuleExternalObjectClassReferenceContext)

	// EnterRuleExternalObjectReference is called when entering the ruleExternalObjectReference production.
	EnterRuleExternalObjectReference(c *RuleExternalObjectReferenceContext)

	// EnterRuleExternalObjectSetReference is called when entering the ruleExternalObjectSetReference production.
	EnterRuleExternalObjectSetReference(c *RuleExternalObjectSetReferenceContext)

	// EnterRuleTypeFieldReference is called when entering the ruleTypeFieldReference production.
	EnterRuleTypeFieldReference(c *RuleTypeFieldReferenceContext)

	// EnterRuleValueFieldReference is called when entering the ruleValueFieldReference production.
	EnterRuleValueFieldReference(c *RuleValueFieldReferenceContext)

	// EnterRuleValueSetFieldReference is called when entering the ruleValueSetFieldReference production.
	EnterRuleValueSetFieldReference(c *RuleValueSetFieldReferenceContext)

	// EnterRuleObjectFieldReference is called when entering the ruleObjectFieldReference production.
	EnterRuleObjectFieldReference(c *RuleObjectFieldReferenceContext)

	// EnterRuleObjectSetFieldReference is called when entering the ruleObjectSetFieldReference production.
	EnterRuleObjectSetFieldReference(c *RuleObjectSetFieldReferenceContext)

	// EnterRuleUsefulObjectClassReference is called when entering the ruleUsefulObjectClassReference production.
	EnterRuleUsefulObjectClassReference(c *RuleUsefulObjectClassReferenceContext)

	// EnterRuleImports is called when entering the ruleImports production.
	EnterRuleImports(c *RuleImportsContext)

	// EnterRuleSymbolsImported is called when entering the ruleSymbolsImported production.
	EnterRuleSymbolsImported(c *RuleSymbolsImportedContext)

	// EnterRuleSymbolsFromModuleList is called when entering the ruleSymbolsFromModuleList production.
	EnterRuleSymbolsFromModuleList(c *RuleSymbolsFromModuleListContext)

	// EnterRuleSymbolsFromModule is called when entering the ruleSymbolsFromModule production.
	EnterRuleSymbolsFromModule(c *RuleSymbolsFromModuleContext)

	// EnterRuleGlobalModuleReference is called when entering the ruleGlobalModuleReference production.
	EnterRuleGlobalModuleReference(c *RuleGlobalModuleReferenceContext)

	// EnterRuleAssignedIdentifier is called when entering the ruleAssignedIdentifier production.
	EnterRuleAssignedIdentifier(c *RuleAssignedIdentifierContext)

	// EnterRuleAssignmentList is called when entering the ruleAssignmentList production.
	EnterRuleAssignmentList(c *RuleAssignmentListContext)

	// EnterRuleAssignment is called when entering the ruleAssignment production.
	EnterRuleAssignment(c *RuleAssignmentContext)

	// EnterRuleTypeAssignment is called when entering the ruleTypeAssignment production.
	EnterRuleTypeAssignment(c *RuleTypeAssignmentContext)

	// EnterRuleValueAssignment is called when entering the ruleValueAssignment production.
	EnterRuleValueAssignment(c *RuleValueAssignmentContext)

	// EnterRuleXMLValueAssignment is called when entering the ruleXMLValueAssignment production.
	EnterRuleXMLValueAssignment(c *RuleXMLValueAssignmentContext)

	// EnterRuleValueSetTypeAssignment is called when entering the ruleValueSetTypeAssignment production.
	EnterRuleValueSetTypeAssignment(c *RuleValueSetTypeAssignmentContext)

	// EnterRuleObjectAssignment is called when entering the ruleObjectAssignment production.
	EnterRuleObjectAssignment(c *RuleObjectAssignmentContext)

	// EnterRuleObjectClassAssignment is called when entering the ruleObjectClassAssignment production.
	EnterRuleObjectClassAssignment(c *RuleObjectClassAssignmentContext)

	// EnterRuleObjectSetAssignment is called when entering the ruleObjectSetAssignment production.
	EnterRuleObjectSetAssignment(c *RuleObjectSetAssignmentContext)

	// EnterRuleParameterizedAssignment is called when entering the ruleParameterizedAssignment production.
	EnterRuleParameterizedAssignment(c *RuleParameterizedAssignmentContext)

	// EnterRuleObjectIdentifierValue is called when entering the ruleObjectIdentifierValue production.
	EnterRuleObjectIdentifierValue(c *RuleObjectIdentifierValueContext)

	// EnterRuleObjIdComponentsList is called when entering the ruleObjIdComponentsList production.
	EnterRuleObjIdComponentsList(c *RuleObjIdComponentsListContext)

	// EnterRuleObjIdComponents is called when entering the ruleObjIdComponents production.
	EnterRuleObjIdComponents(c *RuleObjIdComponentsContext)

	// EnterRuleNumberForm is called when entering the ruleNumberForm production.
	EnterRuleNumberForm(c *RuleNumberFormContext)

	// EnterRuleNameAndNumberForm is called when entering the ruleNameAndNumberForm production.
	EnterRuleNameAndNumberForm(c *RuleNameAndNumberFormContext)

	// EnterRuleDefinedValue is called when entering the ruleDefinedValue production.
	EnterRuleDefinedValue(c *RuleDefinedValueContext)

	// EnterRuleNamedBit is called when entering the ruleNamedBit production.
	EnterRuleNamedBit(c *RuleNamedBitContext)

	// EnterRuleNamedBitList is called when entering the ruleNamedBitList production.
	EnterRuleNamedBitList(c *RuleNamedBitListContext)

	// EnterRuleRestrictedCharacterStringType is called when entering the ruleRestrictedCharacterStringType production.
	EnterRuleRestrictedCharacterStringType(c *RuleRestrictedCharacterStringTypeContext)

	// EnterRuleUnrestrictedCharacterStringType is called when entering the ruleUnrestrictedCharacterStringType production.
	EnterRuleUnrestrictedCharacterStringType(c *RuleUnrestrictedCharacterStringTypeContext)

	// EnterRuleAlternativeTypeList is called when entering the ruleAlternativeTypeList production.
	EnterRuleAlternativeTypeList(c *RuleAlternativeTypeListContext)

	// EnterRuleRootAlternativeTypeList is called when entering the ruleRootAlternativeTypeList production.
	EnterRuleRootAlternativeTypeList(c *RuleRootAlternativeTypeListContext)

	// EnterRuleVersionNumber is called when entering the ruleVersionNumber production.
	EnterRuleVersionNumber(c *RuleVersionNumberContext)

	// EnterRuleExtensionAdditionAlternativesGroup is called when entering the ruleExtensionAdditionAlternativesGroup production.
	EnterRuleExtensionAdditionAlternativesGroup(c *RuleExtensionAdditionAlternativesGroupContext)

	// EnterRuleExtensionAdditionAlternative is called when entering the ruleExtensionAdditionAlternative production.
	EnterRuleExtensionAdditionAlternative(c *RuleExtensionAdditionAlternativeContext)

	// EnterRuleExtensionAdditionAlternativesList is called when entering the ruleExtensionAdditionAlternativesList production.
	EnterRuleExtensionAdditionAlternativesList(c *RuleExtensionAdditionAlternativesListContext)

	// EnterRuleExtensionAdditionAlternatives is called when entering the ruleExtensionAdditionAlternatives production.
	EnterRuleExtensionAdditionAlternatives(c *RuleExtensionAdditionAlternativesContext)

	// EnterRuleEnumerationItem is called when entering the ruleEnumerationItem production.
	EnterRuleEnumerationItem(c *RuleEnumerationItemContext)

	// EnterRuleEnumeration is called when entering the ruleEnumeration production.
	EnterRuleEnumeration(c *RuleEnumerationContext)

	// EnterRuleRootEnumeration is called when entering the ruleRootEnumeration production.
	EnterRuleRootEnumeration(c *RuleRootEnumerationContext)

	// EnterRuleAdditionalEnumeration is called when entering the ruleAdditionalEnumeration production.
	EnterRuleAdditionalEnumeration(c *RuleAdditionalEnumerationContext)

	// EnterRuleAlternativeTypeLists is called when entering the ruleAlternativeTypeLists production.
	EnterRuleAlternativeTypeLists(c *RuleAlternativeTypeListsContext)

	// EnterRuleEnumerations is called when entering the ruleEnumerations production.
	EnterRuleEnumerations(c *RuleEnumerationsContext)

	// EnterRuleNamedNumber is called when entering the ruleNamedNumber production.
	EnterRuleNamedNumber(c *RuleNamedNumberContext)

	// EnterRuleNamedNumberList is called when entering the ruleNamedNumberList production.
	EnterRuleNamedNumberList(c *RuleNamedNumberListContext)

	// EnterRuleComponentType is called when entering the ruleComponentType production.
	EnterRuleComponentType(c *RuleComponentTypeContext)

	// EnterRuleExtensionAdditionGroup is called when entering the ruleExtensionAdditionGroup production.
	EnterRuleExtensionAdditionGroup(c *RuleExtensionAdditionGroupContext)

	// EnterRuleExtensionAddition is called when entering the ruleExtensionAddition production.
	EnterRuleExtensionAddition(c *RuleExtensionAdditionContext)

	// EnterRuleComponentTypeList is called when entering the ruleComponentTypeList production.
	EnterRuleComponentTypeList(c *RuleComponentTypeListContext)

	// EnterRuleExtensionAdditionList is called when entering the ruleExtensionAdditionList production.
	EnterRuleExtensionAdditionList(c *RuleExtensionAdditionListContext)

	// EnterRuleRootComponentTypeList is called when entering the ruleRootComponentTypeList production.
	EnterRuleRootComponentTypeList(c *RuleRootComponentTypeListContext)

	// EnterRuleExtensionAdditions is called when entering the ruleExtensionAdditions production.
	EnterRuleExtensionAdditions(c *RuleExtensionAdditionsContext)

	// EnterRuleExtensionEndMarker is called when entering the ruleExtensionEndMarker production.
	EnterRuleExtensionEndMarker(c *RuleExtensionEndMarkerContext)

	// EnterRuleComponentTypeLists is called when entering the ruleComponentTypeLists production.
	EnterRuleComponentTypeLists(c *RuleComponentTypeListsContext)

	// EnterRuleExtensionAndException is called when entering the ruleExtensionAndException production.
	EnterRuleExtensionAndException(c *RuleExtensionAndExceptionContext)

	// EnterRuleOptionalExtensionMarker is called when entering the ruleOptionalExtensionMarker production.
	EnterRuleOptionalExtensionMarker(c *RuleOptionalExtensionMarkerContext)

	// EnterRuleEncodingReference is called when entering the ruleEncodingReference production.
	EnterRuleEncodingReference(c *RuleEncodingReferenceContext)

	// EnterRuleClass is called when entering the ruleClass production.
	EnterRuleClass(c *RuleClassContext)

	// EnterRuleClassNumber is called when entering the ruleClassNumber production.
	EnterRuleClassNumber(c *RuleClassNumberContext)

	// EnterRuleEncodingInstruction is called when entering the ruleEncodingInstruction production.
	EnterRuleEncodingInstruction(c *RuleEncodingInstructionContext)

	// EnterRuleTag is called when entering the ruleTag production.
	EnterRuleTag(c *RuleTagContext)

	// EnterRuleEncodingPrefix is called when entering the ruleEncodingPrefix production.
	EnterRuleEncodingPrefix(c *RuleEncodingPrefixContext)

	// EnterRuleTaggedType is called when entering the ruleTaggedType production.
	EnterRuleTaggedType(c *RuleTaggedTypeContext)

	// EnterRuleEncodingPrefixedType is called when entering the ruleEncodingPrefixedType production.
	EnterRuleEncodingPrefixedType(c *RuleEncodingPrefixedTypeContext)

	// EnterRuleBitStringType is called when entering the ruleBitStringType production.
	EnterRuleBitStringType(c *RuleBitStringTypeContext)

	// EnterRuleBooleanType is called when entering the ruleBooleanType production.
	EnterRuleBooleanType(c *RuleBooleanTypeContext)

	// EnterRuleCharacterStringType is called when entering the ruleCharacterStringType production.
	EnterRuleCharacterStringType(c *RuleCharacterStringTypeContext)

	// EnterRuleChoiceType is called when entering the ruleChoiceType production.
	EnterRuleChoiceType(c *RuleChoiceTypeContext)

	// EnterRuleDateType is called when entering the ruleDateType production.
	EnterRuleDateType(c *RuleDateTypeContext)

	// EnterRuleDateTimeType is called when entering the ruleDateTimeType production.
	EnterRuleDateTimeType(c *RuleDateTimeTypeContext)

	// EnterRuleDurationType is called when entering the ruleDurationType production.
	EnterRuleDurationType(c *RuleDurationTypeContext)

	// EnterRuleEmbeddedPDVType is called when entering the ruleEmbeddedPDVType production.
	EnterRuleEmbeddedPDVType(c *RuleEmbeddedPDVTypeContext)

	// EnterRuleEnumeratedType is called when entering the ruleEnumeratedType production.
	EnterRuleEnumeratedType(c *RuleEnumeratedTypeContext)

	// EnterRuleExternalType is called when entering the ruleExternalType production.
	EnterRuleExternalType(c *RuleExternalTypeContext)

	// EnterRuleInstanceOfType is called when entering the ruleInstanceOfType production.
	EnterRuleInstanceOfType(c *RuleInstanceOfTypeContext)

	// EnterRuleIntegerType is called when entering the ruleIntegerType production.
	EnterRuleIntegerType(c *RuleIntegerTypeContext)

	// EnterRuleIRIType is called when entering the ruleIRIType production.
	EnterRuleIRIType(c *RuleIRITypeContext)

	// EnterRuleNullType is called when entering the ruleNullType production.
	EnterRuleNullType(c *RuleNullTypeContext)

	// EnterRuleObjectClassFieldType is called when entering the ruleObjectClassFieldType production.
	EnterRuleObjectClassFieldType(c *RuleObjectClassFieldTypeContext)

	// EnterRuleObjectIdentifierType is called when entering the ruleObjectIdentifierType production.
	EnterRuleObjectIdentifierType(c *RuleObjectIdentifierTypeContext)

	// EnterRuleOctetStringType is called when entering the ruleOctetStringType production.
	EnterRuleOctetStringType(c *RuleOctetStringTypeContext)

	// EnterRuleRealType is called when entering the ruleRealType production.
	EnterRuleRealType(c *RuleRealTypeContext)

	// EnterRuleRelativeIRIType is called when entering the ruleRelativeIRIType production.
	EnterRuleRelativeIRIType(c *RuleRelativeIRITypeContext)

	// EnterRuleRelativeOIDType is called when entering the ruleRelativeOIDType production.
	EnterRuleRelativeOIDType(c *RuleRelativeOIDTypeContext)

	// EnterRuleSequenceType is called when entering the ruleSequenceType production.
	EnterRuleSequenceType(c *RuleSequenceTypeContext)

	// EnterRuleSequenceOfType is called when entering the ruleSequenceOfType production.
	EnterRuleSequenceOfType(c *RuleSequenceOfTypeContext)

	// EnterRuleSetType is called when entering the ruleSetType production.
	EnterRuleSetType(c *RuleSetTypeContext)

	// EnterRuleSetOfType is called when entering the ruleSetOfType production.
	EnterRuleSetOfType(c *RuleSetOfTypeContext)

	// EnterRulePrefixedType is called when entering the rulePrefixedType production.
	EnterRulePrefixedType(c *RulePrefixedTypeContext)

	// EnterRuleTimeType is called when entering the ruleTimeType production.
	EnterRuleTimeType(c *RuleTimeTypeContext)

	// EnterRuleTimeOfDayType is called when entering the ruleTimeOfDayType production.
	EnterRuleTimeOfDayType(c *RuleTimeOfDayTypeContext)

	// EnterRuleBuiltinType is called when entering the ruleBuiltinType production.
	EnterRuleBuiltinType(c *RuleBuiltinTypeContext)

	// EnterRuleSimpleDefinedType is called when entering the ruleSimpleDefinedType production.
	EnterRuleSimpleDefinedType(c *RuleSimpleDefinedTypeContext)

	// EnterRuleParameterizedType is called when entering the ruleParameterizedType production.
	EnterRuleParameterizedType(c *RuleParameterizedTypeContext)

	// EnterRuleParameterizedValueSetType is called when entering the ruleParameterizedValueSetType production.
	EnterRuleParameterizedValueSetType(c *RuleParameterizedValueSetTypeContext)

	// EnterRuleDefinedType is called when entering the ruleDefinedType production.
	EnterRuleDefinedType(c *RuleDefinedTypeContext)

	// EnterRuleUsefulType is called when entering the ruleUsefulType production.
	EnterRuleUsefulType(c *RuleUsefulTypeContext)

	// EnterRuleSelectionType is called when entering the ruleSelectionType production.
	EnterRuleSelectionType(c *RuleSelectionTypeContext)

	// EnterRuleTypeFromObject is called when entering the ruleTypeFromObject production.
	EnterRuleTypeFromObject(c *RuleTypeFromObjectContext)

	// EnterRuleValueSetFromObjects is called when entering the ruleValueSetFromObjects production.
	EnterRuleValueSetFromObjects(c *RuleValueSetFromObjectsContext)

	// EnterRuleReferencedType is called when entering the ruleReferencedType production.
	EnterRuleReferencedType(c *RuleReferencedTypeContext)

	// EnterRuleTypeForConstraint is called when entering the ruleTypeForConstraint production.
	EnterRuleTypeForConstraint(c *RuleTypeForConstraintContext)

	// EnterRuleNamedType is called when entering the ruleNamedType production.
	EnterRuleNamedType(c *RuleNamedTypeContext)

	// EnterRuleTypeWithConstraint is called when entering the ruleTypeWithConstraint production.
	EnterRuleTypeWithConstraint(c *RuleTypeWithConstraintContext)

	// EnterRuleConstrainedType is called when entering the ruleConstrainedType production.
	EnterRuleConstrainedType(c *RuleConstrainedTypeContext)

	// EnterRuleType is called when entering the ruleType production.
	EnterRuleType(c *RuleTypeContext)

	// EnterRuleIdentifierList is called when entering the ruleIdentifierList production.
	EnterRuleIdentifierList(c *RuleIdentifierListContext)

	// EnterRuleCharsDefn is called when entering the ruleCharsDefn production.
	EnterRuleCharsDefn(c *RuleCharsDefnContext)

	// EnterRuleCharSyms is called when entering the ruleCharSyms production.
	EnterRuleCharSyms(c *RuleCharSymsContext)

	// EnterRuleGroup is called when entering the ruleGroup production.
	EnterRuleGroup(c *RuleGroupContext)

	// EnterRulePlane is called when entering the rulePlane production.
	EnterRulePlane(c *RulePlaneContext)

	// EnterRuleRow is called when entering the ruleRow production.
	EnterRuleRow(c *RuleRowContext)

	// EnterRuleCell is called when entering the ruleCell production.
	EnterRuleCell(c *RuleCellContext)

	// EnterRuleTableColumn is called when entering the ruleTableColumn production.
	EnterRuleTableColumn(c *RuleTableColumnContext)

	// EnterRuleTableRow is called when entering the ruleTableRow production.
	EnterRuleTableRow(c *RuleTableRowContext)

	// EnterRuleCharacterStringList is called when entering the ruleCharacterStringList production.
	EnterRuleCharacterStringList(c *RuleCharacterStringListContext)

	// EnterRuleQuadruple is called when entering the ruleQuadruple production.
	EnterRuleQuadruple(c *RuleQuadrupleContext)

	// EnterRuleTuple is called when entering the ruleTuple production.
	EnterRuleTuple(c *RuleTupleContext)

	// EnterRuleRestrictedCharacterStringValue is called when entering the ruleRestrictedCharacterStringValue production.
	EnterRuleRestrictedCharacterStringValue(c *RuleRestrictedCharacterStringValueContext)

	// EnterRuleUnrestrictedCharacterStringValue is called when entering the ruleUnrestrictedCharacterStringValue production.
	EnterRuleUnrestrictedCharacterStringValue(c *RuleUnrestrictedCharacterStringValueContext)

	// EnterRuleNumericRealValue is called when entering the ruleNumericRealValue production.
	EnterRuleNumericRealValue(c *RuleNumericRealValueContext)

	// EnterRuleSpecialRealValue is called when entering the ruleSpecialRealValue production.
	EnterRuleSpecialRealValue(c *RuleSpecialRealValueContext)

	// EnterRuleRelativeOIDComponents is called when entering the ruleRelativeOIDComponents production.
	EnterRuleRelativeOIDComponents(c *RuleRelativeOIDComponentsContext)

	// EnterRuleRelativeOIDComponentsList is called when entering the ruleRelativeOIDComponentsList production.
	EnterRuleRelativeOIDComponentsList(c *RuleRelativeOIDComponentsListContext)

	// EnterRuleComponentValueList is called when entering the ruleComponentValueList production.
	EnterRuleComponentValueList(c *RuleComponentValueListContext)

	// EnterRuleValueList is called when entering the ruleValueList production.
	EnterRuleValueList(c *RuleValueListContext)

	// EnterRuleNamedValue is called when entering the ruleNamedValue production.
	EnterRuleNamedValue(c *RuleNamedValueContext)

	// EnterRuleNamedValueList is called when entering the ruleNamedValueList production.
	EnterRuleNamedValueList(c *RuleNamedValueListContext)

	// EnterRuleBitStringValue is called when entering the ruleBitStringValue production.
	EnterRuleBitStringValue(c *RuleBitStringValueContext)

	// EnterRuleBooleanValue is called when entering the ruleBooleanValue production.
	EnterRuleBooleanValue(c *RuleBooleanValueContext)

	// EnterRuleCharacterStringValue is called when entering the ruleCharacterStringValue production.
	EnterRuleCharacterStringValue(c *RuleCharacterStringValueContext)

	// EnterRuleChoiceValue is called when entering the ruleChoiceValue production.
	EnterRuleChoiceValue(c *RuleChoiceValueContext)

	// EnterRuleEmbeddedPDVValue is called when entering the ruleEmbeddedPDVValue production.
	EnterRuleEmbeddedPDVValue(c *RuleEmbeddedPDVValueContext)

	// EnterRuleEnumeratedValue is called when entering the ruleEnumeratedValue production.
	EnterRuleEnumeratedValue(c *RuleEnumeratedValueContext)

	// EnterRuleExternalValue is called when entering the ruleExternalValue production.
	EnterRuleExternalValue(c *RuleExternalValueContext)

	// EnterRuleInstanceOfValue is called when entering the ruleInstanceOfValue production.
	EnterRuleInstanceOfValue(c *RuleInstanceOfValueContext)

	// EnterRuleIntegerValue is called when entering the ruleIntegerValue production.
	EnterRuleIntegerValue(c *RuleIntegerValueContext)

	// EnterRuleNullValue is called when entering the ruleNullValue production.
	EnterRuleNullValue(c *RuleNullValueContext)

	// EnterRuleOctetStringValue is called when entering the ruleOctetStringValue production.
	EnterRuleOctetStringValue(c *RuleOctetStringValueContext)

	// EnterRuleRealValue is called when entering the ruleRealValue production.
	EnterRuleRealValue(c *RuleRealValueContext)

	// EnterRuleRelativeIRIValue is called when entering the ruleRelativeIRIValue production.
	EnterRuleRelativeIRIValue(c *RuleRelativeIRIValueContext)

	// EnterRuleRelativeOIDValue is called when entering the ruleRelativeOIDValue production.
	EnterRuleRelativeOIDValue(c *RuleRelativeOIDValueContext)

	// EnterRuleSequenceValue is called when entering the ruleSequenceValue production.
	EnterRuleSequenceValue(c *RuleSequenceValueContext)

	// EnterRuleSequenceOfValue is called when entering the ruleSequenceOfValue production.
	EnterRuleSequenceOfValue(c *RuleSequenceOfValueContext)

	// EnterRuleSetValue is called when entering the ruleSetValue production.
	EnterRuleSetValue(c *RuleSetValueContext)

	// EnterRuleSetOfValue is called when entering the ruleSetOfValue production.
	EnterRuleSetOfValue(c *RuleSetOfValueContext)

	// EnterRulePrefixedValue is called when entering the rulePrefixedValue production.
	EnterRulePrefixedValue(c *RulePrefixedValueContext)

	// EnterRuleTimeValue is called when entering the ruleTimeValue production.
	EnterRuleTimeValue(c *RuleTimeValueContext)

	// EnterRuleBuiltinValue is called when entering the ruleBuiltinValue production.
	EnterRuleBuiltinValue(c *RuleBuiltinValueContext)

	// EnterRuleValueFromObject is called when entering the ruleValueFromObject production.
	EnterRuleValueFromObject(c *RuleValueFromObjectContext)

	// EnterRuleReferencedValue is called when entering the ruleReferencedValue production.
	EnterRuleReferencedValue(c *RuleReferencedValueContext)

	// EnterRuleOpenTypeFieldVal is called when entering the ruleOpenTypeFieldVal production.
	EnterRuleOpenTypeFieldVal(c *RuleOpenTypeFieldValContext)

	// EnterRuleFixedTypeFieldVal is called when entering the ruleFixedTypeFieldVal production.
	EnterRuleFixedTypeFieldVal(c *RuleFixedTypeFieldValContext)

	// EnterRuleObjectClassFieldValue is called when entering the ruleObjectClassFieldValue production.
	EnterRuleObjectClassFieldValue(c *RuleObjectClassFieldValueContext)

	// EnterRuleValue is called when entering the ruleValue production.
	EnterRuleValue(c *RuleValueContext)

	// EnterRuleValueSet is called when entering the ruleValueSet production.
	EnterRuleValueSet(c *RuleValueSetContext)

	// EnterRuleDefinedObjectClass is called when entering the ruleDefinedObjectClass production.
	EnterRuleDefinedObjectClass(c *RuleDefinedObjectClassContext)

	// EnterRuleObject is called when entering the ruleObject production.
	EnterRuleObject(c *RuleObjectContext)

	// EnterRuleDefinedObject is called when entering the ruleDefinedObject production.
	EnterRuleDefinedObject(c *RuleDefinedObjectContext)

	// EnterRuleDefinedObjectSet is called when entering the ruleDefinedObjectSet production.
	EnterRuleDefinedObjectSet(c *RuleDefinedObjectSetContext)

	// EnterRuleObjectDefn is called when entering the ruleObjectDefn production.
	EnterRuleObjectDefn(c *RuleObjectDefnContext)

	// EnterRuleDefaultSyntax is called when entering the ruleDefaultSyntax production.
	EnterRuleDefaultSyntax(c *RuleDefaultSyntaxContext)

	// EnterRuleDefinedSyntax is called when entering the ruleDefinedSyntax production.
	EnterRuleDefinedSyntax(c *RuleDefinedSyntaxContext)

	// EnterRuleDefinedSyntaxTokenList is called when entering the ruleDefinedSyntaxTokenList production.
	EnterRuleDefinedSyntaxTokenList(c *RuleDefinedSyntaxTokenListContext)

	// EnterRuleDefinedSyntaxToken is called when entering the ruleDefinedSyntaxToken production.
	EnterRuleDefinedSyntaxToken(c *RuleDefinedSyntaxTokenContext)

	// EnterRuleLiteral is called when entering the ruleLiteral production.
	EnterRuleLiteral(c *RuleLiteralContext)

	// EnterRuleWord is called when entering the ruleWord production.
	EnterRuleWord(c *RuleWordContext)

	// EnterRuleFieldSettingList is called when entering the ruleFieldSettingList production.
	EnterRuleFieldSettingList(c *RuleFieldSettingListContext)

	// EnterRuleFieldSetting is called when entering the ruleFieldSetting production.
	EnterRuleFieldSetting(c *RuleFieldSettingContext)

	// EnterRulePrimitiveFieldName is called when entering the rulePrimitiveFieldName production.
	EnterRulePrimitiveFieldName(c *RulePrimitiveFieldNameContext)

	// EnterRuleSetting is called when entering the ruleSetting production.
	EnterRuleSetting(c *RuleSettingContext)

	// EnterRuleObjectFromObject is called when entering the ruleObjectFromObject production.
	EnterRuleObjectFromObject(c *RuleObjectFromObjectContext)

	// EnterRuleObjectSetFromObjects is called when entering the ruleObjectSetFromObjects production.
	EnterRuleObjectSetFromObjects(c *RuleObjectSetFromObjectsContext)

	// EnterRuleReferencedObjects is called when entering the ruleReferencedObjects production.
	EnterRuleReferencedObjects(c *RuleReferencedObjectsContext)

	// EnterRuleFieldName is called when entering the ruleFieldName production.
	EnterRuleFieldName(c *RuleFieldNameContext)

	// EnterRuleParameterizedObject is called when entering the ruleParameterizedObject production.
	EnterRuleParameterizedObject(c *RuleParameterizedObjectContext)

	// EnterRuleObjectClass is called when entering the ruleObjectClass production.
	EnterRuleObjectClass(c *RuleObjectClassContext)

	// EnterRuleObjectClassDefn is called when entering the ruleObjectClassDefn production.
	EnterRuleObjectClassDefn(c *RuleObjectClassDefnContext)

	// EnterRuleFieldSpecList is called when entering the ruleFieldSpecList production.
	EnterRuleFieldSpecList(c *RuleFieldSpecListContext)

	// EnterRuleFieldSpec is called when entering the ruleFieldSpec production.
	EnterRuleFieldSpec(c *RuleFieldSpecContext)

	// EnterRuleTypeFieldSpec is called when entering the ruleTypeFieldSpec production.
	EnterRuleTypeFieldSpec(c *RuleTypeFieldSpecContext)

	// EnterRuleFixedTypeValueFieldSpec is called when entering the ruleFixedTypeValueFieldSpec production.
	EnterRuleFixedTypeValueFieldSpec(c *RuleFixedTypeValueFieldSpecContext)

	// EnterRuleUnique is called when entering the ruleUnique production.
	EnterRuleUnique(c *RuleUniqueContext)

	// EnterRuleVariableTypeValueFieldSpec is called when entering the ruleVariableTypeValueFieldSpec production.
	EnterRuleVariableTypeValueFieldSpec(c *RuleVariableTypeValueFieldSpecContext)

	// EnterRuleFixedTypeValueSetFieldSpec is called when entering the ruleFixedTypeValueSetFieldSpec production.
	EnterRuleFixedTypeValueSetFieldSpec(c *RuleFixedTypeValueSetFieldSpecContext)

	// EnterRuleVariableTypeValueSetFieldSpec is called when entering the ruleVariableTypeValueSetFieldSpec production.
	EnterRuleVariableTypeValueSetFieldSpec(c *RuleVariableTypeValueSetFieldSpecContext)

	// EnterRuleObjectFieldSpec is called when entering the ruleObjectFieldSpec production.
	EnterRuleObjectFieldSpec(c *RuleObjectFieldSpecContext)

	// EnterRuleObjectSetFieldSpec is called when entering the ruleObjectSetFieldSpec production.
	EnterRuleObjectSetFieldSpec(c *RuleObjectSetFieldSpecContext)

	// EnterRuleTypeOptionalitySpec is called when entering the ruleTypeOptionalitySpec production.
	EnterRuleTypeOptionalitySpec(c *RuleTypeOptionalitySpecContext)

	// EnterRuleValueOptionalitySpec is called when entering the ruleValueOptionalitySpec production.
	EnterRuleValueOptionalitySpec(c *RuleValueOptionalitySpecContext)

	// EnterRuleValueSetOptionalitySpec is called when entering the ruleValueSetOptionalitySpec production.
	EnterRuleValueSetOptionalitySpec(c *RuleValueSetOptionalitySpecContext)

	// EnterRuleObjectOptionalitySpec is called when entering the ruleObjectOptionalitySpec production.
	EnterRuleObjectOptionalitySpec(c *RuleObjectOptionalitySpecContext)

	// EnterRuleObjectSetOptionalitySpec is called when entering the ruleObjectSetOptionalitySpec production.
	EnterRuleObjectSetOptionalitySpec(c *RuleObjectSetOptionalitySpecContext)

	// EnterRuleWithSyntaxSpec is called when entering the ruleWithSyntaxSpec production.
	EnterRuleWithSyntaxSpec(c *RuleWithSyntaxSpecContext)

	// EnterRuleSyntaxList is called when entering the ruleSyntaxList production.
	EnterRuleSyntaxList(c *RuleSyntaxListContext)

	// EnterRuleTokenOrGroupSpecList is called when entering the ruleTokenOrGroupSpecList production.
	EnterRuleTokenOrGroupSpecList(c *RuleTokenOrGroupSpecListContext)

	// EnterRuleTokenOrGroupSpec is called when entering the ruleTokenOrGroupSpec production.
	EnterRuleTokenOrGroupSpec(c *RuleTokenOrGroupSpecContext)

	// EnterRuleRequiredToken is called when entering the ruleRequiredToken production.
	EnterRuleRequiredToken(c *RuleRequiredTokenContext)

	// EnterRuleOptionalGroup is called when entering the ruleOptionalGroup production.
	EnterRuleOptionalGroup(c *RuleOptionalGroupContext)

	// EnterRuleParameterizedObjectClass is called when entering the ruleParameterizedObjectClass production.
	EnterRuleParameterizedObjectClass(c *RuleParameterizedObjectClassContext)

	// EnterRuleObjectSet is called when entering the ruleObjectSet production.
	EnterRuleObjectSet(c *RuleObjectSetContext)

	// EnterRuleObjectSetSpec is called when entering the ruleObjectSetSpec production.
	EnterRuleObjectSetSpec(c *RuleObjectSetSpecContext)

	// EnterRuleRootElementSetSpec is called when entering the ruleRootElementSetSpec production.
	EnterRuleRootElementSetSpec(c *RuleRootElementSetSpecContext)

	// EnterRuleAdditionalElementSetSpec is called when entering the ruleAdditionalElementSetSpec production.
	EnterRuleAdditionalElementSetSpec(c *RuleAdditionalElementSetSpecContext)

	// EnterRuleElementSetSpecs is called when entering the ruleElementSetSpecs production.
	EnterRuleElementSetSpecs(c *RuleElementSetSpecsContext)

	// EnterRuleElementSetSpec is called when entering the ruleElementSetSpec production.
	EnterRuleElementSetSpec(c *RuleElementSetSpecContext)

	// EnterRuleUnions is called when entering the ruleUnions production.
	EnterRuleUnions(c *RuleUnionsContext)

	// EnterRuleExclusions is called when entering the ruleExclusions production.
	EnterRuleExclusions(c *RuleExclusionsContext)

	// EnterRuleIntersections is called when entering the ruleIntersections production.
	EnterRuleIntersections(c *RuleIntersectionsContext)

	// EnterRuleUElems is called when entering the ruleUElems production.
	EnterRuleUElems(c *RuleUElemsContext)

	// EnterRuleUnionMark is called when entering the ruleUnionMark production.
	EnterRuleUnionMark(c *RuleUnionMarkContext)

	// EnterRuleLowerEndValue is called when entering the ruleLowerEndValue production.
	EnterRuleLowerEndValue(c *RuleLowerEndValueContext)

	// EnterRuleUpperEndValue is called when entering the ruleUpperEndValue production.
	EnterRuleUpperEndValue(c *RuleUpperEndValueContext)

	// EnterRuleIncludes is called when entering the ruleIncludes production.
	EnterRuleIncludes(c *RuleIncludesContext)

	// EnterRuleLowerEndpoint is called when entering the ruleLowerEndpoint production.
	EnterRuleLowerEndpoint(c *RuleLowerEndpointContext)

	// EnterRuleUpperEndpoint is called when entering the ruleUpperEndpoint production.
	EnterRuleUpperEndpoint(c *RuleUpperEndpointContext)

	// EnterRuleLevel is called when entering the ruleLevel production.
	EnterRuleLevel(c *RuleLevelContext)

	// EnterRuleComponentIdList is called when entering the ruleComponentIdList production.
	EnterRuleComponentIdList(c *RuleComponentIdListContext)

	// EnterRuleAtNotation is called when entering the ruleAtNotation production.
	EnterRuleAtNotation(c *RuleAtNotationContext)

	// EnterRuleAtNotationList is called when entering the ruleAtNotationList production.
	EnterRuleAtNotationList(c *RuleAtNotationListContext)

	// EnterRuleUserDefinedConstraintParameter is called when entering the ruleUserDefinedConstraintParameter production.
	EnterRuleUserDefinedConstraintParameter(c *RuleUserDefinedConstraintParameterContext)

	// EnterRuleUserDefinedConstraintParameterList is called when entering the ruleUserDefinedConstraintParameterList production.
	EnterRuleUserDefinedConstraintParameterList(c *RuleUserDefinedConstraintParameterListContext)

	// EnterRuleUserDefinedConstraint is called when entering the ruleUserDefinedConstraint production.
	EnterRuleUserDefinedConstraint(c *RuleUserDefinedConstraintContext)

	// EnterRuleSimpleTableConstraint is called when entering the ruleSimpleTableConstraint production.
	EnterRuleSimpleTableConstraint(c *RuleSimpleTableConstraintContext)

	// EnterRuleComponentRelationConstraint is called when entering the ruleComponentRelationConstraint production.
	EnterRuleComponentRelationConstraint(c *RuleComponentRelationConstraintContext)

	// EnterRuleTableConstraint is called when entering the ruleTableConstraint production.
	EnterRuleTableConstraint(c *RuleTableConstraintContext)

	// EnterRuleContentsConstraint is called when entering the ruleContentsConstraint production.
	EnterRuleContentsConstraint(c *RuleContentsConstraintContext)

	// EnterRuleSubtypeConstraint is called when entering the ruleSubtypeConstraint production.
	EnterRuleSubtypeConstraint(c *RuleSubtypeConstraintContext)

	// EnterRuleGeneralConstraint is called when entering the ruleGeneralConstraint production.
	EnterRuleGeneralConstraint(c *RuleGeneralConstraintContext)

	// EnterRuleConstraintSpec is called when entering the ruleConstraintSpec production.
	EnterRuleConstraintSpec(c *RuleConstraintSpecContext)

	// EnterRuleSignedNumber is called when entering the ruleSignedNumber production.
	EnterRuleSignedNumber(c *RuleSignedNumberContext)

	// EnterRuleExceptionIdentification is called when entering the ruleExceptionIdentification production.
	EnterRuleExceptionIdentification(c *RuleExceptionIdentificationContext)

	// EnterRuleExceptionSpec is called when entering the ruleExceptionSpec production.
	EnterRuleExceptionSpec(c *RuleExceptionSpecContext)

	// EnterRuleConstraint is called when entering the ruleConstraint production.
	EnterRuleConstraint(c *RuleConstraintContext)

	// EnterRuleSingleTypeConstraint is called when entering the ruleSingleTypeConstraint production.
	EnterRuleSingleTypeConstraint(c *RuleSingleTypeConstraintContext)

	// EnterRuleValueConstraint is called when entering the ruleValueConstraint production.
	EnterRuleValueConstraint(c *RuleValueConstraintContext)

	// EnterRulePresenceConstraint is called when entering the rulePresenceConstraint production.
	EnterRulePresenceConstraint(c *RulePresenceConstraintContext)

	// EnterRuleComponentConstraint is called when entering the ruleComponentConstraint production.
	EnterRuleComponentConstraint(c *RuleComponentConstraintContext)

	// EnterRuleNamedConstraint is called when entering the ruleNamedConstraint production.
	EnterRuleNamedConstraint(c *RuleNamedConstraintContext)

	// EnterRuleTypeConstraints is called when entering the ruleTypeConstraints production.
	EnterRuleTypeConstraints(c *RuleTypeConstraintsContext)

	// EnterRuleFullSpecification is called when entering the ruleFullSpecification production.
	EnterRuleFullSpecification(c *RuleFullSpecificationContext)

	// EnterRulePartialSpecification is called when entering the rulePartialSpecification production.
	EnterRulePartialSpecification(c *RulePartialSpecificationContext)

	// EnterRuleMultipleTypeConstraints is called when entering the ruleMultipleTypeConstraints production.
	EnterRuleMultipleTypeConstraints(c *RuleMultipleTypeConstraintsContext)

	// EnterRuleSimpleString is called when entering the ruleSimpleString production.
	EnterRuleSimpleString(c *RuleSimpleStringContext)

	// EnterRuleSingleValue is called when entering the ruleSingleValue production.
	EnterRuleSingleValue(c *RuleSingleValueContext)

	// EnterRuleContainedSubtype is called when entering the ruleContainedSubtype production.
	EnterRuleContainedSubtype(c *RuleContainedSubtypeContext)

	// EnterRuleValueRange is called when entering the ruleValueRange production.
	EnterRuleValueRange(c *RuleValueRangeContext)

	// EnterRulePermittedAlphabet is called when entering the rulePermittedAlphabet production.
	EnterRulePermittedAlphabet(c *RulePermittedAlphabetContext)

	// EnterRuleSizeConstraint is called when entering the ruleSizeConstraint production.
	EnterRuleSizeConstraint(c *RuleSizeConstraintContext)

	// EnterRuleInnerTypeConstraints is called when entering the ruleInnerTypeConstraints production.
	EnterRuleInnerTypeConstraints(c *RuleInnerTypeConstraintsContext)

	// EnterRulePatternConstraint is called when entering the rulePatternConstraint production.
	EnterRulePatternConstraint(c *RulePatternConstraintContext)

	// EnterRulePropertySettings is called when entering the rulePropertySettings production.
	EnterRulePropertySettings(c *RulePropertySettingsContext)

	// EnterRuleDurationRange is called when entering the ruleDurationRange production.
	EnterRuleDurationRange(c *RuleDurationRangeContext)

	// EnterRuleTimePointRange is called when entering the ruleTimePointRange production.
	EnterRuleTimePointRange(c *RuleTimePointRangeContext)

	// EnterRuleRecurrenceRange is called when entering the ruleRecurrenceRange production.
	EnterRuleRecurrenceRange(c *RuleRecurrenceRangeContext)

	// EnterRuleSubtypeElements is called when entering the ruleSubtypeElements production.
	EnterRuleSubtypeElements(c *RuleSubtypeElementsContext)

	// EnterRuleObjectSetElements is called when entering the ruleObjectSetElements production.
	EnterRuleObjectSetElements(c *RuleObjectSetElementsContext)

	// EnterRuleElements is called when entering the ruleElements production.
	EnterRuleElements(c *RuleElementsContext)

	// EnterRuleElems is called when entering the ruleElems production.
	EnterRuleElems(c *RuleElemsContext)

	// EnterRuleIntersectionElements is called when entering the ruleIntersectionElements production.
	EnterRuleIntersectionElements(c *RuleIntersectionElementsContext)

	// EnterRuleIElems is called when entering the ruleIElems production.
	EnterRuleIElems(c *RuleIElemsContext)

	// EnterRuleIntersectionMark is called when entering the ruleIntersectionMark production.
	EnterRuleIntersectionMark(c *RuleIntersectionMarkContext)

	// EnterRuleSimpleDefinedValue is called when entering the ruleSimpleDefinedValue production.
	EnterRuleSimpleDefinedValue(c *RuleSimpleDefinedValueContext)

	// EnterRuleActualParameterList is called when entering the ruleActualParameterList production.
	EnterRuleActualParameterList(c *RuleActualParameterListContext)

	// EnterRuleActualParameters is called when entering the ruleActualParameters production.
	EnterRuleActualParameters(c *RuleActualParametersContext)

	// EnterRuleActualParameter is called when entering the ruleActualParameter production.
	EnterRuleActualParameter(c *RuleActualParameterContext)

	// EnterRuleParameterizedObjectSet is called when entering the ruleParameterizedObjectSet production.
	EnterRuleParameterizedObjectSet(c *RuleParameterizedObjectSetContext)

	// EnterRuleParameterizedValue is called when entering the ruleParameterizedValue production.
	EnterRuleParameterizedValue(c *RuleParameterizedValueContext)

	// EnterRuleParameterizedTypeAssignment is called when entering the ruleParameterizedTypeAssignment production.
	EnterRuleParameterizedTypeAssignment(c *RuleParameterizedTypeAssignmentContext)

	// EnterRuleParameterizedValueAssignment is called when entering the ruleParameterizedValueAssignment production.
	EnterRuleParameterizedValueAssignment(c *RuleParameterizedValueAssignmentContext)

	// EnterRuleParameterizedValueSetTypeAssignment is called when entering the ruleParameterizedValueSetTypeAssignment production.
	EnterRuleParameterizedValueSetTypeAssignment(c *RuleParameterizedValueSetTypeAssignmentContext)

	// EnterRuleParameterizedObjectClassAssignment is called when entering the ruleParameterizedObjectClassAssignment production.
	EnterRuleParameterizedObjectClassAssignment(c *RuleParameterizedObjectClassAssignmentContext)

	// EnterRuleParameterizedObjectAssignment is called when entering the ruleParameterizedObjectAssignment production.
	EnterRuleParameterizedObjectAssignment(c *RuleParameterizedObjectAssignmentContext)

	// EnterRuleParameterizedObjectSetAssignment is called when entering the ruleParameterizedObjectSetAssignment production.
	EnterRuleParameterizedObjectSetAssignment(c *RuleParameterizedObjectSetAssignmentContext)

	// EnterRuleParameterList is called when entering the ruleParameterList production.
	EnterRuleParameterList(c *RuleParameterListContext)

	// EnterRuleParameters is called when entering the ruleParameters production.
	EnterRuleParameters(c *RuleParametersContext)

	// EnterRuleParameter is called when entering the ruleParameter production.
	EnterRuleParameter(c *RuleParameterContext)

	// EnterRuleParamGovernor is called when entering the ruleParamGovernor production.
	EnterRuleParamGovernor(c *RuleParamGovernorContext)

	// EnterRuleDummyReference is called when entering the ruleDummyReference production.
	EnterRuleDummyReference(c *RuleDummyReferenceContext)

	// EnterRuleGovernor is called when entering the ruleGovernor production.
	EnterRuleGovernor(c *RuleGovernorContext)

	// EnterRuleDummyGovernor is called when entering the ruleDummyGovernor production.
	EnterRuleDummyGovernor(c *RuleDummyGovernorContext)

	// EnterRuleEncodingControlSections is called when entering the ruleEncodingControlSections production.
	EnterRuleEncodingControlSections(c *RuleEncodingControlSectionsContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitRuleModules is called when exiting the ruleModules production.
	ExitRuleModules(c *RuleModulesContext)

	// ExitRuleModuleDefinition is called when exiting the ruleModuleDefinition production.
	ExitRuleModuleDefinition(c *RuleModuleDefinitionContext)

	// ExitRuleModuleReference is called when exiting the ruleModuleReference production.
	ExitRuleModuleReference(c *RuleModuleReferenceContext)

	// ExitRuleModuleIdentifier is called when exiting the ruleModuleIdentifier production.
	ExitRuleModuleIdentifier(c *RuleModuleIdentifierContext)

	// ExitRuleDefinitiveIdentification is called when exiting the ruleDefinitiveIdentification production.
	ExitRuleDefinitiveIdentification(c *RuleDefinitiveIdentificationContext)

	// ExitRuleDefinitiveOID is called when exiting the ruleDefinitiveOID production.
	ExitRuleDefinitiveOID(c *RuleDefinitiveOIDContext)

	// ExitRuleDefinitiveObjIdComponentList is called when exiting the ruleDefinitiveObjIdComponentList production.
	ExitRuleDefinitiveObjIdComponentList(c *RuleDefinitiveObjIdComponentListContext)

	// ExitRuleDefinitiveObjIdComponent is called when exiting the ruleDefinitiveObjIdComponent production.
	ExitRuleDefinitiveObjIdComponent(c *RuleDefinitiveObjIdComponentContext)

	// ExitRuleNameForm is called when exiting the ruleNameForm production.
	ExitRuleNameForm(c *RuleNameFormContext)

	// ExitRuleDefinitiveNumberForm is called when exiting the ruleDefinitiveNumberForm production.
	ExitRuleDefinitiveNumberForm(c *RuleDefinitiveNumberFormContext)

	// ExitRuleDefinitiveNameAndNumberForm is called when exiting the ruleDefinitiveNameAndNumberForm production.
	ExitRuleDefinitiveNameAndNumberForm(c *RuleDefinitiveNameAndNumberFormContext)

	// ExitRuleDefinitiveOIDAndIRI is called when exiting the ruleDefinitiveOIDAndIRI production.
	ExitRuleDefinitiveOIDAndIRI(c *RuleDefinitiveOIDAndIRIContext)

	// ExitRuleFirstArcIdentifier is called when exiting the ruleFirstArcIdentifier production.
	ExitRuleFirstArcIdentifier(c *RuleFirstArcIdentifierContext)

	// ExitRuleSubsequentArcIdentifier is called when exiting the ruleSubsequentArcIdentifier production.
	ExitRuleSubsequentArcIdentifier(c *RuleSubsequentArcIdentifierContext)

	// ExitRuleIRIValue is called when exiting the ruleIRIValue production.
	ExitRuleIRIValue(c *RuleIRIValueContext)

	// ExitRuleEncodingReferenceDefault is called when exiting the ruleEncodingReferenceDefault production.
	ExitRuleEncodingReferenceDefault(c *RuleEncodingReferenceDefaultContext)

	// ExitRuleTagDefault is called when exiting the ruleTagDefault production.
	ExitRuleTagDefault(c *RuleTagDefaultContext)

	// ExitRuleExtensionDefault is called when exiting the ruleExtensionDefault production.
	ExitRuleExtensionDefault(c *RuleExtensionDefaultContext)

	// ExitRuleModuleBody is called when exiting the ruleModuleBody production.
	ExitRuleModuleBody(c *RuleModuleBodyContext)

	// ExitRuleExports is called when exiting the ruleExports production.
	ExitRuleExports(c *RuleExportsContext)

	// ExitRuleSymbolsExported is called when exiting the ruleSymbolsExported production.
	ExitRuleSymbolsExported(c *RuleSymbolsExportedContext)

	// ExitRuleSymbolList is called when exiting the ruleSymbolList production.
	ExitRuleSymbolList(c *RuleSymbolListContext)

	// ExitRuleSymbol is called when exiting the ruleSymbol production.
	ExitRuleSymbol(c *RuleSymbolContext)

	// ExitRuleReference is called when exiting the ruleReference production.
	ExitRuleReference(c *RuleReferenceContext)

	// ExitRuleIdentifier is called when exiting the ruleIdentifier production.
	ExitRuleIdentifier(c *RuleIdentifierContext)

	// ExitRuleTypeReference is called when exiting the ruleTypeReference production.
	ExitRuleTypeReference(c *RuleTypeReferenceContext)

	// ExitRuleValueReference is called when exiting the ruleValueReference production.
	ExitRuleValueReference(c *RuleValueReferenceContext)

	// ExitRuleObjectReference is called when exiting the ruleObjectReference production.
	ExitRuleObjectReference(c *RuleObjectReferenceContext)

	// ExitRuleObjectClassReference is called when exiting the ruleObjectClassReference production.
	ExitRuleObjectClassReference(c *RuleObjectClassReferenceContext)

	// ExitRuleObjectSetReference is called when exiting the ruleObjectSetReference production.
	ExitRuleObjectSetReference(c *RuleObjectSetReferenceContext)

	// ExitRuleParameterizedReference is called when exiting the ruleParameterizedReference production.
	ExitRuleParameterizedReference(c *RuleParameterizedReferenceContext)

	// ExitRuleExternalTypeReference is called when exiting the ruleExternalTypeReference production.
	ExitRuleExternalTypeReference(c *RuleExternalTypeReferenceContext)

	// ExitRuleExternalValueReference is called when exiting the ruleExternalValueReference production.
	ExitRuleExternalValueReference(c *RuleExternalValueReferenceContext)

	// ExitRuleExternalObjectClassReference is called when exiting the ruleExternalObjectClassReference production.
	ExitRuleExternalObjectClassReference(c *RuleExternalObjectClassReferenceContext)

	// ExitRuleExternalObjectReference is called when exiting the ruleExternalObjectReference production.
	ExitRuleExternalObjectReference(c *RuleExternalObjectReferenceContext)

	// ExitRuleExternalObjectSetReference is called when exiting the ruleExternalObjectSetReference production.
	ExitRuleExternalObjectSetReference(c *RuleExternalObjectSetReferenceContext)

	// ExitRuleTypeFieldReference is called when exiting the ruleTypeFieldReference production.
	ExitRuleTypeFieldReference(c *RuleTypeFieldReferenceContext)

	// ExitRuleValueFieldReference is called when exiting the ruleValueFieldReference production.
	ExitRuleValueFieldReference(c *RuleValueFieldReferenceContext)

	// ExitRuleValueSetFieldReference is called when exiting the ruleValueSetFieldReference production.
	ExitRuleValueSetFieldReference(c *RuleValueSetFieldReferenceContext)

	// ExitRuleObjectFieldReference is called when exiting the ruleObjectFieldReference production.
	ExitRuleObjectFieldReference(c *RuleObjectFieldReferenceContext)

	// ExitRuleObjectSetFieldReference is called when exiting the ruleObjectSetFieldReference production.
	ExitRuleObjectSetFieldReference(c *RuleObjectSetFieldReferenceContext)

	// ExitRuleUsefulObjectClassReference is called when exiting the ruleUsefulObjectClassReference production.
	ExitRuleUsefulObjectClassReference(c *RuleUsefulObjectClassReferenceContext)

	// ExitRuleImports is called when exiting the ruleImports production.
	ExitRuleImports(c *RuleImportsContext)

	// ExitRuleSymbolsImported is called when exiting the ruleSymbolsImported production.
	ExitRuleSymbolsImported(c *RuleSymbolsImportedContext)

	// ExitRuleSymbolsFromModuleList is called when exiting the ruleSymbolsFromModuleList production.
	ExitRuleSymbolsFromModuleList(c *RuleSymbolsFromModuleListContext)

	// ExitRuleSymbolsFromModule is called when exiting the ruleSymbolsFromModule production.
	ExitRuleSymbolsFromModule(c *RuleSymbolsFromModuleContext)

	// ExitRuleGlobalModuleReference is called when exiting the ruleGlobalModuleReference production.
	ExitRuleGlobalModuleReference(c *RuleGlobalModuleReferenceContext)

	// ExitRuleAssignedIdentifier is called when exiting the ruleAssignedIdentifier production.
	ExitRuleAssignedIdentifier(c *RuleAssignedIdentifierContext)

	// ExitRuleAssignmentList is called when exiting the ruleAssignmentList production.
	ExitRuleAssignmentList(c *RuleAssignmentListContext)

	// ExitRuleAssignment is called when exiting the ruleAssignment production.
	ExitRuleAssignment(c *RuleAssignmentContext)

	// ExitRuleTypeAssignment is called when exiting the ruleTypeAssignment production.
	ExitRuleTypeAssignment(c *RuleTypeAssignmentContext)

	// ExitRuleValueAssignment is called when exiting the ruleValueAssignment production.
	ExitRuleValueAssignment(c *RuleValueAssignmentContext)

	// ExitRuleXMLValueAssignment is called when exiting the ruleXMLValueAssignment production.
	ExitRuleXMLValueAssignment(c *RuleXMLValueAssignmentContext)

	// ExitRuleValueSetTypeAssignment is called when exiting the ruleValueSetTypeAssignment production.
	ExitRuleValueSetTypeAssignment(c *RuleValueSetTypeAssignmentContext)

	// ExitRuleObjectAssignment is called when exiting the ruleObjectAssignment production.
	ExitRuleObjectAssignment(c *RuleObjectAssignmentContext)

	// ExitRuleObjectClassAssignment is called when exiting the ruleObjectClassAssignment production.
	ExitRuleObjectClassAssignment(c *RuleObjectClassAssignmentContext)

	// ExitRuleObjectSetAssignment is called when exiting the ruleObjectSetAssignment production.
	ExitRuleObjectSetAssignment(c *RuleObjectSetAssignmentContext)

	// ExitRuleParameterizedAssignment is called when exiting the ruleParameterizedAssignment production.
	ExitRuleParameterizedAssignment(c *RuleParameterizedAssignmentContext)

	// ExitRuleObjectIdentifierValue is called when exiting the ruleObjectIdentifierValue production.
	ExitRuleObjectIdentifierValue(c *RuleObjectIdentifierValueContext)

	// ExitRuleObjIdComponentsList is called when exiting the ruleObjIdComponentsList production.
	ExitRuleObjIdComponentsList(c *RuleObjIdComponentsListContext)

	// ExitRuleObjIdComponents is called when exiting the ruleObjIdComponents production.
	ExitRuleObjIdComponents(c *RuleObjIdComponentsContext)

	// ExitRuleNumberForm is called when exiting the ruleNumberForm production.
	ExitRuleNumberForm(c *RuleNumberFormContext)

	// ExitRuleNameAndNumberForm is called when exiting the ruleNameAndNumberForm production.
	ExitRuleNameAndNumberForm(c *RuleNameAndNumberFormContext)

	// ExitRuleDefinedValue is called when exiting the ruleDefinedValue production.
	ExitRuleDefinedValue(c *RuleDefinedValueContext)

	// ExitRuleNamedBit is called when exiting the ruleNamedBit production.
	ExitRuleNamedBit(c *RuleNamedBitContext)

	// ExitRuleNamedBitList is called when exiting the ruleNamedBitList production.
	ExitRuleNamedBitList(c *RuleNamedBitListContext)

	// ExitRuleRestrictedCharacterStringType is called when exiting the ruleRestrictedCharacterStringType production.
	ExitRuleRestrictedCharacterStringType(c *RuleRestrictedCharacterStringTypeContext)

	// ExitRuleUnrestrictedCharacterStringType is called when exiting the ruleUnrestrictedCharacterStringType production.
	ExitRuleUnrestrictedCharacterStringType(c *RuleUnrestrictedCharacterStringTypeContext)

	// ExitRuleAlternativeTypeList is called when exiting the ruleAlternativeTypeList production.
	ExitRuleAlternativeTypeList(c *RuleAlternativeTypeListContext)

	// ExitRuleRootAlternativeTypeList is called when exiting the ruleRootAlternativeTypeList production.
	ExitRuleRootAlternativeTypeList(c *RuleRootAlternativeTypeListContext)

	// ExitRuleVersionNumber is called when exiting the ruleVersionNumber production.
	ExitRuleVersionNumber(c *RuleVersionNumberContext)

	// ExitRuleExtensionAdditionAlternativesGroup is called when exiting the ruleExtensionAdditionAlternativesGroup production.
	ExitRuleExtensionAdditionAlternativesGroup(c *RuleExtensionAdditionAlternativesGroupContext)

	// ExitRuleExtensionAdditionAlternative is called when exiting the ruleExtensionAdditionAlternative production.
	ExitRuleExtensionAdditionAlternative(c *RuleExtensionAdditionAlternativeContext)

	// ExitRuleExtensionAdditionAlternativesList is called when exiting the ruleExtensionAdditionAlternativesList production.
	ExitRuleExtensionAdditionAlternativesList(c *RuleExtensionAdditionAlternativesListContext)

	// ExitRuleExtensionAdditionAlternatives is called when exiting the ruleExtensionAdditionAlternatives production.
	ExitRuleExtensionAdditionAlternatives(c *RuleExtensionAdditionAlternativesContext)

	// ExitRuleEnumerationItem is called when exiting the ruleEnumerationItem production.
	ExitRuleEnumerationItem(c *RuleEnumerationItemContext)

	// ExitRuleEnumeration is called when exiting the ruleEnumeration production.
	ExitRuleEnumeration(c *RuleEnumerationContext)

	// ExitRuleRootEnumeration is called when exiting the ruleRootEnumeration production.
	ExitRuleRootEnumeration(c *RuleRootEnumerationContext)

	// ExitRuleAdditionalEnumeration is called when exiting the ruleAdditionalEnumeration production.
	ExitRuleAdditionalEnumeration(c *RuleAdditionalEnumerationContext)

	// ExitRuleAlternativeTypeLists is called when exiting the ruleAlternativeTypeLists production.
	ExitRuleAlternativeTypeLists(c *RuleAlternativeTypeListsContext)

	// ExitRuleEnumerations is called when exiting the ruleEnumerations production.
	ExitRuleEnumerations(c *RuleEnumerationsContext)

	// ExitRuleNamedNumber is called when exiting the ruleNamedNumber production.
	ExitRuleNamedNumber(c *RuleNamedNumberContext)

	// ExitRuleNamedNumberList is called when exiting the ruleNamedNumberList production.
	ExitRuleNamedNumberList(c *RuleNamedNumberListContext)

	// ExitRuleComponentType is called when exiting the ruleComponentType production.
	ExitRuleComponentType(c *RuleComponentTypeContext)

	// ExitRuleExtensionAdditionGroup is called when exiting the ruleExtensionAdditionGroup production.
	ExitRuleExtensionAdditionGroup(c *RuleExtensionAdditionGroupContext)

	// ExitRuleExtensionAddition is called when exiting the ruleExtensionAddition production.
	ExitRuleExtensionAddition(c *RuleExtensionAdditionContext)

	// ExitRuleComponentTypeList is called when exiting the ruleComponentTypeList production.
	ExitRuleComponentTypeList(c *RuleComponentTypeListContext)

	// ExitRuleExtensionAdditionList is called when exiting the ruleExtensionAdditionList production.
	ExitRuleExtensionAdditionList(c *RuleExtensionAdditionListContext)

	// ExitRuleRootComponentTypeList is called when exiting the ruleRootComponentTypeList production.
	ExitRuleRootComponentTypeList(c *RuleRootComponentTypeListContext)

	// ExitRuleExtensionAdditions is called when exiting the ruleExtensionAdditions production.
	ExitRuleExtensionAdditions(c *RuleExtensionAdditionsContext)

	// ExitRuleExtensionEndMarker is called when exiting the ruleExtensionEndMarker production.
	ExitRuleExtensionEndMarker(c *RuleExtensionEndMarkerContext)

	// ExitRuleComponentTypeLists is called when exiting the ruleComponentTypeLists production.
	ExitRuleComponentTypeLists(c *RuleComponentTypeListsContext)

	// ExitRuleExtensionAndException is called when exiting the ruleExtensionAndException production.
	ExitRuleExtensionAndException(c *RuleExtensionAndExceptionContext)

	// ExitRuleOptionalExtensionMarker is called when exiting the ruleOptionalExtensionMarker production.
	ExitRuleOptionalExtensionMarker(c *RuleOptionalExtensionMarkerContext)

	// ExitRuleEncodingReference is called when exiting the ruleEncodingReference production.
	ExitRuleEncodingReference(c *RuleEncodingReferenceContext)

	// ExitRuleClass is called when exiting the ruleClass production.
	ExitRuleClass(c *RuleClassContext)

	// ExitRuleClassNumber is called when exiting the ruleClassNumber production.
	ExitRuleClassNumber(c *RuleClassNumberContext)

	// ExitRuleEncodingInstruction is called when exiting the ruleEncodingInstruction production.
	ExitRuleEncodingInstruction(c *RuleEncodingInstructionContext)

	// ExitRuleTag is called when exiting the ruleTag production.
	ExitRuleTag(c *RuleTagContext)

	// ExitRuleEncodingPrefix is called when exiting the ruleEncodingPrefix production.
	ExitRuleEncodingPrefix(c *RuleEncodingPrefixContext)

	// ExitRuleTaggedType is called when exiting the ruleTaggedType production.
	ExitRuleTaggedType(c *RuleTaggedTypeContext)

	// ExitRuleEncodingPrefixedType is called when exiting the ruleEncodingPrefixedType production.
	ExitRuleEncodingPrefixedType(c *RuleEncodingPrefixedTypeContext)

	// ExitRuleBitStringType is called when exiting the ruleBitStringType production.
	ExitRuleBitStringType(c *RuleBitStringTypeContext)

	// ExitRuleBooleanType is called when exiting the ruleBooleanType production.
	ExitRuleBooleanType(c *RuleBooleanTypeContext)

	// ExitRuleCharacterStringType is called when exiting the ruleCharacterStringType production.
	ExitRuleCharacterStringType(c *RuleCharacterStringTypeContext)

	// ExitRuleChoiceType is called when exiting the ruleChoiceType production.
	ExitRuleChoiceType(c *RuleChoiceTypeContext)

	// ExitRuleDateType is called when exiting the ruleDateType production.
	ExitRuleDateType(c *RuleDateTypeContext)

	// ExitRuleDateTimeType is called when exiting the ruleDateTimeType production.
	ExitRuleDateTimeType(c *RuleDateTimeTypeContext)

	// ExitRuleDurationType is called when exiting the ruleDurationType production.
	ExitRuleDurationType(c *RuleDurationTypeContext)

	// ExitRuleEmbeddedPDVType is called when exiting the ruleEmbeddedPDVType production.
	ExitRuleEmbeddedPDVType(c *RuleEmbeddedPDVTypeContext)

	// ExitRuleEnumeratedType is called when exiting the ruleEnumeratedType production.
	ExitRuleEnumeratedType(c *RuleEnumeratedTypeContext)

	// ExitRuleExternalType is called when exiting the ruleExternalType production.
	ExitRuleExternalType(c *RuleExternalTypeContext)

	// ExitRuleInstanceOfType is called when exiting the ruleInstanceOfType production.
	ExitRuleInstanceOfType(c *RuleInstanceOfTypeContext)

	// ExitRuleIntegerType is called when exiting the ruleIntegerType production.
	ExitRuleIntegerType(c *RuleIntegerTypeContext)

	// ExitRuleIRIType is called when exiting the ruleIRIType production.
	ExitRuleIRIType(c *RuleIRITypeContext)

	// ExitRuleNullType is called when exiting the ruleNullType production.
	ExitRuleNullType(c *RuleNullTypeContext)

	// ExitRuleObjectClassFieldType is called when exiting the ruleObjectClassFieldType production.
	ExitRuleObjectClassFieldType(c *RuleObjectClassFieldTypeContext)

	// ExitRuleObjectIdentifierType is called when exiting the ruleObjectIdentifierType production.
	ExitRuleObjectIdentifierType(c *RuleObjectIdentifierTypeContext)

	// ExitRuleOctetStringType is called when exiting the ruleOctetStringType production.
	ExitRuleOctetStringType(c *RuleOctetStringTypeContext)

	// ExitRuleRealType is called when exiting the ruleRealType production.
	ExitRuleRealType(c *RuleRealTypeContext)

	// ExitRuleRelativeIRIType is called when exiting the ruleRelativeIRIType production.
	ExitRuleRelativeIRIType(c *RuleRelativeIRITypeContext)

	// ExitRuleRelativeOIDType is called when exiting the ruleRelativeOIDType production.
	ExitRuleRelativeOIDType(c *RuleRelativeOIDTypeContext)

	// ExitRuleSequenceType is called when exiting the ruleSequenceType production.
	ExitRuleSequenceType(c *RuleSequenceTypeContext)

	// ExitRuleSequenceOfType is called when exiting the ruleSequenceOfType production.
	ExitRuleSequenceOfType(c *RuleSequenceOfTypeContext)

	// ExitRuleSetType is called when exiting the ruleSetType production.
	ExitRuleSetType(c *RuleSetTypeContext)

	// ExitRuleSetOfType is called when exiting the ruleSetOfType production.
	ExitRuleSetOfType(c *RuleSetOfTypeContext)

	// ExitRulePrefixedType is called when exiting the rulePrefixedType production.
	ExitRulePrefixedType(c *RulePrefixedTypeContext)

	// ExitRuleTimeType is called when exiting the ruleTimeType production.
	ExitRuleTimeType(c *RuleTimeTypeContext)

	// ExitRuleTimeOfDayType is called when exiting the ruleTimeOfDayType production.
	ExitRuleTimeOfDayType(c *RuleTimeOfDayTypeContext)

	// ExitRuleBuiltinType is called when exiting the ruleBuiltinType production.
	ExitRuleBuiltinType(c *RuleBuiltinTypeContext)

	// ExitRuleSimpleDefinedType is called when exiting the ruleSimpleDefinedType production.
	ExitRuleSimpleDefinedType(c *RuleSimpleDefinedTypeContext)

	// ExitRuleParameterizedType is called when exiting the ruleParameterizedType production.
	ExitRuleParameterizedType(c *RuleParameterizedTypeContext)

	// ExitRuleParameterizedValueSetType is called when exiting the ruleParameterizedValueSetType production.
	ExitRuleParameterizedValueSetType(c *RuleParameterizedValueSetTypeContext)

	// ExitRuleDefinedType is called when exiting the ruleDefinedType production.
	ExitRuleDefinedType(c *RuleDefinedTypeContext)

	// ExitRuleUsefulType is called when exiting the ruleUsefulType production.
	ExitRuleUsefulType(c *RuleUsefulTypeContext)

	// ExitRuleSelectionType is called when exiting the ruleSelectionType production.
	ExitRuleSelectionType(c *RuleSelectionTypeContext)

	// ExitRuleTypeFromObject is called when exiting the ruleTypeFromObject production.
	ExitRuleTypeFromObject(c *RuleTypeFromObjectContext)

	// ExitRuleValueSetFromObjects is called when exiting the ruleValueSetFromObjects production.
	ExitRuleValueSetFromObjects(c *RuleValueSetFromObjectsContext)

	// ExitRuleReferencedType is called when exiting the ruleReferencedType production.
	ExitRuleReferencedType(c *RuleReferencedTypeContext)

	// ExitRuleTypeForConstraint is called when exiting the ruleTypeForConstraint production.
	ExitRuleTypeForConstraint(c *RuleTypeForConstraintContext)

	// ExitRuleNamedType is called when exiting the ruleNamedType production.
	ExitRuleNamedType(c *RuleNamedTypeContext)

	// ExitRuleTypeWithConstraint is called when exiting the ruleTypeWithConstraint production.
	ExitRuleTypeWithConstraint(c *RuleTypeWithConstraintContext)

	// ExitRuleConstrainedType is called when exiting the ruleConstrainedType production.
	ExitRuleConstrainedType(c *RuleConstrainedTypeContext)

	// ExitRuleType is called when exiting the ruleType production.
	ExitRuleType(c *RuleTypeContext)

	// ExitRuleIdentifierList is called when exiting the ruleIdentifierList production.
	ExitRuleIdentifierList(c *RuleIdentifierListContext)

	// ExitRuleCharsDefn is called when exiting the ruleCharsDefn production.
	ExitRuleCharsDefn(c *RuleCharsDefnContext)

	// ExitRuleCharSyms is called when exiting the ruleCharSyms production.
	ExitRuleCharSyms(c *RuleCharSymsContext)

	// ExitRuleGroup is called when exiting the ruleGroup production.
	ExitRuleGroup(c *RuleGroupContext)

	// ExitRulePlane is called when exiting the rulePlane production.
	ExitRulePlane(c *RulePlaneContext)

	// ExitRuleRow is called when exiting the ruleRow production.
	ExitRuleRow(c *RuleRowContext)

	// ExitRuleCell is called when exiting the ruleCell production.
	ExitRuleCell(c *RuleCellContext)

	// ExitRuleTableColumn is called when exiting the ruleTableColumn production.
	ExitRuleTableColumn(c *RuleTableColumnContext)

	// ExitRuleTableRow is called when exiting the ruleTableRow production.
	ExitRuleTableRow(c *RuleTableRowContext)

	// ExitRuleCharacterStringList is called when exiting the ruleCharacterStringList production.
	ExitRuleCharacterStringList(c *RuleCharacterStringListContext)

	// ExitRuleQuadruple is called when exiting the ruleQuadruple production.
	ExitRuleQuadruple(c *RuleQuadrupleContext)

	// ExitRuleTuple is called when exiting the ruleTuple production.
	ExitRuleTuple(c *RuleTupleContext)

	// ExitRuleRestrictedCharacterStringValue is called when exiting the ruleRestrictedCharacterStringValue production.
	ExitRuleRestrictedCharacterStringValue(c *RuleRestrictedCharacterStringValueContext)

	// ExitRuleUnrestrictedCharacterStringValue is called when exiting the ruleUnrestrictedCharacterStringValue production.
	ExitRuleUnrestrictedCharacterStringValue(c *RuleUnrestrictedCharacterStringValueContext)

	// ExitRuleNumericRealValue is called when exiting the ruleNumericRealValue production.
	ExitRuleNumericRealValue(c *RuleNumericRealValueContext)

	// ExitRuleSpecialRealValue is called when exiting the ruleSpecialRealValue production.
	ExitRuleSpecialRealValue(c *RuleSpecialRealValueContext)

	// ExitRuleRelativeOIDComponents is called when exiting the ruleRelativeOIDComponents production.
	ExitRuleRelativeOIDComponents(c *RuleRelativeOIDComponentsContext)

	// ExitRuleRelativeOIDComponentsList is called when exiting the ruleRelativeOIDComponentsList production.
	ExitRuleRelativeOIDComponentsList(c *RuleRelativeOIDComponentsListContext)

	// ExitRuleComponentValueList is called when exiting the ruleComponentValueList production.
	ExitRuleComponentValueList(c *RuleComponentValueListContext)

	// ExitRuleValueList is called when exiting the ruleValueList production.
	ExitRuleValueList(c *RuleValueListContext)

	// ExitRuleNamedValue is called when exiting the ruleNamedValue production.
	ExitRuleNamedValue(c *RuleNamedValueContext)

	// ExitRuleNamedValueList is called when exiting the ruleNamedValueList production.
	ExitRuleNamedValueList(c *RuleNamedValueListContext)

	// ExitRuleBitStringValue is called when exiting the ruleBitStringValue production.
	ExitRuleBitStringValue(c *RuleBitStringValueContext)

	// ExitRuleBooleanValue is called when exiting the ruleBooleanValue production.
	ExitRuleBooleanValue(c *RuleBooleanValueContext)

	// ExitRuleCharacterStringValue is called when exiting the ruleCharacterStringValue production.
	ExitRuleCharacterStringValue(c *RuleCharacterStringValueContext)

	// ExitRuleChoiceValue is called when exiting the ruleChoiceValue production.
	ExitRuleChoiceValue(c *RuleChoiceValueContext)

	// ExitRuleEmbeddedPDVValue is called when exiting the ruleEmbeddedPDVValue production.
	ExitRuleEmbeddedPDVValue(c *RuleEmbeddedPDVValueContext)

	// ExitRuleEnumeratedValue is called when exiting the ruleEnumeratedValue production.
	ExitRuleEnumeratedValue(c *RuleEnumeratedValueContext)

	// ExitRuleExternalValue is called when exiting the ruleExternalValue production.
	ExitRuleExternalValue(c *RuleExternalValueContext)

	// ExitRuleInstanceOfValue is called when exiting the ruleInstanceOfValue production.
	ExitRuleInstanceOfValue(c *RuleInstanceOfValueContext)

	// ExitRuleIntegerValue is called when exiting the ruleIntegerValue production.
	ExitRuleIntegerValue(c *RuleIntegerValueContext)

	// ExitRuleNullValue is called when exiting the ruleNullValue production.
	ExitRuleNullValue(c *RuleNullValueContext)

	// ExitRuleOctetStringValue is called when exiting the ruleOctetStringValue production.
	ExitRuleOctetStringValue(c *RuleOctetStringValueContext)

	// ExitRuleRealValue is called when exiting the ruleRealValue production.
	ExitRuleRealValue(c *RuleRealValueContext)

	// ExitRuleRelativeIRIValue is called when exiting the ruleRelativeIRIValue production.
	ExitRuleRelativeIRIValue(c *RuleRelativeIRIValueContext)

	// ExitRuleRelativeOIDValue is called when exiting the ruleRelativeOIDValue production.
	ExitRuleRelativeOIDValue(c *RuleRelativeOIDValueContext)

	// ExitRuleSequenceValue is called when exiting the ruleSequenceValue production.
	ExitRuleSequenceValue(c *RuleSequenceValueContext)

	// ExitRuleSequenceOfValue is called when exiting the ruleSequenceOfValue production.
	ExitRuleSequenceOfValue(c *RuleSequenceOfValueContext)

	// ExitRuleSetValue is called when exiting the ruleSetValue production.
	ExitRuleSetValue(c *RuleSetValueContext)

	// ExitRuleSetOfValue is called when exiting the ruleSetOfValue production.
	ExitRuleSetOfValue(c *RuleSetOfValueContext)

	// ExitRulePrefixedValue is called when exiting the rulePrefixedValue production.
	ExitRulePrefixedValue(c *RulePrefixedValueContext)

	// ExitRuleTimeValue is called when exiting the ruleTimeValue production.
	ExitRuleTimeValue(c *RuleTimeValueContext)

	// ExitRuleBuiltinValue is called when exiting the ruleBuiltinValue production.
	ExitRuleBuiltinValue(c *RuleBuiltinValueContext)

	// ExitRuleValueFromObject is called when exiting the ruleValueFromObject production.
	ExitRuleValueFromObject(c *RuleValueFromObjectContext)

	// ExitRuleReferencedValue is called when exiting the ruleReferencedValue production.
	ExitRuleReferencedValue(c *RuleReferencedValueContext)

	// ExitRuleOpenTypeFieldVal is called when exiting the ruleOpenTypeFieldVal production.
	ExitRuleOpenTypeFieldVal(c *RuleOpenTypeFieldValContext)

	// ExitRuleFixedTypeFieldVal is called when exiting the ruleFixedTypeFieldVal production.
	ExitRuleFixedTypeFieldVal(c *RuleFixedTypeFieldValContext)

	// ExitRuleObjectClassFieldValue is called when exiting the ruleObjectClassFieldValue production.
	ExitRuleObjectClassFieldValue(c *RuleObjectClassFieldValueContext)

	// ExitRuleValue is called when exiting the ruleValue production.
	ExitRuleValue(c *RuleValueContext)

	// ExitRuleValueSet is called when exiting the ruleValueSet production.
	ExitRuleValueSet(c *RuleValueSetContext)

	// ExitRuleDefinedObjectClass is called when exiting the ruleDefinedObjectClass production.
	ExitRuleDefinedObjectClass(c *RuleDefinedObjectClassContext)

	// ExitRuleObject is called when exiting the ruleObject production.
	ExitRuleObject(c *RuleObjectContext)

	// ExitRuleDefinedObject is called when exiting the ruleDefinedObject production.
	ExitRuleDefinedObject(c *RuleDefinedObjectContext)

	// ExitRuleDefinedObjectSet is called when exiting the ruleDefinedObjectSet production.
	ExitRuleDefinedObjectSet(c *RuleDefinedObjectSetContext)

	// ExitRuleObjectDefn is called when exiting the ruleObjectDefn production.
	ExitRuleObjectDefn(c *RuleObjectDefnContext)

	// ExitRuleDefaultSyntax is called when exiting the ruleDefaultSyntax production.
	ExitRuleDefaultSyntax(c *RuleDefaultSyntaxContext)

	// ExitRuleDefinedSyntax is called when exiting the ruleDefinedSyntax production.
	ExitRuleDefinedSyntax(c *RuleDefinedSyntaxContext)

	// ExitRuleDefinedSyntaxTokenList is called when exiting the ruleDefinedSyntaxTokenList production.
	ExitRuleDefinedSyntaxTokenList(c *RuleDefinedSyntaxTokenListContext)

	// ExitRuleDefinedSyntaxToken is called when exiting the ruleDefinedSyntaxToken production.
	ExitRuleDefinedSyntaxToken(c *RuleDefinedSyntaxTokenContext)

	// ExitRuleLiteral is called when exiting the ruleLiteral production.
	ExitRuleLiteral(c *RuleLiteralContext)

	// ExitRuleWord is called when exiting the ruleWord production.
	ExitRuleWord(c *RuleWordContext)

	// ExitRuleFieldSettingList is called when exiting the ruleFieldSettingList production.
	ExitRuleFieldSettingList(c *RuleFieldSettingListContext)

	// ExitRuleFieldSetting is called when exiting the ruleFieldSetting production.
	ExitRuleFieldSetting(c *RuleFieldSettingContext)

	// ExitRulePrimitiveFieldName is called when exiting the rulePrimitiveFieldName production.
	ExitRulePrimitiveFieldName(c *RulePrimitiveFieldNameContext)

	// ExitRuleSetting is called when exiting the ruleSetting production.
	ExitRuleSetting(c *RuleSettingContext)

	// ExitRuleObjectFromObject is called when exiting the ruleObjectFromObject production.
	ExitRuleObjectFromObject(c *RuleObjectFromObjectContext)

	// ExitRuleObjectSetFromObjects is called when exiting the ruleObjectSetFromObjects production.
	ExitRuleObjectSetFromObjects(c *RuleObjectSetFromObjectsContext)

	// ExitRuleReferencedObjects is called when exiting the ruleReferencedObjects production.
	ExitRuleReferencedObjects(c *RuleReferencedObjectsContext)

	// ExitRuleFieldName is called when exiting the ruleFieldName production.
	ExitRuleFieldName(c *RuleFieldNameContext)

	// ExitRuleParameterizedObject is called when exiting the ruleParameterizedObject production.
	ExitRuleParameterizedObject(c *RuleParameterizedObjectContext)

	// ExitRuleObjectClass is called when exiting the ruleObjectClass production.
	ExitRuleObjectClass(c *RuleObjectClassContext)

	// ExitRuleObjectClassDefn is called when exiting the ruleObjectClassDefn production.
	ExitRuleObjectClassDefn(c *RuleObjectClassDefnContext)

	// ExitRuleFieldSpecList is called when exiting the ruleFieldSpecList production.
	ExitRuleFieldSpecList(c *RuleFieldSpecListContext)

	// ExitRuleFieldSpec is called when exiting the ruleFieldSpec production.
	ExitRuleFieldSpec(c *RuleFieldSpecContext)

	// ExitRuleTypeFieldSpec is called when exiting the ruleTypeFieldSpec production.
	ExitRuleTypeFieldSpec(c *RuleTypeFieldSpecContext)

	// ExitRuleFixedTypeValueFieldSpec is called when exiting the ruleFixedTypeValueFieldSpec production.
	ExitRuleFixedTypeValueFieldSpec(c *RuleFixedTypeValueFieldSpecContext)

	// ExitRuleUnique is called when exiting the ruleUnique production.
	ExitRuleUnique(c *RuleUniqueContext)

	// ExitRuleVariableTypeValueFieldSpec is called when exiting the ruleVariableTypeValueFieldSpec production.
	ExitRuleVariableTypeValueFieldSpec(c *RuleVariableTypeValueFieldSpecContext)

	// ExitRuleFixedTypeValueSetFieldSpec is called when exiting the ruleFixedTypeValueSetFieldSpec production.
	ExitRuleFixedTypeValueSetFieldSpec(c *RuleFixedTypeValueSetFieldSpecContext)

	// ExitRuleVariableTypeValueSetFieldSpec is called when exiting the ruleVariableTypeValueSetFieldSpec production.
	ExitRuleVariableTypeValueSetFieldSpec(c *RuleVariableTypeValueSetFieldSpecContext)

	// ExitRuleObjectFieldSpec is called when exiting the ruleObjectFieldSpec production.
	ExitRuleObjectFieldSpec(c *RuleObjectFieldSpecContext)

	// ExitRuleObjectSetFieldSpec is called when exiting the ruleObjectSetFieldSpec production.
	ExitRuleObjectSetFieldSpec(c *RuleObjectSetFieldSpecContext)

	// ExitRuleTypeOptionalitySpec is called when exiting the ruleTypeOptionalitySpec production.
	ExitRuleTypeOptionalitySpec(c *RuleTypeOptionalitySpecContext)

	// ExitRuleValueOptionalitySpec is called when exiting the ruleValueOptionalitySpec production.
	ExitRuleValueOptionalitySpec(c *RuleValueOptionalitySpecContext)

	// ExitRuleValueSetOptionalitySpec is called when exiting the ruleValueSetOptionalitySpec production.
	ExitRuleValueSetOptionalitySpec(c *RuleValueSetOptionalitySpecContext)

	// ExitRuleObjectOptionalitySpec is called when exiting the ruleObjectOptionalitySpec production.
	ExitRuleObjectOptionalitySpec(c *RuleObjectOptionalitySpecContext)

	// ExitRuleObjectSetOptionalitySpec is called when exiting the ruleObjectSetOptionalitySpec production.
	ExitRuleObjectSetOptionalitySpec(c *RuleObjectSetOptionalitySpecContext)

	// ExitRuleWithSyntaxSpec is called when exiting the ruleWithSyntaxSpec production.
	ExitRuleWithSyntaxSpec(c *RuleWithSyntaxSpecContext)

	// ExitRuleSyntaxList is called when exiting the ruleSyntaxList production.
	ExitRuleSyntaxList(c *RuleSyntaxListContext)

	// ExitRuleTokenOrGroupSpecList is called when exiting the ruleTokenOrGroupSpecList production.
	ExitRuleTokenOrGroupSpecList(c *RuleTokenOrGroupSpecListContext)

	// ExitRuleTokenOrGroupSpec is called when exiting the ruleTokenOrGroupSpec production.
	ExitRuleTokenOrGroupSpec(c *RuleTokenOrGroupSpecContext)

	// ExitRuleRequiredToken is called when exiting the ruleRequiredToken production.
	ExitRuleRequiredToken(c *RuleRequiredTokenContext)

	// ExitRuleOptionalGroup is called when exiting the ruleOptionalGroup production.
	ExitRuleOptionalGroup(c *RuleOptionalGroupContext)

	// ExitRuleParameterizedObjectClass is called when exiting the ruleParameterizedObjectClass production.
	ExitRuleParameterizedObjectClass(c *RuleParameterizedObjectClassContext)

	// ExitRuleObjectSet is called when exiting the ruleObjectSet production.
	ExitRuleObjectSet(c *RuleObjectSetContext)

	// ExitRuleObjectSetSpec is called when exiting the ruleObjectSetSpec production.
	ExitRuleObjectSetSpec(c *RuleObjectSetSpecContext)

	// ExitRuleRootElementSetSpec is called when exiting the ruleRootElementSetSpec production.
	ExitRuleRootElementSetSpec(c *RuleRootElementSetSpecContext)

	// ExitRuleAdditionalElementSetSpec is called when exiting the ruleAdditionalElementSetSpec production.
	ExitRuleAdditionalElementSetSpec(c *RuleAdditionalElementSetSpecContext)

	// ExitRuleElementSetSpecs is called when exiting the ruleElementSetSpecs production.
	ExitRuleElementSetSpecs(c *RuleElementSetSpecsContext)

	// ExitRuleElementSetSpec is called when exiting the ruleElementSetSpec production.
	ExitRuleElementSetSpec(c *RuleElementSetSpecContext)

	// ExitRuleUnions is called when exiting the ruleUnions production.
	ExitRuleUnions(c *RuleUnionsContext)

	// ExitRuleExclusions is called when exiting the ruleExclusions production.
	ExitRuleExclusions(c *RuleExclusionsContext)

	// ExitRuleIntersections is called when exiting the ruleIntersections production.
	ExitRuleIntersections(c *RuleIntersectionsContext)

	// ExitRuleUElems is called when exiting the ruleUElems production.
	ExitRuleUElems(c *RuleUElemsContext)

	// ExitRuleUnionMark is called when exiting the ruleUnionMark production.
	ExitRuleUnionMark(c *RuleUnionMarkContext)

	// ExitRuleLowerEndValue is called when exiting the ruleLowerEndValue production.
	ExitRuleLowerEndValue(c *RuleLowerEndValueContext)

	// ExitRuleUpperEndValue is called when exiting the ruleUpperEndValue production.
	ExitRuleUpperEndValue(c *RuleUpperEndValueContext)

	// ExitRuleIncludes is called when exiting the ruleIncludes production.
	ExitRuleIncludes(c *RuleIncludesContext)

	// ExitRuleLowerEndpoint is called when exiting the ruleLowerEndpoint production.
	ExitRuleLowerEndpoint(c *RuleLowerEndpointContext)

	// ExitRuleUpperEndpoint is called when exiting the ruleUpperEndpoint production.
	ExitRuleUpperEndpoint(c *RuleUpperEndpointContext)

	// ExitRuleLevel is called when exiting the ruleLevel production.
	ExitRuleLevel(c *RuleLevelContext)

	// ExitRuleComponentIdList is called when exiting the ruleComponentIdList production.
	ExitRuleComponentIdList(c *RuleComponentIdListContext)

	// ExitRuleAtNotation is called when exiting the ruleAtNotation production.
	ExitRuleAtNotation(c *RuleAtNotationContext)

	// ExitRuleAtNotationList is called when exiting the ruleAtNotationList production.
	ExitRuleAtNotationList(c *RuleAtNotationListContext)

	// ExitRuleUserDefinedConstraintParameter is called when exiting the ruleUserDefinedConstraintParameter production.
	ExitRuleUserDefinedConstraintParameter(c *RuleUserDefinedConstraintParameterContext)

	// ExitRuleUserDefinedConstraintParameterList is called when exiting the ruleUserDefinedConstraintParameterList production.
	ExitRuleUserDefinedConstraintParameterList(c *RuleUserDefinedConstraintParameterListContext)

	// ExitRuleUserDefinedConstraint is called when exiting the ruleUserDefinedConstraint production.
	ExitRuleUserDefinedConstraint(c *RuleUserDefinedConstraintContext)

	// ExitRuleSimpleTableConstraint is called when exiting the ruleSimpleTableConstraint production.
	ExitRuleSimpleTableConstraint(c *RuleSimpleTableConstraintContext)

	// ExitRuleComponentRelationConstraint is called when exiting the ruleComponentRelationConstraint production.
	ExitRuleComponentRelationConstraint(c *RuleComponentRelationConstraintContext)

	// ExitRuleTableConstraint is called when exiting the ruleTableConstraint production.
	ExitRuleTableConstraint(c *RuleTableConstraintContext)

	// ExitRuleContentsConstraint is called when exiting the ruleContentsConstraint production.
	ExitRuleContentsConstraint(c *RuleContentsConstraintContext)

	// ExitRuleSubtypeConstraint is called when exiting the ruleSubtypeConstraint production.
	ExitRuleSubtypeConstraint(c *RuleSubtypeConstraintContext)

	// ExitRuleGeneralConstraint is called when exiting the ruleGeneralConstraint production.
	ExitRuleGeneralConstraint(c *RuleGeneralConstraintContext)

	// ExitRuleConstraintSpec is called when exiting the ruleConstraintSpec production.
	ExitRuleConstraintSpec(c *RuleConstraintSpecContext)

	// ExitRuleSignedNumber is called when exiting the ruleSignedNumber production.
	ExitRuleSignedNumber(c *RuleSignedNumberContext)

	// ExitRuleExceptionIdentification is called when exiting the ruleExceptionIdentification production.
	ExitRuleExceptionIdentification(c *RuleExceptionIdentificationContext)

	// ExitRuleExceptionSpec is called when exiting the ruleExceptionSpec production.
	ExitRuleExceptionSpec(c *RuleExceptionSpecContext)

	// ExitRuleConstraint is called when exiting the ruleConstraint production.
	ExitRuleConstraint(c *RuleConstraintContext)

	// ExitRuleSingleTypeConstraint is called when exiting the ruleSingleTypeConstraint production.
	ExitRuleSingleTypeConstraint(c *RuleSingleTypeConstraintContext)

	// ExitRuleValueConstraint is called when exiting the ruleValueConstraint production.
	ExitRuleValueConstraint(c *RuleValueConstraintContext)

	// ExitRulePresenceConstraint is called when exiting the rulePresenceConstraint production.
	ExitRulePresenceConstraint(c *RulePresenceConstraintContext)

	// ExitRuleComponentConstraint is called when exiting the ruleComponentConstraint production.
	ExitRuleComponentConstraint(c *RuleComponentConstraintContext)

	// ExitRuleNamedConstraint is called when exiting the ruleNamedConstraint production.
	ExitRuleNamedConstraint(c *RuleNamedConstraintContext)

	// ExitRuleTypeConstraints is called when exiting the ruleTypeConstraints production.
	ExitRuleTypeConstraints(c *RuleTypeConstraintsContext)

	// ExitRuleFullSpecification is called when exiting the ruleFullSpecification production.
	ExitRuleFullSpecification(c *RuleFullSpecificationContext)

	// ExitRulePartialSpecification is called when exiting the rulePartialSpecification production.
	ExitRulePartialSpecification(c *RulePartialSpecificationContext)

	// ExitRuleMultipleTypeConstraints is called when exiting the ruleMultipleTypeConstraints production.
	ExitRuleMultipleTypeConstraints(c *RuleMultipleTypeConstraintsContext)

	// ExitRuleSimpleString is called when exiting the ruleSimpleString production.
	ExitRuleSimpleString(c *RuleSimpleStringContext)

	// ExitRuleSingleValue is called when exiting the ruleSingleValue production.
	ExitRuleSingleValue(c *RuleSingleValueContext)

	// ExitRuleContainedSubtype is called when exiting the ruleContainedSubtype production.
	ExitRuleContainedSubtype(c *RuleContainedSubtypeContext)

	// ExitRuleValueRange is called when exiting the ruleValueRange production.
	ExitRuleValueRange(c *RuleValueRangeContext)

	// ExitRulePermittedAlphabet is called when exiting the rulePermittedAlphabet production.
	ExitRulePermittedAlphabet(c *RulePermittedAlphabetContext)

	// ExitRuleSizeConstraint is called when exiting the ruleSizeConstraint production.
	ExitRuleSizeConstraint(c *RuleSizeConstraintContext)

	// ExitRuleInnerTypeConstraints is called when exiting the ruleInnerTypeConstraints production.
	ExitRuleInnerTypeConstraints(c *RuleInnerTypeConstraintsContext)

	// ExitRulePatternConstraint is called when exiting the rulePatternConstraint production.
	ExitRulePatternConstraint(c *RulePatternConstraintContext)

	// ExitRulePropertySettings is called when exiting the rulePropertySettings production.
	ExitRulePropertySettings(c *RulePropertySettingsContext)

	// ExitRuleDurationRange is called when exiting the ruleDurationRange production.
	ExitRuleDurationRange(c *RuleDurationRangeContext)

	// ExitRuleTimePointRange is called when exiting the ruleTimePointRange production.
	ExitRuleTimePointRange(c *RuleTimePointRangeContext)

	// ExitRuleRecurrenceRange is called when exiting the ruleRecurrenceRange production.
	ExitRuleRecurrenceRange(c *RuleRecurrenceRangeContext)

	// ExitRuleSubtypeElements is called when exiting the ruleSubtypeElements production.
	ExitRuleSubtypeElements(c *RuleSubtypeElementsContext)

	// ExitRuleObjectSetElements is called when exiting the ruleObjectSetElements production.
	ExitRuleObjectSetElements(c *RuleObjectSetElementsContext)

	// ExitRuleElements is called when exiting the ruleElements production.
	ExitRuleElements(c *RuleElementsContext)

	// ExitRuleElems is called when exiting the ruleElems production.
	ExitRuleElems(c *RuleElemsContext)

	// ExitRuleIntersectionElements is called when exiting the ruleIntersectionElements production.
	ExitRuleIntersectionElements(c *RuleIntersectionElementsContext)

	// ExitRuleIElems is called when exiting the ruleIElems production.
	ExitRuleIElems(c *RuleIElemsContext)

	// ExitRuleIntersectionMark is called when exiting the ruleIntersectionMark production.
	ExitRuleIntersectionMark(c *RuleIntersectionMarkContext)

	// ExitRuleSimpleDefinedValue is called when exiting the ruleSimpleDefinedValue production.
	ExitRuleSimpleDefinedValue(c *RuleSimpleDefinedValueContext)

	// ExitRuleActualParameterList is called when exiting the ruleActualParameterList production.
	ExitRuleActualParameterList(c *RuleActualParameterListContext)

	// ExitRuleActualParameters is called when exiting the ruleActualParameters production.
	ExitRuleActualParameters(c *RuleActualParametersContext)

	// ExitRuleActualParameter is called when exiting the ruleActualParameter production.
	ExitRuleActualParameter(c *RuleActualParameterContext)

	// ExitRuleParameterizedObjectSet is called when exiting the ruleParameterizedObjectSet production.
	ExitRuleParameterizedObjectSet(c *RuleParameterizedObjectSetContext)

	// ExitRuleParameterizedValue is called when exiting the ruleParameterizedValue production.
	ExitRuleParameterizedValue(c *RuleParameterizedValueContext)

	// ExitRuleParameterizedTypeAssignment is called when exiting the ruleParameterizedTypeAssignment production.
	ExitRuleParameterizedTypeAssignment(c *RuleParameterizedTypeAssignmentContext)

	// ExitRuleParameterizedValueAssignment is called when exiting the ruleParameterizedValueAssignment production.
	ExitRuleParameterizedValueAssignment(c *RuleParameterizedValueAssignmentContext)

	// ExitRuleParameterizedValueSetTypeAssignment is called when exiting the ruleParameterizedValueSetTypeAssignment production.
	ExitRuleParameterizedValueSetTypeAssignment(c *RuleParameterizedValueSetTypeAssignmentContext)

	// ExitRuleParameterizedObjectClassAssignment is called when exiting the ruleParameterizedObjectClassAssignment production.
	ExitRuleParameterizedObjectClassAssignment(c *RuleParameterizedObjectClassAssignmentContext)

	// ExitRuleParameterizedObjectAssignment is called when exiting the ruleParameterizedObjectAssignment production.
	ExitRuleParameterizedObjectAssignment(c *RuleParameterizedObjectAssignmentContext)

	// ExitRuleParameterizedObjectSetAssignment is called when exiting the ruleParameterizedObjectSetAssignment production.
	ExitRuleParameterizedObjectSetAssignment(c *RuleParameterizedObjectSetAssignmentContext)

	// ExitRuleParameterList is called when exiting the ruleParameterList production.
	ExitRuleParameterList(c *RuleParameterListContext)

	// ExitRuleParameters is called when exiting the ruleParameters production.
	ExitRuleParameters(c *RuleParametersContext)

	// ExitRuleParameter is called when exiting the ruleParameter production.
	ExitRuleParameter(c *RuleParameterContext)

	// ExitRuleParamGovernor is called when exiting the ruleParamGovernor production.
	ExitRuleParamGovernor(c *RuleParamGovernorContext)

	// ExitRuleDummyReference is called when exiting the ruleDummyReference production.
	ExitRuleDummyReference(c *RuleDummyReferenceContext)

	// ExitRuleGovernor is called when exiting the ruleGovernor production.
	ExitRuleGovernor(c *RuleGovernorContext)

	// ExitRuleDummyGovernor is called when exiting the ruleDummyGovernor production.
	ExitRuleDummyGovernor(c *RuleDummyGovernorContext)

	// ExitRuleEncodingControlSections is called when exiting the ruleEncodingControlSections production.
	ExitRuleEncodingControlSections(c *RuleEncodingControlSectionsContext)
}
