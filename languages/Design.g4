grammar Design;

designIt: commentDeclaration* declaration declaration* EOF;

declaration
    : designSystemDeclaration
//    | commentBlockDeclaration
    | designBlockDeclaration
    | expressDeclaration
    | templateBlockDeclaration
    | componentBlockDeclaration
    | layerBlockDeclaration
    | codeBlockDeclaration
    ;

designSystemDeclaration: DESIGN_SYSTEM ':' IDENTIFIER;

// Block

commentBlockDeclaration: commentDeclaration?;

designBlockDeclaration  : DESIGN IDENTIFIER LBRACE designBodyDeclaration RBRACE;
templateBlockDeclaration: TEMPLATE IDENTIFIER LBRACE templateBodyDeclaration RBRACE;
componentBlockDeclaration: COMPONENT IDENTIFIER LBRACE componentBodyDeclaration RBRACE;
layerBlockDeclaration: LAYER IDENTIFIER LBRACE layerBodyDeclaration RBRACE;
codeBlockDeclaration: '```' HTML_STRING* '```';

designBodyDeclaration: express*;
templateBodyDeclaration: express*;
componentBodyDeclaration: express*;
layerBodyDeclaration: express*;

expressDeclaration: express;

express: equalExpress
    | useExpress
    | valueExpress
    | templateExpress
    | layerExpress
    ;

equalExpress:    IDENTIFIER '='  IDENTIFIER;
useExpress:      IDENTIFIER '->' IDENTIFIER;
valueExpress:    IDENTIFIER ':'  IDENTIFIER;
layerExpress:    LAYER IDENTIFIER ;
templateExpress: TEMPLATE IDENTIFIER;

layer: LAYER;

commentDeclaration: '#' Space? IDENTIFIER;

// Design Keywords

DESIGN_SYSTEM: 'DesignSystem';
DESIGN: 'Design';
PROJECT: 'Project';


// Atomic Keywords

Page: 'Page';           // 页面, Pages
LAYER: 'Layer';         // 模板, Templates
Function: 'Function';   // 组织, Organisms
Library: 'Library';     // 份子, Molecules
Unit: 'Unit';           // 原子, Atoms

// Layout

TEMPLATE:  'Template';
COMPONENT: 'Component';

Position: 'position';

// Flow

Flow: 'Flow';

// Behavior

Behavior: 'behavior';

// Fragment

WS:                 [ \t\r\n\u000C]+ -> channel(HIDDEN);
COMMENT:            '/*' .*? '*/'    -> channel(HIDDEN);
LINE_COMMENT:       '//' ~[\r\n]*    -> channel(HIDDEN);

EmptyLine:          NewLine Space+ NewLine -> skip;
Space :             [ \t];
NewLine :           '\r\n' | '\n' | '\r';

LBRACE:             '{';
RBRACE:             '}';
Quote: '"';

IDENTIFIER:        '#'? LetterOrDigit LetterOrDigit*;

STRING_LITERAL: Letter CodeLetter*;

fragment COLON: ':';

fragment Digits
    : [0-9] ([0-9_]* [0-9])?
    ;

fragment LetterOrDigit
    : Letter
    | [0-9]
    ;

fragment CodeLetter
    : [a-zA-Z$_] // these are the "java letters" below 0x7F
    | [0-9]
    | ~[\u0000-\u007F\uD800-\uDBFF] // covers all characters above 0x7F which are not a surrogate
    | [\uD800-\uDBFF] [\uDC00-\uDFFF] // covers UTF-16 surrogate pairs encodings for U+10000 to U+10FFFF
    ;

fragment Letter
    : [a-zA-Z$_] // these are the "java letters" below 0x7F
    | ~[\u0000-\u007F\uD800-\uDBFF] // covers all characters above 0x7F which are not a surrogate
    | [\uD800-\uDBFF] [\uDC00-\uDFFF] // covers UTF-16 surrogate pairs encodings for U+10000 to U+10FFFF
    ;


HTML_STRING
   : '<' ( TAG | ~ [<>] )* '>'
   | '</' ( TAG | ~ [<>] )* '>'
   ;

fragment TAG
   : '<' .*? '>'
   ;
