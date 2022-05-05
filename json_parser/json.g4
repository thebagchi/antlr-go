grammar json;

@parser::header {

}

start returns [interface{} v]
   : p1=value EOF {
       $v = $p1.v
   }
   ;

dict returns [map[string]interface{} v]
   : CURLY_START p1=pairs CURLY_END {
       $v = $p1.v
   }
   | CURLY_START CURLY_START {
       $v = map[string]interface{} {
           // EMPTY
       }
   }
   ;

pairs returns [map[string]interface{} v]
   : p1=pairs COMMA p2=pair {
       $v = $p1.v
       for k, v := range $p2.v {
          $v[k] = v
       }
   }
   | p3=pair {
       $v = $p3.v
   }
   ;

pair returns [map[string]interface{} v]
   : p1=STRING COLON p2=value {
       fmt.Println($p1)
       fmt.Println($p2.v)
       $v = map[string]interface{} {
           $p1.text : $p2.v,
       }
   }
   ;

list returns[[]interface{} v]
   : SQUARE_START p1=items SQUARE_END {
       $v = $p1.v
   }
   | SQUARE_START SQUARE_END {
       $v = []interface{} {
           // EMPTY
       }
   }
   ;

items returns[[]interface{} v]
   : p1=items COMMA p2=value {
       $v = $p1.v
       $v = append($v, $p2.v)
   }
   | p3=value {
       $v = []interface{} {
           $p3.v,
       }
   }
   ;

value returns [interface{} v]
   : p1=STRING {
       $v = $p1.text
   }
   | p2=NUMBER {
       v, _ := strconv.ParseFloat($p2.text, 64)
       $v = v
   }
   | p3=dict {
       $v = $p3.v
   }
   | p4=list {
       $v = $p4.v
   }
   | p5=boolean {
       $v = $p5.v
   }
   | null {
       $v = nil
   }
   ;

boolean returns [bool v]
   : 'true' {
       $v = true
   }
   | 'false' {
       $v = false
   }
   ;

null
   : 'null'
   ;

CURLY_START
   : '{'
   ;

CURLY_END
   : '}'
   ;

COLON
   : ':'
   ;

SQUARE_START
   : '['
   ;

SQUARE_END
   : ']'
   ;

COMMA
   : ','
   ;

STRING
   : '"' (ESC | SAFECODEPOINT)* '"'
   ;

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;

fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;

fragment HEX
   : [0-9a-fA-F]
   ;

fragment SAFECODEPOINT
   : ~ ["\\\u0000-\u001F]
   ;

NUMBER
   : '-'? INT ('.' [0-9] +)? EXP?
   ;

fragment INT
   : '0' | [1-9] [0-9]*
   ;

fragment EXP
   : [Ee] [+\-]? INT
   ;

WHITESPACE
   : [ \t\n\r] + -> skip
   ;