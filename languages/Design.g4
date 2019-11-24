grammar Design;

start: (comment | configDeclaration | decalartions)* EOF;

comment: IDENTIFIER;

configDeclaration: configKey COLON configValue;

configKey: IDENTIFIER;
configValue
    : DIGITS_IDENTIFIER unit?
    | DECIMAL_LITERAL unit?
    | FLOAT_LITERAL unit?
    | IDENTIFIER (',' IDENTIFIER)?
    | STRING_LITERAL
    ;

unit: 'rem'
    | 'px'
    | 'em'
    ;

decalartions
    : configDeclaration
    | flowDeclaration
    | pageDeclaration
    | styleDeclaration
    | componentDeclaration
    | libraryDeclaration
    | layoutDecalaration
    ;

// Flow
flowDeclaration: FLOW IDENTIFIER LBRACE flowBodyDeclaration* RBRACE;

flowBodyDeclaration
    : seeDeclaration
    | doDeclaration
    | reactDeclaration
    ;

seeDeclaration: SEE (IDENTIFIER | STRING_LITERAL DOT componentName);
doDeclaration: DO LBRACK actionName RBRACK STRING_LITERAL DOT componentName ;
reactDeclaration: REACT sceneName? COLON actionKey animateDeclaration?;

animateDeclaration: WITHTEXT ANIMATE LPAREN animateName RPAREN;

actionKey: GOTO_KEY componentName | SHOW_KEY STRING_LITERAL DOT componentName;

actionName: IDENTIFIER;
componentValue: IDENTIFIER;
componentName: IDENTIFIER;
sceneName: IDENTIFIER;
animateName: IDENTIFIER;

GOTO_KEY: 'goto' | 'GOTO' | '跳转';
SHOW_KEY: 'show' | 'SHOW' | '展示';

FLOW: 'flow' | '流' ;
SEE: 'see' | 'SEE' | '看到';
DO: 'do' | 'DO' | '做';
REACT: 'react' | 'REACT' | '响应';

WITHTEXT: 'with' | 'WITH' | '使用';
ANIMATE: 'animate' | 'ANIMATE' | '动画';

//PAGE

pageDeclaration: PAGE IDENTIFIER LBRACE componentBodyDeclaration* RBRACE;
componentDeclaration: COMPONENT IDENTIFIER LBRACE componentBodyDeclaration* RBRACE;

componentBodyDeclaration
    : componentName (',' componentName)*
    | IDENTIFIER COLON configValue
    ;

layoutDecalaration: LAYOUT IDENTIFIER LBRACE layoutBodyDeclaration RBRACE;


REPEAT: 'repeat';
REPEAT_TIMES: INTEGER;

layoutBodyDeclaration: layoutRow*;

layoutRow
    : '-' '-'*
    | layoutLine layoutLine*  '|'
    ;

layoutLine: '|' componentUseDeclaration;

componentUseDeclaration
    : DECIMAL_LITERAL
    | POSITION
    | componentName (LPAREN (DIGITS_IDENTIFIER | POSITION | STRING_LITERAL) RPAREN)?
    | STRING_LITERAL
    ;

LAYOUT: 'layout' | 'Layout' | '布局';
POSITION: 'LEFT' | 'RIGHT' | 'TOP' | 'BOTTOM';

PAGE: 'page' | 'PAGE' | '页面';
COMPONENT: 'component' | 'COMPONENT' | '组件';

// STYLE

styleDeclaration: STYLE styleName LBRACE styleBody RBRACE;

styleName: IDENTIFIER;
styleBody: (configDeclaration ';')*;

STYLE: 'style' | 'STYLE' | 'CSS' | 'css';

// LIBRARY


libraryDeclaration: LIBRARY libraryName LBRACE libraryBody RBRACE;
libraryBody: express*;

express
    : configKey '=' configValue ';'
    | IDENTIFIER LBRACK libraryCall RBRACK
    ;

libraryCall: libraryName DOT IDENTIFIER (',' libraryName DOT IDENTIFIER)*;

libraryName: IDENTIFIER;

LIBRARY: 'library' | 'LIBRARY' | '库';


// WORD

STRING_LITERAL:     '"' (~["\\\r\n] | EscapeSequence)* '"';
WS:                 [ \t\r\n\u000C]+ -> channel(HIDDEN);
COMMENT:            '/*' .*? '*/'    -> channel(HIDDEN);
LINE_COMMENT:       '//' ~[\r\n]*    -> channel(HIDDEN);

EmptyLine:          NewLine Space+ NewLine -> skip;
Space :             [ \t];
NewLine :           '\r\n' | '\n' | '\r';

LPAREN:             '(';
RPAREN:             ')';
LBRACE:             '{';
RBRACE:             '}';
LBRACK:             '[';
RBRACK:             ']';
Quote:              '"';
SingleQuote:        '\'';
COLON:              ':';
DOT:                '.';
COMMA:              ',';


LETTER:             Letter;
IDENTIFIER:         Letter LetterOrDigit*;
DIGITS:             Digits;
DIGITS_IDENTIFIER:  LetterOrDigit LetterOrDigit*;

DECIMAL_LITERAL:    ('0' | [1-9] (Digits? | '_'+ Digits)) [lL]?;

fragment DIGIT
    :'0'..'9'
    ;

fragment ExponentPart
    : [eE] [+-]? Digits
    ;


FLOAT_LITERAL
    : (Digits '.' Digits? | '.' Digits) ExponentPart? [fFdD]?
    | Digits (ExponentPart [fFdD]? | [fFdD])
    ;

fragment INTEGER
    : DIGIT+
    ;

fragment EscapeSequence
    : '\\' [btnfr"'\\]
    | '\\' ([0-3]? [0-7])? [0-7]
    | '\\' 'u'+ HexDigit HexDigit HexDigit HexDigit
    ;

fragment HexDigit
    : [0-9a-fA-F]
    ;

fragment Digits
    : [0-9] ([0-9_]* [0-9])?
    | ('0' .. '9') + ('.' ('0' .. '9') +)?
    ;

fragment LetterOrDigit
    : Letter
    | Digits
    ;

fragment Letter
    : [a-zA-Z$_] // these are the "java letters" below 0x7F
    | ~[\u0000-\u007F\uD800-\uDBFF] // covers all characters above 0x7F which are not a surrogate
    | [\uD800-\uDBFF] [\uDC00-\uDFFF] // covers UTF-16 surrogate pairs encodings for U+10000 to U+10FFFF
    ;
