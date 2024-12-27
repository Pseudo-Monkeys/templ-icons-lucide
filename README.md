Templ Icon Components (Lucide Icons)
====================================

![some templ icons](https://imgur.com/a/OmqgSW7)
Welcome to Icon Components Library for [Templ](https://templ.guide) in GOLang.
It aims to allow easy access to a multitude of icons in your templ components.

## Import them and use them like any other templ components with our custom props object:

```go
import "github.com/Pseudo-Monkeys/templ-icons-lucide/lucide"
import "github.com/Pseudo-Monkeys/templ-icons-lucide/lib"

templ Demo(){
    @lucide.ChevronLeft(lib.IconProps{})
}
```

## You can add your own attributes to the svg, just use the IconProps

```go
import "github.com/Pseudo-Monkeys/templ-icons-lucide/lucide"
import "github.com/Pseudo-Monkeys/templ-icons-lucide/lib"

templ Demo(){
    @lucide.ChevronLeft(lib.IconProps{Size:"24", Fill:"red", StrokeWidth:"12"})
}
```
