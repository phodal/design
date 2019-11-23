grammar Design;

start: (comment | configDecalartion | decalartions)* EOF;

comment: IDENTIFIER;

configDecalartion: configKey COLON configValue;

configKey: IDENTIFIER;
configValue
    : IDENTIFIER
    | IDENTIFIER (',' IDENTIFIER)?
    ;

decalartions
    : configDecalartion
    | flowDecalartion
    | pageDecalartion
    | styleDecalartion
    | componentDecalartion
    | libraryDecalartion
    ;

// Flow
flowDecalartion: FLOW IDENTIFIER LBRACE flowBodyDecalartion* RBRACE;

flowBodyDecalartion
    : seeDecalartion
    | doDecalartion
    | reactDecalartion
    ;

seeDecalartion: SEE (IDENTIFIER | STRING_LITERAL DOT componentName);
doDecalartion: DO LBRACK actionName RBRACK STRING_LITERAL DOT componentName ;
reactDecalartion: REACT sceneName? COLON actionKey animateDecalartion?;

animateDecalartion: WITHTEXT ANIMATE LPAREN animateName RPAREN;

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

pageDecalartion: PAGE IDENTIFIER LBRACE componentBodyDecalartion* RBRACE;
componentDecalartion: COMPONENT IDENTIFIER LBRACE componentBodyDecalartion* RBRACE;

componentBodyDecalartion: IDENTIFIER (COLON configValue | layoutDecalaration);

layoutDecalaration: LBRACE layoutBodyDecalartion* RBRACE;

layoutBodyDecalartion: '|' (emptyLine | layoutLine);

emptyLine:  '-' ('|' | '-' )*;
layoutLine: ('|' | componentUseDeclaration)*;

componentUseDeclaration
    : GridSize
    | componentName (LPAREN (GridSize | STRING_LITERAL) RPAREN)?
    | STRING_LITERAL
    ;


GridSize: Digits | POSITION;

POSITION: 'LEFT' | 'RIGHT' | 'TOP' | 'BOTTOM';

PAGE: 'page' | 'PAGE' | '页面';
COMPONENT: 'component' | 'COMPONENT' | '组件';

// STYLE

styleDecalartion: STYLE styleName LBRACE styleBody RBRACE;

styleName: IDENTIFIER;
styleBody: (configDecalartion ';')*;

STYLE: 'style' | 'STYLE' | 'CSS' | 'css';

// LIBRARY


libraryDecalartion: LIBRARAY LBRACE libraryBody RBRACE;
libraryBody: STRING_LITERAL;

LIBRARAY: 'libraray' | 'LIBRARAY' | '库';


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
DIGITS:             Digits;
IDENTIFIER:         Letter LetterOrDigit*;
DIGITS_IDENTIFIER:  LetterOrDigit LetterOrDigit*;

CONFIG_VALUE: DECIMAL_LITERAL | Letter LetterOrDigit*;

DECIMAL_LITERAL: ('0' | [1-9] (Digits? | '_'+ Digits));

fragment DIGIT
    :'0'..'9'
    ;

fragment INTEGER
    :DIGIT+
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
