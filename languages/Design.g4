grammar Design;

designIt
    : (comment | comment declaration)* (NewLine | EOF)
    ;

declaration
    :
    | designSystemDeclaration
    | commentBlockDeclaration
    | designBlockDeclaration
    ;

designSystemDeclaration: DesignSystem;

// Block

commentBlockDeclaration: comment*;

designBlockDeclaration: Design IDENTIFIER LBRACE designBodyDeclaration RBRACE;

designBodyDeclaration: express*;

express: equalExpress
    | useExpress
    | valueExpress
    | templateExpress
    ;

equalExpress: IDENTIFIER '=' IDENTIFIER;
useExpress: IDENTIFIER '->' IDENTIFIER ;
valueExpress: IDENTIFIER ':' IDENTIFIER ;
templateExpress: Template IDENTIFIER '```' Code '```' ;

layer: Layer;

comment: '#' Space* .*? NewLine;

// Design Keywords

DesignSystem: 'DesignSystem';
Design: 'design';
Project: 'project';


// Atomic Keywords

Page: 'page';           // 页面, Pages
Layer: 'layer';         // 模板, Templates
Function: 'function';   // 组织, Organisms
Library: 'library';     // 份子, Molecules
Unit: 'unit';           // 原子, Atoms

// Layout

Template: 'template';
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
NewLine :           '\r\n' | '\n';

LBRACE:             '{';
RBRACE:             '}';


IDENTIFIER:         Letter LetterOrDigit*;
Code:               Letter LetterOrDigit*;

fragment Digits
    : [0-9] ([0-9_]* [0-9])?
    ;

fragment LetterOrDigit
    : Letter
    | [0-9]
    ;

fragment Letter
    : [a-zA-Z$_] // these are the "java letters" below 0x7F
    | ~[\u0000-\u007F\uD800-\uDBFF] // covers all characters above 0x7F which are not a surrogate
    | [\uD800-\uDBFF] [\uDC00-\uDFFF] // covers UTF-16 surrogate pairs encodings for U+10000 to U+10FFFF
    ;