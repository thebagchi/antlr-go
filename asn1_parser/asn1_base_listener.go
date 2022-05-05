// Code generated from asn1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package asn1_parser // asn1
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Baseasn1Listener is a complete listener for a parse tree produced by asn1Parser.
type Baseasn1Listener struct{}

var _ asn1Listener = &Baseasn1Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Baseasn1Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Baseasn1Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Baseasn1Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Baseasn1Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *Baseasn1Listener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *Baseasn1Listener) ExitStart(ctx *StartContext) {}

// EnterRuleModules is called when production ruleModules is entered.
func (s *Baseasn1Listener) EnterRuleModules(ctx *RuleModulesContext) {}

// ExitRuleModules is called when production ruleModules is exited.
func (s *Baseasn1Listener) ExitRuleModules(ctx *RuleModulesContext) {}

// EnterRuleModuleDefinition is called when production ruleModuleDefinition is entered.
func (s *Baseasn1Listener) EnterRuleModuleDefinition(ctx *RuleModuleDefinitionContext) {}

// ExitRuleModuleDefinition is called when production ruleModuleDefinition is exited.
func (s *Baseasn1Listener) ExitRuleModuleDefinition(ctx *RuleModuleDefinitionContext) {}

// EnterRuleModuleReference is called when production ruleModuleReference is entered.
func (s *Baseasn1Listener) EnterRuleModuleReference(ctx *RuleModuleReferenceContext) {}

// ExitRuleModuleReference is called when production ruleModuleReference is exited.
func (s *Baseasn1Listener) ExitRuleModuleReference(ctx *RuleModuleReferenceContext) {}

// EnterRuleModuleIdentifier is called when production ruleModuleIdentifier is entered.
func (s *Baseasn1Listener) EnterRuleModuleIdentifier(ctx *RuleModuleIdentifierContext) {}

// ExitRuleModuleIdentifier is called when production ruleModuleIdentifier is exited.
func (s *Baseasn1Listener) ExitRuleModuleIdentifier(ctx *RuleModuleIdentifierContext) {}

// EnterRuleDefinitiveIdentification is called when production ruleDefinitiveIdentification is entered.
func (s *Baseasn1Listener) EnterRuleDefinitiveIdentification(ctx *RuleDefinitiveIdentificationContext) {
}

// ExitRuleDefinitiveIdentification is called when production ruleDefinitiveIdentification is exited.
func (s *Baseasn1Listener) ExitRuleDefinitiveIdentification(ctx *RuleDefinitiveIdentificationContext) {
}

// EnterRuleDefinitiveOID is called when production ruleDefinitiveOID is entered.
func (s *Baseasn1Listener) EnterRuleDefinitiveOID(ctx *RuleDefinitiveOIDContext) {}

// ExitRuleDefinitiveOID is called when production ruleDefinitiveOID is exited.
func (s *Baseasn1Listener) ExitRuleDefinitiveOID(ctx *RuleDefinitiveOIDContext) {}

// EnterRuleDefinitiveObjIdComponentList is called when production ruleDefinitiveObjIdComponentList is entered.
func (s *Baseasn1Listener) EnterRuleDefinitiveObjIdComponentList(ctx *RuleDefinitiveObjIdComponentListContext) {
}

// ExitRuleDefinitiveObjIdComponentList is called when production ruleDefinitiveObjIdComponentList is exited.
func (s *Baseasn1Listener) ExitRuleDefinitiveObjIdComponentList(ctx *RuleDefinitiveObjIdComponentListContext) {
}

// EnterRuleDefinitiveObjIdComponent is called when production ruleDefinitiveObjIdComponent is entered.
func (s *Baseasn1Listener) EnterRuleDefinitiveObjIdComponent(ctx *RuleDefinitiveObjIdComponentContext) {
}

// ExitRuleDefinitiveObjIdComponent is called when production ruleDefinitiveObjIdComponent is exited.
func (s *Baseasn1Listener) ExitRuleDefinitiveObjIdComponent(ctx *RuleDefinitiveObjIdComponentContext) {
}

// EnterRuleNameForm is called when production ruleNameForm is entered.
func (s *Baseasn1Listener) EnterRuleNameForm(ctx *RuleNameFormContext) {}

// ExitRuleNameForm is called when production ruleNameForm is exited.
func (s *Baseasn1Listener) ExitRuleNameForm(ctx *RuleNameFormContext) {}

// EnterRuleDefinitiveNumberForm is called when production ruleDefinitiveNumberForm is entered.
func (s *Baseasn1Listener) EnterRuleDefinitiveNumberForm(ctx *RuleDefinitiveNumberFormContext) {}

// ExitRuleDefinitiveNumberForm is called when production ruleDefinitiveNumberForm is exited.
func (s *Baseasn1Listener) ExitRuleDefinitiveNumberForm(ctx *RuleDefinitiveNumberFormContext) {}

// EnterRuleDefinitiveNameAndNumberForm is called when production ruleDefinitiveNameAndNumberForm is entered.
func (s *Baseasn1Listener) EnterRuleDefinitiveNameAndNumberForm(ctx *RuleDefinitiveNameAndNumberFormContext) {
}

// ExitRuleDefinitiveNameAndNumberForm is called when production ruleDefinitiveNameAndNumberForm is exited.
func (s *Baseasn1Listener) ExitRuleDefinitiveNameAndNumberForm(ctx *RuleDefinitiveNameAndNumberFormContext) {
}

// EnterRuleDefinitiveOIDAndIRI is called when production ruleDefinitiveOIDAndIRI is entered.
func (s *Baseasn1Listener) EnterRuleDefinitiveOIDAndIRI(ctx *RuleDefinitiveOIDAndIRIContext) {}

// ExitRuleDefinitiveOIDAndIRI is called when production ruleDefinitiveOIDAndIRI is exited.
func (s *Baseasn1Listener) ExitRuleDefinitiveOIDAndIRI(ctx *RuleDefinitiveOIDAndIRIContext) {}

// EnterRuleFirstArcIdentifier is called when production ruleFirstArcIdentifier is entered.
func (s *Baseasn1Listener) EnterRuleFirstArcIdentifier(ctx *RuleFirstArcIdentifierContext) {}

// ExitRuleFirstArcIdentifier is called when production ruleFirstArcIdentifier is exited.
func (s *Baseasn1Listener) ExitRuleFirstArcIdentifier(ctx *RuleFirstArcIdentifierContext) {}

// EnterRuleSubsequentArcIdentifier is called when production ruleSubsequentArcIdentifier is entered.
func (s *Baseasn1Listener) EnterRuleSubsequentArcIdentifier(ctx *RuleSubsequentArcIdentifierContext) {
}

// ExitRuleSubsequentArcIdentifier is called when production ruleSubsequentArcIdentifier is exited.
func (s *Baseasn1Listener) ExitRuleSubsequentArcIdentifier(ctx *RuleSubsequentArcIdentifierContext) {}

// EnterRuleIRIValue is called when production ruleIRIValue is entered.
func (s *Baseasn1Listener) EnterRuleIRIValue(ctx *RuleIRIValueContext) {}

// ExitRuleIRIValue is called when production ruleIRIValue is exited.
func (s *Baseasn1Listener) ExitRuleIRIValue(ctx *RuleIRIValueContext) {}

// EnterRuleEncodingReferenceDefault is called when production ruleEncodingReferenceDefault is entered.
func (s *Baseasn1Listener) EnterRuleEncodingReferenceDefault(ctx *RuleEncodingReferenceDefaultContext) {
}

// ExitRuleEncodingReferenceDefault is called when production ruleEncodingReferenceDefault is exited.
func (s *Baseasn1Listener) ExitRuleEncodingReferenceDefault(ctx *RuleEncodingReferenceDefaultContext) {
}

// EnterRuleTagDefault is called when production ruleTagDefault is entered.
func (s *Baseasn1Listener) EnterRuleTagDefault(ctx *RuleTagDefaultContext) {}

// ExitRuleTagDefault is called when production ruleTagDefault is exited.
func (s *Baseasn1Listener) ExitRuleTagDefault(ctx *RuleTagDefaultContext) {}

// EnterRuleExtensionDefault is called when production ruleExtensionDefault is entered.
func (s *Baseasn1Listener) EnterRuleExtensionDefault(ctx *RuleExtensionDefaultContext) {}

// ExitRuleExtensionDefault is called when production ruleExtensionDefault is exited.
func (s *Baseasn1Listener) ExitRuleExtensionDefault(ctx *RuleExtensionDefaultContext) {}

// EnterRuleModuleBody is called when production ruleModuleBody is entered.
func (s *Baseasn1Listener) EnterRuleModuleBody(ctx *RuleModuleBodyContext) {}

// ExitRuleModuleBody is called when production ruleModuleBody is exited.
func (s *Baseasn1Listener) ExitRuleModuleBody(ctx *RuleModuleBodyContext) {}

// EnterRuleExports is called when production ruleExports is entered.
func (s *Baseasn1Listener) EnterRuleExports(ctx *RuleExportsContext) {}

// ExitRuleExports is called when production ruleExports is exited.
func (s *Baseasn1Listener) ExitRuleExports(ctx *RuleExportsContext) {}

// EnterRuleSymbolsExported is called when production ruleSymbolsExported is entered.
func (s *Baseasn1Listener) EnterRuleSymbolsExported(ctx *RuleSymbolsExportedContext) {}

// ExitRuleSymbolsExported is called when production ruleSymbolsExported is exited.
func (s *Baseasn1Listener) ExitRuleSymbolsExported(ctx *RuleSymbolsExportedContext) {}

// EnterRuleSymbolList is called when production ruleSymbolList is entered.
func (s *Baseasn1Listener) EnterRuleSymbolList(ctx *RuleSymbolListContext) {}

// ExitRuleSymbolList is called when production ruleSymbolList is exited.
func (s *Baseasn1Listener) ExitRuleSymbolList(ctx *RuleSymbolListContext) {}

// EnterRuleSymbol is called when production ruleSymbol is entered.
func (s *Baseasn1Listener) EnterRuleSymbol(ctx *RuleSymbolContext) {}

// ExitRuleSymbol is called when production ruleSymbol is exited.
func (s *Baseasn1Listener) ExitRuleSymbol(ctx *RuleSymbolContext) {}

// EnterRuleReference is called when production ruleReference is entered.
func (s *Baseasn1Listener) EnterRuleReference(ctx *RuleReferenceContext) {}

// ExitRuleReference is called when production ruleReference is exited.
func (s *Baseasn1Listener) ExitRuleReference(ctx *RuleReferenceContext) {}

// EnterRuleIdentifier is called when production ruleIdentifier is entered.
func (s *Baseasn1Listener) EnterRuleIdentifier(ctx *RuleIdentifierContext) {}

// ExitRuleIdentifier is called when production ruleIdentifier is exited.
func (s *Baseasn1Listener) ExitRuleIdentifier(ctx *RuleIdentifierContext) {}

// EnterRuleTypeReference is called when production ruleTypeReference is entered.
func (s *Baseasn1Listener) EnterRuleTypeReference(ctx *RuleTypeReferenceContext) {}

// ExitRuleTypeReference is called when production ruleTypeReference is exited.
func (s *Baseasn1Listener) ExitRuleTypeReference(ctx *RuleTypeReferenceContext) {}

// EnterRuleValueReference is called when production ruleValueReference is entered.
func (s *Baseasn1Listener) EnterRuleValueReference(ctx *RuleValueReferenceContext) {}

// ExitRuleValueReference is called when production ruleValueReference is exited.
func (s *Baseasn1Listener) ExitRuleValueReference(ctx *RuleValueReferenceContext) {}

// EnterRuleObjectReference is called when production ruleObjectReference is entered.
func (s *Baseasn1Listener) EnterRuleObjectReference(ctx *RuleObjectReferenceContext) {}

// ExitRuleObjectReference is called when production ruleObjectReference is exited.
func (s *Baseasn1Listener) ExitRuleObjectReference(ctx *RuleObjectReferenceContext) {}

// EnterRuleObjectClassReference is called when production ruleObjectClassReference is entered.
func (s *Baseasn1Listener) EnterRuleObjectClassReference(ctx *RuleObjectClassReferenceContext) {}

// ExitRuleObjectClassReference is called when production ruleObjectClassReference is exited.
func (s *Baseasn1Listener) ExitRuleObjectClassReference(ctx *RuleObjectClassReferenceContext) {}

// EnterRuleObjectSetReference is called when production ruleObjectSetReference is entered.
func (s *Baseasn1Listener) EnterRuleObjectSetReference(ctx *RuleObjectSetReferenceContext) {}

// ExitRuleObjectSetReference is called when production ruleObjectSetReference is exited.
func (s *Baseasn1Listener) ExitRuleObjectSetReference(ctx *RuleObjectSetReferenceContext) {}

// EnterRuleParameterizedReference is called when production ruleParameterizedReference is entered.
func (s *Baseasn1Listener) EnterRuleParameterizedReference(ctx *RuleParameterizedReferenceContext) {}

// ExitRuleParameterizedReference is called when production ruleParameterizedReference is exited.
func (s *Baseasn1Listener) ExitRuleParameterizedReference(ctx *RuleParameterizedReferenceContext) {}

// EnterRuleExternalTypeReference is called when production ruleExternalTypeReference is entered.
func (s *Baseasn1Listener) EnterRuleExternalTypeReference(ctx *RuleExternalTypeReferenceContext) {}

// ExitRuleExternalTypeReference is called when production ruleExternalTypeReference is exited.
func (s *Baseasn1Listener) ExitRuleExternalTypeReference(ctx *RuleExternalTypeReferenceContext) {}

// EnterRuleExternalValueReference is called when production ruleExternalValueReference is entered.
func (s *Baseasn1Listener) EnterRuleExternalValueReference(ctx *RuleExternalValueReferenceContext) {}

// ExitRuleExternalValueReference is called when production ruleExternalValueReference is exited.
func (s *Baseasn1Listener) ExitRuleExternalValueReference(ctx *RuleExternalValueReferenceContext) {}

// EnterRuleExternalObjectClassReference is called when production ruleExternalObjectClassReference is entered.
func (s *Baseasn1Listener) EnterRuleExternalObjectClassReference(ctx *RuleExternalObjectClassReferenceContext) {
}

// ExitRuleExternalObjectClassReference is called when production ruleExternalObjectClassReference is exited.
func (s *Baseasn1Listener) ExitRuleExternalObjectClassReference(ctx *RuleExternalObjectClassReferenceContext) {
}

// EnterRuleExternalObjectReference is called when production ruleExternalObjectReference is entered.
func (s *Baseasn1Listener) EnterRuleExternalObjectReference(ctx *RuleExternalObjectReferenceContext) {
}

// ExitRuleExternalObjectReference is called when production ruleExternalObjectReference is exited.
func (s *Baseasn1Listener) ExitRuleExternalObjectReference(ctx *RuleExternalObjectReferenceContext) {}

// EnterRuleExternalObjectSetReference is called when production ruleExternalObjectSetReference is entered.
func (s *Baseasn1Listener) EnterRuleExternalObjectSetReference(ctx *RuleExternalObjectSetReferenceContext) {
}

// ExitRuleExternalObjectSetReference is called when production ruleExternalObjectSetReference is exited.
func (s *Baseasn1Listener) ExitRuleExternalObjectSetReference(ctx *RuleExternalObjectSetReferenceContext) {
}

// EnterRuleTypeFieldReference is called when production ruleTypeFieldReference is entered.
func (s *Baseasn1Listener) EnterRuleTypeFieldReference(ctx *RuleTypeFieldReferenceContext) {}

// ExitRuleTypeFieldReference is called when production ruleTypeFieldReference is exited.
func (s *Baseasn1Listener) ExitRuleTypeFieldReference(ctx *RuleTypeFieldReferenceContext) {}

// EnterRuleValueFieldReference is called when production ruleValueFieldReference is entered.
func (s *Baseasn1Listener) EnterRuleValueFieldReference(ctx *RuleValueFieldReferenceContext) {}

// ExitRuleValueFieldReference is called when production ruleValueFieldReference is exited.
func (s *Baseasn1Listener) ExitRuleValueFieldReference(ctx *RuleValueFieldReferenceContext) {}

// EnterRuleValueSetFieldReference is called when production ruleValueSetFieldReference is entered.
func (s *Baseasn1Listener) EnterRuleValueSetFieldReference(ctx *RuleValueSetFieldReferenceContext) {}

// ExitRuleValueSetFieldReference is called when production ruleValueSetFieldReference is exited.
func (s *Baseasn1Listener) ExitRuleValueSetFieldReference(ctx *RuleValueSetFieldReferenceContext) {}

// EnterRuleObjectFieldReference is called when production ruleObjectFieldReference is entered.
func (s *Baseasn1Listener) EnterRuleObjectFieldReference(ctx *RuleObjectFieldReferenceContext) {}

// ExitRuleObjectFieldReference is called when production ruleObjectFieldReference is exited.
func (s *Baseasn1Listener) ExitRuleObjectFieldReference(ctx *RuleObjectFieldReferenceContext) {}

// EnterRuleObjectSetFieldReference is called when production ruleObjectSetFieldReference is entered.
func (s *Baseasn1Listener) EnterRuleObjectSetFieldReference(ctx *RuleObjectSetFieldReferenceContext) {
}

// ExitRuleObjectSetFieldReference is called when production ruleObjectSetFieldReference is exited.
func (s *Baseasn1Listener) ExitRuleObjectSetFieldReference(ctx *RuleObjectSetFieldReferenceContext) {}

// EnterRuleUsefulObjectClassReference is called when production ruleUsefulObjectClassReference is entered.
func (s *Baseasn1Listener) EnterRuleUsefulObjectClassReference(ctx *RuleUsefulObjectClassReferenceContext) {
}

// ExitRuleUsefulObjectClassReference is called when production ruleUsefulObjectClassReference is exited.
func (s *Baseasn1Listener) ExitRuleUsefulObjectClassReference(ctx *RuleUsefulObjectClassReferenceContext) {
}

// EnterRuleImports is called when production ruleImports is entered.
func (s *Baseasn1Listener) EnterRuleImports(ctx *RuleImportsContext) {}

// ExitRuleImports is called when production ruleImports is exited.
func (s *Baseasn1Listener) ExitRuleImports(ctx *RuleImportsContext) {}

// EnterRuleSymbolsImported is called when production ruleSymbolsImported is entered.
func (s *Baseasn1Listener) EnterRuleSymbolsImported(ctx *RuleSymbolsImportedContext) {}

// ExitRuleSymbolsImported is called when production ruleSymbolsImported is exited.
func (s *Baseasn1Listener) ExitRuleSymbolsImported(ctx *RuleSymbolsImportedContext) {}

// EnterRuleSymbolsFromModuleList is called when production ruleSymbolsFromModuleList is entered.
func (s *Baseasn1Listener) EnterRuleSymbolsFromModuleList(ctx *RuleSymbolsFromModuleListContext) {}

// ExitRuleSymbolsFromModuleList is called when production ruleSymbolsFromModuleList is exited.
func (s *Baseasn1Listener) ExitRuleSymbolsFromModuleList(ctx *RuleSymbolsFromModuleListContext) {}

// EnterRuleSymbolsFromModule is called when production ruleSymbolsFromModule is entered.
func (s *Baseasn1Listener) EnterRuleSymbolsFromModule(ctx *RuleSymbolsFromModuleContext) {}

// ExitRuleSymbolsFromModule is called when production ruleSymbolsFromModule is exited.
func (s *Baseasn1Listener) ExitRuleSymbolsFromModule(ctx *RuleSymbolsFromModuleContext) {}

// EnterRuleGlobalModuleReference is called when production ruleGlobalModuleReference is entered.
func (s *Baseasn1Listener) EnterRuleGlobalModuleReference(ctx *RuleGlobalModuleReferenceContext) {}

// ExitRuleGlobalModuleReference is called when production ruleGlobalModuleReference is exited.
func (s *Baseasn1Listener) ExitRuleGlobalModuleReference(ctx *RuleGlobalModuleReferenceContext) {}

// EnterRuleAssignedIdentifier is called when production ruleAssignedIdentifier is entered.
func (s *Baseasn1Listener) EnterRuleAssignedIdentifier(ctx *RuleAssignedIdentifierContext) {}

// ExitRuleAssignedIdentifier is called when production ruleAssignedIdentifier is exited.
func (s *Baseasn1Listener) ExitRuleAssignedIdentifier(ctx *RuleAssignedIdentifierContext) {}

// EnterRuleAssignmentList is called when production ruleAssignmentList is entered.
func (s *Baseasn1Listener) EnterRuleAssignmentList(ctx *RuleAssignmentListContext) {}

// ExitRuleAssignmentList is called when production ruleAssignmentList is exited.
func (s *Baseasn1Listener) ExitRuleAssignmentList(ctx *RuleAssignmentListContext) {}

// EnterRuleAssignment is called when production ruleAssignment is entered.
func (s *Baseasn1Listener) EnterRuleAssignment(ctx *RuleAssignmentContext) {}

// ExitRuleAssignment is called when production ruleAssignment is exited.
func (s *Baseasn1Listener) ExitRuleAssignment(ctx *RuleAssignmentContext) {}

// EnterRuleTypeAssignment is called when production ruleTypeAssignment is entered.
func (s *Baseasn1Listener) EnterRuleTypeAssignment(ctx *RuleTypeAssignmentContext) {}

// ExitRuleTypeAssignment is called when production ruleTypeAssignment is exited.
func (s *Baseasn1Listener) ExitRuleTypeAssignment(ctx *RuleTypeAssignmentContext) {}

// EnterRuleValueAssignment is called when production ruleValueAssignment is entered.
func (s *Baseasn1Listener) EnterRuleValueAssignment(ctx *RuleValueAssignmentContext) {}

// ExitRuleValueAssignment is called when production ruleValueAssignment is exited.
func (s *Baseasn1Listener) ExitRuleValueAssignment(ctx *RuleValueAssignmentContext) {}

// EnterRuleXMLValueAssignment is called when production ruleXMLValueAssignment is entered.
func (s *Baseasn1Listener) EnterRuleXMLValueAssignment(ctx *RuleXMLValueAssignmentContext) {}

// ExitRuleXMLValueAssignment is called when production ruleXMLValueAssignment is exited.
func (s *Baseasn1Listener) ExitRuleXMLValueAssignment(ctx *RuleXMLValueAssignmentContext) {}

// EnterRuleValueSetTypeAssignment is called when production ruleValueSetTypeAssignment is entered.
func (s *Baseasn1Listener) EnterRuleValueSetTypeAssignment(ctx *RuleValueSetTypeAssignmentContext) {}

// ExitRuleValueSetTypeAssignment is called when production ruleValueSetTypeAssignment is exited.
func (s *Baseasn1Listener) ExitRuleValueSetTypeAssignment(ctx *RuleValueSetTypeAssignmentContext) {}

// EnterRuleObjectAssignment is called when production ruleObjectAssignment is entered.
func (s *Baseasn1Listener) EnterRuleObjectAssignment(ctx *RuleObjectAssignmentContext) {}

// ExitRuleObjectAssignment is called when production ruleObjectAssignment is exited.
func (s *Baseasn1Listener) ExitRuleObjectAssignment(ctx *RuleObjectAssignmentContext) {}

// EnterRuleObjectClassAssignment is called when production ruleObjectClassAssignment is entered.
func (s *Baseasn1Listener) EnterRuleObjectClassAssignment(ctx *RuleObjectClassAssignmentContext) {}

// ExitRuleObjectClassAssignment is called when production ruleObjectClassAssignment is exited.
func (s *Baseasn1Listener) ExitRuleObjectClassAssignment(ctx *RuleObjectClassAssignmentContext) {}

// EnterRuleObjectSetAssignment is called when production ruleObjectSetAssignment is entered.
func (s *Baseasn1Listener) EnterRuleObjectSetAssignment(ctx *RuleObjectSetAssignmentContext) {}

// ExitRuleObjectSetAssignment is called when production ruleObjectSetAssignment is exited.
func (s *Baseasn1Listener) ExitRuleObjectSetAssignment(ctx *RuleObjectSetAssignmentContext) {}

// EnterRuleParameterizedAssignment is called when production ruleParameterizedAssignment is entered.
func (s *Baseasn1Listener) EnterRuleParameterizedAssignment(ctx *RuleParameterizedAssignmentContext) {
}

// ExitRuleParameterizedAssignment is called when production ruleParameterizedAssignment is exited.
func (s *Baseasn1Listener) ExitRuleParameterizedAssignment(ctx *RuleParameterizedAssignmentContext) {}

// EnterRuleObjectIdentifierValue is called when production ruleObjectIdentifierValue is entered.
func (s *Baseasn1Listener) EnterRuleObjectIdentifierValue(ctx *RuleObjectIdentifierValueContext) {}

// ExitRuleObjectIdentifierValue is called when production ruleObjectIdentifierValue is exited.
func (s *Baseasn1Listener) ExitRuleObjectIdentifierValue(ctx *RuleObjectIdentifierValueContext) {}

// EnterRuleObjIdComponentsList is called when production ruleObjIdComponentsList is entered.
func (s *Baseasn1Listener) EnterRuleObjIdComponentsList(ctx *RuleObjIdComponentsListContext) {}

// ExitRuleObjIdComponentsList is called when production ruleObjIdComponentsList is exited.
func (s *Baseasn1Listener) ExitRuleObjIdComponentsList(ctx *RuleObjIdComponentsListContext) {}

// EnterRuleObjIdComponents is called when production ruleObjIdComponents is entered.
func (s *Baseasn1Listener) EnterRuleObjIdComponents(ctx *RuleObjIdComponentsContext) {}

// ExitRuleObjIdComponents is called when production ruleObjIdComponents is exited.
func (s *Baseasn1Listener) ExitRuleObjIdComponents(ctx *RuleObjIdComponentsContext) {}

// EnterRuleNumberForm is called when production ruleNumberForm is entered.
func (s *Baseasn1Listener) EnterRuleNumberForm(ctx *RuleNumberFormContext) {}

// ExitRuleNumberForm is called when production ruleNumberForm is exited.
func (s *Baseasn1Listener) ExitRuleNumberForm(ctx *RuleNumberFormContext) {}

// EnterRuleNameAndNumberForm is called when production ruleNameAndNumberForm is entered.
func (s *Baseasn1Listener) EnterRuleNameAndNumberForm(ctx *RuleNameAndNumberFormContext) {}

// ExitRuleNameAndNumberForm is called when production ruleNameAndNumberForm is exited.
func (s *Baseasn1Listener) ExitRuleNameAndNumberForm(ctx *RuleNameAndNumberFormContext) {}

// EnterRuleDefinedValue is called when production ruleDefinedValue is entered.
func (s *Baseasn1Listener) EnterRuleDefinedValue(ctx *RuleDefinedValueContext) {}

// ExitRuleDefinedValue is called when production ruleDefinedValue is exited.
func (s *Baseasn1Listener) ExitRuleDefinedValue(ctx *RuleDefinedValueContext) {}

// EnterRuleNamedBit is called when production ruleNamedBit is entered.
func (s *Baseasn1Listener) EnterRuleNamedBit(ctx *RuleNamedBitContext) {}

// ExitRuleNamedBit is called when production ruleNamedBit is exited.
func (s *Baseasn1Listener) ExitRuleNamedBit(ctx *RuleNamedBitContext) {}

// EnterRuleNamedBitList is called when production ruleNamedBitList is entered.
func (s *Baseasn1Listener) EnterRuleNamedBitList(ctx *RuleNamedBitListContext) {}

// ExitRuleNamedBitList is called when production ruleNamedBitList is exited.
func (s *Baseasn1Listener) ExitRuleNamedBitList(ctx *RuleNamedBitListContext) {}

// EnterRuleRestrictedCharacterStringType is called when production ruleRestrictedCharacterStringType is entered.
func (s *Baseasn1Listener) EnterRuleRestrictedCharacterStringType(ctx *RuleRestrictedCharacterStringTypeContext) {
}

// ExitRuleRestrictedCharacterStringType is called when production ruleRestrictedCharacterStringType is exited.
func (s *Baseasn1Listener) ExitRuleRestrictedCharacterStringType(ctx *RuleRestrictedCharacterStringTypeContext) {
}

// EnterRuleUnrestrictedCharacterStringType is called when production ruleUnrestrictedCharacterStringType is entered.
func (s *Baseasn1Listener) EnterRuleUnrestrictedCharacterStringType(ctx *RuleUnrestrictedCharacterStringTypeContext) {
}

// ExitRuleUnrestrictedCharacterStringType is called when production ruleUnrestrictedCharacterStringType is exited.
func (s *Baseasn1Listener) ExitRuleUnrestrictedCharacterStringType(ctx *RuleUnrestrictedCharacterStringTypeContext) {
}

// EnterRuleAlternativeTypeList is called when production ruleAlternativeTypeList is entered.
func (s *Baseasn1Listener) EnterRuleAlternativeTypeList(ctx *RuleAlternativeTypeListContext) {}

// ExitRuleAlternativeTypeList is called when production ruleAlternativeTypeList is exited.
func (s *Baseasn1Listener) ExitRuleAlternativeTypeList(ctx *RuleAlternativeTypeListContext) {}

// EnterRuleRootAlternativeTypeList is called when production ruleRootAlternativeTypeList is entered.
func (s *Baseasn1Listener) EnterRuleRootAlternativeTypeList(ctx *RuleRootAlternativeTypeListContext) {
}

// ExitRuleRootAlternativeTypeList is called when production ruleRootAlternativeTypeList is exited.
func (s *Baseasn1Listener) ExitRuleRootAlternativeTypeList(ctx *RuleRootAlternativeTypeListContext) {}

// EnterRuleVersionNumber is called when production ruleVersionNumber is entered.
func (s *Baseasn1Listener) EnterRuleVersionNumber(ctx *RuleVersionNumberContext) {}

// ExitRuleVersionNumber is called when production ruleVersionNumber is exited.
func (s *Baseasn1Listener) ExitRuleVersionNumber(ctx *RuleVersionNumberContext) {}

// EnterRuleExtensionAdditionAlternativesGroup is called when production ruleExtensionAdditionAlternativesGroup is entered.
func (s *Baseasn1Listener) EnterRuleExtensionAdditionAlternativesGroup(ctx *RuleExtensionAdditionAlternativesGroupContext) {
}

// ExitRuleExtensionAdditionAlternativesGroup is called when production ruleExtensionAdditionAlternativesGroup is exited.
func (s *Baseasn1Listener) ExitRuleExtensionAdditionAlternativesGroup(ctx *RuleExtensionAdditionAlternativesGroupContext) {
}

// EnterRuleExtensionAdditionAlternative is called when production ruleExtensionAdditionAlternative is entered.
func (s *Baseasn1Listener) EnterRuleExtensionAdditionAlternative(ctx *RuleExtensionAdditionAlternativeContext) {
}

// ExitRuleExtensionAdditionAlternative is called when production ruleExtensionAdditionAlternative is exited.
func (s *Baseasn1Listener) ExitRuleExtensionAdditionAlternative(ctx *RuleExtensionAdditionAlternativeContext) {
}

// EnterRuleExtensionAdditionAlternativesList is called when production ruleExtensionAdditionAlternativesList is entered.
func (s *Baseasn1Listener) EnterRuleExtensionAdditionAlternativesList(ctx *RuleExtensionAdditionAlternativesListContext) {
}

// ExitRuleExtensionAdditionAlternativesList is called when production ruleExtensionAdditionAlternativesList is exited.
func (s *Baseasn1Listener) ExitRuleExtensionAdditionAlternativesList(ctx *RuleExtensionAdditionAlternativesListContext) {
}

// EnterRuleExtensionAdditionAlternatives is called when production ruleExtensionAdditionAlternatives is entered.
func (s *Baseasn1Listener) EnterRuleExtensionAdditionAlternatives(ctx *RuleExtensionAdditionAlternativesContext) {
}

// ExitRuleExtensionAdditionAlternatives is called when production ruleExtensionAdditionAlternatives is exited.
func (s *Baseasn1Listener) ExitRuleExtensionAdditionAlternatives(ctx *RuleExtensionAdditionAlternativesContext) {
}

// EnterRuleEnumerationItem is called when production ruleEnumerationItem is entered.
func (s *Baseasn1Listener) EnterRuleEnumerationItem(ctx *RuleEnumerationItemContext) {}

// ExitRuleEnumerationItem is called when production ruleEnumerationItem is exited.
func (s *Baseasn1Listener) ExitRuleEnumerationItem(ctx *RuleEnumerationItemContext) {}

// EnterRuleEnumeration is called when production ruleEnumeration is entered.
func (s *Baseasn1Listener) EnterRuleEnumeration(ctx *RuleEnumerationContext) {}

// ExitRuleEnumeration is called when production ruleEnumeration is exited.
func (s *Baseasn1Listener) ExitRuleEnumeration(ctx *RuleEnumerationContext) {}

// EnterRuleRootEnumeration is called when production ruleRootEnumeration is entered.
func (s *Baseasn1Listener) EnterRuleRootEnumeration(ctx *RuleRootEnumerationContext) {}

// ExitRuleRootEnumeration is called when production ruleRootEnumeration is exited.
func (s *Baseasn1Listener) ExitRuleRootEnumeration(ctx *RuleRootEnumerationContext) {}

// EnterRuleAdditionalEnumeration is called when production ruleAdditionalEnumeration is entered.
func (s *Baseasn1Listener) EnterRuleAdditionalEnumeration(ctx *RuleAdditionalEnumerationContext) {}

// ExitRuleAdditionalEnumeration is called when production ruleAdditionalEnumeration is exited.
func (s *Baseasn1Listener) ExitRuleAdditionalEnumeration(ctx *RuleAdditionalEnumerationContext) {}

// EnterRuleAlternativeTypeLists is called when production ruleAlternativeTypeLists is entered.
func (s *Baseasn1Listener) EnterRuleAlternativeTypeLists(ctx *RuleAlternativeTypeListsContext) {}

// ExitRuleAlternativeTypeLists is called when production ruleAlternativeTypeLists is exited.
func (s *Baseasn1Listener) ExitRuleAlternativeTypeLists(ctx *RuleAlternativeTypeListsContext) {}

// EnterRuleEnumerations is called when production ruleEnumerations is entered.
func (s *Baseasn1Listener) EnterRuleEnumerations(ctx *RuleEnumerationsContext) {}

// ExitRuleEnumerations is called when production ruleEnumerations is exited.
func (s *Baseasn1Listener) ExitRuleEnumerations(ctx *RuleEnumerationsContext) {}

// EnterRuleNamedNumber is called when production ruleNamedNumber is entered.
func (s *Baseasn1Listener) EnterRuleNamedNumber(ctx *RuleNamedNumberContext) {}

// ExitRuleNamedNumber is called when production ruleNamedNumber is exited.
func (s *Baseasn1Listener) ExitRuleNamedNumber(ctx *RuleNamedNumberContext) {}

// EnterRuleNamedNumberList is called when production ruleNamedNumberList is entered.
func (s *Baseasn1Listener) EnterRuleNamedNumberList(ctx *RuleNamedNumberListContext) {}

// ExitRuleNamedNumberList is called when production ruleNamedNumberList is exited.
func (s *Baseasn1Listener) ExitRuleNamedNumberList(ctx *RuleNamedNumberListContext) {}

// EnterRuleComponentType is called when production ruleComponentType is entered.
func (s *Baseasn1Listener) EnterRuleComponentType(ctx *RuleComponentTypeContext) {}

// ExitRuleComponentType is called when production ruleComponentType is exited.
func (s *Baseasn1Listener) ExitRuleComponentType(ctx *RuleComponentTypeContext) {}

// EnterRuleExtensionAdditionGroup is called when production ruleExtensionAdditionGroup is entered.
func (s *Baseasn1Listener) EnterRuleExtensionAdditionGroup(ctx *RuleExtensionAdditionGroupContext) {}

// ExitRuleExtensionAdditionGroup is called when production ruleExtensionAdditionGroup is exited.
func (s *Baseasn1Listener) ExitRuleExtensionAdditionGroup(ctx *RuleExtensionAdditionGroupContext) {}

// EnterRuleExtensionAddition is called when production ruleExtensionAddition is entered.
func (s *Baseasn1Listener) EnterRuleExtensionAddition(ctx *RuleExtensionAdditionContext) {}

// ExitRuleExtensionAddition is called when production ruleExtensionAddition is exited.
func (s *Baseasn1Listener) ExitRuleExtensionAddition(ctx *RuleExtensionAdditionContext) {}

// EnterRuleComponentTypeList is called when production ruleComponentTypeList is entered.
func (s *Baseasn1Listener) EnterRuleComponentTypeList(ctx *RuleComponentTypeListContext) {}

// ExitRuleComponentTypeList is called when production ruleComponentTypeList is exited.
func (s *Baseasn1Listener) ExitRuleComponentTypeList(ctx *RuleComponentTypeListContext) {}

// EnterRuleExtensionAdditionList is called when production ruleExtensionAdditionList is entered.
func (s *Baseasn1Listener) EnterRuleExtensionAdditionList(ctx *RuleExtensionAdditionListContext) {}

// ExitRuleExtensionAdditionList is called when production ruleExtensionAdditionList is exited.
func (s *Baseasn1Listener) ExitRuleExtensionAdditionList(ctx *RuleExtensionAdditionListContext) {}

// EnterRuleRootComponentTypeList is called when production ruleRootComponentTypeList is entered.
func (s *Baseasn1Listener) EnterRuleRootComponentTypeList(ctx *RuleRootComponentTypeListContext) {}

// ExitRuleRootComponentTypeList is called when production ruleRootComponentTypeList is exited.
func (s *Baseasn1Listener) ExitRuleRootComponentTypeList(ctx *RuleRootComponentTypeListContext) {}

// EnterRuleExtensionAdditions is called when production ruleExtensionAdditions is entered.
func (s *Baseasn1Listener) EnterRuleExtensionAdditions(ctx *RuleExtensionAdditionsContext) {}

// ExitRuleExtensionAdditions is called when production ruleExtensionAdditions is exited.
func (s *Baseasn1Listener) ExitRuleExtensionAdditions(ctx *RuleExtensionAdditionsContext) {}

// EnterRuleExtensionEndMarker is called when production ruleExtensionEndMarker is entered.
func (s *Baseasn1Listener) EnterRuleExtensionEndMarker(ctx *RuleExtensionEndMarkerContext) {}

// ExitRuleExtensionEndMarker is called when production ruleExtensionEndMarker is exited.
func (s *Baseasn1Listener) ExitRuleExtensionEndMarker(ctx *RuleExtensionEndMarkerContext) {}

// EnterRuleComponentTypeLists is called when production ruleComponentTypeLists is entered.
func (s *Baseasn1Listener) EnterRuleComponentTypeLists(ctx *RuleComponentTypeListsContext) {}

// ExitRuleComponentTypeLists is called when production ruleComponentTypeLists is exited.
func (s *Baseasn1Listener) ExitRuleComponentTypeLists(ctx *RuleComponentTypeListsContext) {}

// EnterRuleExtensionAndException is called when production ruleExtensionAndException is entered.
func (s *Baseasn1Listener) EnterRuleExtensionAndException(ctx *RuleExtensionAndExceptionContext) {}

// ExitRuleExtensionAndException is called when production ruleExtensionAndException is exited.
func (s *Baseasn1Listener) ExitRuleExtensionAndException(ctx *RuleExtensionAndExceptionContext) {}

// EnterRuleOptionalExtensionMarker is called when production ruleOptionalExtensionMarker is entered.
func (s *Baseasn1Listener) EnterRuleOptionalExtensionMarker(ctx *RuleOptionalExtensionMarkerContext) {
}

// ExitRuleOptionalExtensionMarker is called when production ruleOptionalExtensionMarker is exited.
func (s *Baseasn1Listener) ExitRuleOptionalExtensionMarker(ctx *RuleOptionalExtensionMarkerContext) {}

// EnterRuleEncodingReference is called when production ruleEncodingReference is entered.
func (s *Baseasn1Listener) EnterRuleEncodingReference(ctx *RuleEncodingReferenceContext) {}

// ExitRuleEncodingReference is called when production ruleEncodingReference is exited.
func (s *Baseasn1Listener) ExitRuleEncodingReference(ctx *RuleEncodingReferenceContext) {}

// EnterRuleClass is called when production ruleClass is entered.
func (s *Baseasn1Listener) EnterRuleClass(ctx *RuleClassContext) {}

// ExitRuleClass is called when production ruleClass is exited.
func (s *Baseasn1Listener) ExitRuleClass(ctx *RuleClassContext) {}

// EnterRuleClassNumber is called when production ruleClassNumber is entered.
func (s *Baseasn1Listener) EnterRuleClassNumber(ctx *RuleClassNumberContext) {}

// ExitRuleClassNumber is called when production ruleClassNumber is exited.
func (s *Baseasn1Listener) ExitRuleClassNumber(ctx *RuleClassNumberContext) {}

// EnterRuleEncodingInstruction is called when production ruleEncodingInstruction is entered.
func (s *Baseasn1Listener) EnterRuleEncodingInstruction(ctx *RuleEncodingInstructionContext) {}

// ExitRuleEncodingInstruction is called when production ruleEncodingInstruction is exited.
func (s *Baseasn1Listener) ExitRuleEncodingInstruction(ctx *RuleEncodingInstructionContext) {}

// EnterRuleTag is called when production ruleTag is entered.
func (s *Baseasn1Listener) EnterRuleTag(ctx *RuleTagContext) {}

// ExitRuleTag is called when production ruleTag is exited.
func (s *Baseasn1Listener) ExitRuleTag(ctx *RuleTagContext) {}

// EnterRuleEncodingPrefix is called when production ruleEncodingPrefix is entered.
func (s *Baseasn1Listener) EnterRuleEncodingPrefix(ctx *RuleEncodingPrefixContext) {}

// ExitRuleEncodingPrefix is called when production ruleEncodingPrefix is exited.
func (s *Baseasn1Listener) ExitRuleEncodingPrefix(ctx *RuleEncodingPrefixContext) {}

// EnterRuleTaggedType is called when production ruleTaggedType is entered.
func (s *Baseasn1Listener) EnterRuleTaggedType(ctx *RuleTaggedTypeContext) {}

// ExitRuleTaggedType is called when production ruleTaggedType is exited.
func (s *Baseasn1Listener) ExitRuleTaggedType(ctx *RuleTaggedTypeContext) {}

// EnterRuleEncodingPrefixedType is called when production ruleEncodingPrefixedType is entered.
func (s *Baseasn1Listener) EnterRuleEncodingPrefixedType(ctx *RuleEncodingPrefixedTypeContext) {}

// ExitRuleEncodingPrefixedType is called when production ruleEncodingPrefixedType is exited.
func (s *Baseasn1Listener) ExitRuleEncodingPrefixedType(ctx *RuleEncodingPrefixedTypeContext) {}

// EnterRuleBitStringType is called when production ruleBitStringType is entered.
func (s *Baseasn1Listener) EnterRuleBitStringType(ctx *RuleBitStringTypeContext) {}

// ExitRuleBitStringType is called when production ruleBitStringType is exited.
func (s *Baseasn1Listener) ExitRuleBitStringType(ctx *RuleBitStringTypeContext) {}

// EnterRuleBooleanType is called when production ruleBooleanType is entered.
func (s *Baseasn1Listener) EnterRuleBooleanType(ctx *RuleBooleanTypeContext) {}

// ExitRuleBooleanType is called when production ruleBooleanType is exited.
func (s *Baseasn1Listener) ExitRuleBooleanType(ctx *RuleBooleanTypeContext) {}

// EnterRuleCharacterStringType is called when production ruleCharacterStringType is entered.
func (s *Baseasn1Listener) EnterRuleCharacterStringType(ctx *RuleCharacterStringTypeContext) {}

// ExitRuleCharacterStringType is called when production ruleCharacterStringType is exited.
func (s *Baseasn1Listener) ExitRuleCharacterStringType(ctx *RuleCharacterStringTypeContext) {}

// EnterRuleChoiceType is called when production ruleChoiceType is entered.
func (s *Baseasn1Listener) EnterRuleChoiceType(ctx *RuleChoiceTypeContext) {}

// ExitRuleChoiceType is called when production ruleChoiceType is exited.
func (s *Baseasn1Listener) ExitRuleChoiceType(ctx *RuleChoiceTypeContext) {}

// EnterRuleDateType is called when production ruleDateType is entered.
func (s *Baseasn1Listener) EnterRuleDateType(ctx *RuleDateTypeContext) {}

// ExitRuleDateType is called when production ruleDateType is exited.
func (s *Baseasn1Listener) ExitRuleDateType(ctx *RuleDateTypeContext) {}

// EnterRuleDateTimeType is called when production ruleDateTimeType is entered.
func (s *Baseasn1Listener) EnterRuleDateTimeType(ctx *RuleDateTimeTypeContext) {}

// ExitRuleDateTimeType is called when production ruleDateTimeType is exited.
func (s *Baseasn1Listener) ExitRuleDateTimeType(ctx *RuleDateTimeTypeContext) {}

// EnterRuleDurationType is called when production ruleDurationType is entered.
func (s *Baseasn1Listener) EnterRuleDurationType(ctx *RuleDurationTypeContext) {}

// ExitRuleDurationType is called when production ruleDurationType is exited.
func (s *Baseasn1Listener) ExitRuleDurationType(ctx *RuleDurationTypeContext) {}

// EnterRuleEmbeddedPDVType is called when production ruleEmbeddedPDVType is entered.
func (s *Baseasn1Listener) EnterRuleEmbeddedPDVType(ctx *RuleEmbeddedPDVTypeContext) {}

// ExitRuleEmbeddedPDVType is called when production ruleEmbeddedPDVType is exited.
func (s *Baseasn1Listener) ExitRuleEmbeddedPDVType(ctx *RuleEmbeddedPDVTypeContext) {}

// EnterRuleEnumeratedType is called when production ruleEnumeratedType is entered.
func (s *Baseasn1Listener) EnterRuleEnumeratedType(ctx *RuleEnumeratedTypeContext) {}

// ExitRuleEnumeratedType is called when production ruleEnumeratedType is exited.
func (s *Baseasn1Listener) ExitRuleEnumeratedType(ctx *RuleEnumeratedTypeContext) {}

// EnterRuleExternalType is called when production ruleExternalType is entered.
func (s *Baseasn1Listener) EnterRuleExternalType(ctx *RuleExternalTypeContext) {}

// ExitRuleExternalType is called when production ruleExternalType is exited.
func (s *Baseasn1Listener) ExitRuleExternalType(ctx *RuleExternalTypeContext) {}

// EnterRuleInstanceOfType is called when production ruleInstanceOfType is entered.
func (s *Baseasn1Listener) EnterRuleInstanceOfType(ctx *RuleInstanceOfTypeContext) {}

// ExitRuleInstanceOfType is called when production ruleInstanceOfType is exited.
func (s *Baseasn1Listener) ExitRuleInstanceOfType(ctx *RuleInstanceOfTypeContext) {}

// EnterRuleIntegerType is called when production ruleIntegerType is entered.
func (s *Baseasn1Listener) EnterRuleIntegerType(ctx *RuleIntegerTypeContext) {}

// ExitRuleIntegerType is called when production ruleIntegerType is exited.
func (s *Baseasn1Listener) ExitRuleIntegerType(ctx *RuleIntegerTypeContext) {}

// EnterRuleIRIType is called when production ruleIRIType is entered.
func (s *Baseasn1Listener) EnterRuleIRIType(ctx *RuleIRITypeContext) {}

// ExitRuleIRIType is called when production ruleIRIType is exited.
func (s *Baseasn1Listener) ExitRuleIRIType(ctx *RuleIRITypeContext) {}

// EnterRuleNullType is called when production ruleNullType is entered.
func (s *Baseasn1Listener) EnterRuleNullType(ctx *RuleNullTypeContext) {}

// ExitRuleNullType is called when production ruleNullType is exited.
func (s *Baseasn1Listener) ExitRuleNullType(ctx *RuleNullTypeContext) {}

// EnterRuleObjectClassFieldType is called when production ruleObjectClassFieldType is entered.
func (s *Baseasn1Listener) EnterRuleObjectClassFieldType(ctx *RuleObjectClassFieldTypeContext) {}

// ExitRuleObjectClassFieldType is called when production ruleObjectClassFieldType is exited.
func (s *Baseasn1Listener) ExitRuleObjectClassFieldType(ctx *RuleObjectClassFieldTypeContext) {}

// EnterRuleObjectIdentifierType is called when production ruleObjectIdentifierType is entered.
func (s *Baseasn1Listener) EnterRuleObjectIdentifierType(ctx *RuleObjectIdentifierTypeContext) {}

// ExitRuleObjectIdentifierType is called when production ruleObjectIdentifierType is exited.
func (s *Baseasn1Listener) ExitRuleObjectIdentifierType(ctx *RuleObjectIdentifierTypeContext) {}

// EnterRuleOctetStringType is called when production ruleOctetStringType is entered.
func (s *Baseasn1Listener) EnterRuleOctetStringType(ctx *RuleOctetStringTypeContext) {}

// ExitRuleOctetStringType is called when production ruleOctetStringType is exited.
func (s *Baseasn1Listener) ExitRuleOctetStringType(ctx *RuleOctetStringTypeContext) {}

// EnterRuleRealType is called when production ruleRealType is entered.
func (s *Baseasn1Listener) EnterRuleRealType(ctx *RuleRealTypeContext) {}

// ExitRuleRealType is called when production ruleRealType is exited.
func (s *Baseasn1Listener) ExitRuleRealType(ctx *RuleRealTypeContext) {}

// EnterRuleRelativeIRIType is called when production ruleRelativeIRIType is entered.
func (s *Baseasn1Listener) EnterRuleRelativeIRIType(ctx *RuleRelativeIRITypeContext) {}

// ExitRuleRelativeIRIType is called when production ruleRelativeIRIType is exited.
func (s *Baseasn1Listener) ExitRuleRelativeIRIType(ctx *RuleRelativeIRITypeContext) {}

// EnterRuleRelativeOIDType is called when production ruleRelativeOIDType is entered.
func (s *Baseasn1Listener) EnterRuleRelativeOIDType(ctx *RuleRelativeOIDTypeContext) {}

// ExitRuleRelativeOIDType is called when production ruleRelativeOIDType is exited.
func (s *Baseasn1Listener) ExitRuleRelativeOIDType(ctx *RuleRelativeOIDTypeContext) {}

// EnterRuleSequenceType is called when production ruleSequenceType is entered.
func (s *Baseasn1Listener) EnterRuleSequenceType(ctx *RuleSequenceTypeContext) {}

// ExitRuleSequenceType is called when production ruleSequenceType is exited.
func (s *Baseasn1Listener) ExitRuleSequenceType(ctx *RuleSequenceTypeContext) {}

// EnterRuleSequenceOfType is called when production ruleSequenceOfType is entered.
func (s *Baseasn1Listener) EnterRuleSequenceOfType(ctx *RuleSequenceOfTypeContext) {}

// ExitRuleSequenceOfType is called when production ruleSequenceOfType is exited.
func (s *Baseasn1Listener) ExitRuleSequenceOfType(ctx *RuleSequenceOfTypeContext) {}

// EnterRuleSetType is called when production ruleSetType is entered.
func (s *Baseasn1Listener) EnterRuleSetType(ctx *RuleSetTypeContext) {}

// ExitRuleSetType is called when production ruleSetType is exited.
func (s *Baseasn1Listener) ExitRuleSetType(ctx *RuleSetTypeContext) {}

// EnterRuleSetOfType is called when production ruleSetOfType is entered.
func (s *Baseasn1Listener) EnterRuleSetOfType(ctx *RuleSetOfTypeContext) {}

// ExitRuleSetOfType is called when production ruleSetOfType is exited.
func (s *Baseasn1Listener) ExitRuleSetOfType(ctx *RuleSetOfTypeContext) {}

// EnterRulePrefixedType is called when production rulePrefixedType is entered.
func (s *Baseasn1Listener) EnterRulePrefixedType(ctx *RulePrefixedTypeContext) {}

// ExitRulePrefixedType is called when production rulePrefixedType is exited.
func (s *Baseasn1Listener) ExitRulePrefixedType(ctx *RulePrefixedTypeContext) {}

// EnterRuleTimeType is called when production ruleTimeType is entered.
func (s *Baseasn1Listener) EnterRuleTimeType(ctx *RuleTimeTypeContext) {}

// ExitRuleTimeType is called when production ruleTimeType is exited.
func (s *Baseasn1Listener) ExitRuleTimeType(ctx *RuleTimeTypeContext) {}

// EnterRuleTimeOfDayType is called when production ruleTimeOfDayType is entered.
func (s *Baseasn1Listener) EnterRuleTimeOfDayType(ctx *RuleTimeOfDayTypeContext) {}

// ExitRuleTimeOfDayType is called when production ruleTimeOfDayType is exited.
func (s *Baseasn1Listener) ExitRuleTimeOfDayType(ctx *RuleTimeOfDayTypeContext) {}

// EnterRuleBuiltinType is called when production ruleBuiltinType is entered.
func (s *Baseasn1Listener) EnterRuleBuiltinType(ctx *RuleBuiltinTypeContext) {}

// ExitRuleBuiltinType is called when production ruleBuiltinType is exited.
func (s *Baseasn1Listener) ExitRuleBuiltinType(ctx *RuleBuiltinTypeContext) {}

// EnterRuleSimpleDefinedType is called when production ruleSimpleDefinedType is entered.
func (s *Baseasn1Listener) EnterRuleSimpleDefinedType(ctx *RuleSimpleDefinedTypeContext) {}

// ExitRuleSimpleDefinedType is called when production ruleSimpleDefinedType is exited.
func (s *Baseasn1Listener) ExitRuleSimpleDefinedType(ctx *RuleSimpleDefinedTypeContext) {}

// EnterRuleParameterizedType is called when production ruleParameterizedType is entered.
func (s *Baseasn1Listener) EnterRuleParameterizedType(ctx *RuleParameterizedTypeContext) {}

// ExitRuleParameterizedType is called when production ruleParameterizedType is exited.
func (s *Baseasn1Listener) ExitRuleParameterizedType(ctx *RuleParameterizedTypeContext) {}

// EnterRuleParameterizedValueSetType is called when production ruleParameterizedValueSetType is entered.
func (s *Baseasn1Listener) EnterRuleParameterizedValueSetType(ctx *RuleParameterizedValueSetTypeContext) {
}

// ExitRuleParameterizedValueSetType is called when production ruleParameterizedValueSetType is exited.
func (s *Baseasn1Listener) ExitRuleParameterizedValueSetType(ctx *RuleParameterizedValueSetTypeContext) {
}

// EnterRuleDefinedType is called when production ruleDefinedType is entered.
func (s *Baseasn1Listener) EnterRuleDefinedType(ctx *RuleDefinedTypeContext) {}

// ExitRuleDefinedType is called when production ruleDefinedType is exited.
func (s *Baseasn1Listener) ExitRuleDefinedType(ctx *RuleDefinedTypeContext) {}

// EnterRuleUsefulType is called when production ruleUsefulType is entered.
func (s *Baseasn1Listener) EnterRuleUsefulType(ctx *RuleUsefulTypeContext) {}

// ExitRuleUsefulType is called when production ruleUsefulType is exited.
func (s *Baseasn1Listener) ExitRuleUsefulType(ctx *RuleUsefulTypeContext) {}

// EnterRuleSelectionType is called when production ruleSelectionType is entered.
func (s *Baseasn1Listener) EnterRuleSelectionType(ctx *RuleSelectionTypeContext) {}

// ExitRuleSelectionType is called when production ruleSelectionType is exited.
func (s *Baseasn1Listener) ExitRuleSelectionType(ctx *RuleSelectionTypeContext) {}

// EnterRuleTypeFromObject is called when production ruleTypeFromObject is entered.
func (s *Baseasn1Listener) EnterRuleTypeFromObject(ctx *RuleTypeFromObjectContext) {}

// ExitRuleTypeFromObject is called when production ruleTypeFromObject is exited.
func (s *Baseasn1Listener) ExitRuleTypeFromObject(ctx *RuleTypeFromObjectContext) {}

// EnterRuleValueSetFromObjects is called when production ruleValueSetFromObjects is entered.
func (s *Baseasn1Listener) EnterRuleValueSetFromObjects(ctx *RuleValueSetFromObjectsContext) {}

// ExitRuleValueSetFromObjects is called when production ruleValueSetFromObjects is exited.
func (s *Baseasn1Listener) ExitRuleValueSetFromObjects(ctx *RuleValueSetFromObjectsContext) {}

// EnterRuleReferencedType is called when production ruleReferencedType is entered.
func (s *Baseasn1Listener) EnterRuleReferencedType(ctx *RuleReferencedTypeContext) {}

// ExitRuleReferencedType is called when production ruleReferencedType is exited.
func (s *Baseasn1Listener) ExitRuleReferencedType(ctx *RuleReferencedTypeContext) {}

// EnterRuleTypeForConstraint is called when production ruleTypeForConstraint is entered.
func (s *Baseasn1Listener) EnterRuleTypeForConstraint(ctx *RuleTypeForConstraintContext) {}

// ExitRuleTypeForConstraint is called when production ruleTypeForConstraint is exited.
func (s *Baseasn1Listener) ExitRuleTypeForConstraint(ctx *RuleTypeForConstraintContext) {}

// EnterRuleNamedType is called when production ruleNamedType is entered.
func (s *Baseasn1Listener) EnterRuleNamedType(ctx *RuleNamedTypeContext) {}

// ExitRuleNamedType is called when production ruleNamedType is exited.
func (s *Baseasn1Listener) ExitRuleNamedType(ctx *RuleNamedTypeContext) {}

// EnterRuleTypeWithConstraint is called when production ruleTypeWithConstraint is entered.
func (s *Baseasn1Listener) EnterRuleTypeWithConstraint(ctx *RuleTypeWithConstraintContext) {}

// ExitRuleTypeWithConstraint is called when production ruleTypeWithConstraint is exited.
func (s *Baseasn1Listener) ExitRuleTypeWithConstraint(ctx *RuleTypeWithConstraintContext) {}

// EnterRuleConstrainedType is called when production ruleConstrainedType is entered.
func (s *Baseasn1Listener) EnterRuleConstrainedType(ctx *RuleConstrainedTypeContext) {}

// ExitRuleConstrainedType is called when production ruleConstrainedType is exited.
func (s *Baseasn1Listener) ExitRuleConstrainedType(ctx *RuleConstrainedTypeContext) {}

// EnterRuleType is called when production ruleType is entered.
func (s *Baseasn1Listener) EnterRuleType(ctx *RuleTypeContext) {}

// ExitRuleType is called when production ruleType is exited.
func (s *Baseasn1Listener) ExitRuleType(ctx *RuleTypeContext) {}

// EnterRuleIdentifierList is called when production ruleIdentifierList is entered.
func (s *Baseasn1Listener) EnterRuleIdentifierList(ctx *RuleIdentifierListContext) {}

// ExitRuleIdentifierList is called when production ruleIdentifierList is exited.
func (s *Baseasn1Listener) ExitRuleIdentifierList(ctx *RuleIdentifierListContext) {}

// EnterRuleCharsDefn is called when production ruleCharsDefn is entered.
func (s *Baseasn1Listener) EnterRuleCharsDefn(ctx *RuleCharsDefnContext) {}

// ExitRuleCharsDefn is called when production ruleCharsDefn is exited.
func (s *Baseasn1Listener) ExitRuleCharsDefn(ctx *RuleCharsDefnContext) {}

// EnterRuleCharSyms is called when production ruleCharSyms is entered.
func (s *Baseasn1Listener) EnterRuleCharSyms(ctx *RuleCharSymsContext) {}

// ExitRuleCharSyms is called when production ruleCharSyms is exited.
func (s *Baseasn1Listener) ExitRuleCharSyms(ctx *RuleCharSymsContext) {}

// EnterRuleGroup is called when production ruleGroup is entered.
func (s *Baseasn1Listener) EnterRuleGroup(ctx *RuleGroupContext) {}

// ExitRuleGroup is called when production ruleGroup is exited.
func (s *Baseasn1Listener) ExitRuleGroup(ctx *RuleGroupContext) {}

// EnterRulePlane is called when production rulePlane is entered.
func (s *Baseasn1Listener) EnterRulePlane(ctx *RulePlaneContext) {}

// ExitRulePlane is called when production rulePlane is exited.
func (s *Baseasn1Listener) ExitRulePlane(ctx *RulePlaneContext) {}

// EnterRuleRow is called when production ruleRow is entered.
func (s *Baseasn1Listener) EnterRuleRow(ctx *RuleRowContext) {}

// ExitRuleRow is called when production ruleRow is exited.
func (s *Baseasn1Listener) ExitRuleRow(ctx *RuleRowContext) {}

// EnterRuleCell is called when production ruleCell is entered.
func (s *Baseasn1Listener) EnterRuleCell(ctx *RuleCellContext) {}

// ExitRuleCell is called when production ruleCell is exited.
func (s *Baseasn1Listener) ExitRuleCell(ctx *RuleCellContext) {}

// EnterRuleTableColumn is called when production ruleTableColumn is entered.
func (s *Baseasn1Listener) EnterRuleTableColumn(ctx *RuleTableColumnContext) {}

// ExitRuleTableColumn is called when production ruleTableColumn is exited.
func (s *Baseasn1Listener) ExitRuleTableColumn(ctx *RuleTableColumnContext) {}

// EnterRuleTableRow is called when production ruleTableRow is entered.
func (s *Baseasn1Listener) EnterRuleTableRow(ctx *RuleTableRowContext) {}

// ExitRuleTableRow is called when production ruleTableRow is exited.
func (s *Baseasn1Listener) ExitRuleTableRow(ctx *RuleTableRowContext) {}

// EnterRuleCharacterStringList is called when production ruleCharacterStringList is entered.
func (s *Baseasn1Listener) EnterRuleCharacterStringList(ctx *RuleCharacterStringListContext) {}

// ExitRuleCharacterStringList is called when production ruleCharacterStringList is exited.
func (s *Baseasn1Listener) ExitRuleCharacterStringList(ctx *RuleCharacterStringListContext) {}

// EnterRuleQuadruple is called when production ruleQuadruple is entered.
func (s *Baseasn1Listener) EnterRuleQuadruple(ctx *RuleQuadrupleContext) {}

// ExitRuleQuadruple is called when production ruleQuadruple is exited.
func (s *Baseasn1Listener) ExitRuleQuadruple(ctx *RuleQuadrupleContext) {}

// EnterRuleTuple is called when production ruleTuple is entered.
func (s *Baseasn1Listener) EnterRuleTuple(ctx *RuleTupleContext) {}

// ExitRuleTuple is called when production ruleTuple is exited.
func (s *Baseasn1Listener) ExitRuleTuple(ctx *RuleTupleContext) {}

// EnterRuleRestrictedCharacterStringValue is called when production ruleRestrictedCharacterStringValue is entered.
func (s *Baseasn1Listener) EnterRuleRestrictedCharacterStringValue(ctx *RuleRestrictedCharacterStringValueContext) {
}

// ExitRuleRestrictedCharacterStringValue is called when production ruleRestrictedCharacterStringValue is exited.
func (s *Baseasn1Listener) ExitRuleRestrictedCharacterStringValue(ctx *RuleRestrictedCharacterStringValueContext) {
}

// EnterRuleUnrestrictedCharacterStringValue is called when production ruleUnrestrictedCharacterStringValue is entered.
func (s *Baseasn1Listener) EnterRuleUnrestrictedCharacterStringValue(ctx *RuleUnrestrictedCharacterStringValueContext) {
}

// ExitRuleUnrestrictedCharacterStringValue is called when production ruleUnrestrictedCharacterStringValue is exited.
func (s *Baseasn1Listener) ExitRuleUnrestrictedCharacterStringValue(ctx *RuleUnrestrictedCharacterStringValueContext) {
}

// EnterRuleNumericRealValue is called when production ruleNumericRealValue is entered.
func (s *Baseasn1Listener) EnterRuleNumericRealValue(ctx *RuleNumericRealValueContext) {}

// ExitRuleNumericRealValue is called when production ruleNumericRealValue is exited.
func (s *Baseasn1Listener) ExitRuleNumericRealValue(ctx *RuleNumericRealValueContext) {}

// EnterRuleSpecialRealValue is called when production ruleSpecialRealValue is entered.
func (s *Baseasn1Listener) EnterRuleSpecialRealValue(ctx *RuleSpecialRealValueContext) {}

// ExitRuleSpecialRealValue is called when production ruleSpecialRealValue is exited.
func (s *Baseasn1Listener) ExitRuleSpecialRealValue(ctx *RuleSpecialRealValueContext) {}

// EnterRuleRelativeOIDComponents is called when production ruleRelativeOIDComponents is entered.
func (s *Baseasn1Listener) EnterRuleRelativeOIDComponents(ctx *RuleRelativeOIDComponentsContext) {}

// ExitRuleRelativeOIDComponents is called when production ruleRelativeOIDComponents is exited.
func (s *Baseasn1Listener) ExitRuleRelativeOIDComponents(ctx *RuleRelativeOIDComponentsContext) {}

// EnterRuleRelativeOIDComponentsList is called when production ruleRelativeOIDComponentsList is entered.
func (s *Baseasn1Listener) EnterRuleRelativeOIDComponentsList(ctx *RuleRelativeOIDComponentsListContext) {
}

// ExitRuleRelativeOIDComponentsList is called when production ruleRelativeOIDComponentsList is exited.
func (s *Baseasn1Listener) ExitRuleRelativeOIDComponentsList(ctx *RuleRelativeOIDComponentsListContext) {
}

// EnterRuleComponentValueList is called when production ruleComponentValueList is entered.
func (s *Baseasn1Listener) EnterRuleComponentValueList(ctx *RuleComponentValueListContext) {}

// ExitRuleComponentValueList is called when production ruleComponentValueList is exited.
func (s *Baseasn1Listener) ExitRuleComponentValueList(ctx *RuleComponentValueListContext) {}

// EnterRuleValueList is called when production ruleValueList is entered.
func (s *Baseasn1Listener) EnterRuleValueList(ctx *RuleValueListContext) {}

// ExitRuleValueList is called when production ruleValueList is exited.
func (s *Baseasn1Listener) ExitRuleValueList(ctx *RuleValueListContext) {}

// EnterRuleNamedValue is called when production ruleNamedValue is entered.
func (s *Baseasn1Listener) EnterRuleNamedValue(ctx *RuleNamedValueContext) {}

// ExitRuleNamedValue is called when production ruleNamedValue is exited.
func (s *Baseasn1Listener) ExitRuleNamedValue(ctx *RuleNamedValueContext) {}

// EnterRuleNamedValueList is called when production ruleNamedValueList is entered.
func (s *Baseasn1Listener) EnterRuleNamedValueList(ctx *RuleNamedValueListContext) {}

// ExitRuleNamedValueList is called when production ruleNamedValueList is exited.
func (s *Baseasn1Listener) ExitRuleNamedValueList(ctx *RuleNamedValueListContext) {}

// EnterRuleBitStringValue is called when production ruleBitStringValue is entered.
func (s *Baseasn1Listener) EnterRuleBitStringValue(ctx *RuleBitStringValueContext) {}

// ExitRuleBitStringValue is called when production ruleBitStringValue is exited.
func (s *Baseasn1Listener) ExitRuleBitStringValue(ctx *RuleBitStringValueContext) {}

// EnterRuleBooleanValue is called when production ruleBooleanValue is entered.
func (s *Baseasn1Listener) EnterRuleBooleanValue(ctx *RuleBooleanValueContext) {}

// ExitRuleBooleanValue is called when production ruleBooleanValue is exited.
func (s *Baseasn1Listener) ExitRuleBooleanValue(ctx *RuleBooleanValueContext) {}

// EnterRuleCharacterStringValue is called when production ruleCharacterStringValue is entered.
func (s *Baseasn1Listener) EnterRuleCharacterStringValue(ctx *RuleCharacterStringValueContext) {}

// ExitRuleCharacterStringValue is called when production ruleCharacterStringValue is exited.
func (s *Baseasn1Listener) ExitRuleCharacterStringValue(ctx *RuleCharacterStringValueContext) {}

// EnterRuleChoiceValue is called when production ruleChoiceValue is entered.
func (s *Baseasn1Listener) EnterRuleChoiceValue(ctx *RuleChoiceValueContext) {}

// ExitRuleChoiceValue is called when production ruleChoiceValue is exited.
func (s *Baseasn1Listener) ExitRuleChoiceValue(ctx *RuleChoiceValueContext) {}

// EnterRuleEmbeddedPDVValue is called when production ruleEmbeddedPDVValue is entered.
func (s *Baseasn1Listener) EnterRuleEmbeddedPDVValue(ctx *RuleEmbeddedPDVValueContext) {}

// ExitRuleEmbeddedPDVValue is called when production ruleEmbeddedPDVValue is exited.
func (s *Baseasn1Listener) ExitRuleEmbeddedPDVValue(ctx *RuleEmbeddedPDVValueContext) {}

// EnterRuleEnumeratedValue is called when production ruleEnumeratedValue is entered.
func (s *Baseasn1Listener) EnterRuleEnumeratedValue(ctx *RuleEnumeratedValueContext) {}

// ExitRuleEnumeratedValue is called when production ruleEnumeratedValue is exited.
func (s *Baseasn1Listener) ExitRuleEnumeratedValue(ctx *RuleEnumeratedValueContext) {}

// EnterRuleExternalValue is called when production ruleExternalValue is entered.
func (s *Baseasn1Listener) EnterRuleExternalValue(ctx *RuleExternalValueContext) {}

// ExitRuleExternalValue is called when production ruleExternalValue is exited.
func (s *Baseasn1Listener) ExitRuleExternalValue(ctx *RuleExternalValueContext) {}

// EnterRuleInstanceOfValue is called when production ruleInstanceOfValue is entered.
func (s *Baseasn1Listener) EnterRuleInstanceOfValue(ctx *RuleInstanceOfValueContext) {}

// ExitRuleInstanceOfValue is called when production ruleInstanceOfValue is exited.
func (s *Baseasn1Listener) ExitRuleInstanceOfValue(ctx *RuleInstanceOfValueContext) {}

// EnterRuleIntegerValue is called when production ruleIntegerValue is entered.
func (s *Baseasn1Listener) EnterRuleIntegerValue(ctx *RuleIntegerValueContext) {}

// ExitRuleIntegerValue is called when production ruleIntegerValue is exited.
func (s *Baseasn1Listener) ExitRuleIntegerValue(ctx *RuleIntegerValueContext) {}

// EnterRuleNullValue is called when production ruleNullValue is entered.
func (s *Baseasn1Listener) EnterRuleNullValue(ctx *RuleNullValueContext) {}

// ExitRuleNullValue is called when production ruleNullValue is exited.
func (s *Baseasn1Listener) ExitRuleNullValue(ctx *RuleNullValueContext) {}

// EnterRuleOctetStringValue is called when production ruleOctetStringValue is entered.
func (s *Baseasn1Listener) EnterRuleOctetStringValue(ctx *RuleOctetStringValueContext) {}

// ExitRuleOctetStringValue is called when production ruleOctetStringValue is exited.
func (s *Baseasn1Listener) ExitRuleOctetStringValue(ctx *RuleOctetStringValueContext) {}

// EnterRuleRealValue is called when production ruleRealValue is entered.
func (s *Baseasn1Listener) EnterRuleRealValue(ctx *RuleRealValueContext) {}

// ExitRuleRealValue is called when production ruleRealValue is exited.
func (s *Baseasn1Listener) ExitRuleRealValue(ctx *RuleRealValueContext) {}

// EnterRuleRelativeIRIValue is called when production ruleRelativeIRIValue is entered.
func (s *Baseasn1Listener) EnterRuleRelativeIRIValue(ctx *RuleRelativeIRIValueContext) {}

// ExitRuleRelativeIRIValue is called when production ruleRelativeIRIValue is exited.
func (s *Baseasn1Listener) ExitRuleRelativeIRIValue(ctx *RuleRelativeIRIValueContext) {}

// EnterRuleRelativeOIDValue is called when production ruleRelativeOIDValue is entered.
func (s *Baseasn1Listener) EnterRuleRelativeOIDValue(ctx *RuleRelativeOIDValueContext) {}

// ExitRuleRelativeOIDValue is called when production ruleRelativeOIDValue is exited.
func (s *Baseasn1Listener) ExitRuleRelativeOIDValue(ctx *RuleRelativeOIDValueContext) {}

// EnterRuleSequenceValue is called when production ruleSequenceValue is entered.
func (s *Baseasn1Listener) EnterRuleSequenceValue(ctx *RuleSequenceValueContext) {}

// ExitRuleSequenceValue is called when production ruleSequenceValue is exited.
func (s *Baseasn1Listener) ExitRuleSequenceValue(ctx *RuleSequenceValueContext) {}

// EnterRuleSequenceOfValue is called when production ruleSequenceOfValue is entered.
func (s *Baseasn1Listener) EnterRuleSequenceOfValue(ctx *RuleSequenceOfValueContext) {}

// ExitRuleSequenceOfValue is called when production ruleSequenceOfValue is exited.
func (s *Baseasn1Listener) ExitRuleSequenceOfValue(ctx *RuleSequenceOfValueContext) {}

// EnterRuleSetValue is called when production ruleSetValue is entered.
func (s *Baseasn1Listener) EnterRuleSetValue(ctx *RuleSetValueContext) {}

// ExitRuleSetValue is called when production ruleSetValue is exited.
func (s *Baseasn1Listener) ExitRuleSetValue(ctx *RuleSetValueContext) {}

// EnterRuleSetOfValue is called when production ruleSetOfValue is entered.
func (s *Baseasn1Listener) EnterRuleSetOfValue(ctx *RuleSetOfValueContext) {}

// ExitRuleSetOfValue is called when production ruleSetOfValue is exited.
func (s *Baseasn1Listener) ExitRuleSetOfValue(ctx *RuleSetOfValueContext) {}

// EnterRulePrefixedValue is called when production rulePrefixedValue is entered.
func (s *Baseasn1Listener) EnterRulePrefixedValue(ctx *RulePrefixedValueContext) {}

// ExitRulePrefixedValue is called when production rulePrefixedValue is exited.
func (s *Baseasn1Listener) ExitRulePrefixedValue(ctx *RulePrefixedValueContext) {}

// EnterRuleTimeValue is called when production ruleTimeValue is entered.
func (s *Baseasn1Listener) EnterRuleTimeValue(ctx *RuleTimeValueContext) {}

// ExitRuleTimeValue is called when production ruleTimeValue is exited.
func (s *Baseasn1Listener) ExitRuleTimeValue(ctx *RuleTimeValueContext) {}

// EnterRuleBuiltinValue is called when production ruleBuiltinValue is entered.
func (s *Baseasn1Listener) EnterRuleBuiltinValue(ctx *RuleBuiltinValueContext) {}

// ExitRuleBuiltinValue is called when production ruleBuiltinValue is exited.
func (s *Baseasn1Listener) ExitRuleBuiltinValue(ctx *RuleBuiltinValueContext) {}

// EnterRuleValueFromObject is called when production ruleValueFromObject is entered.
func (s *Baseasn1Listener) EnterRuleValueFromObject(ctx *RuleValueFromObjectContext) {}

// ExitRuleValueFromObject is called when production ruleValueFromObject is exited.
func (s *Baseasn1Listener) ExitRuleValueFromObject(ctx *RuleValueFromObjectContext) {}

// EnterRuleReferencedValue is called when production ruleReferencedValue is entered.
func (s *Baseasn1Listener) EnterRuleReferencedValue(ctx *RuleReferencedValueContext) {}

// ExitRuleReferencedValue is called when production ruleReferencedValue is exited.
func (s *Baseasn1Listener) ExitRuleReferencedValue(ctx *RuleReferencedValueContext) {}

// EnterRuleOpenTypeFieldVal is called when production ruleOpenTypeFieldVal is entered.
func (s *Baseasn1Listener) EnterRuleOpenTypeFieldVal(ctx *RuleOpenTypeFieldValContext) {}

// ExitRuleOpenTypeFieldVal is called when production ruleOpenTypeFieldVal is exited.
func (s *Baseasn1Listener) ExitRuleOpenTypeFieldVal(ctx *RuleOpenTypeFieldValContext) {}

// EnterRuleFixedTypeFieldVal is called when production ruleFixedTypeFieldVal is entered.
func (s *Baseasn1Listener) EnterRuleFixedTypeFieldVal(ctx *RuleFixedTypeFieldValContext) {}

// ExitRuleFixedTypeFieldVal is called when production ruleFixedTypeFieldVal is exited.
func (s *Baseasn1Listener) ExitRuleFixedTypeFieldVal(ctx *RuleFixedTypeFieldValContext) {}

// EnterRuleObjectClassFieldValue is called when production ruleObjectClassFieldValue is entered.
func (s *Baseasn1Listener) EnterRuleObjectClassFieldValue(ctx *RuleObjectClassFieldValueContext) {}

// ExitRuleObjectClassFieldValue is called when production ruleObjectClassFieldValue is exited.
func (s *Baseasn1Listener) ExitRuleObjectClassFieldValue(ctx *RuleObjectClassFieldValueContext) {}

// EnterRuleValue is called when production ruleValue is entered.
func (s *Baseasn1Listener) EnterRuleValue(ctx *RuleValueContext) {}

// ExitRuleValue is called when production ruleValue is exited.
func (s *Baseasn1Listener) ExitRuleValue(ctx *RuleValueContext) {}

// EnterRuleValueSet is called when production ruleValueSet is entered.
func (s *Baseasn1Listener) EnterRuleValueSet(ctx *RuleValueSetContext) {}

// ExitRuleValueSet is called when production ruleValueSet is exited.
func (s *Baseasn1Listener) ExitRuleValueSet(ctx *RuleValueSetContext) {}

// EnterRuleDefinedObjectClass is called when production ruleDefinedObjectClass is entered.
func (s *Baseasn1Listener) EnterRuleDefinedObjectClass(ctx *RuleDefinedObjectClassContext) {}

// ExitRuleDefinedObjectClass is called when production ruleDefinedObjectClass is exited.
func (s *Baseasn1Listener) ExitRuleDefinedObjectClass(ctx *RuleDefinedObjectClassContext) {}

// EnterRuleObject is called when production ruleObject is entered.
func (s *Baseasn1Listener) EnterRuleObject(ctx *RuleObjectContext) {}

// ExitRuleObject is called when production ruleObject is exited.
func (s *Baseasn1Listener) ExitRuleObject(ctx *RuleObjectContext) {}

// EnterRuleDefinedObject is called when production ruleDefinedObject is entered.
func (s *Baseasn1Listener) EnterRuleDefinedObject(ctx *RuleDefinedObjectContext) {}

// ExitRuleDefinedObject is called when production ruleDefinedObject is exited.
func (s *Baseasn1Listener) ExitRuleDefinedObject(ctx *RuleDefinedObjectContext) {}

// EnterRuleDefinedObjectSet is called when production ruleDefinedObjectSet is entered.
func (s *Baseasn1Listener) EnterRuleDefinedObjectSet(ctx *RuleDefinedObjectSetContext) {}

// ExitRuleDefinedObjectSet is called when production ruleDefinedObjectSet is exited.
func (s *Baseasn1Listener) ExitRuleDefinedObjectSet(ctx *RuleDefinedObjectSetContext) {}

// EnterRuleObjectDefn is called when production ruleObjectDefn is entered.
func (s *Baseasn1Listener) EnterRuleObjectDefn(ctx *RuleObjectDefnContext) {}

// ExitRuleObjectDefn is called when production ruleObjectDefn is exited.
func (s *Baseasn1Listener) ExitRuleObjectDefn(ctx *RuleObjectDefnContext) {}

// EnterRuleDefaultSyntax is called when production ruleDefaultSyntax is entered.
func (s *Baseasn1Listener) EnterRuleDefaultSyntax(ctx *RuleDefaultSyntaxContext) {}

// ExitRuleDefaultSyntax is called when production ruleDefaultSyntax is exited.
func (s *Baseasn1Listener) ExitRuleDefaultSyntax(ctx *RuleDefaultSyntaxContext) {}

// EnterRuleDefinedSyntax is called when production ruleDefinedSyntax is entered.
func (s *Baseasn1Listener) EnterRuleDefinedSyntax(ctx *RuleDefinedSyntaxContext) {}

// ExitRuleDefinedSyntax is called when production ruleDefinedSyntax is exited.
func (s *Baseasn1Listener) ExitRuleDefinedSyntax(ctx *RuleDefinedSyntaxContext) {}

// EnterRuleDefinedSyntaxTokenList is called when production ruleDefinedSyntaxTokenList is entered.
func (s *Baseasn1Listener) EnterRuleDefinedSyntaxTokenList(ctx *RuleDefinedSyntaxTokenListContext) {}

// ExitRuleDefinedSyntaxTokenList is called when production ruleDefinedSyntaxTokenList is exited.
func (s *Baseasn1Listener) ExitRuleDefinedSyntaxTokenList(ctx *RuleDefinedSyntaxTokenListContext) {}

// EnterRuleDefinedSyntaxToken is called when production ruleDefinedSyntaxToken is entered.
func (s *Baseasn1Listener) EnterRuleDefinedSyntaxToken(ctx *RuleDefinedSyntaxTokenContext) {}

// ExitRuleDefinedSyntaxToken is called when production ruleDefinedSyntaxToken is exited.
func (s *Baseasn1Listener) ExitRuleDefinedSyntaxToken(ctx *RuleDefinedSyntaxTokenContext) {}

// EnterRuleLiteral is called when production ruleLiteral is entered.
func (s *Baseasn1Listener) EnterRuleLiteral(ctx *RuleLiteralContext) {}

// ExitRuleLiteral is called when production ruleLiteral is exited.
func (s *Baseasn1Listener) ExitRuleLiteral(ctx *RuleLiteralContext) {}

// EnterRuleWord is called when production ruleWord is entered.
func (s *Baseasn1Listener) EnterRuleWord(ctx *RuleWordContext) {}

// ExitRuleWord is called when production ruleWord is exited.
func (s *Baseasn1Listener) ExitRuleWord(ctx *RuleWordContext) {}

// EnterRuleFieldSettingList is called when production ruleFieldSettingList is entered.
func (s *Baseasn1Listener) EnterRuleFieldSettingList(ctx *RuleFieldSettingListContext) {}

// ExitRuleFieldSettingList is called when production ruleFieldSettingList is exited.
func (s *Baseasn1Listener) ExitRuleFieldSettingList(ctx *RuleFieldSettingListContext) {}

// EnterRuleFieldSetting is called when production ruleFieldSetting is entered.
func (s *Baseasn1Listener) EnterRuleFieldSetting(ctx *RuleFieldSettingContext) {}

// ExitRuleFieldSetting is called when production ruleFieldSetting is exited.
func (s *Baseasn1Listener) ExitRuleFieldSetting(ctx *RuleFieldSettingContext) {}

// EnterRulePrimitiveFieldName is called when production rulePrimitiveFieldName is entered.
func (s *Baseasn1Listener) EnterRulePrimitiveFieldName(ctx *RulePrimitiveFieldNameContext) {}

// ExitRulePrimitiveFieldName is called when production rulePrimitiveFieldName is exited.
func (s *Baseasn1Listener) ExitRulePrimitiveFieldName(ctx *RulePrimitiveFieldNameContext) {}

// EnterRuleSetting is called when production ruleSetting is entered.
func (s *Baseasn1Listener) EnterRuleSetting(ctx *RuleSettingContext) {}

// ExitRuleSetting is called when production ruleSetting is exited.
func (s *Baseasn1Listener) ExitRuleSetting(ctx *RuleSettingContext) {}

// EnterRuleObjectFromObject is called when production ruleObjectFromObject is entered.
func (s *Baseasn1Listener) EnterRuleObjectFromObject(ctx *RuleObjectFromObjectContext) {}

// ExitRuleObjectFromObject is called when production ruleObjectFromObject is exited.
func (s *Baseasn1Listener) ExitRuleObjectFromObject(ctx *RuleObjectFromObjectContext) {}

// EnterRuleObjectSetFromObjects is called when production ruleObjectSetFromObjects is entered.
func (s *Baseasn1Listener) EnterRuleObjectSetFromObjects(ctx *RuleObjectSetFromObjectsContext) {}

// ExitRuleObjectSetFromObjects is called when production ruleObjectSetFromObjects is exited.
func (s *Baseasn1Listener) ExitRuleObjectSetFromObjects(ctx *RuleObjectSetFromObjectsContext) {}

// EnterRuleReferencedObjects is called when production ruleReferencedObjects is entered.
func (s *Baseasn1Listener) EnterRuleReferencedObjects(ctx *RuleReferencedObjectsContext) {}

// ExitRuleReferencedObjects is called when production ruleReferencedObjects is exited.
func (s *Baseasn1Listener) ExitRuleReferencedObjects(ctx *RuleReferencedObjectsContext) {}

// EnterRuleFieldName is called when production ruleFieldName is entered.
func (s *Baseasn1Listener) EnterRuleFieldName(ctx *RuleFieldNameContext) {}

// ExitRuleFieldName is called when production ruleFieldName is exited.
func (s *Baseasn1Listener) ExitRuleFieldName(ctx *RuleFieldNameContext) {}

// EnterRuleParameterizedObject is called when production ruleParameterizedObject is entered.
func (s *Baseasn1Listener) EnterRuleParameterizedObject(ctx *RuleParameterizedObjectContext) {}

// ExitRuleParameterizedObject is called when production ruleParameterizedObject is exited.
func (s *Baseasn1Listener) ExitRuleParameterizedObject(ctx *RuleParameterizedObjectContext) {}

// EnterRuleObjectClass is called when production ruleObjectClass is entered.
func (s *Baseasn1Listener) EnterRuleObjectClass(ctx *RuleObjectClassContext) {}

// ExitRuleObjectClass is called when production ruleObjectClass is exited.
func (s *Baseasn1Listener) ExitRuleObjectClass(ctx *RuleObjectClassContext) {}

// EnterRuleObjectClassDefn is called when production ruleObjectClassDefn is entered.
func (s *Baseasn1Listener) EnterRuleObjectClassDefn(ctx *RuleObjectClassDefnContext) {}

// ExitRuleObjectClassDefn is called when production ruleObjectClassDefn is exited.
func (s *Baseasn1Listener) ExitRuleObjectClassDefn(ctx *RuleObjectClassDefnContext) {}

// EnterRuleFieldSpecList is called when production ruleFieldSpecList is entered.
func (s *Baseasn1Listener) EnterRuleFieldSpecList(ctx *RuleFieldSpecListContext) {}

// ExitRuleFieldSpecList is called when production ruleFieldSpecList is exited.
func (s *Baseasn1Listener) ExitRuleFieldSpecList(ctx *RuleFieldSpecListContext) {}

// EnterRuleFieldSpec is called when production ruleFieldSpec is entered.
func (s *Baseasn1Listener) EnterRuleFieldSpec(ctx *RuleFieldSpecContext) {}

// ExitRuleFieldSpec is called when production ruleFieldSpec is exited.
func (s *Baseasn1Listener) ExitRuleFieldSpec(ctx *RuleFieldSpecContext) {}

// EnterRuleTypeFieldSpec is called when production ruleTypeFieldSpec is entered.
func (s *Baseasn1Listener) EnterRuleTypeFieldSpec(ctx *RuleTypeFieldSpecContext) {}

// ExitRuleTypeFieldSpec is called when production ruleTypeFieldSpec is exited.
func (s *Baseasn1Listener) ExitRuleTypeFieldSpec(ctx *RuleTypeFieldSpecContext) {}

// EnterRuleFixedTypeValueFieldSpec is called when production ruleFixedTypeValueFieldSpec is entered.
func (s *Baseasn1Listener) EnterRuleFixedTypeValueFieldSpec(ctx *RuleFixedTypeValueFieldSpecContext) {
}

// ExitRuleFixedTypeValueFieldSpec is called when production ruleFixedTypeValueFieldSpec is exited.
func (s *Baseasn1Listener) ExitRuleFixedTypeValueFieldSpec(ctx *RuleFixedTypeValueFieldSpecContext) {}

// EnterRuleUnique is called when production ruleUnique is entered.
func (s *Baseasn1Listener) EnterRuleUnique(ctx *RuleUniqueContext) {}

// ExitRuleUnique is called when production ruleUnique is exited.
func (s *Baseasn1Listener) ExitRuleUnique(ctx *RuleUniqueContext) {}

// EnterRuleVariableTypeValueFieldSpec is called when production ruleVariableTypeValueFieldSpec is entered.
func (s *Baseasn1Listener) EnterRuleVariableTypeValueFieldSpec(ctx *RuleVariableTypeValueFieldSpecContext) {
}

// ExitRuleVariableTypeValueFieldSpec is called when production ruleVariableTypeValueFieldSpec is exited.
func (s *Baseasn1Listener) ExitRuleVariableTypeValueFieldSpec(ctx *RuleVariableTypeValueFieldSpecContext) {
}

// EnterRuleFixedTypeValueSetFieldSpec is called when production ruleFixedTypeValueSetFieldSpec is entered.
func (s *Baseasn1Listener) EnterRuleFixedTypeValueSetFieldSpec(ctx *RuleFixedTypeValueSetFieldSpecContext) {
}

// ExitRuleFixedTypeValueSetFieldSpec is called when production ruleFixedTypeValueSetFieldSpec is exited.
func (s *Baseasn1Listener) ExitRuleFixedTypeValueSetFieldSpec(ctx *RuleFixedTypeValueSetFieldSpecContext) {
}

// EnterRuleVariableTypeValueSetFieldSpec is called when production ruleVariableTypeValueSetFieldSpec is entered.
func (s *Baseasn1Listener) EnterRuleVariableTypeValueSetFieldSpec(ctx *RuleVariableTypeValueSetFieldSpecContext) {
}

// ExitRuleVariableTypeValueSetFieldSpec is called when production ruleVariableTypeValueSetFieldSpec is exited.
func (s *Baseasn1Listener) ExitRuleVariableTypeValueSetFieldSpec(ctx *RuleVariableTypeValueSetFieldSpecContext) {
}

// EnterRuleObjectFieldSpec is called when production ruleObjectFieldSpec is entered.
func (s *Baseasn1Listener) EnterRuleObjectFieldSpec(ctx *RuleObjectFieldSpecContext) {}

// ExitRuleObjectFieldSpec is called when production ruleObjectFieldSpec is exited.
func (s *Baseasn1Listener) ExitRuleObjectFieldSpec(ctx *RuleObjectFieldSpecContext) {}

// EnterRuleObjectSetFieldSpec is called when production ruleObjectSetFieldSpec is entered.
func (s *Baseasn1Listener) EnterRuleObjectSetFieldSpec(ctx *RuleObjectSetFieldSpecContext) {}

// ExitRuleObjectSetFieldSpec is called when production ruleObjectSetFieldSpec is exited.
func (s *Baseasn1Listener) ExitRuleObjectSetFieldSpec(ctx *RuleObjectSetFieldSpecContext) {}

// EnterRuleTypeOptionalitySpec is called when production ruleTypeOptionalitySpec is entered.
func (s *Baseasn1Listener) EnterRuleTypeOptionalitySpec(ctx *RuleTypeOptionalitySpecContext) {}

// ExitRuleTypeOptionalitySpec is called when production ruleTypeOptionalitySpec is exited.
func (s *Baseasn1Listener) ExitRuleTypeOptionalitySpec(ctx *RuleTypeOptionalitySpecContext) {}

// EnterRuleValueOptionalitySpec is called when production ruleValueOptionalitySpec is entered.
func (s *Baseasn1Listener) EnterRuleValueOptionalitySpec(ctx *RuleValueOptionalitySpecContext) {}

// ExitRuleValueOptionalitySpec is called when production ruleValueOptionalitySpec is exited.
func (s *Baseasn1Listener) ExitRuleValueOptionalitySpec(ctx *RuleValueOptionalitySpecContext) {}

// EnterRuleValueSetOptionalitySpec is called when production ruleValueSetOptionalitySpec is entered.
func (s *Baseasn1Listener) EnterRuleValueSetOptionalitySpec(ctx *RuleValueSetOptionalitySpecContext) {
}

// ExitRuleValueSetOptionalitySpec is called when production ruleValueSetOptionalitySpec is exited.
func (s *Baseasn1Listener) ExitRuleValueSetOptionalitySpec(ctx *RuleValueSetOptionalitySpecContext) {}

// EnterRuleObjectOptionalitySpec is called when production ruleObjectOptionalitySpec is entered.
func (s *Baseasn1Listener) EnterRuleObjectOptionalitySpec(ctx *RuleObjectOptionalitySpecContext) {}

// ExitRuleObjectOptionalitySpec is called when production ruleObjectOptionalitySpec is exited.
func (s *Baseasn1Listener) ExitRuleObjectOptionalitySpec(ctx *RuleObjectOptionalitySpecContext) {}

// EnterRuleObjectSetOptionalitySpec is called when production ruleObjectSetOptionalitySpec is entered.
func (s *Baseasn1Listener) EnterRuleObjectSetOptionalitySpec(ctx *RuleObjectSetOptionalitySpecContext) {
}

// ExitRuleObjectSetOptionalitySpec is called when production ruleObjectSetOptionalitySpec is exited.
func (s *Baseasn1Listener) ExitRuleObjectSetOptionalitySpec(ctx *RuleObjectSetOptionalitySpecContext) {
}

// EnterRuleWithSyntaxSpec is called when production ruleWithSyntaxSpec is entered.
func (s *Baseasn1Listener) EnterRuleWithSyntaxSpec(ctx *RuleWithSyntaxSpecContext) {}

// ExitRuleWithSyntaxSpec is called when production ruleWithSyntaxSpec is exited.
func (s *Baseasn1Listener) ExitRuleWithSyntaxSpec(ctx *RuleWithSyntaxSpecContext) {}

// EnterRuleSyntaxList is called when production ruleSyntaxList is entered.
func (s *Baseasn1Listener) EnterRuleSyntaxList(ctx *RuleSyntaxListContext) {}

// ExitRuleSyntaxList is called when production ruleSyntaxList is exited.
func (s *Baseasn1Listener) ExitRuleSyntaxList(ctx *RuleSyntaxListContext) {}

// EnterRuleTokenOrGroupSpecList is called when production ruleTokenOrGroupSpecList is entered.
func (s *Baseasn1Listener) EnterRuleTokenOrGroupSpecList(ctx *RuleTokenOrGroupSpecListContext) {}

// ExitRuleTokenOrGroupSpecList is called when production ruleTokenOrGroupSpecList is exited.
func (s *Baseasn1Listener) ExitRuleTokenOrGroupSpecList(ctx *RuleTokenOrGroupSpecListContext) {}

// EnterRuleTokenOrGroupSpec is called when production ruleTokenOrGroupSpec is entered.
func (s *Baseasn1Listener) EnterRuleTokenOrGroupSpec(ctx *RuleTokenOrGroupSpecContext) {}

// ExitRuleTokenOrGroupSpec is called when production ruleTokenOrGroupSpec is exited.
func (s *Baseasn1Listener) ExitRuleTokenOrGroupSpec(ctx *RuleTokenOrGroupSpecContext) {}

// EnterRuleRequiredToken is called when production ruleRequiredToken is entered.
func (s *Baseasn1Listener) EnterRuleRequiredToken(ctx *RuleRequiredTokenContext) {}

// ExitRuleRequiredToken is called when production ruleRequiredToken is exited.
func (s *Baseasn1Listener) ExitRuleRequiredToken(ctx *RuleRequiredTokenContext) {}

// EnterRuleOptionalGroup is called when production ruleOptionalGroup is entered.
func (s *Baseasn1Listener) EnterRuleOptionalGroup(ctx *RuleOptionalGroupContext) {}

// ExitRuleOptionalGroup is called when production ruleOptionalGroup is exited.
func (s *Baseasn1Listener) ExitRuleOptionalGroup(ctx *RuleOptionalGroupContext) {}

// EnterRuleParameterizedObjectClass is called when production ruleParameterizedObjectClass is entered.
func (s *Baseasn1Listener) EnterRuleParameterizedObjectClass(ctx *RuleParameterizedObjectClassContext) {
}

// ExitRuleParameterizedObjectClass is called when production ruleParameterizedObjectClass is exited.
func (s *Baseasn1Listener) ExitRuleParameterizedObjectClass(ctx *RuleParameterizedObjectClassContext) {
}

// EnterRuleObjectSet is called when production ruleObjectSet is entered.
func (s *Baseasn1Listener) EnterRuleObjectSet(ctx *RuleObjectSetContext) {}

// ExitRuleObjectSet is called when production ruleObjectSet is exited.
func (s *Baseasn1Listener) ExitRuleObjectSet(ctx *RuleObjectSetContext) {}

// EnterRuleObjectSetSpec is called when production ruleObjectSetSpec is entered.
func (s *Baseasn1Listener) EnterRuleObjectSetSpec(ctx *RuleObjectSetSpecContext) {}

// ExitRuleObjectSetSpec is called when production ruleObjectSetSpec is exited.
func (s *Baseasn1Listener) ExitRuleObjectSetSpec(ctx *RuleObjectSetSpecContext) {}

// EnterRuleRootElementSetSpec is called when production ruleRootElementSetSpec is entered.
func (s *Baseasn1Listener) EnterRuleRootElementSetSpec(ctx *RuleRootElementSetSpecContext) {}

// ExitRuleRootElementSetSpec is called when production ruleRootElementSetSpec is exited.
func (s *Baseasn1Listener) ExitRuleRootElementSetSpec(ctx *RuleRootElementSetSpecContext) {}

// EnterRuleAdditionalElementSetSpec is called when production ruleAdditionalElementSetSpec is entered.
func (s *Baseasn1Listener) EnterRuleAdditionalElementSetSpec(ctx *RuleAdditionalElementSetSpecContext) {
}

// ExitRuleAdditionalElementSetSpec is called when production ruleAdditionalElementSetSpec is exited.
func (s *Baseasn1Listener) ExitRuleAdditionalElementSetSpec(ctx *RuleAdditionalElementSetSpecContext) {
}

// EnterRuleElementSetSpecs is called when production ruleElementSetSpecs is entered.
func (s *Baseasn1Listener) EnterRuleElementSetSpecs(ctx *RuleElementSetSpecsContext) {}

// ExitRuleElementSetSpecs is called when production ruleElementSetSpecs is exited.
func (s *Baseasn1Listener) ExitRuleElementSetSpecs(ctx *RuleElementSetSpecsContext) {}

// EnterRuleElementSetSpec is called when production ruleElementSetSpec is entered.
func (s *Baseasn1Listener) EnterRuleElementSetSpec(ctx *RuleElementSetSpecContext) {}

// ExitRuleElementSetSpec is called when production ruleElementSetSpec is exited.
func (s *Baseasn1Listener) ExitRuleElementSetSpec(ctx *RuleElementSetSpecContext) {}

// EnterRuleUnions is called when production ruleUnions is entered.
func (s *Baseasn1Listener) EnterRuleUnions(ctx *RuleUnionsContext) {}

// ExitRuleUnions is called when production ruleUnions is exited.
func (s *Baseasn1Listener) ExitRuleUnions(ctx *RuleUnionsContext) {}

// EnterRuleExclusions is called when production ruleExclusions is entered.
func (s *Baseasn1Listener) EnterRuleExclusions(ctx *RuleExclusionsContext) {}

// ExitRuleExclusions is called when production ruleExclusions is exited.
func (s *Baseasn1Listener) ExitRuleExclusions(ctx *RuleExclusionsContext) {}

// EnterRuleIntersections is called when production ruleIntersections is entered.
func (s *Baseasn1Listener) EnterRuleIntersections(ctx *RuleIntersectionsContext) {}

// ExitRuleIntersections is called when production ruleIntersections is exited.
func (s *Baseasn1Listener) ExitRuleIntersections(ctx *RuleIntersectionsContext) {}

// EnterRuleUElems is called when production ruleUElems is entered.
func (s *Baseasn1Listener) EnterRuleUElems(ctx *RuleUElemsContext) {}

// ExitRuleUElems is called when production ruleUElems is exited.
func (s *Baseasn1Listener) ExitRuleUElems(ctx *RuleUElemsContext) {}

// EnterRuleUnionMark is called when production ruleUnionMark is entered.
func (s *Baseasn1Listener) EnterRuleUnionMark(ctx *RuleUnionMarkContext) {}

// ExitRuleUnionMark is called when production ruleUnionMark is exited.
func (s *Baseasn1Listener) ExitRuleUnionMark(ctx *RuleUnionMarkContext) {}

// EnterRuleLowerEndValue is called when production ruleLowerEndValue is entered.
func (s *Baseasn1Listener) EnterRuleLowerEndValue(ctx *RuleLowerEndValueContext) {}

// ExitRuleLowerEndValue is called when production ruleLowerEndValue is exited.
func (s *Baseasn1Listener) ExitRuleLowerEndValue(ctx *RuleLowerEndValueContext) {}

// EnterRuleUpperEndValue is called when production ruleUpperEndValue is entered.
func (s *Baseasn1Listener) EnterRuleUpperEndValue(ctx *RuleUpperEndValueContext) {}

// ExitRuleUpperEndValue is called when production ruleUpperEndValue is exited.
func (s *Baseasn1Listener) ExitRuleUpperEndValue(ctx *RuleUpperEndValueContext) {}

// EnterRuleIncludes is called when production ruleIncludes is entered.
func (s *Baseasn1Listener) EnterRuleIncludes(ctx *RuleIncludesContext) {}

// ExitRuleIncludes is called when production ruleIncludes is exited.
func (s *Baseasn1Listener) ExitRuleIncludes(ctx *RuleIncludesContext) {}

// EnterRuleLowerEndpoint is called when production ruleLowerEndpoint is entered.
func (s *Baseasn1Listener) EnterRuleLowerEndpoint(ctx *RuleLowerEndpointContext) {}

// ExitRuleLowerEndpoint is called when production ruleLowerEndpoint is exited.
func (s *Baseasn1Listener) ExitRuleLowerEndpoint(ctx *RuleLowerEndpointContext) {}

// EnterRuleUpperEndpoint is called when production ruleUpperEndpoint is entered.
func (s *Baseasn1Listener) EnterRuleUpperEndpoint(ctx *RuleUpperEndpointContext) {}

// ExitRuleUpperEndpoint is called when production ruleUpperEndpoint is exited.
func (s *Baseasn1Listener) ExitRuleUpperEndpoint(ctx *RuleUpperEndpointContext) {}

// EnterRuleLevel is called when production ruleLevel is entered.
func (s *Baseasn1Listener) EnterRuleLevel(ctx *RuleLevelContext) {}

// ExitRuleLevel is called when production ruleLevel is exited.
func (s *Baseasn1Listener) ExitRuleLevel(ctx *RuleLevelContext) {}

// EnterRuleComponentIdList is called when production ruleComponentIdList is entered.
func (s *Baseasn1Listener) EnterRuleComponentIdList(ctx *RuleComponentIdListContext) {}

// ExitRuleComponentIdList is called when production ruleComponentIdList is exited.
func (s *Baseasn1Listener) ExitRuleComponentIdList(ctx *RuleComponentIdListContext) {}

// EnterRuleAtNotation is called when production ruleAtNotation is entered.
func (s *Baseasn1Listener) EnterRuleAtNotation(ctx *RuleAtNotationContext) {}

// ExitRuleAtNotation is called when production ruleAtNotation is exited.
func (s *Baseasn1Listener) ExitRuleAtNotation(ctx *RuleAtNotationContext) {}

// EnterRuleAtNotationList is called when production ruleAtNotationList is entered.
func (s *Baseasn1Listener) EnterRuleAtNotationList(ctx *RuleAtNotationListContext) {}

// ExitRuleAtNotationList is called when production ruleAtNotationList is exited.
func (s *Baseasn1Listener) ExitRuleAtNotationList(ctx *RuleAtNotationListContext) {}

// EnterRuleUserDefinedConstraintParameter is called when production ruleUserDefinedConstraintParameter is entered.
func (s *Baseasn1Listener) EnterRuleUserDefinedConstraintParameter(ctx *RuleUserDefinedConstraintParameterContext) {
}

// ExitRuleUserDefinedConstraintParameter is called when production ruleUserDefinedConstraintParameter is exited.
func (s *Baseasn1Listener) ExitRuleUserDefinedConstraintParameter(ctx *RuleUserDefinedConstraintParameterContext) {
}

// EnterRuleUserDefinedConstraintParameterList is called when production ruleUserDefinedConstraintParameterList is entered.
func (s *Baseasn1Listener) EnterRuleUserDefinedConstraintParameterList(ctx *RuleUserDefinedConstraintParameterListContext) {
}

// ExitRuleUserDefinedConstraintParameterList is called when production ruleUserDefinedConstraintParameterList is exited.
func (s *Baseasn1Listener) ExitRuleUserDefinedConstraintParameterList(ctx *RuleUserDefinedConstraintParameterListContext) {
}

// EnterRuleUserDefinedConstraint is called when production ruleUserDefinedConstraint is entered.
func (s *Baseasn1Listener) EnterRuleUserDefinedConstraint(ctx *RuleUserDefinedConstraintContext) {}

// ExitRuleUserDefinedConstraint is called when production ruleUserDefinedConstraint is exited.
func (s *Baseasn1Listener) ExitRuleUserDefinedConstraint(ctx *RuleUserDefinedConstraintContext) {}

// EnterRuleSimpleTableConstraint is called when production ruleSimpleTableConstraint is entered.
func (s *Baseasn1Listener) EnterRuleSimpleTableConstraint(ctx *RuleSimpleTableConstraintContext) {}

// ExitRuleSimpleTableConstraint is called when production ruleSimpleTableConstraint is exited.
func (s *Baseasn1Listener) ExitRuleSimpleTableConstraint(ctx *RuleSimpleTableConstraintContext) {}

// EnterRuleComponentRelationConstraint is called when production ruleComponentRelationConstraint is entered.
func (s *Baseasn1Listener) EnterRuleComponentRelationConstraint(ctx *RuleComponentRelationConstraintContext) {
}

// ExitRuleComponentRelationConstraint is called when production ruleComponentRelationConstraint is exited.
func (s *Baseasn1Listener) ExitRuleComponentRelationConstraint(ctx *RuleComponentRelationConstraintContext) {
}

// EnterRuleTableConstraint is called when production ruleTableConstraint is entered.
func (s *Baseasn1Listener) EnterRuleTableConstraint(ctx *RuleTableConstraintContext) {}

// ExitRuleTableConstraint is called when production ruleTableConstraint is exited.
func (s *Baseasn1Listener) ExitRuleTableConstraint(ctx *RuleTableConstraintContext) {}

// EnterRuleContentsConstraint is called when production ruleContentsConstraint is entered.
func (s *Baseasn1Listener) EnterRuleContentsConstraint(ctx *RuleContentsConstraintContext) {}

// ExitRuleContentsConstraint is called when production ruleContentsConstraint is exited.
func (s *Baseasn1Listener) ExitRuleContentsConstraint(ctx *RuleContentsConstraintContext) {}

// EnterRuleSubtypeConstraint is called when production ruleSubtypeConstraint is entered.
func (s *Baseasn1Listener) EnterRuleSubtypeConstraint(ctx *RuleSubtypeConstraintContext) {}

// ExitRuleSubtypeConstraint is called when production ruleSubtypeConstraint is exited.
func (s *Baseasn1Listener) ExitRuleSubtypeConstraint(ctx *RuleSubtypeConstraintContext) {}

// EnterRuleGeneralConstraint is called when production ruleGeneralConstraint is entered.
func (s *Baseasn1Listener) EnterRuleGeneralConstraint(ctx *RuleGeneralConstraintContext) {}

// ExitRuleGeneralConstraint is called when production ruleGeneralConstraint is exited.
func (s *Baseasn1Listener) ExitRuleGeneralConstraint(ctx *RuleGeneralConstraintContext) {}

// EnterRuleConstraintSpec is called when production ruleConstraintSpec is entered.
func (s *Baseasn1Listener) EnterRuleConstraintSpec(ctx *RuleConstraintSpecContext) {}

// ExitRuleConstraintSpec is called when production ruleConstraintSpec is exited.
func (s *Baseasn1Listener) ExitRuleConstraintSpec(ctx *RuleConstraintSpecContext) {}

// EnterRuleSignedNumber is called when production ruleSignedNumber is entered.
func (s *Baseasn1Listener) EnterRuleSignedNumber(ctx *RuleSignedNumberContext) {}

// ExitRuleSignedNumber is called when production ruleSignedNumber is exited.
func (s *Baseasn1Listener) ExitRuleSignedNumber(ctx *RuleSignedNumberContext) {}

// EnterRuleExceptionIdentification is called when production ruleExceptionIdentification is entered.
func (s *Baseasn1Listener) EnterRuleExceptionIdentification(ctx *RuleExceptionIdentificationContext) {
}

// ExitRuleExceptionIdentification is called when production ruleExceptionIdentification is exited.
func (s *Baseasn1Listener) ExitRuleExceptionIdentification(ctx *RuleExceptionIdentificationContext) {}

// EnterRuleExceptionSpec is called when production ruleExceptionSpec is entered.
func (s *Baseasn1Listener) EnterRuleExceptionSpec(ctx *RuleExceptionSpecContext) {}

// ExitRuleExceptionSpec is called when production ruleExceptionSpec is exited.
func (s *Baseasn1Listener) ExitRuleExceptionSpec(ctx *RuleExceptionSpecContext) {}

// EnterRuleConstraint is called when production ruleConstraint is entered.
func (s *Baseasn1Listener) EnterRuleConstraint(ctx *RuleConstraintContext) {}

// ExitRuleConstraint is called when production ruleConstraint is exited.
func (s *Baseasn1Listener) ExitRuleConstraint(ctx *RuleConstraintContext) {}

// EnterRuleSingleTypeConstraint is called when production ruleSingleTypeConstraint is entered.
func (s *Baseasn1Listener) EnterRuleSingleTypeConstraint(ctx *RuleSingleTypeConstraintContext) {}

// ExitRuleSingleTypeConstraint is called when production ruleSingleTypeConstraint is exited.
func (s *Baseasn1Listener) ExitRuleSingleTypeConstraint(ctx *RuleSingleTypeConstraintContext) {}

// EnterRuleValueConstraint is called when production ruleValueConstraint is entered.
func (s *Baseasn1Listener) EnterRuleValueConstraint(ctx *RuleValueConstraintContext) {}

// ExitRuleValueConstraint is called when production ruleValueConstraint is exited.
func (s *Baseasn1Listener) ExitRuleValueConstraint(ctx *RuleValueConstraintContext) {}

// EnterRulePresenceConstraint is called when production rulePresenceConstraint is entered.
func (s *Baseasn1Listener) EnterRulePresenceConstraint(ctx *RulePresenceConstraintContext) {}

// ExitRulePresenceConstraint is called when production rulePresenceConstraint is exited.
func (s *Baseasn1Listener) ExitRulePresenceConstraint(ctx *RulePresenceConstraintContext) {}

// EnterRuleComponentConstraint is called when production ruleComponentConstraint is entered.
func (s *Baseasn1Listener) EnterRuleComponentConstraint(ctx *RuleComponentConstraintContext) {}

// ExitRuleComponentConstraint is called when production ruleComponentConstraint is exited.
func (s *Baseasn1Listener) ExitRuleComponentConstraint(ctx *RuleComponentConstraintContext) {}

// EnterRuleNamedConstraint is called when production ruleNamedConstraint is entered.
func (s *Baseasn1Listener) EnterRuleNamedConstraint(ctx *RuleNamedConstraintContext) {}

// ExitRuleNamedConstraint is called when production ruleNamedConstraint is exited.
func (s *Baseasn1Listener) ExitRuleNamedConstraint(ctx *RuleNamedConstraintContext) {}

// EnterRuleTypeConstraints is called when production ruleTypeConstraints is entered.
func (s *Baseasn1Listener) EnterRuleTypeConstraints(ctx *RuleTypeConstraintsContext) {}

// ExitRuleTypeConstraints is called when production ruleTypeConstraints is exited.
func (s *Baseasn1Listener) ExitRuleTypeConstraints(ctx *RuleTypeConstraintsContext) {}

// EnterRuleFullSpecification is called when production ruleFullSpecification is entered.
func (s *Baseasn1Listener) EnterRuleFullSpecification(ctx *RuleFullSpecificationContext) {}

// ExitRuleFullSpecification is called when production ruleFullSpecification is exited.
func (s *Baseasn1Listener) ExitRuleFullSpecification(ctx *RuleFullSpecificationContext) {}

// EnterRulePartialSpecification is called when production rulePartialSpecification is entered.
func (s *Baseasn1Listener) EnterRulePartialSpecification(ctx *RulePartialSpecificationContext) {}

// ExitRulePartialSpecification is called when production rulePartialSpecification is exited.
func (s *Baseasn1Listener) ExitRulePartialSpecification(ctx *RulePartialSpecificationContext) {}

// EnterRuleMultipleTypeConstraints is called when production ruleMultipleTypeConstraints is entered.
func (s *Baseasn1Listener) EnterRuleMultipleTypeConstraints(ctx *RuleMultipleTypeConstraintsContext) {
}

// ExitRuleMultipleTypeConstraints is called when production ruleMultipleTypeConstraints is exited.
func (s *Baseasn1Listener) ExitRuleMultipleTypeConstraints(ctx *RuleMultipleTypeConstraintsContext) {}

// EnterRuleSimpleString is called when production ruleSimpleString is entered.
func (s *Baseasn1Listener) EnterRuleSimpleString(ctx *RuleSimpleStringContext) {}

// ExitRuleSimpleString is called when production ruleSimpleString is exited.
func (s *Baseasn1Listener) ExitRuleSimpleString(ctx *RuleSimpleStringContext) {}

// EnterRuleSingleValue is called when production ruleSingleValue is entered.
func (s *Baseasn1Listener) EnterRuleSingleValue(ctx *RuleSingleValueContext) {}

// ExitRuleSingleValue is called when production ruleSingleValue is exited.
func (s *Baseasn1Listener) ExitRuleSingleValue(ctx *RuleSingleValueContext) {}

// EnterRuleContainedSubtype is called when production ruleContainedSubtype is entered.
func (s *Baseasn1Listener) EnterRuleContainedSubtype(ctx *RuleContainedSubtypeContext) {}

// ExitRuleContainedSubtype is called when production ruleContainedSubtype is exited.
func (s *Baseasn1Listener) ExitRuleContainedSubtype(ctx *RuleContainedSubtypeContext) {}

// EnterRuleValueRange is called when production ruleValueRange is entered.
func (s *Baseasn1Listener) EnterRuleValueRange(ctx *RuleValueRangeContext) {}

// ExitRuleValueRange is called when production ruleValueRange is exited.
func (s *Baseasn1Listener) ExitRuleValueRange(ctx *RuleValueRangeContext) {}

// EnterRulePermittedAlphabet is called when production rulePermittedAlphabet is entered.
func (s *Baseasn1Listener) EnterRulePermittedAlphabet(ctx *RulePermittedAlphabetContext) {}

// ExitRulePermittedAlphabet is called when production rulePermittedAlphabet is exited.
func (s *Baseasn1Listener) ExitRulePermittedAlphabet(ctx *RulePermittedAlphabetContext) {}

// EnterRuleSizeConstraint is called when production ruleSizeConstraint is entered.
func (s *Baseasn1Listener) EnterRuleSizeConstraint(ctx *RuleSizeConstraintContext) {}

// ExitRuleSizeConstraint is called when production ruleSizeConstraint is exited.
func (s *Baseasn1Listener) ExitRuleSizeConstraint(ctx *RuleSizeConstraintContext) {}

// EnterRuleInnerTypeConstraints is called when production ruleInnerTypeConstraints is entered.
func (s *Baseasn1Listener) EnterRuleInnerTypeConstraints(ctx *RuleInnerTypeConstraintsContext) {}

// ExitRuleInnerTypeConstraints is called when production ruleInnerTypeConstraints is exited.
func (s *Baseasn1Listener) ExitRuleInnerTypeConstraints(ctx *RuleInnerTypeConstraintsContext) {}

// EnterRulePatternConstraint is called when production rulePatternConstraint is entered.
func (s *Baseasn1Listener) EnterRulePatternConstraint(ctx *RulePatternConstraintContext) {}

// ExitRulePatternConstraint is called when production rulePatternConstraint is exited.
func (s *Baseasn1Listener) ExitRulePatternConstraint(ctx *RulePatternConstraintContext) {}

// EnterRulePropertySettings is called when production rulePropertySettings is entered.
func (s *Baseasn1Listener) EnterRulePropertySettings(ctx *RulePropertySettingsContext) {}

// ExitRulePropertySettings is called when production rulePropertySettings is exited.
func (s *Baseasn1Listener) ExitRulePropertySettings(ctx *RulePropertySettingsContext) {}

// EnterRuleDurationRange is called when production ruleDurationRange is entered.
func (s *Baseasn1Listener) EnterRuleDurationRange(ctx *RuleDurationRangeContext) {}

// ExitRuleDurationRange is called when production ruleDurationRange is exited.
func (s *Baseasn1Listener) ExitRuleDurationRange(ctx *RuleDurationRangeContext) {}

// EnterRuleTimePointRange is called when production ruleTimePointRange is entered.
func (s *Baseasn1Listener) EnterRuleTimePointRange(ctx *RuleTimePointRangeContext) {}

// ExitRuleTimePointRange is called when production ruleTimePointRange is exited.
func (s *Baseasn1Listener) ExitRuleTimePointRange(ctx *RuleTimePointRangeContext) {}

// EnterRuleRecurrenceRange is called when production ruleRecurrenceRange is entered.
func (s *Baseasn1Listener) EnterRuleRecurrenceRange(ctx *RuleRecurrenceRangeContext) {}

// ExitRuleRecurrenceRange is called when production ruleRecurrenceRange is exited.
func (s *Baseasn1Listener) ExitRuleRecurrenceRange(ctx *RuleRecurrenceRangeContext) {}

// EnterRuleSubtypeElements is called when production ruleSubtypeElements is entered.
func (s *Baseasn1Listener) EnterRuleSubtypeElements(ctx *RuleSubtypeElementsContext) {}

// ExitRuleSubtypeElements is called when production ruleSubtypeElements is exited.
func (s *Baseasn1Listener) ExitRuleSubtypeElements(ctx *RuleSubtypeElementsContext) {}

// EnterRuleObjectSetElements is called when production ruleObjectSetElements is entered.
func (s *Baseasn1Listener) EnterRuleObjectSetElements(ctx *RuleObjectSetElementsContext) {}

// ExitRuleObjectSetElements is called when production ruleObjectSetElements is exited.
func (s *Baseasn1Listener) ExitRuleObjectSetElements(ctx *RuleObjectSetElementsContext) {}

// EnterRuleElements is called when production ruleElements is entered.
func (s *Baseasn1Listener) EnterRuleElements(ctx *RuleElementsContext) {}

// ExitRuleElements is called when production ruleElements is exited.
func (s *Baseasn1Listener) ExitRuleElements(ctx *RuleElementsContext) {}

// EnterRuleElems is called when production ruleElems is entered.
func (s *Baseasn1Listener) EnterRuleElems(ctx *RuleElemsContext) {}

// ExitRuleElems is called when production ruleElems is exited.
func (s *Baseasn1Listener) ExitRuleElems(ctx *RuleElemsContext) {}

// EnterRuleIntersectionElements is called when production ruleIntersectionElements is entered.
func (s *Baseasn1Listener) EnterRuleIntersectionElements(ctx *RuleIntersectionElementsContext) {}

// ExitRuleIntersectionElements is called when production ruleIntersectionElements is exited.
func (s *Baseasn1Listener) ExitRuleIntersectionElements(ctx *RuleIntersectionElementsContext) {}

// EnterRuleIElems is called when production ruleIElems is entered.
func (s *Baseasn1Listener) EnterRuleIElems(ctx *RuleIElemsContext) {}

// ExitRuleIElems is called when production ruleIElems is exited.
func (s *Baseasn1Listener) ExitRuleIElems(ctx *RuleIElemsContext) {}

// EnterRuleIntersectionMark is called when production ruleIntersectionMark is entered.
func (s *Baseasn1Listener) EnterRuleIntersectionMark(ctx *RuleIntersectionMarkContext) {}

// ExitRuleIntersectionMark is called when production ruleIntersectionMark is exited.
func (s *Baseasn1Listener) ExitRuleIntersectionMark(ctx *RuleIntersectionMarkContext) {}

// EnterRuleSimpleDefinedValue is called when production ruleSimpleDefinedValue is entered.
func (s *Baseasn1Listener) EnterRuleSimpleDefinedValue(ctx *RuleSimpleDefinedValueContext) {}

// ExitRuleSimpleDefinedValue is called when production ruleSimpleDefinedValue is exited.
func (s *Baseasn1Listener) ExitRuleSimpleDefinedValue(ctx *RuleSimpleDefinedValueContext) {}

// EnterRuleActualParameterList is called when production ruleActualParameterList is entered.
func (s *Baseasn1Listener) EnterRuleActualParameterList(ctx *RuleActualParameterListContext) {}

// ExitRuleActualParameterList is called when production ruleActualParameterList is exited.
func (s *Baseasn1Listener) ExitRuleActualParameterList(ctx *RuleActualParameterListContext) {}

// EnterRuleActualParameters is called when production ruleActualParameters is entered.
func (s *Baseasn1Listener) EnterRuleActualParameters(ctx *RuleActualParametersContext) {}

// ExitRuleActualParameters is called when production ruleActualParameters is exited.
func (s *Baseasn1Listener) ExitRuleActualParameters(ctx *RuleActualParametersContext) {}

// EnterRuleActualParameter is called when production ruleActualParameter is entered.
func (s *Baseasn1Listener) EnterRuleActualParameter(ctx *RuleActualParameterContext) {}

// ExitRuleActualParameter is called when production ruleActualParameter is exited.
func (s *Baseasn1Listener) ExitRuleActualParameter(ctx *RuleActualParameterContext) {}

// EnterRuleParameterizedObjectSet is called when production ruleParameterizedObjectSet is entered.
func (s *Baseasn1Listener) EnterRuleParameterizedObjectSet(ctx *RuleParameterizedObjectSetContext) {}

// ExitRuleParameterizedObjectSet is called when production ruleParameterizedObjectSet is exited.
func (s *Baseasn1Listener) ExitRuleParameterizedObjectSet(ctx *RuleParameterizedObjectSetContext) {}

// EnterRuleParameterizedValue is called when production ruleParameterizedValue is entered.
func (s *Baseasn1Listener) EnterRuleParameterizedValue(ctx *RuleParameterizedValueContext) {}

// ExitRuleParameterizedValue is called when production ruleParameterizedValue is exited.
func (s *Baseasn1Listener) ExitRuleParameterizedValue(ctx *RuleParameterizedValueContext) {}

// EnterRuleParameterizedTypeAssignment is called when production ruleParameterizedTypeAssignment is entered.
func (s *Baseasn1Listener) EnterRuleParameterizedTypeAssignment(ctx *RuleParameterizedTypeAssignmentContext) {
}

// ExitRuleParameterizedTypeAssignment is called when production ruleParameterizedTypeAssignment is exited.
func (s *Baseasn1Listener) ExitRuleParameterizedTypeAssignment(ctx *RuleParameterizedTypeAssignmentContext) {
}

// EnterRuleParameterizedValueAssignment is called when production ruleParameterizedValueAssignment is entered.
func (s *Baseasn1Listener) EnterRuleParameterizedValueAssignment(ctx *RuleParameterizedValueAssignmentContext) {
}

// ExitRuleParameterizedValueAssignment is called when production ruleParameterizedValueAssignment is exited.
func (s *Baseasn1Listener) ExitRuleParameterizedValueAssignment(ctx *RuleParameterizedValueAssignmentContext) {
}

// EnterRuleParameterizedValueSetTypeAssignment is called when production ruleParameterizedValueSetTypeAssignment is entered.
func (s *Baseasn1Listener) EnterRuleParameterizedValueSetTypeAssignment(ctx *RuleParameterizedValueSetTypeAssignmentContext) {
}

// ExitRuleParameterizedValueSetTypeAssignment is called when production ruleParameterizedValueSetTypeAssignment is exited.
func (s *Baseasn1Listener) ExitRuleParameterizedValueSetTypeAssignment(ctx *RuleParameterizedValueSetTypeAssignmentContext) {
}

// EnterRuleParameterizedObjectClassAssignment is called when production ruleParameterizedObjectClassAssignment is entered.
func (s *Baseasn1Listener) EnterRuleParameterizedObjectClassAssignment(ctx *RuleParameterizedObjectClassAssignmentContext) {
}

// ExitRuleParameterizedObjectClassAssignment is called when production ruleParameterizedObjectClassAssignment is exited.
func (s *Baseasn1Listener) ExitRuleParameterizedObjectClassAssignment(ctx *RuleParameterizedObjectClassAssignmentContext) {
}

// EnterRuleParameterizedObjectAssignment is called when production ruleParameterizedObjectAssignment is entered.
func (s *Baseasn1Listener) EnterRuleParameterizedObjectAssignment(ctx *RuleParameterizedObjectAssignmentContext) {
}

// ExitRuleParameterizedObjectAssignment is called when production ruleParameterizedObjectAssignment is exited.
func (s *Baseasn1Listener) ExitRuleParameterizedObjectAssignment(ctx *RuleParameterizedObjectAssignmentContext) {
}

// EnterRuleParameterizedObjectSetAssignment is called when production ruleParameterizedObjectSetAssignment is entered.
func (s *Baseasn1Listener) EnterRuleParameterizedObjectSetAssignment(ctx *RuleParameterizedObjectSetAssignmentContext) {
}

// ExitRuleParameterizedObjectSetAssignment is called when production ruleParameterizedObjectSetAssignment is exited.
func (s *Baseasn1Listener) ExitRuleParameterizedObjectSetAssignment(ctx *RuleParameterizedObjectSetAssignmentContext) {
}

// EnterRuleParameterList is called when production ruleParameterList is entered.
func (s *Baseasn1Listener) EnterRuleParameterList(ctx *RuleParameterListContext) {}

// ExitRuleParameterList is called when production ruleParameterList is exited.
func (s *Baseasn1Listener) ExitRuleParameterList(ctx *RuleParameterListContext) {}

// EnterRuleParameters is called when production ruleParameters is entered.
func (s *Baseasn1Listener) EnterRuleParameters(ctx *RuleParametersContext) {}

// ExitRuleParameters is called when production ruleParameters is exited.
func (s *Baseasn1Listener) ExitRuleParameters(ctx *RuleParametersContext) {}

// EnterRuleParameter is called when production ruleParameter is entered.
func (s *Baseasn1Listener) EnterRuleParameter(ctx *RuleParameterContext) {}

// ExitRuleParameter is called when production ruleParameter is exited.
func (s *Baseasn1Listener) ExitRuleParameter(ctx *RuleParameterContext) {}

// EnterRuleParamGovernor is called when production ruleParamGovernor is entered.
func (s *Baseasn1Listener) EnterRuleParamGovernor(ctx *RuleParamGovernorContext) {}

// ExitRuleParamGovernor is called when production ruleParamGovernor is exited.
func (s *Baseasn1Listener) ExitRuleParamGovernor(ctx *RuleParamGovernorContext) {}

// EnterRuleDummyReference is called when production ruleDummyReference is entered.
func (s *Baseasn1Listener) EnterRuleDummyReference(ctx *RuleDummyReferenceContext) {}

// ExitRuleDummyReference is called when production ruleDummyReference is exited.
func (s *Baseasn1Listener) ExitRuleDummyReference(ctx *RuleDummyReferenceContext) {}

// EnterRuleGovernor is called when production ruleGovernor is entered.
func (s *Baseasn1Listener) EnterRuleGovernor(ctx *RuleGovernorContext) {}

// ExitRuleGovernor is called when production ruleGovernor is exited.
func (s *Baseasn1Listener) ExitRuleGovernor(ctx *RuleGovernorContext) {}

// EnterRuleDummyGovernor is called when production ruleDummyGovernor is entered.
func (s *Baseasn1Listener) EnterRuleDummyGovernor(ctx *RuleDummyGovernorContext) {}

// ExitRuleDummyGovernor is called when production ruleDummyGovernor is exited.
func (s *Baseasn1Listener) ExitRuleDummyGovernor(ctx *RuleDummyGovernorContext) {}

// EnterRuleEncodingControlSections is called when production ruleEncodingControlSections is entered.
func (s *Baseasn1Listener) EnterRuleEncodingControlSections(ctx *RuleEncodingControlSectionsContext) {
}

// ExitRuleEncodingControlSections is called when production ruleEncodingControlSections is exited.
func (s *Baseasn1Listener) ExitRuleEncodingControlSections(ctx *RuleEncodingControlSectionsContext) {}
