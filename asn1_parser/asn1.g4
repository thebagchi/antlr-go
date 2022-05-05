grammar asn1;

@parser::header {

}

@parser::members {

}

// Grammar Rules
// =============

start
  : ruleModules EOF
  ;

ruleModules
  : ruleModules ruleModuleDefinition
  | ruleModuleDefinition
  ;

ruleModuleDefinition
  : ruleModuleIdentifier
    DEFINITIONS_SYM
    ruleEncodingReferenceDefault
    ruleTagDefault
    ruleExtensionDefault
    ASSIGNMENT
    BEGIN_SYM
    ruleModuleBody
    ruleEncodingControlSections
    END_SYM
  ;

ruleModuleReference
  : UCASE_ID
  | LCASE_ID
  ;

ruleModuleIdentifier
  : ruleModuleReference
    ruleDefinitiveIdentification
  ;

ruleDefinitiveIdentification
  : ruleDefinitiveOID
  | ruleDefinitiveOIDAndIRI
  | /* empty */
  ;

ruleDefinitiveOID
  : CURLY_START ruleDefinitiveObjIdComponentList CURLY_END
  ;

ruleDefinitiveObjIdComponentList
  : ruleDefinitiveObjIdComponentList ruleDefinitiveObjIdComponent
  | ruleDefinitiveObjIdComponent
  ;

ruleDefinitiveObjIdComponent
  : ruleNameForm
  | ruleDefinitiveNumberForm
  | ruleDefinitiveNameAndNumberForm
  ;

ruleNameForm
  : LCASE_ID
  ;

ruleDefinitiveNumberForm
  : NUMBER
  ;

ruleDefinitiveNameAndNumberForm
  : LCASE_ID ROUND_START ruleDefinitiveNumberForm ROUND_END
  ;

ruleDefinitiveOIDAndIRI
  : ruleDefinitiveOID ruleIRIValue
  ;

ruleFirstArcIdentifier
  : /* TODO */
  ;

ruleSubsequentArcIdentifier
  : /* TODO */
  ;

ruleIRIValue
  : CHAR_STRING
//| DOUBLE_QUOTE ruleFirstArcIdentifier ruleSubsequentArcIdentifier DOUBLE_QUOTE
  ;

ruleEncodingReferenceDefault
  : UCASE_ID INSTRUCTIONS_SYM
  | /* EMPTY */
  ;

ruleTagDefault
  : EXPLICIT_SYM TAGS_SYM
  | IMPLICIT_SYM TAGS_SYM
  | AUTOMATIC_SYM TAGS_SYM
  | /* EMPTY */
  ;

ruleExtensionDefault
  : EXTENSIBILITY_SYM IMPLIED_SYM
  | /* EMPTY */
  ;

ruleModuleBody
  : ruleExports ruleImports ruleAssignmentList
  | /* EMPTY */
  ;

ruleExports
  : EXPORTS_SYM ruleSymbolsExported SEMI_COMMA
  | EXPORTS_SYM ALL_SYM SEMI_COMMA
  | /* EMPTY */
  ;

ruleSymbolsExported
  : ruleSymbolList
  | /* EMPTY */
  ;

ruleSymbolList
  : ruleSymbolList COMMA ruleSymbol
  | ruleSymbol
  ;

ruleSymbol
  : ruleReference
  | ruleParameterizedReference
  ;

ruleReference
  : ruleTypeReference
  | ruleValueReference
  | ruleObjectClassReference
  | ruleObjectReference
  | ruleObjectSetReference
  ;

ruleIdentifier
  : LCASE_ID
  ;

ruleTypeReference
  : UCASE_ID
  ;

ruleValueReference
  : LCASE_ID
  ;

ruleObjectReference
  : LCASE_ID
  ;

ruleObjectClassReference
  : UCASE_ID
  ;

ruleObjectSetReference
  : UCASE_ID
  ;

ruleParameterizedReference
  : ruleReference
  | ruleReference CURLY_START CURLY_END
  ;

ruleExternalTypeReference
  : ruleModuleReference DOT ruleTypeReference
  ;

ruleExternalValueReference
  : ruleModuleReference DOT ruleValueReference
  ;

ruleExternalObjectClassReference
  : ruleModuleReference DOT ruleObjectClassReference
  ;

ruleExternalObjectReference
  : ruleModuleReference DOT ruleObjectReference
  ;

ruleExternalObjectSetReference
  : ruleModuleReference DOT ruleObjectSetReference
  ;

ruleTypeFieldReference
  : AND UCASE_ID
  ;

ruleValueFieldReference
  : AND LCASE_ID
  ;

ruleValueSetFieldReference
  : AND UCASE_ID
  ;

ruleObjectFieldReference
  : AND LCASE_ID
  ;

ruleObjectSetFieldReference
  : AND UCASE_ID
  ;

ruleUsefulObjectClassReference
  : TYPE_IDENTIFIER_SYM
  | ABSTRACT_SYNTAX_SYM
  ;

ruleImports
  : IMPORTS_SYM ruleSymbolsImported SEMI_COMMA
  | /* EMPTY */
  ;

ruleSymbolsImported
  : ruleSymbolsFromModuleList
  | /* EMPTY */
  ;

ruleSymbolsFromModuleList
  : ruleSymbolsFromModuleList ruleSymbolsFromModule
  | ruleSymbolsFromModule
  ;

ruleSymbolsFromModule
  : ruleSymbolList FROM_SYM ruleGlobalModuleReference
  ;

ruleGlobalModuleReference
  : ruleModuleReference ruleAssignedIdentifier
  ;

ruleAssignedIdentifier
  : ruleObjectIdentifierValue
  | ruleDefinedValue
  | /* EMPTY */
  ;

ruleAssignmentList
  : ruleAssignmentList ruleAssignment
  | ruleAssignment
  | /* EMPTY */
  ;

ruleAssignment
  : ruleTypeAssignment
  | ruleValueAssignment
//| ruleXMLValueAssignment
  | ruleValueSetTypeAssignment
  | ruleObjectClassAssignment
  | ruleObjectAssignment
  | ruleObjectSetAssignment
  | ruleParameterizedAssignment
  ;

ruleTypeAssignment
  : ruleTypeReference ASSIGNMENT ruleType
  ;

ruleValueAssignment
  : ruleValueReference ruleType ASSIGNMENT ruleValue
  ;

ruleXMLValueAssignment
  : /* TODO */
  ;

ruleValueSetTypeAssignment
  : ruleTypeReference ruleType ASSIGNMENT ruleValueSet
  ;

ruleObjectAssignment
  : ruleObjectReference ruleDefinedObjectClass ASSIGNMENT ruleObject
  ;

ruleObjectClassAssignment
  : ruleObjectClassReference ASSIGNMENT ruleObjectClass
  ;

ruleObjectSetAssignment
  : ruleObjectSetReference ruleDefinedObjectClass ASSIGNMENT ruleObjectSet
  ;

ruleParameterizedAssignment
  : ruleParameterizedTypeAssignment
  | ruleParameterizedValueAssignment
  | ruleParameterizedValueSetTypeAssignment
  | ruleParameterizedObjectClassAssignment
  | ruleParameterizedObjectAssignment
  | ruleParameterizedObjectSetAssignment
  ;

ruleObjectIdentifierValue
  : CURLY_START ruleObjIdComponentsList CURLY_END
  | CURLY_START ruleDefinedValue ruleObjIdComponentsList CURLY_END
  ;

ruleObjIdComponentsList
  : ruleObjIdComponentsList ruleObjIdComponents
  | ruleObjIdComponents
  ;

ruleObjIdComponents
  : ruleNameForm
  | ruleNumberForm
  | ruleNameAndNumberForm
  | ruleDefinedValue
  ;

ruleNumberForm
  : NUMBER
  | ruleDefinedValue
  ;

ruleNameAndNumberForm
  : ruleIdentifier CURLY_START ruleNumberForm CURLY_END
  ;

ruleDefinedValue
  : ruleExternalValueReference
  | ruleValueReference
  | ruleParameterizedValue
  ;

ruleNamedBit
  : ruleIdentifier ROUND_START NUMBER ROUND_END
  | ruleIdentifier ROUND_START ruleDefinedValue ROUND_END
  ;

ruleNamedBitList
  : ruleNamedBitList COMMA ruleNamedBit
  | ruleNamedBit
  ;

ruleRestrictedCharacterStringType
  : BMP_STRING_SYM
  | GENERAL_STRING_SYM
  | GRAPHIC_STRING_SYM
  | IA5_STRING_SYM
  | ISO646_STRING_SYM
  | NUMERIC_STRING_SYM
  | PRINTABLE_STRING_SYM
  | TELETEX_STRING_SYM
  | T61_STRING_SYM
  | UNIVERSAL_STRING_SYM
  | UTF8_STRING_SYM
  | VIDEOTEX_STRING_SYM
  | VISIBLE_STRING_SYM
  ;

ruleUnrestrictedCharacterStringType
  : CHARACTER_SYM STRING_SYM
  ;

ruleAlternativeTypeList
  : ruleAlternativeTypeList COMMA ruleNamedType
  | ruleNamedType
  ;

ruleRootAlternativeTypeList
  : ruleAlternativeTypeList
  ;

ruleVersionNumber
  : NUMBER COLON
  | /* EMPTY */
  ;

ruleExtensionAdditionAlternativesGroup
  : VERSION_START ruleVersionNumber ruleAlternativeTypeList VERSION_END
  ;

ruleExtensionAdditionAlternative
  : ruleExtensionAdditionAlternativesGroup
  | ruleNamedType
  ;

ruleExtensionAdditionAlternativesList
  : ruleExtensionAdditionAlternativesList COMMA ruleExtensionAdditionAlternative
  | ruleExtensionAdditionAlternative
  ;

ruleExtensionAdditionAlternatives
  : COMMA ruleExtensionAdditionAlternativesList
  | /* EMPTY */
  ;

ruleEnumerationItem
  : ruleNamedNumber
  | ruleIdentifier
  ;

ruleEnumeration
  : ruleEnumeration COMMA ruleEnumerationItem
  | ruleEnumerationItem
  ;

ruleRootEnumeration
  : ruleEnumeration
  ;

ruleAdditionalEnumeration
  : ruleEnumeration
  ;

ruleAlternativeTypeLists
  : ruleRootAlternativeTypeList COMMA ruleExtensionAndException ruleExtensionAdditionAlternatives ruleOptionalExtensionMarker
  | ruleRootAlternativeTypeList
  ;

ruleEnumerations
  : ruleRootEnumeration COMMA ELLIPSIS ruleExceptionSpec COMMA ruleAdditionalEnumeration
  | ruleRootEnumeration COMMA ELLIPSIS ruleExceptionSpec
  | ruleRootEnumeration
  ;

ruleNamedNumber
  : ruleIdentifier ROUND_START ruleSignedNumber ROUND_END
  | ruleIdentifier ROUND_START ruleDefinedValue ROUND_END
  ;

ruleNamedNumberList
  : ruleNamedNumberList COMMA ruleNamedNumber
  | ruleNamedNumber
  ;

ruleComponentType
  : COMPONENTS_SYM OF_SYM ruleType
  | ruleNamedType DEFAULT_SYM ruleValue
  | ruleNamedType OPTIONAL_SYM
  | ruleNamedType
  ;

ruleExtensionAdditionGroup
  : VERSION_START ruleVersionNumber ruleComponentTypeList VERSION_END
  ;

ruleExtensionAddition
  : ruleComponentType
  | ruleExtensionAdditionGroup
  ;

ruleComponentTypeList
  : ruleComponentTypeList COMMA ruleComponentType
  | ruleComponentType
  ;

ruleExtensionAdditionList
  : ruleExtensionAdditionList COMMA ruleExtensionAddition
  | ruleExtensionAddition
  ;

ruleRootComponentTypeList
  : ruleComponentTypeList
  ;

ruleExtensionAdditions
  : COMMA ruleExtensionAdditionList
  | /* EMPTY */
  ;

ruleExtensionEndMarker
  : COMMA ELLIPSIS
  ;

ruleComponentTypeLists
  : ruleRootComponentTypeList ruleExtensionAdditions ruleExtensionEndMarker COMMA ruleRootComponentTypeList
  | ruleRootComponentTypeList ruleExtensionAdditions ruleOptionalExtensionMarker
  | ruleRootComponentTypeList COMMA ruleExtensionAndException ruleExtensionAdditions ruleExtensionEndMarker COMMA ruleRootComponentTypeList
  | ruleRootComponentTypeList COMMA ruleExtensionAndException ruleExtensionAdditions ruleOptionalExtensionMarker
  | ruleRootComponentTypeList
  ;

ruleExtensionAndException
  : ELLIPSIS ruleExceptionSpec
  | ELLIPSIS
  ;

ruleOptionalExtensionMarker
  : COMMA ELLIPSIS
  | /* EMPTY */
  ;

ruleEncodingReference
  : UCASE_ID COLON
  | /* EMPTY */
  ;

ruleClass
  : UNIVERSAL_SYM
  | APPLICATION_SYM
  | PRIVATE_SYM
  | /* EMPTY */
  ;

ruleClassNumber
  : NUMBER
  | ruleDefinedValue
  ;

ruleEncodingInstruction
  : LCASE_ID
  | UCASE_ID
  ;

ruleTag
  : SQUARE_START ruleEncodingReference ruleClass ruleClassNumber SQUARE_END
  ;

ruleEncodingPrefix
  : SQUARE_START ruleEncodingReference ruleEncodingInstruction SQUARE_END
  ;

ruleTaggedType
  : ruleTag IMPLICIT_SYM ruleType
  | ruleTag EXPLICIT_SYM ruleType
  ;

ruleEncodingPrefixedType
  : ruleEncodingPrefix ruleType
  ;

ruleBitStringType
  : BIT_SYM STRING_SYM CURLY_START ruleNamedBitList CURLY_END
  | BIT_SYM STRING_SYM
  ;

ruleBooleanType
  : BOOLEAN_SYM
  ;

ruleCharacterStringType
  : ruleRestrictedCharacterStringType
  | ruleUnrestrictedCharacterStringType
  ;

ruleChoiceType
  : CHOICE_SYM CURLY_START ruleAlternativeTypeLists CURLY_END
  ;

ruleDateType
  : DATE_SYM
  ;

ruleDateTimeType
  : DATE_TIME_SYM
  ;

ruleDurationType
  : DURATION_SYM
  ;

ruleEmbeddedPDVType
  : EMBEDDED_SYM PDV_SYM
  ;

ruleEnumeratedType
  : ENUMERATED_SYM CURLY_START ruleEnumerations CURLY_END
  ;

ruleExternalType
  : EXTERNAL_SYM
  ;

ruleInstanceOfType
  : INSTANCE_SYM OF_SYM ruleDefinedObjectClass
  ;


ruleIntegerType
  : INTEGER_SYM CURLY_START ruleNamedNumberList CURLY_END
  | INTEGER_SYM
  ;

ruleIRIType
  : OID_IRI_SYM
  ;

ruleNullType
  : NULL_SYM
  ;

ruleObjectClassFieldType
  : ruleDefinedObjectClass DOT ruleFieldName
  ;

ruleObjectIdentifierType
  : OBJECT_SYM IDENTIFIER_SYM
  ;

ruleOctetStringType
  : OCTET_SYM STRING_SYM
  ;

ruleRealType
  : REAL_SYM
  ;

ruleRelativeIRIType
  : RELATIVE_OID_IRI_SYM
  ;

ruleRelativeOIDType
  : RELATIVE_OID_SYM
  ;

ruleSequenceType
  : SEQUENCE_SYM CURLY_START CURLY_END
  | SEQUENCE_SYM CURLY_START ruleExtensionAndException ruleOptionalExtensionMarker CURLY_END
  | SEQUENCE_SYM CURLY_START ruleComponentTypeLists CURLY_END
  ;

ruleSequenceOfType
  : SEQUENCE_SYM OF_SYM ruleNamedType
  | SEQUENCE_SYM OF_SYM ruleType
  ;

ruleSetType
  : SET_SYM CURLY_START ruleComponentTypeLists CURLY_END
  | SET_SYM CURLY_START ruleExtensionAndException ruleOptionalExtensionMarker CURLY_END
  | SET_SYM CURLY_START CURLY_END
  ;

ruleSetOfType
  : SET_SYM OF_SYM ruleNamedType
  | SET_SYM OF_SYM ruleType
  ;

rulePrefixedType
  : ruleTaggedType
  | ruleEncodingPrefixedType
  ;

ruleTimeType
  : TIME_SYM
  ;

ruleTimeOfDayType
  : TIME_OF_DAY_SYM
  ;

ruleBuiltinType
  : ruleBitStringType
  | ruleBooleanType
  | ruleCharacterStringType
  | ruleChoiceType
  | ruleDateType
  | ruleDateTimeType
  | ruleDurationType
  | ruleEmbeddedPDVType
  | ruleEnumeratedType
  | ruleExternalType
  | ruleInstanceOfType
  | ruleIntegerType
  | ruleIRIType
  | ruleNullType
  | ruleObjectClassFieldType
  | ruleObjectIdentifierType
  | ruleOctetStringType
  | ruleRealType
  | ruleRelativeIRIType
  | ruleRelativeOIDType
  | ruleSequenceType
  | ruleSequenceOfType
  | ruleSetType
  | ruleSetOfType
//| rulePrefixedType
  | ruleTimeType
  | ruleTimeOfDayType
  ;

ruleSimpleDefinedType
  : ruleExternalTypeReference
  | ruleTypeReference
  ;

ruleParameterizedType
  : ruleSimpleDefinedType ruleActualParameterList
  ;

ruleParameterizedValueSetType
  : ruleSimpleDefinedType ruleActualParameterList
  ;

ruleDefinedType
  : ruleExternalTypeReference
  | ruleTypeReference
  | ruleParameterizedType
  | ruleParameterizedValueSetType
  ;

ruleUsefulType
  : GENERALIZED_TIME_SYM
  | UTC_TIME_SYM
  ;

ruleSelectionType
  : ruleIdentifier LESS ruleType
  ;

ruleTypeFromObject
  : ruleReferencedObjects DOT ruleFieldName
  ;

ruleValueSetFromObjects
  : ruleReferencedObjects DOT ruleFieldName
  ;

ruleReferencedType
  : ruleDefinedType
  | ruleUsefulType
  | ruleSelectionType
  | ruleTypeFromObject
  | ruleValueSetFromObjects
  ;

ruleTypeForConstraint
  : ruleBuiltinType
  | ruleReferencedType
  ;

ruleNamedType
  : ruleIdentifier ruleType
  ;

ruleTypeWithConstraint
  : SET_SYM ruleConstraint OF_SYM ruleType
  | SET_SYM ruleSizeConstraint OF_SYM ruleType
  | SEQUENCE_SYM ruleConstraint OF_SYM ruleType
  | SEQUENCE_SYM ruleSizeConstraint OF_SYM ruleType
  | SET_SYM ruleConstraint OF_SYM ruleNamedType
  | SET_SYM ruleSizeConstraint OF_SYM ruleNamedType
  | SEQUENCE_SYM ruleConstraint OF_SYM ruleNamedType
  | SEQUENCE_SYM ruleSizeConstraint OF_SYM ruleNamedType
  ;

ruleConstrainedType
  : ruleTypeForConstraint ruleConstraint
  | ruleTypeWithConstraint
  ;

ruleType
  : ruleBuiltinType
  | ruleReferencedType
  | ruleConstrainedType
  ;

ruleIdentifierList
  : ruleIdentifierList COMMA ruleIdentifier
  | ruleIdentifier
  ;

ruleCharsDefn
  : CHAR_STRING
  | ruleQuadruple
  | ruleTuple
  | ruleDefinedValue
  ;

ruleCharSyms
  : ruleCharSyms COMMA ruleCharsDefn
  | ruleCharsDefn
  ;

ruleGroup
  : NUMBER
  ;

rulePlane
  : NUMBER
  ;

ruleRow
  : NUMBER
  ;

ruleCell
  : NUMBER
  ;

ruleTableColumn
  : NUMBER
  ;

ruleTableRow
  : NUMBER
  ;

ruleCharacterStringList
  : CURLY_START ruleCharSyms CURLY_END
  ;

ruleQuadruple
  : CURLY_START ruleGroup COMMA rulePlane COMMA ruleRow COMMA ruleCell CURLY_END
  ;

ruleTuple
  : CURLY_START ruleTableColumn COMMA ruleTableRow CURLY_END
  ;

ruleRestrictedCharacterStringValue
  : CHAR_STRING
  | ruleCharacterStringList
  | ruleQuadruple
  | ruleTuple
  ;

ruleUnrestrictedCharacterStringValue
  : ruleSequenceValue
  ;

ruleNumericRealValue
  : FLOAT
  | MINUS FLOAT
  | ruleSequenceValue
  ;

ruleSpecialRealValue
  : PLUS_INFINITY_SYM
  | MINUS_INFINITY_SYM
  | NOT_A_NUMBER_SYM
  ;

ruleRelativeOIDComponents
  : ruleNumberForm
  | ruleNameAndNumberForm
  | ruleDefinedValue
  ;

ruleRelativeOIDComponentsList
  : ruleRelativeOIDComponentsList ruleRelativeOIDComponents
  | ruleRelativeOIDComponents
  ;

ruleComponentValueList
  : ruleComponentValueList COMMA ruleNamedValue
  | ruleNamedValue
  ;

ruleValueList
  : ruleValueList COMMA ruleValue
  | ruleValue
  ;

ruleNamedValue
  : ruleIdentifier ruleValue
  ;

ruleNamedValueList
  : ruleNamedValueList COMMA ruleNamedValue
  | ruleNamedValue
  ;

ruleBitStringValue
  : BIN_STRING
  | HEX_STRING
  | CURLY_START ruleIdentifierList CURLY_END
  | CURLY_START CURLY_END
  | CONTAINING_SYM ruleValue
  ;

ruleBooleanValue
  : TRUE_SYM
  | FALSE_SYM
  ;

ruleCharacterStringValue
  : ruleRestrictedCharacterStringValue
  | ruleUnrestrictedCharacterStringValue
  ;

ruleChoiceValue
  : ruleIdentifier COLON ruleValue
  ;

ruleEmbeddedPDVValue
  : ruleSequenceValue
  ;

ruleEnumeratedValue
  : ruleIdentifier
  ;

ruleExternalValue
  : ruleSequenceValue
  ;

ruleInstanceOfValue
  : ruleValue
  ;

ruleIntegerValue
  : ruleSignedNumber
  | ruleIdentifier
  ;

ruleNullValue
  : NULL_SYM
  ;

ruleOctetStringValue
  : BIN_STRING
  | HEX_STRING
  | CONTAINING_SYM ruleValue
  ;

ruleRealValue
  : ruleNumericRealValue
  | ruleSpecialRealValue
  ;

ruleRelativeIRIValue
  : /* TODO */
  ;

ruleRelativeOIDValue
  : CURLY_START ruleRelativeOIDComponentsList CURLY_END
  ;

ruleSequenceValue
  : CURLY_START ruleComponentValueList CURLY_END
  ;

ruleSequenceOfValue
  : CURLY_START ruleValueList CURLY_END
  | CURLY_START ruleNamedValueList CURLY_END
  | CURLY_START CURLY_END
  ;

ruleSetValue
  : CURLY_START ruleComponentValueList CURLY_END
  | CURLY_START CURLY_END
  ;

ruleSetOfValue
  : CURLY_START ruleValueList CURLY_END
  | CURLY_START ruleNamedValueList CURLY_END
  | CURLY_START CURLY_END
  ;

rulePrefixedValue
  : ruleValue
  ;

ruleTimeValue
  : CHAR_STRING
  ;

ruleBuiltinValue
  : ruleBitStringValue
  | ruleBooleanValue
  | ruleCharacterStringValue
  | ruleChoiceValue
//| ruleEmbeddedPDVValue
  | ruleEnumeratedValue
//| ruleExternalValue
//| ruleInstanceOfValue
  | ruleIntegerValue
  | ruleIRIValue
  | ruleNullValue
  | ruleObjectIdentifierValue
  | ruleOctetStringValue
  | ruleRealValue
//| ruleRelativeIRIValue
  | ruleRelativeOIDValue
  | ruleSequenceValue
  | ruleSequenceOfValue
  | ruleSetValue
  | ruleSetOfValue
//| rulePrefixedValue
  | ruleTimeValue
  ;

ruleValueFromObject
  : ruleReferencedObjects DOT ruleFieldName
  ;

ruleReferencedValue
  : ruleDefinedValue
  | ruleValueFromObject
  ;

ruleOpenTypeFieldVal
  : ruleType COLON ruleValue
  ;

ruleFixedTypeFieldVal
  : ruleBuiltinValue
  | ruleReferencedValue
  ;

ruleObjectClassFieldValue
  : ruleOpenTypeFieldVal
  | ruleFixedTypeFieldVal
  ;

ruleValue
  : ruleBuiltinValue
  | ruleReferencedValue
  | ruleObjectClassFieldValue
  ;

ruleValueSet
  : CURLY_START ruleElementSetSpecs CURLY_END
  ;

ruleDefinedObjectClass
  : ruleExternalObjectClassReference
  | ruleObjectClassReference
  | ruleUsefulObjectClassReference
  ;

ruleObject
  : ruleDefinedObject
  | ruleObjectDefn
  | ruleObjectFromObject
  | ruleParameterizedObject
  ;

ruleDefinedObject
  : ruleExternalObjectReference
  | ruleObjectReference
  ;

ruleDefinedObjectSet
  : ruleExternalObjectSetReference
  | ruleObjectSetReference
  ;

ruleObjectDefn
  : ruleDefaultSyntax
  | ruleDefinedSyntax
  ;

ruleDefaultSyntax
  : CURLY_START ruleFieldSettingList CURLY_END
  ;

ruleDefinedSyntax
  : CURLY_START ruleDefinedSyntaxTokenList CURLY_END
  ;

ruleDefinedSyntaxTokenList
  : ruleDefinedSyntaxTokenList ruleDefinedSyntaxToken
  | ruleDefinedSyntaxToken
  | /* EMPTY */
  ;

ruleDefinedSyntaxToken
  : ruleLiteral
  | ruleSetting
  ;

ruleLiteral
  : ruleWord
  | COMMA
  ;

// TODO: Note this rule might require more tokens handling
ruleWord
  : LCASE_ID
  | UCASE_ID
  ;

ruleFieldSettingList
  : ruleFieldSettingList COMMA ruleFieldSetting
  | ruleFieldSetting
  | /* EMPTY */
  ;

ruleFieldSetting
  : rulePrimitiveFieldName ruleSetting
  ;

rulePrimitiveFieldName
  : ruleTypeFieldReference
  | ruleValueFieldReference
  | ruleValueSetFieldReference
  | ruleObjectFieldReference
  | ruleObjectSetFieldReference
  ;

ruleSetting
  : ruleType
  | ruleValue
  | ruleValueSet
  | ruleObject
  | ruleObjectSet
  ;

ruleObjectFromObject
  : ruleReferencedObjects DOT ruleFieldName
  ;

ruleObjectSetFromObjects
  : ruleReferencedObjects DOT ruleFieldName
  ;

ruleReferencedObjects
  : ruleDefinedObject
  | ruleParameterizedObject
  | ruleDefinedObjectSet
  | ruleParameterizedObjectSet
  ;

ruleFieldName
  : ruleFieldName DOT rulePrimitiveFieldName
  | rulePrimitiveFieldName
  ;

ruleParameterizedObject
  : ruleDefinedObject ruleActualParameterList
  ;

ruleObjectClass
  : ruleDefinedObjectClass
  | ruleObjectClassDefn
  | ruleParameterizedObjectClass
  ;

ruleObjectClassDefn
  : CLASS_SYM CURLY_START ruleFieldSpecList CURLY_END ruleWithSyntaxSpec
  ;

ruleFieldSpecList
  : ruleFieldSpecList COMMA ruleFieldSpec
  | ruleFieldSpec
  ;

ruleFieldSpec
  : ruleTypeFieldSpec
  | ruleFixedTypeValueFieldSpec
  | ruleVariableTypeValueFieldSpec
  | ruleFixedTypeValueSetFieldSpec
  | ruleVariableTypeValueSetFieldSpec
  | ruleObjectFieldSpec
  | ruleObjectSetFieldSpec
  ;

ruleTypeFieldSpec
  : ruleTypeFieldReference ruleTypeOptionalitySpec
  ;

ruleFixedTypeValueFieldSpec
  : ruleValueFieldReference ruleType ruleUnique ruleValueOptionalitySpec
  ;

ruleUnique
  : UNIQUE_SYM
  | /* EMPTY */
  ;

ruleVariableTypeValueFieldSpec
  : ruleValueFieldReference ruleFieldName ruleValueOptionalitySpec
  ;

ruleFixedTypeValueSetFieldSpec
  : ruleValueSetFieldReference ruleType ruleValueSetOptionalitySpec
  ;

ruleVariableTypeValueSetFieldSpec
  : ruleValueSetFieldReference ruleFieldName ruleValueSetOptionalitySpec
  ;

ruleObjectFieldSpec
  : ruleObjectFieldReference ruleDefinedObjectClass ruleObjectOptionalitySpec
  ;

ruleObjectSetFieldSpec
  : ruleObjectSetFieldReference ruleDefinedObjectClass ruleObjectSetOptionalitySpec
  ;

ruleTypeOptionalitySpec
  : OPTIONAL_SYM
  | DEFAULT_SYM
  | /* EMPTY */
  ;

ruleValueOptionalitySpec
  : OPTIONAL_SYM
  | DEFAULT_SYM ruleValue
  | /* EMPTY */
  ;

ruleValueSetOptionalitySpec
  : OPTIONAL_SYM
  | DEFAULT_SYM ruleValueSet
  | /* EMPTY */
  ;

ruleObjectOptionalitySpec
  : OPTIONAL_SYM
  | DEFAULT_SYM ruleObject
  | /* EMPTY */
  ;

ruleObjectSetOptionalitySpec
  : OPTIONAL_SYM
  | DEFAULT_SYM ruleObjectSet
  | /* EMPTY */
  ;

ruleWithSyntaxSpec
  : WITH_SYM SYNTAX_SYM ruleSyntaxList
  ;

ruleSyntaxList
  : CURLY_START ruleTokenOrGroupSpecList CURLY_END
  ;

ruleTokenOrGroupSpecList
  : ruleTokenOrGroupSpecList ruleTokenOrGroupSpec
  | ruleTokenOrGroupSpec
  | /* EMPTY */
  ;

ruleTokenOrGroupSpec
  : ruleRequiredToken
  | ruleOptionalGroup
  ;

ruleRequiredToken
  : ruleLiteral
  | rulePrimitiveFieldName
  ;

ruleOptionalGroup
  : SQUARE_START ruleTokenOrGroupSpecList SQUARE_END
  ;

ruleParameterizedObjectClass
  : ruleDefinedObjectClass ruleActualParameterList
  ;

ruleObjectSet
  : CURLY_START ruleObjectSetSpec CURLY_END
  ;

ruleObjectSetSpec
  : ruleRootElementSetSpec
  | ruleRootElementSetSpec COMMA ELLIPSIS
  | ELLIPSIS
  | ELLIPSIS COMMA ruleAdditionalElementSetSpec
  | ruleRootElementSetSpec COMMA ELLIPSIS COMMA ruleAdditionalElementSetSpec
  ;

ruleRootElementSetSpec
  : ruleElementSetSpec
  ;

ruleAdditionalElementSetSpec
  : ruleElementSetSpec
  ;

ruleElementSetSpecs
  : ruleRootElementSetSpec
  | ruleRootElementSetSpec COMMA ELLIPSIS
  | ruleRootElementSetSpec COMMA ELLIPSIS COMMA ruleAdditionalElementSetSpec
  ;

ruleElementSetSpec
  : ruleUnions
  | ALL_SYM ruleExclusions
  ;

ruleUnions
  : ruleIntersections
//| ruleUElems ruleUnionMark ruleIntersections
  | ruleUnions ruleUnionMark ruleIntersections
  ;

ruleExclusions
  : EXCEPT_SYM ruleElements
  ;

ruleIntersections
  : ruleIntersectionElements
//| ruleIElems ruleIntersectionMark ruleIntersectionElements
  | ruleIntersections ruleIntersectionMark ruleIntersectionElements
  ;

ruleUElems
  : ruleUnions
  ;

ruleUnionMark
  : PIPE
  | UNION_SYM
  ;

ruleLowerEndValue
  : ruleValue
  | MIN_SYM
  ;

ruleUpperEndValue
  : ruleValue
  | MAX_SYM
  ;

ruleIncludes
  : INCLUDES_SYM
  | /* EMPTY */
  ;

ruleLowerEndpoint
  : ruleLowerEndValue
  | ruleLowerEndValue LESS
  ;

ruleUpperEndpoint
  : ruleUpperEndValue
  | LESS ruleUpperEndValue
  ;

ruleLevel
  : DOT ruleLevel
  | /* EMPTY */
  ;

ruleComponentIdList
  : ruleComponentIdList DOT ruleIdentifier
  | ruleIdentifier
  ;

ruleAtNotation
  : AT ruleComponentIdList
  | AT DOT ruleLevel ruleComponentIdList
  ;

ruleAtNotationList
  : ruleAtNotationList COMMA ruleAtNotation
  | ruleAtNotation
  ;

ruleUserDefinedConstraintParameter
  : ruleGovernor COLON ruleValue
  | ruleGovernor COLON ruleObject
  | ruleDefinedObjectSet
  | ruleType
  | ruleDefinedObjectClass
  ;

ruleUserDefinedConstraintParameterList
  : ruleUserDefinedConstraintParameterList COMMA ruleUserDefinedConstraintParameter
  | ruleUserDefinedConstraintParameter
  | /* EMPTY */
  ;

ruleUserDefinedConstraint
  : CONSTRAINED_SYM BY_SYM CURLY_START ruleUserDefinedConstraintParameterList CURLY_END
  ;

ruleSimpleTableConstraint
  : ruleObjectSet
  ;

ruleComponentRelationConstraint
  : CURLY_START ruleDefinedObjectSet CURLY_END CURLY_START ruleAtNotationList CURLY_END
  ;

ruleTableConstraint
  : ruleSimpleTableConstraint
  | ruleComponentRelationConstraint
  ;

ruleContentsConstraint
  : CONTAINING_SYM ruleType ENCODED_SYM BY_SYM ruleValue
  | CONTAINING_SYM ruleType
  | ENCODED_SYM BY_SYM ruleValue
  ;

ruleSubtypeConstraint
  : ruleElementSetSpecs
  ;

ruleGeneralConstraint
  : ruleUserDefinedConstraint
  | ruleTableConstraint
  | ruleContentsConstraint
  ;

ruleConstraintSpec
  : ruleSubtypeConstraint
  | ruleGeneralConstraint
  ;

ruleSignedNumber
  : NUMBER
  | MINUS NUMBER
  ;

ruleExceptionIdentification
  : ruleSignedNumber
  | ruleDefinedValue
  | ruleType COLON ruleValue
  ;

ruleExceptionSpec
  : EXCLAMATION ruleExceptionIdentification
  | /* EMPTY */
  ;

ruleConstraint
  : ROUND_START ruleConstraintSpec ruleExceptionSpec ROUND_END
  ;

ruleSingleTypeConstraint
  : ruleConstraint
  ;

ruleValueConstraint
  : ruleConstraint
  | /* EMPTY */
  ;

rulePresenceConstraint
  : PRESENT_SYM
  | ABSENT_SYM
  | OPTIONAL_SYM
  | /* EMPTY */
  ;

ruleComponentConstraint
  : ruleValueConstraint rulePresenceConstraint
  ;

ruleNamedConstraint
  : ruleIdentifier ruleComponentConstraint
  ;

ruleTypeConstraints
  : ruleTypeConstraints COMMA ruleNamedConstraint
  | ruleNamedConstraint
  ;

ruleFullSpecification
  : CURLY_START ruleTypeConstraints CURLY_END
  ;

rulePartialSpecification
  : CURLY_START ELLIPSIS COMMA ruleTypeConstraints CURLY_END
  ;

ruleMultipleTypeConstraints
  : ruleFullSpecification
  | rulePartialSpecification
  ;

ruleSimpleString
  : CHAR_STRING
  ;

ruleSingleValue
  : ruleValue
  ;

ruleContainedSubtype
  : ruleIncludes ruleType
  ;

ruleValueRange
  : ruleLowerEndpoint RANGE ruleUpperEndpoint
  ;

rulePermittedAlphabet
  : FROM_SYM ruleConstraint
  ;

ruleSizeConstraint
  : SIZE_SYM ruleConstraint
  ;

ruleInnerTypeConstraints
  : WITH_SYM COMPONENT_SYM ruleSingleTypeConstraint
  | WITH_SYM COMPONENTS_SYM ruleMultipleTypeConstraints
  ;

rulePatternConstraint
  : PATTERN_SYM ruleValue
  ;

rulePropertySettings
  : SETTINGS_SYM ruleSimpleString
  ;

ruleDurationRange
  : ruleValueRange
  ;

ruleTimePointRange
  : ruleValueRange
  ;

ruleRecurrenceRange
  : ruleValueRange
  ;

ruleSubtypeElements
  : ruleSingleValue
  | ruleContainedSubtype
  | ruleValueRange
  | rulePermittedAlphabet
  | ruleSizeConstraint
  | ruleInnerTypeConstraints
  | rulePatternConstraint
  | rulePropertySettings
//| ruleDurationRange
//| ruleTimePointRange
//| ruleRecurrenceRange
  ;

ruleObjectSetElements
  : ruleObject
  | ruleDefinedObjectSet
  | ruleObjectSetFromObjects
  | ruleParameterizedObjectSet
  ;

ruleElements
  : ruleSubtypeElements
  | ruleObjectSetElements
  | ROUND_START ruleElementSetSpec ROUND_END
  ;

ruleElems
  : ruleElements
  ;

ruleIntersectionElements
  : ruleElements
  | ruleElems ruleExclusions
  ;

ruleIElems
  : ruleIntersections
  ;

ruleIntersectionMark
  : CARAT
  | INTERSECTION_SYM
  ;

ruleSimpleDefinedValue
  : ruleExternalValueReference
  | ruleValueReference
  ;

ruleActualParameterList
  : CURLY_START ruleActualParameters CURLY_END
  ;

ruleActualParameters
  : ruleActualParameters COMMA ruleActualParameter
  | ruleActualParameter
  ;

ruleActualParameter
  : ruleType
  | ruleValue
  | ruleValueSet
  | ruleDefinedObjectClass
  | ruleObject
  | ruleObjectSet
  ;

ruleParameterizedObjectSet
  : ruleDefinedObjectSet ruleActualParameterList
  ;

ruleParameterizedValue
  : ruleSimpleDefinedValue ruleActualParameterList
  ;

ruleParameterizedTypeAssignment
  : ruleTypeReference ruleParameterList ASSIGNMENT ruleType
  ;

ruleParameterizedValueAssignment
  : ruleValueReference ruleParameterList ruleType ASSIGNMENT ruleValue
  ;

ruleParameterizedValueSetTypeAssignment
  : ruleTypeReference ruleParameterList ruleType ASSIGNMENT ruleValueSet
  ;

ruleParameterizedObjectClassAssignment
  : ruleObjectClassReference ruleParameterList ASSIGNMENT ruleObjectClass
  ;

ruleParameterizedObjectAssignment
  : ruleObjectReference ruleParameterList ruleDefinedObjectClass ASSIGNMENT ruleObject
  ;

ruleParameterizedObjectSetAssignment
  : ruleObjectSetReference ruleParameterList ruleDefinedObjectClass ASSIGNMENT ruleObjectSet
  ;

ruleParameterList
  : CURLY_START ruleParameters CURLY_END
  ;

ruleParameters
  : ruleParameters COMMA ruleParameter
  | ruleParameter
  ;

ruleParameter
  : ruleParamGovernor COLON ruleDummyReference
  | ruleDummyReference
  ;

ruleParamGovernor
  : ruleGovernor
  | ruleDummyGovernor
  ;

ruleDummyReference
  : ruleReference
  ;

ruleGovernor
  : ruleType
  | ruleDefinedObjectClass
  ;

ruleDummyGovernor
  : ruleDummyReference
  ;

ruleEncodingControlSections
  : /* TODO */
  ;

// Lexer Rules
// =============
ABSENT_SYM            : 'ABSENT';
ABSTRACT_SYNTAX_SYM   : 'ABSTRACT-SYNTAX';
ALL_SYM               : 'ALL';
APPLICATION_SYM       : 'APPLICATION';
AUTOMATIC_SYM         : 'AUTOMATIC';
BEGIN_SYM             : 'BEGIN';
BIT_SYM               : 'BIT';
BMP_STRING_SYM        : 'BMPString';
BOOLEAN_SYM           : 'BOOLEAN';
BY_SYM                : 'BY';
CHARACTER_SYM         : 'CHARACTER';
CHOICE_SYM            : 'CHOICE';
CLASS_SYM             : 'CLASS';
COMPONENT_SYM         : 'COMPONENT';
COMPONENTS_SYM        : 'COMPONENTS';
CONSTRAINED_SYM       : 'CONSTRAINED';
CONTAINING_SYM        : 'CONTAINING';
DATE_SYM              : 'DATE';
DATE_TIME_SYM         : 'DATE-TIME';
DEFAULT_SYM           : 'DEFAULT';
DEFINITIONS_SYM       : 'DEFINITIONS';
DURATION_SYM          : 'DURATION';
EMBEDDED_SYM          : 'EMBEDDED';
ENCODED_SYM           : 'ENCODED';
ENCODING_CONTROL_SYM  : 'ENCODING-CONTROL';
END_SYM               : 'END';
ENUMERATED_SYM        : 'ENUMERATED';
EXCEPT_SYM            : 'EXCEPT';
EXPLICIT_SYM          : 'EXPLICIT';
EXPORTS_SYM           : 'EXPORTS';
EXTENSIBILITY_SYM     : 'EXTENSIBILITY';
EXTERNAL_SYM          : 'EXTERNAL';
FALSE_SYM             : 'FALSE';
FROM_SYM              : 'FROM';
GENERALIZED_TIME_SYM  : 'GeneralizedTime';
GENERAL_STRING_SYM    : 'GeneralString';
GRAPHIC_STRING_SYM    : 'GraphicString';
IA5_STRING_SYM        : 'IA5String';
IDENTIFIER_SYM        : 'IDENTIFIER';
IMPLICIT_SYM          : 'IMPLICIT';
IMPLIED_SYM           : 'IMPLIED';
IMPORTS_SYM           : 'IMPORTS';
INCLUDES_SYM          : 'INCLUDES';
INSTANCE_SYM          : 'INSTANCE';
INSTRUCTIONS_SYM      : 'INSTRUCTIONS';
INTEGER_SYM           : 'INTEGER';
INTERSECTION_SYM      : 'INTERSECTION';
ISO646_STRING_SYM     : 'ISO646String';
MAX_SYM               : 'MAX';
MIN_SYM               : 'MIN';
MINUS_INFINITY_SYM    : 'MINUS-INFINITY';
NOT_A_NUMBER_SYM      : 'NOT-A-NUMBER';
NULL_SYM              : 'NULL';
NUMERIC_STRING_SYM    : 'NumericString';
OBJECT_SYM            : 'OBJECT';
OBJECT_DESCRIPTOR_SYM : 'ObjectDescriptor';
OCTET_SYM             : 'OCTET';
OF_SYM                : 'OF';
OID_IRI_SYM           : 'OID-IRI';
OPTIONAL_SYM          : 'OPTIONAL';
PATTERN_SYM           : 'PATTERN';
PDV_SYM               : 'PDV';
PLUS_INFINITY_SYM     : 'PLUS-INFINITY';
PRESENT_SYM           : 'PRESENT';
PRINTABLE_STRING_SYM  : 'PrintableString';
PRIVATE_SYM           : 'PRIVATE';
REAL_SYM              : 'REAL';
RELATIVE_OID_SYM      : 'RELATIVE-OID';
RELATIVE_OID_IRI_SYM  : 'RELATIVE-OID-IRI';
SEQUENCE_SYM          : 'SEQUENCE';
SET_SYM               : 'SET';
SETTINGS_SYM          : 'SETTINGS';
SIZE_SYM              : 'SIZE';
STRING_SYM            : 'STRING';
SYNTAX_SYM            : 'SYNTAX';
T61_STRING_SYM        : 'T61String';
TAGS_SYM              : 'TAGS';
TELETEX_STRING_SYM    : 'TeletexString';
TIME_SYM              : 'TIME';
TIME_OF_DAY_SYM       : 'TIME-OF-DAY';
TRUE_SYM              : 'TRUE';
TYPE_IDENTIFIER_SYM   : 'TYPE-IDENTIFIER';
UNION_SYM             : 'UNION';
UNIQUE_SYM            : 'UNIQUE';
UNIVERSAL_SYM         : 'UNIVERSAL';
UNIVERSAL_STRING_SYM  : 'UniversalString';
UTC_TIME_SYM          : 'UTCTime';
UTF8_STRING_SYM       : 'UTF8String';
VIDEOTEX_STRING_SYM   : 'VideotexString';
VISIBLE_STRING_SYM    : 'VisibleString';
WITH_SYM              : 'WITH';
TAG_SYM               : 'TAG';
XER_SYM               : 'XER';
PER_SYM               : 'PER';

ASSIGNMENT            : '::=';
ELLIPSIS              : '...';
RANGE                 : '..';
VERSION_START         : '[[';
VERSION_END           : ']]';
DASH_COMMENT          : '--';
BLOCK_START           : '/*';
BLOCK_END             : '*/';
AND                   : '&';
CURLY_START           : '{';
CURLY_END             : '}';
SQUARE_START          : '[';
SQUARE_END            : ']';
ROUND_START           : '(';
ROUND_END             : ')';
COLON                 : ':';
GREATER               : '>';
LESS                  : '<';
COMMA                 : ',';
SLASH                 : '/';
PIPE                  : '|';
EXCLAMATION           : '!';
AT                    : '@';
CARAT                 : '^';
SEMI_COMMA            : ';';
DOT                   : '.';
EQUALS                : '=';
MINUS                 : '-';
SINGLE_QUOTE          : '\'';
DOUBLE_QUOTE          : '"';

fragment DIGIT        : '0'..'9';
fragment EXPONENT     : ('e' | 'E') ( '+' | '-' )? DIGIT+;
fragment HEXDIGIT     : ('0'..'9'|'a'..'f'|'A'..'F');
fragment BINDIGIT     : ('0'..'1');
fragment HYPHEN       : '-';
fragment UC_LETTER    : 'A'..'Z';
fragment LC_LETTER    : 'a'..'z';
fragment LETTER       : 'A'..'Z' | 'a'..'z';
fragment BACKSLASH    : '\\';

fragment ESC_SEQUENCE
  : BACKSLASH ('a'|'b'|'f'|'n'|'r'|'t'|'v'|'"'|'\''|'\\'|'?'|'`')
  ;

FLOAT
  : ( DIGIT+ ('.' DIGIT+) EXPONENT?
    | DIGIT+ EXPONENT
    | '.' DIGIT+ EXPONENT?
    )
  ;

NUMBER
  : DIGIT+;

HEX_STRING
  : SINGLE_QUOTE HEXDIGIT+ SINGLE_QUOTE 'H';

BIN_STRING
  : SINGLE_QUOTE BINDIGIT+ SINGLE_QUOTE 'H';

CHAR_STRING
  : DOUBLE_QUOTE (ESC_SEQUENCE | ~('\\'|'"'|'\n'|'\r'))* DOUBLE_QUOTE
  ;

WHITESPACE
  : ( '\t' | ' ' | '\r' | '\n'| '\u000C' )+ -> channel(HIDDEN)
  ;

//CAPS_ID
//  : (UC_LETTER) (UC_LETTER | DIGIT | HYPHEN)*
//  ;

UCASE_ID
  : (UC_LETTER) (LC_LETTER | UC_LETTER | DIGIT | HYPHEN)*
  ;

LCASE_ID
  : (LC_LETTER) (LC_LETTER | UC_LETTER | DIGIT | HYPHEN)*
  ;

//OID_LABEL
//  : (SLASH | LC_LETTER | UC_LETTER | DOT | DIGIT | HYPHEN)+
//  ;

LINE_COMMENT
  : DASH_COMMENT (.*? (DASH_COMMENT | EOF | '\r'? '\n')) -> channel(HIDDEN)
  ;

BLOCK_COMMENT
  : BLOCK_START (BLOCK_COMMENT | .)*? BLOCK_END -> channel(HIDDEN)
  ;