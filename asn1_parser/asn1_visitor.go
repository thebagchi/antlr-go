// Code generated from asn1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package asn1_parser // asn1
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by asn1Parser.
type asn1Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by asn1Parser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleModules.
	VisitRuleModules(ctx *RuleModulesContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleModuleDefinition.
	VisitRuleModuleDefinition(ctx *RuleModuleDefinitionContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleModuleReference.
	VisitRuleModuleReference(ctx *RuleModuleReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleModuleIdentifier.
	VisitRuleModuleIdentifier(ctx *RuleModuleIdentifierContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefinitiveIdentification.
	VisitRuleDefinitiveIdentification(ctx *RuleDefinitiveIdentificationContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefinitiveOID.
	VisitRuleDefinitiveOID(ctx *RuleDefinitiveOIDContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefinitiveObjIdComponentList.
	VisitRuleDefinitiveObjIdComponentList(ctx *RuleDefinitiveObjIdComponentListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefinitiveObjIdComponent.
	VisitRuleDefinitiveObjIdComponent(ctx *RuleDefinitiveObjIdComponentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleNameForm.
	VisitRuleNameForm(ctx *RuleNameFormContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefinitiveNumberForm.
	VisitRuleDefinitiveNumberForm(ctx *RuleDefinitiveNumberFormContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefinitiveNameAndNumberForm.
	VisitRuleDefinitiveNameAndNumberForm(ctx *RuleDefinitiveNameAndNumberFormContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefinitiveOIDAndIRI.
	VisitRuleDefinitiveOIDAndIRI(ctx *RuleDefinitiveOIDAndIRIContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleFirstArcIdentifier.
	VisitRuleFirstArcIdentifier(ctx *RuleFirstArcIdentifierContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSubsequentArcIdentifier.
	VisitRuleSubsequentArcIdentifier(ctx *RuleSubsequentArcIdentifierContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleIRIValue.
	VisitRuleIRIValue(ctx *RuleIRIValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleEncodingReferenceDefault.
	VisitRuleEncodingReferenceDefault(ctx *RuleEncodingReferenceDefaultContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTagDefault.
	VisitRuleTagDefault(ctx *RuleTagDefaultContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExtensionDefault.
	VisitRuleExtensionDefault(ctx *RuleExtensionDefaultContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleModuleBody.
	VisitRuleModuleBody(ctx *RuleModuleBodyContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExports.
	VisitRuleExports(ctx *RuleExportsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSymbolsExported.
	VisitRuleSymbolsExported(ctx *RuleSymbolsExportedContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSymbolList.
	VisitRuleSymbolList(ctx *RuleSymbolListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSymbol.
	VisitRuleSymbol(ctx *RuleSymbolContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleReference.
	VisitRuleReference(ctx *RuleReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleIdentifier.
	VisitRuleIdentifier(ctx *RuleIdentifierContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTypeReference.
	VisitRuleTypeReference(ctx *RuleTypeReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleValueReference.
	VisitRuleValueReference(ctx *RuleValueReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectReference.
	VisitRuleObjectReference(ctx *RuleObjectReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectClassReference.
	VisitRuleObjectClassReference(ctx *RuleObjectClassReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectSetReference.
	VisitRuleObjectSetReference(ctx *RuleObjectSetReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameterizedReference.
	VisitRuleParameterizedReference(ctx *RuleParameterizedReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExternalTypeReference.
	VisitRuleExternalTypeReference(ctx *RuleExternalTypeReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExternalValueReference.
	VisitRuleExternalValueReference(ctx *RuleExternalValueReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExternalObjectClassReference.
	VisitRuleExternalObjectClassReference(ctx *RuleExternalObjectClassReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExternalObjectReference.
	VisitRuleExternalObjectReference(ctx *RuleExternalObjectReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExternalObjectSetReference.
	VisitRuleExternalObjectSetReference(ctx *RuleExternalObjectSetReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTypeFieldReference.
	VisitRuleTypeFieldReference(ctx *RuleTypeFieldReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleValueFieldReference.
	VisitRuleValueFieldReference(ctx *RuleValueFieldReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleValueSetFieldReference.
	VisitRuleValueSetFieldReference(ctx *RuleValueSetFieldReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectFieldReference.
	VisitRuleObjectFieldReference(ctx *RuleObjectFieldReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectSetFieldReference.
	VisitRuleObjectSetFieldReference(ctx *RuleObjectSetFieldReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleUsefulObjectClassReference.
	VisitRuleUsefulObjectClassReference(ctx *RuleUsefulObjectClassReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleImports.
	VisitRuleImports(ctx *RuleImportsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSymbolsImported.
	VisitRuleSymbolsImported(ctx *RuleSymbolsImportedContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSymbolsFromModuleList.
	VisitRuleSymbolsFromModuleList(ctx *RuleSymbolsFromModuleListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSymbolsFromModule.
	VisitRuleSymbolsFromModule(ctx *RuleSymbolsFromModuleContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleGlobalModuleReference.
	VisitRuleGlobalModuleReference(ctx *RuleGlobalModuleReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleAssignedIdentifier.
	VisitRuleAssignedIdentifier(ctx *RuleAssignedIdentifierContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleAssignmentList.
	VisitRuleAssignmentList(ctx *RuleAssignmentListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleAssignment.
	VisitRuleAssignment(ctx *RuleAssignmentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTypeAssignment.
	VisitRuleTypeAssignment(ctx *RuleTypeAssignmentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleValueAssignment.
	VisitRuleValueAssignment(ctx *RuleValueAssignmentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleXMLValueAssignment.
	VisitRuleXMLValueAssignment(ctx *RuleXMLValueAssignmentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleValueSetTypeAssignment.
	VisitRuleValueSetTypeAssignment(ctx *RuleValueSetTypeAssignmentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectAssignment.
	VisitRuleObjectAssignment(ctx *RuleObjectAssignmentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectClassAssignment.
	VisitRuleObjectClassAssignment(ctx *RuleObjectClassAssignmentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectSetAssignment.
	VisitRuleObjectSetAssignment(ctx *RuleObjectSetAssignmentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameterizedAssignment.
	VisitRuleParameterizedAssignment(ctx *RuleParameterizedAssignmentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectIdentifierValue.
	VisitRuleObjectIdentifierValue(ctx *RuleObjectIdentifierValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjIdComponentsList.
	VisitRuleObjIdComponentsList(ctx *RuleObjIdComponentsListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjIdComponents.
	VisitRuleObjIdComponents(ctx *RuleObjIdComponentsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleNumberForm.
	VisitRuleNumberForm(ctx *RuleNumberFormContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleNameAndNumberForm.
	VisitRuleNameAndNumberForm(ctx *RuleNameAndNumberFormContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefinedValue.
	VisitRuleDefinedValue(ctx *RuleDefinedValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleNamedBit.
	VisitRuleNamedBit(ctx *RuleNamedBitContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleNamedBitList.
	VisitRuleNamedBitList(ctx *RuleNamedBitListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRestrictedCharacterStringType.
	VisitRuleRestrictedCharacterStringType(ctx *RuleRestrictedCharacterStringTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleUnrestrictedCharacterStringType.
	VisitRuleUnrestrictedCharacterStringType(ctx *RuleUnrestrictedCharacterStringTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleAlternativeTypeList.
	VisitRuleAlternativeTypeList(ctx *RuleAlternativeTypeListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRootAlternativeTypeList.
	VisitRuleRootAlternativeTypeList(ctx *RuleRootAlternativeTypeListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleVersionNumber.
	VisitRuleVersionNumber(ctx *RuleVersionNumberContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExtensionAdditionAlternativesGroup.
	VisitRuleExtensionAdditionAlternativesGroup(ctx *RuleExtensionAdditionAlternativesGroupContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExtensionAdditionAlternative.
	VisitRuleExtensionAdditionAlternative(ctx *RuleExtensionAdditionAlternativeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExtensionAdditionAlternativesList.
	VisitRuleExtensionAdditionAlternativesList(ctx *RuleExtensionAdditionAlternativesListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExtensionAdditionAlternatives.
	VisitRuleExtensionAdditionAlternatives(ctx *RuleExtensionAdditionAlternativesContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleEnumerationItem.
	VisitRuleEnumerationItem(ctx *RuleEnumerationItemContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleEnumeration.
	VisitRuleEnumeration(ctx *RuleEnumerationContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRootEnumeration.
	VisitRuleRootEnumeration(ctx *RuleRootEnumerationContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleAdditionalEnumeration.
	VisitRuleAdditionalEnumeration(ctx *RuleAdditionalEnumerationContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleAlternativeTypeLists.
	VisitRuleAlternativeTypeLists(ctx *RuleAlternativeTypeListsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleEnumerations.
	VisitRuleEnumerations(ctx *RuleEnumerationsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleNamedNumber.
	VisitRuleNamedNumber(ctx *RuleNamedNumberContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleNamedNumberList.
	VisitRuleNamedNumberList(ctx *RuleNamedNumberListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleComponentType.
	VisitRuleComponentType(ctx *RuleComponentTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExtensionAdditionGroup.
	VisitRuleExtensionAdditionGroup(ctx *RuleExtensionAdditionGroupContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExtensionAddition.
	VisitRuleExtensionAddition(ctx *RuleExtensionAdditionContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleComponentTypeList.
	VisitRuleComponentTypeList(ctx *RuleComponentTypeListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExtensionAdditionList.
	VisitRuleExtensionAdditionList(ctx *RuleExtensionAdditionListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRootComponentTypeList.
	VisitRuleRootComponentTypeList(ctx *RuleRootComponentTypeListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExtensionAdditions.
	VisitRuleExtensionAdditions(ctx *RuleExtensionAdditionsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExtensionEndMarker.
	VisitRuleExtensionEndMarker(ctx *RuleExtensionEndMarkerContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleComponentTypeLists.
	VisitRuleComponentTypeLists(ctx *RuleComponentTypeListsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExtensionAndException.
	VisitRuleExtensionAndException(ctx *RuleExtensionAndExceptionContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleOptionalExtensionMarker.
	VisitRuleOptionalExtensionMarker(ctx *RuleOptionalExtensionMarkerContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleEncodingReference.
	VisitRuleEncodingReference(ctx *RuleEncodingReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleClass.
	VisitRuleClass(ctx *RuleClassContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleClassNumber.
	VisitRuleClassNumber(ctx *RuleClassNumberContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleEncodingInstruction.
	VisitRuleEncodingInstruction(ctx *RuleEncodingInstructionContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTag.
	VisitRuleTag(ctx *RuleTagContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleEncodingPrefix.
	VisitRuleEncodingPrefix(ctx *RuleEncodingPrefixContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTaggedType.
	VisitRuleTaggedType(ctx *RuleTaggedTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleEncodingPrefixedType.
	VisitRuleEncodingPrefixedType(ctx *RuleEncodingPrefixedTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleBitStringType.
	VisitRuleBitStringType(ctx *RuleBitStringTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleBooleanType.
	VisitRuleBooleanType(ctx *RuleBooleanTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleCharacterStringType.
	VisitRuleCharacterStringType(ctx *RuleCharacterStringTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleChoiceType.
	VisitRuleChoiceType(ctx *RuleChoiceTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDateType.
	VisitRuleDateType(ctx *RuleDateTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDateTimeType.
	VisitRuleDateTimeType(ctx *RuleDateTimeTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDurationType.
	VisitRuleDurationType(ctx *RuleDurationTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleEmbeddedPDVType.
	VisitRuleEmbeddedPDVType(ctx *RuleEmbeddedPDVTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleEnumeratedType.
	VisitRuleEnumeratedType(ctx *RuleEnumeratedTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExternalType.
	VisitRuleExternalType(ctx *RuleExternalTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleInstanceOfType.
	VisitRuleInstanceOfType(ctx *RuleInstanceOfTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleIntegerType.
	VisitRuleIntegerType(ctx *RuleIntegerTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleIRIType.
	VisitRuleIRIType(ctx *RuleIRITypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleNullType.
	VisitRuleNullType(ctx *RuleNullTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectClassFieldType.
	VisitRuleObjectClassFieldType(ctx *RuleObjectClassFieldTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectIdentifierType.
	VisitRuleObjectIdentifierType(ctx *RuleObjectIdentifierTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleOctetStringType.
	VisitRuleOctetStringType(ctx *RuleOctetStringTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRealType.
	VisitRuleRealType(ctx *RuleRealTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRelativeIRIType.
	VisitRuleRelativeIRIType(ctx *RuleRelativeIRITypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRelativeOIDType.
	VisitRuleRelativeOIDType(ctx *RuleRelativeOIDTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSequenceType.
	VisitRuleSequenceType(ctx *RuleSequenceTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSequenceOfType.
	VisitRuleSequenceOfType(ctx *RuleSequenceOfTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSetType.
	VisitRuleSetType(ctx *RuleSetTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSetOfType.
	VisitRuleSetOfType(ctx *RuleSetOfTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#rulePrefixedType.
	VisitRulePrefixedType(ctx *RulePrefixedTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTimeType.
	VisitRuleTimeType(ctx *RuleTimeTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTimeOfDayType.
	VisitRuleTimeOfDayType(ctx *RuleTimeOfDayTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleBuiltinType.
	VisitRuleBuiltinType(ctx *RuleBuiltinTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSimpleDefinedType.
	VisitRuleSimpleDefinedType(ctx *RuleSimpleDefinedTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameterizedType.
	VisitRuleParameterizedType(ctx *RuleParameterizedTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameterizedValueSetType.
	VisitRuleParameterizedValueSetType(ctx *RuleParameterizedValueSetTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefinedType.
	VisitRuleDefinedType(ctx *RuleDefinedTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleUsefulType.
	VisitRuleUsefulType(ctx *RuleUsefulTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSelectionType.
	VisitRuleSelectionType(ctx *RuleSelectionTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTypeFromObject.
	VisitRuleTypeFromObject(ctx *RuleTypeFromObjectContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleValueSetFromObjects.
	VisitRuleValueSetFromObjects(ctx *RuleValueSetFromObjectsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleReferencedType.
	VisitRuleReferencedType(ctx *RuleReferencedTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTypeForConstraint.
	VisitRuleTypeForConstraint(ctx *RuleTypeForConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleNamedType.
	VisitRuleNamedType(ctx *RuleNamedTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTypeWithConstraint.
	VisitRuleTypeWithConstraint(ctx *RuleTypeWithConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleConstrainedType.
	VisitRuleConstrainedType(ctx *RuleConstrainedTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleType.
	VisitRuleType(ctx *RuleTypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleIdentifierList.
	VisitRuleIdentifierList(ctx *RuleIdentifierListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleCharsDefn.
	VisitRuleCharsDefn(ctx *RuleCharsDefnContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleCharSyms.
	VisitRuleCharSyms(ctx *RuleCharSymsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleGroup.
	VisitRuleGroup(ctx *RuleGroupContext) interface{}

	// Visit a parse tree produced by asn1Parser#rulePlane.
	VisitRulePlane(ctx *RulePlaneContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRow.
	VisitRuleRow(ctx *RuleRowContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleCell.
	VisitRuleCell(ctx *RuleCellContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTableColumn.
	VisitRuleTableColumn(ctx *RuleTableColumnContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTableRow.
	VisitRuleTableRow(ctx *RuleTableRowContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleCharacterStringList.
	VisitRuleCharacterStringList(ctx *RuleCharacterStringListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleQuadruple.
	VisitRuleQuadruple(ctx *RuleQuadrupleContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTuple.
	VisitRuleTuple(ctx *RuleTupleContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRestrictedCharacterStringValue.
	VisitRuleRestrictedCharacterStringValue(ctx *RuleRestrictedCharacterStringValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleUnrestrictedCharacterStringValue.
	VisitRuleUnrestrictedCharacterStringValue(ctx *RuleUnrestrictedCharacterStringValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleNumericRealValue.
	VisitRuleNumericRealValue(ctx *RuleNumericRealValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSpecialRealValue.
	VisitRuleSpecialRealValue(ctx *RuleSpecialRealValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRelativeOIDComponents.
	VisitRuleRelativeOIDComponents(ctx *RuleRelativeOIDComponentsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRelativeOIDComponentsList.
	VisitRuleRelativeOIDComponentsList(ctx *RuleRelativeOIDComponentsListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleComponentValueList.
	VisitRuleComponentValueList(ctx *RuleComponentValueListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleValueList.
	VisitRuleValueList(ctx *RuleValueListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleNamedValue.
	VisitRuleNamedValue(ctx *RuleNamedValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleNamedValueList.
	VisitRuleNamedValueList(ctx *RuleNamedValueListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleBitStringValue.
	VisitRuleBitStringValue(ctx *RuleBitStringValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleBooleanValue.
	VisitRuleBooleanValue(ctx *RuleBooleanValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleCharacterStringValue.
	VisitRuleCharacterStringValue(ctx *RuleCharacterStringValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleChoiceValue.
	VisitRuleChoiceValue(ctx *RuleChoiceValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleEmbeddedPDVValue.
	VisitRuleEmbeddedPDVValue(ctx *RuleEmbeddedPDVValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleEnumeratedValue.
	VisitRuleEnumeratedValue(ctx *RuleEnumeratedValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExternalValue.
	VisitRuleExternalValue(ctx *RuleExternalValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleInstanceOfValue.
	VisitRuleInstanceOfValue(ctx *RuleInstanceOfValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleIntegerValue.
	VisitRuleIntegerValue(ctx *RuleIntegerValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleNullValue.
	VisitRuleNullValue(ctx *RuleNullValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleOctetStringValue.
	VisitRuleOctetStringValue(ctx *RuleOctetStringValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRealValue.
	VisitRuleRealValue(ctx *RuleRealValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRelativeIRIValue.
	VisitRuleRelativeIRIValue(ctx *RuleRelativeIRIValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRelativeOIDValue.
	VisitRuleRelativeOIDValue(ctx *RuleRelativeOIDValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSequenceValue.
	VisitRuleSequenceValue(ctx *RuleSequenceValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSequenceOfValue.
	VisitRuleSequenceOfValue(ctx *RuleSequenceOfValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSetValue.
	VisitRuleSetValue(ctx *RuleSetValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSetOfValue.
	VisitRuleSetOfValue(ctx *RuleSetOfValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#rulePrefixedValue.
	VisitRulePrefixedValue(ctx *RulePrefixedValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTimeValue.
	VisitRuleTimeValue(ctx *RuleTimeValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleBuiltinValue.
	VisitRuleBuiltinValue(ctx *RuleBuiltinValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleValueFromObject.
	VisitRuleValueFromObject(ctx *RuleValueFromObjectContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleReferencedValue.
	VisitRuleReferencedValue(ctx *RuleReferencedValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleOpenTypeFieldVal.
	VisitRuleOpenTypeFieldVal(ctx *RuleOpenTypeFieldValContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleFixedTypeFieldVal.
	VisitRuleFixedTypeFieldVal(ctx *RuleFixedTypeFieldValContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectClassFieldValue.
	VisitRuleObjectClassFieldValue(ctx *RuleObjectClassFieldValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleValue.
	VisitRuleValue(ctx *RuleValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleValueSet.
	VisitRuleValueSet(ctx *RuleValueSetContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefinedObjectClass.
	VisitRuleDefinedObjectClass(ctx *RuleDefinedObjectClassContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObject.
	VisitRuleObject(ctx *RuleObjectContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefinedObject.
	VisitRuleDefinedObject(ctx *RuleDefinedObjectContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefinedObjectSet.
	VisitRuleDefinedObjectSet(ctx *RuleDefinedObjectSetContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectDefn.
	VisitRuleObjectDefn(ctx *RuleObjectDefnContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefaultSyntax.
	VisitRuleDefaultSyntax(ctx *RuleDefaultSyntaxContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefinedSyntax.
	VisitRuleDefinedSyntax(ctx *RuleDefinedSyntaxContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefinedSyntaxTokenList.
	VisitRuleDefinedSyntaxTokenList(ctx *RuleDefinedSyntaxTokenListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDefinedSyntaxToken.
	VisitRuleDefinedSyntaxToken(ctx *RuleDefinedSyntaxTokenContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleLiteral.
	VisitRuleLiteral(ctx *RuleLiteralContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleWord.
	VisitRuleWord(ctx *RuleWordContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleFieldSettingList.
	VisitRuleFieldSettingList(ctx *RuleFieldSettingListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleFieldSetting.
	VisitRuleFieldSetting(ctx *RuleFieldSettingContext) interface{}

	// Visit a parse tree produced by asn1Parser#rulePrimitiveFieldName.
	VisitRulePrimitiveFieldName(ctx *RulePrimitiveFieldNameContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSetting.
	VisitRuleSetting(ctx *RuleSettingContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectFromObject.
	VisitRuleObjectFromObject(ctx *RuleObjectFromObjectContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectSetFromObjects.
	VisitRuleObjectSetFromObjects(ctx *RuleObjectSetFromObjectsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleReferencedObjects.
	VisitRuleReferencedObjects(ctx *RuleReferencedObjectsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleFieldName.
	VisitRuleFieldName(ctx *RuleFieldNameContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameterizedObject.
	VisitRuleParameterizedObject(ctx *RuleParameterizedObjectContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectClass.
	VisitRuleObjectClass(ctx *RuleObjectClassContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectClassDefn.
	VisitRuleObjectClassDefn(ctx *RuleObjectClassDefnContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleFieldSpecList.
	VisitRuleFieldSpecList(ctx *RuleFieldSpecListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleFieldSpec.
	VisitRuleFieldSpec(ctx *RuleFieldSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTypeFieldSpec.
	VisitRuleTypeFieldSpec(ctx *RuleTypeFieldSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleFixedTypeValueFieldSpec.
	VisitRuleFixedTypeValueFieldSpec(ctx *RuleFixedTypeValueFieldSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleUnique.
	VisitRuleUnique(ctx *RuleUniqueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleVariableTypeValueFieldSpec.
	VisitRuleVariableTypeValueFieldSpec(ctx *RuleVariableTypeValueFieldSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleFixedTypeValueSetFieldSpec.
	VisitRuleFixedTypeValueSetFieldSpec(ctx *RuleFixedTypeValueSetFieldSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleVariableTypeValueSetFieldSpec.
	VisitRuleVariableTypeValueSetFieldSpec(ctx *RuleVariableTypeValueSetFieldSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectFieldSpec.
	VisitRuleObjectFieldSpec(ctx *RuleObjectFieldSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectSetFieldSpec.
	VisitRuleObjectSetFieldSpec(ctx *RuleObjectSetFieldSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTypeOptionalitySpec.
	VisitRuleTypeOptionalitySpec(ctx *RuleTypeOptionalitySpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleValueOptionalitySpec.
	VisitRuleValueOptionalitySpec(ctx *RuleValueOptionalitySpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleValueSetOptionalitySpec.
	VisitRuleValueSetOptionalitySpec(ctx *RuleValueSetOptionalitySpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectOptionalitySpec.
	VisitRuleObjectOptionalitySpec(ctx *RuleObjectOptionalitySpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectSetOptionalitySpec.
	VisitRuleObjectSetOptionalitySpec(ctx *RuleObjectSetOptionalitySpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleWithSyntaxSpec.
	VisitRuleWithSyntaxSpec(ctx *RuleWithSyntaxSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSyntaxList.
	VisitRuleSyntaxList(ctx *RuleSyntaxListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTokenOrGroupSpecList.
	VisitRuleTokenOrGroupSpecList(ctx *RuleTokenOrGroupSpecListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTokenOrGroupSpec.
	VisitRuleTokenOrGroupSpec(ctx *RuleTokenOrGroupSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRequiredToken.
	VisitRuleRequiredToken(ctx *RuleRequiredTokenContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleOptionalGroup.
	VisitRuleOptionalGroup(ctx *RuleOptionalGroupContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameterizedObjectClass.
	VisitRuleParameterizedObjectClass(ctx *RuleParameterizedObjectClassContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectSet.
	VisitRuleObjectSet(ctx *RuleObjectSetContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectSetSpec.
	VisitRuleObjectSetSpec(ctx *RuleObjectSetSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRootElementSetSpec.
	VisitRuleRootElementSetSpec(ctx *RuleRootElementSetSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleAdditionalElementSetSpec.
	VisitRuleAdditionalElementSetSpec(ctx *RuleAdditionalElementSetSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleElementSetSpecs.
	VisitRuleElementSetSpecs(ctx *RuleElementSetSpecsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleElementSetSpec.
	VisitRuleElementSetSpec(ctx *RuleElementSetSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleUnions.
	VisitRuleUnions(ctx *RuleUnionsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExclusions.
	VisitRuleExclusions(ctx *RuleExclusionsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleIntersections.
	VisitRuleIntersections(ctx *RuleIntersectionsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleUElems.
	VisitRuleUElems(ctx *RuleUElemsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleUnionMark.
	VisitRuleUnionMark(ctx *RuleUnionMarkContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleLowerEndValue.
	VisitRuleLowerEndValue(ctx *RuleLowerEndValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleUpperEndValue.
	VisitRuleUpperEndValue(ctx *RuleUpperEndValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleIncludes.
	VisitRuleIncludes(ctx *RuleIncludesContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleLowerEndpoint.
	VisitRuleLowerEndpoint(ctx *RuleLowerEndpointContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleUpperEndpoint.
	VisitRuleUpperEndpoint(ctx *RuleUpperEndpointContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleLevel.
	VisitRuleLevel(ctx *RuleLevelContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleComponentIdList.
	VisitRuleComponentIdList(ctx *RuleComponentIdListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleAtNotation.
	VisitRuleAtNotation(ctx *RuleAtNotationContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleAtNotationList.
	VisitRuleAtNotationList(ctx *RuleAtNotationListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleUserDefinedConstraintParameter.
	VisitRuleUserDefinedConstraintParameter(ctx *RuleUserDefinedConstraintParameterContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleUserDefinedConstraintParameterList.
	VisitRuleUserDefinedConstraintParameterList(ctx *RuleUserDefinedConstraintParameterListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleUserDefinedConstraint.
	VisitRuleUserDefinedConstraint(ctx *RuleUserDefinedConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSimpleTableConstraint.
	VisitRuleSimpleTableConstraint(ctx *RuleSimpleTableConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleComponentRelationConstraint.
	VisitRuleComponentRelationConstraint(ctx *RuleComponentRelationConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTableConstraint.
	VisitRuleTableConstraint(ctx *RuleTableConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleContentsConstraint.
	VisitRuleContentsConstraint(ctx *RuleContentsConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSubtypeConstraint.
	VisitRuleSubtypeConstraint(ctx *RuleSubtypeConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleGeneralConstraint.
	VisitRuleGeneralConstraint(ctx *RuleGeneralConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleConstraintSpec.
	VisitRuleConstraintSpec(ctx *RuleConstraintSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSignedNumber.
	VisitRuleSignedNumber(ctx *RuleSignedNumberContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExceptionIdentification.
	VisitRuleExceptionIdentification(ctx *RuleExceptionIdentificationContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleExceptionSpec.
	VisitRuleExceptionSpec(ctx *RuleExceptionSpecContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleConstraint.
	VisitRuleConstraint(ctx *RuleConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSingleTypeConstraint.
	VisitRuleSingleTypeConstraint(ctx *RuleSingleTypeConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleValueConstraint.
	VisitRuleValueConstraint(ctx *RuleValueConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#rulePresenceConstraint.
	VisitRulePresenceConstraint(ctx *RulePresenceConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleComponentConstraint.
	VisitRuleComponentConstraint(ctx *RuleComponentConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleNamedConstraint.
	VisitRuleNamedConstraint(ctx *RuleNamedConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTypeConstraints.
	VisitRuleTypeConstraints(ctx *RuleTypeConstraintsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleFullSpecification.
	VisitRuleFullSpecification(ctx *RuleFullSpecificationContext) interface{}

	// Visit a parse tree produced by asn1Parser#rulePartialSpecification.
	VisitRulePartialSpecification(ctx *RulePartialSpecificationContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleMultipleTypeConstraints.
	VisitRuleMultipleTypeConstraints(ctx *RuleMultipleTypeConstraintsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSimpleString.
	VisitRuleSimpleString(ctx *RuleSimpleStringContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSingleValue.
	VisitRuleSingleValue(ctx *RuleSingleValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleContainedSubtype.
	VisitRuleContainedSubtype(ctx *RuleContainedSubtypeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleValueRange.
	VisitRuleValueRange(ctx *RuleValueRangeContext) interface{}

	// Visit a parse tree produced by asn1Parser#rulePermittedAlphabet.
	VisitRulePermittedAlphabet(ctx *RulePermittedAlphabetContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSizeConstraint.
	VisitRuleSizeConstraint(ctx *RuleSizeConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleInnerTypeConstraints.
	VisitRuleInnerTypeConstraints(ctx *RuleInnerTypeConstraintsContext) interface{}

	// Visit a parse tree produced by asn1Parser#rulePatternConstraint.
	VisitRulePatternConstraint(ctx *RulePatternConstraintContext) interface{}

	// Visit a parse tree produced by asn1Parser#rulePropertySettings.
	VisitRulePropertySettings(ctx *RulePropertySettingsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDurationRange.
	VisitRuleDurationRange(ctx *RuleDurationRangeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleTimePointRange.
	VisitRuleTimePointRange(ctx *RuleTimePointRangeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleRecurrenceRange.
	VisitRuleRecurrenceRange(ctx *RuleRecurrenceRangeContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSubtypeElements.
	VisitRuleSubtypeElements(ctx *RuleSubtypeElementsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleObjectSetElements.
	VisitRuleObjectSetElements(ctx *RuleObjectSetElementsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleElements.
	VisitRuleElements(ctx *RuleElementsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleElems.
	VisitRuleElems(ctx *RuleElemsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleIntersectionElements.
	VisitRuleIntersectionElements(ctx *RuleIntersectionElementsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleIElems.
	VisitRuleIElems(ctx *RuleIElemsContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleIntersectionMark.
	VisitRuleIntersectionMark(ctx *RuleIntersectionMarkContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleSimpleDefinedValue.
	VisitRuleSimpleDefinedValue(ctx *RuleSimpleDefinedValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleActualParameterList.
	VisitRuleActualParameterList(ctx *RuleActualParameterListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleActualParameters.
	VisitRuleActualParameters(ctx *RuleActualParametersContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleActualParameter.
	VisitRuleActualParameter(ctx *RuleActualParameterContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameterizedObjectSet.
	VisitRuleParameterizedObjectSet(ctx *RuleParameterizedObjectSetContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameterizedValue.
	VisitRuleParameterizedValue(ctx *RuleParameterizedValueContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameterizedTypeAssignment.
	VisitRuleParameterizedTypeAssignment(ctx *RuleParameterizedTypeAssignmentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameterizedValueAssignment.
	VisitRuleParameterizedValueAssignment(ctx *RuleParameterizedValueAssignmentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameterizedValueSetTypeAssignment.
	VisitRuleParameterizedValueSetTypeAssignment(ctx *RuleParameterizedValueSetTypeAssignmentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameterizedObjectClassAssignment.
	VisitRuleParameterizedObjectClassAssignment(ctx *RuleParameterizedObjectClassAssignmentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameterizedObjectAssignment.
	VisitRuleParameterizedObjectAssignment(ctx *RuleParameterizedObjectAssignmentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameterizedObjectSetAssignment.
	VisitRuleParameterizedObjectSetAssignment(ctx *RuleParameterizedObjectSetAssignmentContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameterList.
	VisitRuleParameterList(ctx *RuleParameterListContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameters.
	VisitRuleParameters(ctx *RuleParametersContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParameter.
	VisitRuleParameter(ctx *RuleParameterContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleParamGovernor.
	VisitRuleParamGovernor(ctx *RuleParamGovernorContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDummyReference.
	VisitRuleDummyReference(ctx *RuleDummyReferenceContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleGovernor.
	VisitRuleGovernor(ctx *RuleGovernorContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleDummyGovernor.
	VisitRuleDummyGovernor(ctx *RuleDummyGovernorContext) interface{}

	// Visit a parse tree produced by asn1Parser#ruleEncodingControlSections.
	VisitRuleEncodingControlSections(ctx *RuleEncodingControlSectionsContext) interface{}
}
