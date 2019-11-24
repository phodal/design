# design

Design as Code

## DSL

Task Flows / User Flows Example

```flow
flow login {
    SEE HomePage
    DO [Click] "Login".Button
        REACT Success: SHOW "Login Success".Toast with ANIMATE(bounce)
        REACT Failure: SHOW "Login Failure".Dialog

    SEE "Login Failure".Dialog
    DO [Click] "ForgotPassword".Button
        REACT: GOTO ForgotPasswordPage

    SEE ForgotPasswordPage
    DO [Click] "RESET PASSWORD".Button
        REACT: SHOW "Please Check Email".Message
}
```

Elements
 
```
page HomePage {
    LayoutGrid: 12x
    LayoutId: HomePage
    Router: "/home"
}

component Navigation {
    LayoutId: Navigation
}

component TitleComponent {}
component ImageComponent {
    Size: 1080px
}
component BlogList {
    BlogDetail, Space8
    BlogDetail, Space8
    BlogDetail, Space8
    BlogDetail, Space8
    BlogDetail, Space8
    BlogDetail, Space8
}
```

Layout Examples:

```
Layout HomePage {
--------------------------
|      Navigation(RIGHT) |
--------------------------
| Empty(2x) | TitleComponent | Empty(2x) |
--------------------------
|     ImageComponent     |
--------------------------
| BlogList(8x)  | Archives(2x) |
--------------------------
| Footer                 |
--------------------------
}

Layout Navigation {
--------------------------------------
| "home" |"detail" | Button("Login") |
--------------------------------------
}
```

Library Define Examples

```
library FontSize {
    H1 = 18px;
    H2 = 16px;
    H3 = 14px;
    H4 = 12px;
    H5 = 10px;
}

library Color {
    Primary = 4.3rem;
    Secondary = 3.4rem;
}

library Button {
    Default [
        FontSize.H2, Color.Primary
    ]
    Primary [
        FontSize.H2, Color.Primary
    ]
}
```


License
---

[![Phodal's Idea](http://brand.phodal.com/shields/idea-small.svg)](http://ideas.phodal.com/)

@ 2019 A [Phodal Huang](https://www.phodal.com)'s [Idea](http://github.com/phodal/ideas).  This code is distributed under the MIT license. See `LICENSE` in this directory.
