// Generated from .\prql.g4 by ANTLR 4.9.2
// jshint ignore: start
import antlr4 from "antlr4";
import prqlListener from "./prqlListener.js";

const serializedATN = [
  "\u0003\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786",
  "\u5964\u00034\u0140\u0004\u0002\t\u0002\u0004\u0003\t\u0003\u0004\u0004",
  "\t\u0004\u0004\u0005\t\u0005\u0004\u0006\t\u0006\u0004\u0007\t\u0007",
  "\u0004\b\t\b\u0004\t\t\t\u0004\n\t\n\u0004\u000b\t\u000b\u0004\f\t\f",
  "\u0004\r\t\r\u0004\u000e\t\u000e\u0004\u000f\t\u000f\u0004\u0010\t\u0010",
  "\u0004\u0011\t\u0011\u0004\u0012\t\u0012\u0004\u0013\t\u0013\u0004\u0014",
  "\t\u0014\u0004\u0015\t\u0015\u0004\u0016\t\u0016\u0004\u0017\t\u0017",
  "\u0004\u0018\t\u0018\u0004\u0019\t\u0019\u0004\u001a\t\u001a\u0004\u001b",
  "\t\u001b\u0003\u0002\u0003\u0002\u0003\u0003\u0007\u0003:\n\u0003\f",
  "\u0003\u000e\u0003=\u000b\u0003\u0003\u0003\u0005\u0003@\n\u0003\u0003",
  "\u0003\u0007\u0003C\n\u0003\f\u0003\u000e\u0003F\u000b\u0003\u0003\u0003",
  "\u0003\u0003\u0003\u0003\u0005\u0003K\n\u0003\u0003\u0003\u0007\u0003",
  "N\n\u0003\f\u0003\u000e\u0003Q\u000b\u0003\u0007\u0003S\n\u0003\f\u0003",
  "\u000e\u0003V\u000b\u0003\u0003\u0003\u0003\u0003\u0003\u0004\u0003",
  "\u0004\u0007\u0004\\\n\u0004\f\u0004\u000e\u0004_\u000b\u0004\u0003",
  "\u0004\u0003\u0004\u0003\u0005\u0003\u0005\u0003\u0005\u0003\u0005\u0003",
  "\u0005\u0003\u0005\u0003\u0006\u0003\u0006\u0005\u0006k\n\u0006\u0003",
  "\u0007\u0007\u0007n\n\u0007\f\u0007\u000e\u0007q\u000b\u0007\u0003\b",
  "\u0003\b\u0005\bu\n\b\u0003\b\u0005\bx\n\b\u0003\t\u0003\t\u0003\t\u0003",
  "\t\u0007\t~\n\t\f\t\u000e\t\u0081\u000b\t\u0003\t\u0003\t\u0003\n\u0003",
  "\n\u0005\n\u0087\n\n\u0003\u000b\u0003\u000b\u0003\f\u0003\f\u0003\f",
  "\u0003\f\u0003\f\u0003\r\u0003\r\u0003\r\u0003\r\u0007\r\u0094\n\r\f",
  "\r\u000e\r\u0097\u000b\r\u0003\r\u0003\r\u0005\r\u009b\n\r\u0003\u000e",
  "\u0003\u000e\u0003\u000e\u0007\u000e\u00a0\n\u000e\f\u000e\u000e\u000e",
  "\u00a3\u000b\u000e\u0003\u000f\u0003\u000f\u0007\u000f\u00a7\n\u000f",
  "\f\u000f\u000e\u000f\u00aa\u000b\u000f\u0003\u000f\u0003\u000f\u0003",
  "\u0010\u0003\u0010\u0007\u0010\u00b0\n\u0010\f\u0010\u000e\u0010\u00b3",
  "\u000b\u0010\u0003\u0011\u0003\u0011\u0003\u0011\u0005\u0011\u00b8\n",
  "\u0011\u0003\u0012\u0003\u0012\u0003\u0012\u0003\u0012\u0005\u0012\u00be",
  "\n\u0012\u0003\u0013\u0003\u0013\u0003\u0013\u0003\u0013\u0003\u0014",
  "\u0003\u0014\u0003\u0014\u0003\u0014\u0003\u0015\u0003\u0015\u0005\u0015",
  "\u00ca\n\u0015\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003",
  "\u0016\u0003\u0016\u0005\u0016\u00d2\n\u0016\u0003\u0016\u0003\u0016",
  "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016",
  "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016",
  "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016",
  "\u0003\u0016\u0007\u0016\u00e9\n\u0016\f\u0016\u000e\u0016\u00ec\u000b",
  "\u0016\u0003\u0017\u0003\u0017\u0003\u0017\u0003\u0017\u0003\u0017\u0005",
  "\u0017\u00f3\n\u0017\u0003\u0018\u0003\u0018\u0003\u0018\u0003\u0018",
  "\u0005\u0018\u00f9\n\u0018\u0003\u0019\u0003\u0019\u0003\u0019\u0003",
  "\u0019\u0003\u0019\u0003\u0019\u0003\u0019\u0003\u0019\u0003\u0019\u0003",
  "\u0019\u0003\u0019\u0005\u0019\u0106\n\u0019\u0003\u001a\u0003\u001a",
  "\u0007\u001a\u010a\n\u001a\f\u001a\u000e\u001a\u010d\u000b\u001a\u0003",
  "\u001a\u0003\u001a\u0005\u001a\u0111\n\u001a\u0003\u001a\u0003\u001a",
  "\u0007\u001a\u0115\n\u001a\f\u001a\u000e\u001a\u0118\u000b\u001a\u0003",
  "\u001a\u0003\u001a\u0005\u001a\u011c\n\u001a\u0007\u001a\u011e\n\u001a",
  "\f\u001a\u000e\u001a\u0121\u000b\u001a\u0003\u001a\u0005\u001a\u0124",
  "\n\u001a\u0003\u001a\u0005\u001a\u0127\n\u001a\u0005\u001a\u0129\n\u001a",
  "\u0003\u001a\u0003\u001a\u0003\u001b\u0003\u001b\u0007\u001b\u012f\n",
  "\u001b\f\u001b\u000e\u001b\u0132\u000b\u001b\u0003\u001b\u0003\u001b",
  "\u0005\u001b\u0136\n\u001b\u0003\u001b\u0007\u001b\u0139\n\u001b\f\u001b",
  "\u000e\u001b\u013c\u000b\u001b\u0003\u001b\u0003\u001b\u0003\u001b\u0002",
  "\u0003*\u001c\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016",
  '\u0018\u001a\u001c\u001e "$&(*,.024\u0002\n\u0003\u000212\u0004\u0002',
  "  11\u0004\u0002\n\n\f\r\u0003\u0002\b\t\u0004\u0002\u000f\u0012\u0019",
  "\u001a\u0003\u0002%&\u0004\u0002\b\t''\u0003\u0002,-\u0002\u015c\u0002",
  "6\u0003\u0002\u0002\u0002\u0004;\u0003\u0002\u0002\u0002\u0006Y\u0003",
  "\u0002\u0002\u0002\bb\u0003\u0002\u0002\u0002\nh\u0003\u0002\u0002\u0002",
  "\fo\u0003\u0002\u0002\u0002\u000et\u0003\u0002\u0002\u0002\u0010y\u0003",
  "\u0002\u0002\u0002\u0012\u0084\u0003\u0002\u0002\u0002\u0014\u0088\u0003",
  "\u0002\u0002\u0002\u0016\u008a\u0003\u0002\u0002\u0002\u0018\u008f\u0003",
  "\u0002\u0002\u0002\u001a\u009c\u0003\u0002\u0002\u0002\u001c\u00a4\u0003",
  "\u0002\u0002\u0002\u001e\u00ad\u0003\u0002\u0002\u0002 \u00b7\u0003",
  '\u0002\u0002\u0002"\u00b9\u0003\u0002\u0002\u0002$\u00bf\u0003\u0002',
  "\u0002\u0002&\u00c3\u0003\u0002\u0002\u0002(\u00c9\u0003\u0002\u0002",
  "\u0002*\u00d1\u0003\u0002\u0002\u0002,\u00f2\u0003\u0002\u0002\u0002",
  ".\u00f4\u0003\u0002\u0002\u00020\u0105\u0003\u0002\u0002\u00022\u0107",
  "\u0003\u0002\u0002\u00024\u012c\u0003\u0002\u0002\u000267\t\u0002\u0002",
  "\u00027\u0003\u0003\u0002\u0002\u00028:\u0005\u0002\u0002\u000298\u0003",
  "\u0002\u0002\u0002:=\u0003\u0002\u0002\u0002;9\u0003\u0002\u0002\u0002",
  ";<\u0003\u0002\u0002\u0002<?\u0003\u0002\u0002\u0002=;\u0003\u0002\u0002",
  "\u0002>@\u0005\u0006\u0004\u0002?>\u0003\u0002\u0002\u0002?@\u0003\u0002",
  "\u0002\u0002@D\u0003\u0002\u0002\u0002AC\u0005\u0002\u0002\u0002BA\u0003",
  "\u0002\u0002\u0002CF\u0003\u0002\u0002\u0002DB\u0003\u0002\u0002\u0002",
  "DE\u0003\u0002\u0002\u0002ET\u0003\u0002\u0002\u0002FD\u0003\u0002\u0002",
  "\u0002GK\u0005\b\u0005\u0002HK\u0005\u0014\u000b\u0002IK\u0005\u0018",
  "\r\u0002JG\u0003\u0002\u0002\u0002JH\u0003\u0002\u0002\u0002JI\u0003",
  "\u0002\u0002\u0002KO\u0003\u0002\u0002\u0002LN\u0005\u0002\u0002\u0002",
  "ML\u0003\u0002\u0002\u0002NQ\u0003\u0002\u0002\u0002OM\u0003\u0002\u0002",
  "\u0002OP\u0003\u0002\u0002\u0002PS\u0003\u0002\u0002\u0002QO\u0003\u0002",
  "\u0002\u0002RJ\u0003\u0002\u0002\u0002SV\u0003\u0002\u0002\u0002TR\u0003",
  "\u0002\u0002\u0002TU\u0003\u0002\u0002\u0002UW\u0003\u0002\u0002\u0002",
  "VT\u0003\u0002\u0002\u0002WX\u0007\u0002\u0002\u0003X\u0005\u0003\u0002",
  '\u0002\u0002Y]\u0007\u0004\u0002\u0002Z\\\u0005"\u0012\u0002[Z\u0003',
  "\u0002\u0002\u0002\\_\u0003\u0002\u0002\u0002][\u0003\u0002\u0002\u0002",
  "]^\u0003\u0002\u0002\u0002^`\u0003\u0002\u0002\u0002_]\u0003\u0002\u0002",
  "\u0002`a\u0005\u0002\u0002\u0002a\u0007\u0003\u0002\u0002\u0002bc\u0007",
  "\u0003\u0002\u0002cd\u0005\n\u0006\u0002de\u0005\f\u0007\u0002ef\u0007",
  "\u0006\u0002\u0002fg\u0005*\u0016\u0002g\t\u0003\u0002\u0002\u0002h",
  "j\u0007-\u0002\u0002ik\u0005\u0010\t\u0002ji\u0003\u0002\u0002\u0002",
  "jk\u0003\u0002\u0002\u0002k\u000b\u0003\u0002\u0002\u0002ln\u0005\u000e",
  "\b\u0002ml\u0003\u0002\u0002\u0002nq\u0003\u0002\u0002\u0002om\u0003",
  "\u0002\u0002\u0002op\u0003\u0002\u0002\u0002p\r\u0003\u0002\u0002\u0002",
  'qo\u0003\u0002\u0002\u0002ru\u0005"\u0012\u0002su\u0007-\u0002\u0002',
  "tr\u0003\u0002\u0002\u0002ts\u0003\u0002\u0002\u0002uw\u0003\u0002\u0002",
  "\u0002vx\u0005\u0010\t\u0002wv\u0003\u0002\u0002\u0002wx\u0003\u0002",
  "\u0002\u0002x\u000f\u0003\u0002\u0002\u0002yz\u0007\u0019\u0002\u0002",
  "z{\u0005\u0012\n\u0002{\u007f\u0007\u0013\u0002\u0002|~\u0005\u0012",
  "\n\u0002}|\u0003\u0002\u0002\u0002~\u0081\u0003\u0002\u0002\u0002\u007f",
  "}\u0003\u0002\u0002\u0002\u007f\u0080\u0003\u0002\u0002\u0002\u0080",
  "\u0082\u0003\u0002\u0002\u0002\u0081\u007f\u0003\u0002\u0002\u0002\u0082",
  "\u0083\u0007\u001a\u0002\u0002\u0083\u0011\u0003\u0002\u0002\u0002\u0084",
  "\u0086\u0007-\u0002\u0002\u0085\u0087\u0005\u0010\t\u0002\u0086\u0085",
  "\u0003\u0002\u0002\u0002\u0086\u0087\u0003\u0002\u0002\u0002\u0087\u0013",
  "\u0003\u0002\u0002\u0002\u0088\u0089\u0005\u0016\f\u0002\u0089\u0015",
  "\u0003\u0002\u0002\u0002\u008a\u008b\u0007\u0005\u0002\u0002\u008b\u008c",
  "\u0007-\u0002\u0002\u008c\u008d\u0007\u0007\u0002\u0002\u008d\u008e",
  "\u0005*\u0016\u0002\u008e\u0017\u0003\u0002\u0002\u0002\u008f\u0095",
  "\u0005(\u0015\u0002\u0090\u0091\u0005\u0002\u0002\u0002\u0091\u0092",
  "\u0005\u001e\u0010\u0002\u0092\u0094\u0003\u0002\u0002\u0002\u0093\u0090",
  "\u0003\u0002\u0002\u0002\u0094\u0097\u0003\u0002\u0002\u0002\u0095\u0093",
  "\u0003\u0002\u0002\u0002\u0095\u0096\u0003\u0002\u0002\u0002\u0096\u009a",
  "\u0003\u0002\u0002\u0002\u0097\u0095\u0003\u0002\u0002\u0002\u0098\u009b",
  "\u0005\u0002\u0002\u0002\u0099\u009b\u0007\u0002\u0002\u0003\u009a\u0098",
  "\u0003\u0002\u0002\u0002\u009a\u0099\u0003\u0002\u0002\u0002\u009b\u0019",
  "\u0003\u0002\u0002\u0002\u009c\u00a1\u0005(\u0015\u0002\u009d\u009e",
  "\u0007\u0013\u0002\u0002\u009e\u00a0\u0005\u001e\u0010\u0002\u009f\u009d",
  "\u0003\u0002\u0002\u0002\u00a0\u00a3\u0003\u0002\u0002\u0002\u00a1\u009f",
  "\u0003\u0002\u0002\u0002\u00a1\u00a2\u0003\u0002\u0002\u0002\u00a2\u001b",
  "\u0003\u0002\u0002\u0002\u00a3\u00a1\u0003\u0002\u0002\u0002\u00a4\u00a8",
  "\u0007 \u0002\u0002\u00a5\u00a7\n\u0003\u0002\u0002\u00a6\u00a5\u0003",
  "\u0002\u0002\u0002\u00a7\u00aa\u0003\u0002\u0002\u0002\u00a8\u00a6\u0003",
  "\u0002\u0002\u0002\u00a8\u00a9\u0003\u0002\u0002\u0002\u00a9\u00ab\u0003",
  "\u0002\u0002\u0002\u00aa\u00a8\u0003\u0002\u0002\u0002\u00ab\u00ac\u0007",
  " \u0002\u0002\u00ac\u001d\u0003\u0002\u0002\u0002\u00ad\u00b1\u0007",
  "-\u0002\u0002\u00ae\u00b0\u0005 \u0011\u0002\u00af\u00ae\u0003\u0002",
  "\u0002\u0002\u00b0\u00b3\u0003\u0002\u0002\u0002\u00b1\u00af\u0003\u0002",
  "\u0002\u0002\u00b1\u00b2\u0003\u0002\u0002\u0002\u00b2\u001f\u0003\u0002",
  '\u0002\u0002\u00b3\u00b1\u0003\u0002\u0002\u0002\u00b4\u00b8\u0005"',
  "\u0012\u0002\u00b5\u00b8\u0005$\u0013\u0002\u00b6\u00b8\u0005*\u0016",
  "\u0002\u00b7\u00b4\u0003\u0002\u0002\u0002\u00b7\u00b5\u0003\u0002\u0002",
  "\u0002\u00b7\u00b6\u0003\u0002\u0002\u0002\u00b8!\u0003\u0002\u0002",
  "\u0002\u00b9\u00ba\u0007-\u0002\u0002\u00ba\u00bd\u0007\u0014\u0002",
  "\u0002\u00bb\u00be\u0005$\u0013\u0002\u00bc\u00be\u0005*\u0016\u0002",
  "\u00bd\u00bb\u0003\u0002\u0002\u0002\u00bd\u00bc\u0003\u0002\u0002\u0002",
  "\u00be#\u0003\u0002\u0002\u0002\u00bf\u00c0\u0007-\u0002\u0002\u00c0",
  "\u00c1\u0007\u0007\u0002\u0002\u00c1\u00c2\u0005*\u0016\u0002\u00c2",
  "%\u0003\u0002\u0002\u0002\u00c3\u00c4\u0007-\u0002\u0002\u00c4\u00c5",
  "\u0007\u0007\u0002\u0002\u00c5\u00c6\u0005(\u0015\u0002\u00c6'\u0003",
  "\u0002\u0002\u0002\u00c7\u00ca\u0005*\u0016\u0002\u00c8\u00ca\u0005",
  "\u001e\u0010\u0002\u00c9\u00c7\u0003\u0002\u0002\u0002\u00c9\u00c8\u0003",
  "\u0002\u0002\u0002\u00ca)\u0003\u0002\u0002\u0002\u00cb\u00cc\b\u0016",
  "\u0001\u0002\u00cc\u00cd\u0007\u001d\u0002\u0002\u00cd\u00ce\u0005*",
  "\u0016\u0002\u00ce\u00cf\u0007\u001e\u0002\u0002\u00cf\u00d2\u0003\u0002",
  "\u0002\u0002\u00d0\u00d2\u0005,\u0017\u0002\u00d1\u00cb\u0003\u0002",
  "\u0002\u0002\u00d1\u00d0\u0003\u0002\u0002\u0002\u00d2\u00ea\u0003\u0002",
  "\u0002\u0002\u00d3\u00d4\f\u000b\u0002\u0002\u00d4\u00d5\t\u0004\u0002",
  "\u0002\u00d5\u00e9\u0005*\u0016\f\u00d6\u00d7\f\n\u0002\u0002\u00d7",
  "\u00d8\t\u0005\u0002\u0002\u00d8\u00e9\u0005*\u0016\u000b\u00d9\u00da",
  "\f\t\u0002\u0002\u00da\u00db\u0007\u000b\u0002\u0002\u00db\u00e9\u0005",
  "*\u0016\n\u00dc\u00dd\f\b\u0002\u0002\u00dd\u00de\u0007\u000e\u0002",
  "\u0002\u00de\u00e9\u0005*\u0016\t\u00df\u00e0\f\u0007\u0002\u0002\u00e0",
  "\u00e1\t\u0006\u0002\u0002\u00e1\u00e9\u0005*\u0016\b\u00e2\u00e3\f",
  "\u0006\u0002\u0002\u00e3\u00e4\u0007(\u0002\u0002\u00e4\u00e9\u0005",
  "*\u0016\u0007\u00e5\u00e6\f\u0005\u0002\u0002\u00e6\u00e7\t\u0007\u0002",
  "\u0002\u00e7\u00e9\u0005*\u0016\u0006\u00e8\u00d3\u0003\u0002\u0002",
  "\u0002\u00e8\u00d6\u0003\u0002\u0002\u0002\u00e8\u00d9\u0003\u0002\u0002",
  "\u0002\u00e8\u00dc\u0003\u0002\u0002\u0002\u00e8\u00df\u0003\u0002\u0002",
  "\u0002\u00e8\u00e2\u0003\u0002\u0002\u0002\u00e8\u00e5\u0003\u0002\u0002",
  "\u0002\u00e9\u00ec\u0003\u0002\u0002\u0002\u00ea\u00e8\u0003\u0002\u0002",
  "\u0002\u00ea\u00eb\u0003\u0002\u0002\u0002\u00eb+\u0003\u0002\u0002",
  "\u0002\u00ec\u00ea\u0003\u0002\u0002\u0002\u00ed\u00f3\u00050\u0019",
  "\u0002\u00ee\u00f3\u0005\u001c\u000f\u0002\u00ef\u00f3\u0005.\u0018",
  "\u0002\u00f0\u00f3\u00052\u001a\u0002\u00f1\u00f3\u00054\u001b\u0002",
  "\u00f2\u00ed\u0003\u0002\u0002\u0002\u00f2\u00ee\u0003\u0002\u0002\u0002",
  "\u00f2\u00ef\u0003\u0002\u0002\u0002\u00f2\u00f0\u0003\u0002\u0002\u0002",
  "\u00f2\u00f1\u0003\u0002\u0002\u0002\u00f3-\u0003\u0002\u0002\u0002",
  "\u00f4\u00f8\t\b\u0002\u0002\u00f5\u00f9\u00054\u001b\u0002\u00f6\u00f9",
  "\u00050\u0019\u0002\u00f7\u00f9\u0007-\u0002\u0002\u00f8\u00f5\u0003",
  "\u0002\u0002\u0002\u00f8\u00f6\u0003\u0002\u0002\u0002\u00f8\u00f7\u0003",
  "\u0002\u0002\u0002\u00f9/\u0003\u0002\u0002\u0002\u00fa\u0106\u0007",
  "-\u0002\u0002\u00fb\u0106\u0007)\u0002\u0002\u00fc\u0106\u0007*\u0002",
  "\u0002\u00fd\u0106\u00074\u0002\u0002\u00fe\u0106\u0007+\u0002\u0002",
  "\u00ff\u0106\u0007,\u0002\u0002\u0100\u0101\u0007,\u0002\u0002\u0101",
  "\u0106\u00073\u0002\u0002\u0102\u0103\t\t\u0002\u0002\u0103\u0104\u0007",
  "\u0018\u0002\u0002\u0104\u0106\t\t\u0002\u0002\u0105\u00fa\u0003\u0002",
  "\u0002\u0002\u0105\u00fb\u0003\u0002\u0002\u0002\u0105\u00fc\u0003\u0002",
  "\u0002\u0002\u0105\u00fd\u0003\u0002\u0002\u0002\u0105\u00fe\u0003\u0002",
  "\u0002\u0002\u0105\u00ff\u0003\u0002\u0002\u0002\u0105\u0100\u0003\u0002",
  "\u0002\u0002\u0105\u0102\u0003\u0002\u0002\u0002\u01061\u0003\u0002",
  "\u0002\u0002\u0107\u0128\u0007\u001b\u0002\u0002\u0108\u010a\u0005\u0002",
  "\u0002\u0002\u0109\u0108\u0003\u0002\u0002\u0002\u010a\u010d\u0003\u0002",
  "\u0002\u0002\u010b\u0109\u0003\u0002\u0002\u0002\u010b\u010c\u0003\u0002",
  "\u0002\u0002\u010c\u0110\u0003\u0002\u0002\u0002\u010d\u010b\u0003\u0002",
  "\u0002\u0002\u010e\u0111\u0005&\u0014\u0002\u010f\u0111\u0005(\u0015",
  "\u0002\u0110\u010e\u0003\u0002\u0002\u0002\u0110\u010f\u0003\u0002\u0002",
  "\u0002\u0111\u011f\u0003\u0002\u0002\u0002\u0112\u0116\u0007\u0015\u0002",
  "\u0002\u0113\u0115\u0005\u0002\u0002\u0002\u0114\u0113\u0003\u0002\u0002",
  "\u0002\u0115\u0118\u0003\u0002\u0002\u0002\u0116\u0114\u0003\u0002\u0002",
  "\u0002\u0116\u0117\u0003\u0002\u0002\u0002\u0117\u011b\u0003\u0002\u0002",
  "\u0002\u0118\u0116\u0003\u0002\u0002\u0002\u0119\u011c\u0005&\u0014",
  "\u0002\u011a\u011c\u0005(\u0015\u0002\u011b\u0119\u0003\u0002\u0002",
  "\u0002\u011b\u011a\u0003\u0002\u0002\u0002\u011c\u011e\u0003\u0002\u0002",
  "\u0002\u011d\u0112\u0003\u0002\u0002\u0002\u011e\u0121\u0003\u0002\u0002",
  "\u0002\u011f\u011d\u0003\u0002\u0002\u0002\u011f\u0120\u0003\u0002\u0002",
  "\u0002\u0120\u0123\u0003\u0002\u0002\u0002\u0121\u011f\u0003\u0002\u0002",
  "\u0002\u0122\u0124\u0007\u0015\u0002\u0002\u0123\u0122\u0003\u0002\u0002",
  "\u0002\u0123\u0124\u0003\u0002\u0002\u0002\u0124\u0126\u0003\u0002\u0002",
  "\u0002\u0125\u0127\u0005\u0002\u0002\u0002\u0126\u0125\u0003\u0002\u0002",
  "\u0002\u0126\u0127\u0003\u0002\u0002\u0002\u0127\u0129\u0003\u0002\u0002",
  "\u0002\u0128\u010b\u0003\u0002\u0002\u0002\u0128\u0129\u0003\u0002\u0002",
  "\u0002\u0129\u012a\u0003\u0002\u0002\u0002\u012a\u012b\u0007\u001c\u0002",
  "\u0002\u012b3\u0003\u0002\u0002\u0002\u012c\u0130\u0007\u001d\u0002",
  "\u0002\u012d\u012f\u0005\u0002\u0002\u0002\u012e\u012d\u0003\u0002\u0002",
  "\u0002\u012f\u0132\u0003\u0002\u0002\u0002\u0130\u012e\u0003\u0002\u0002",
  "\u0002\u0130\u0131\u0003\u0002\u0002\u0002\u0131\u0135\u0003\u0002\u0002",
  "\u0002\u0132\u0130\u0003\u0002\u0002\u0002\u0133\u0136\u0005\u0018\r",
  "\u0002\u0134\u0136\u0005\u001a\u000e\u0002\u0135\u0133\u0003\u0002\u0002",
  "\u0002\u0135\u0134\u0003\u0002\u0002\u0002\u0136\u013a\u0003\u0002\u0002",
  "\u0002\u0137\u0139\u0005\u0002\u0002\u0002\u0138\u0137\u0003\u0002\u0002",
  "\u0002\u0139\u013c\u0003\u0002\u0002\u0002\u013a\u0138\u0003\u0002\u0002",
  "\u0002\u013a\u013b\u0003\u0002\u0002\u0002\u013b\u013d\u0003\u0002\u0002",
  "\u0002\u013c\u013a\u0003\u0002\u0002\u0002\u013d\u013e\u0007\u001e\u0002",
  "\u0002\u013e5\u0003\u0002\u0002\u0002(;?DJOT]jotw\u007f\u0086\u0095",
  "\u009a\u00a1\u00a8\u00b1\u00b7\u00bd\u00c9\u00d1\u00e8\u00ea\u00f2\u00f8",
  "\u0105\u010b\u0110\u0116\u011b\u011f\u0123\u0126\u0128\u0130\u0135\u013a",
].join("");

const atn = new antlr4.atn.ATNDeserializer().deserialize(serializedATN);

const decisionsToDFA = atn.decisionToState.map(
  (ds, index) => new antlr4.dfa.DFA(ds, index)
);

const sharedContextCache = new antlr4.PredictionContextCache();

export default class prqlParser extends antlr4.Parser {
  static grammarFileName = "prql.g4";
  static literalNames = [
    null,
    "'func'",
    "'prql'",
    "'let'",
    "'->'",
    "'='",
    "'+'",
    "'-'",
    "'*'",
    "'**'",
    "'/'",
    "'%'",
    "'~'",
    "'=='",
    "'!='",
    "'<='",
    "'>='",
    "'|'",
    "':'",
    "','",
    "'.'",
    "'$'",
    "'..'",
    "'<'",
    "'>'",
    "'['",
    "']'",
    "'('",
    "')'",
    "'_'",
    "'`'",
    "'\"'",
    "'''",
    '\'"""\'',
    "'''''",
    "'and'",
    "'or'",
    "'not'",
    "'??'",
    "'null'",
  ];
  static symbolicNames = [
    null,
    "FUNC",
    "PRQL",
    "LET",
    "ARROW",
    "ASSIGN",
    "PLUS",
    "MINUS",
    "STAR",
    "POW",
    "DIV",
    "MOD",
    "MODEL",
    "EQ",
    "NE",
    "LE",
    "GE",
    "BAR",
    "COLON",
    "COMMA",
    "DOT",
    "DOLLAR",
    "RANGE",
    "LANG",
    "RANG",
    "LBRACKET",
    "RBRACKET",
    "LPAREN",
    "RPAREN",
    "UNDERSCORE",
    "BACKTICK",
    "DOUBLE_QUOTE",
    "SINGLE_QUOTE",
    "TRIPLE_DOUBLE_QUOTE",
    "TRIPLE_SINGLE_QUOTE",
    "AND",
    "OR",
    "NOT",
    "COALESCE",
    "NULL_",
    "BOOLEAN",
    "INTEGER",
    "FLOAT",
    "IDENT",
    "IDENT_START",
    "IDENT_NEXT",
    "WHITESPACE",
    "NEWLINE",
    "COMMENT",
    "INTERVAL_KIND",
    "STRING",
  ];
  static ruleNames = [
    "nl",
    "program",
    "programIntro",
    "funcDef",
    "funcDefName",
    "funcDefParams",
    "funcDefParam",
    "typeDef",
    "typeTerm",
    "stmt",
    "assignStmt",
    "pipeline",
    "inlinePipeline",
    "identBacktick",
    "funcCall",
    "funcCallParam",
    "namedArg",
    "assign",
    "assignCall",
    "exprCall",
    "expr",
    "term",
    "exprUnary",
    "literal",
    "list",
    "nestedPipeline",
  ];

  constructor(input) {
    super(input);
    this._interp = new antlr4.atn.ParserATNSimulator(
      this,
      atn,
      decisionsToDFA,
      sharedContextCache
    );
    this.ruleNames = prqlParser.ruleNames;
    this.literalNames = prqlParser.literalNames;
    this.symbolicNames = prqlParser.symbolicNames;
  }

  get atn() {
    return atn;
  }

  sempred(localctx, ruleIndex, predIndex) {
    switch (ruleIndex) {
      case 20:
        return this.expr_sempred(localctx, predIndex);
      default:
        throw "No predicate with index:" + ruleIndex;
    }
  }

  expr_sempred(localctx, predIndex) {
    switch (predIndex) {
      case 0:
        return this.precpred(this._ctx, 9);
      case 1:
        return this.precpred(this._ctx, 8);
      case 2:
        return this.precpred(this._ctx, 7);
      case 3:
        return this.precpred(this._ctx, 6);
      case 4:
        return this.precpred(this._ctx, 5);
      case 5:
        return this.precpred(this._ctx, 4);
      case 6:
        return this.precpred(this._ctx, 3);
      default:
        throw "No predicate with index:" + predIndex;
    }
  }

  nl() {
    let localctx = new NlContext(this, this._ctx, this.state);
    this.enterRule(localctx, 0, prqlParser.RULE_nl);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 52;
      _la = this._input.LA(1);
      if (!(_la === prqlParser.NEWLINE || _la === prqlParser.COMMENT)) {
        this._errHandler.recoverInline(this);
      } else {
        this._errHandler.reportMatch(this);
        this.consume();
      }
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  program() {
    let localctx = new ProgramContext(this, this._ctx, this.state);
    this.enterRule(localctx, 2, prqlParser.RULE_program);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 57;
      this._errHandler.sync(this);
      var _alt = this._interp.adaptivePredict(this._input, 0, this._ctx);
      while (_alt != 2 && _alt != antlr4.atn.ATN.INVALID_ALT_NUMBER) {
        if (_alt === 1) {
          this.state = 54;
          this.nl();
        }
        this.state = 59;
        this._errHandler.sync(this);
        _alt = this._interp.adaptivePredict(this._input, 0, this._ctx);
      }

      this.state = 61;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      if (_la === prqlParser.PRQL) {
        this.state = 60;
        this.programIntro();
      }

      this.state = 66;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (_la === prqlParser.NEWLINE || _la === prqlParser.COMMENT) {
        this.state = 63;
        this.nl();
        this.state = 68;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
      }
      this.state = 82;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (
        ((_la & ~0x1f) == 0 &&
          ((1 << _la) &
            ((1 << prqlParser.FUNC) |
              (1 << prqlParser.LET) |
              (1 << prqlParser.PLUS) |
              (1 << prqlParser.MINUS) |
              (1 << prqlParser.LBRACKET) |
              (1 << prqlParser.LPAREN) |
              (1 << prqlParser.BACKTICK))) !==
            0) ||
        (((_la - 37) & ~0x1f) == 0 &&
          ((1 << (_la - 37)) &
            ((1 << (prqlParser.NOT - 37)) |
              (1 << (prqlParser.NULL_ - 37)) |
              (1 << (prqlParser.BOOLEAN - 37)) |
              (1 << (prqlParser.INTEGER - 37)) |
              (1 << (prqlParser.FLOAT - 37)) |
              (1 << (prqlParser.IDENT - 37)) |
              (1 << (prqlParser.STRING - 37)))) !==
            0)
      ) {
        this.state = 72;
        this._errHandler.sync(this);
        switch (this._input.LA(1)) {
          case prqlParser.FUNC:
            this.state = 69;
            this.funcDef();
            break;
          case prqlParser.LET:
            this.state = 70;
            this.stmt();
            break;
          case prqlParser.PLUS:
          case prqlParser.MINUS:
          case prqlParser.LBRACKET:
          case prqlParser.LPAREN:
          case prqlParser.BACKTICK:
          case prqlParser.NOT:
          case prqlParser.NULL_:
          case prqlParser.BOOLEAN:
          case prqlParser.INTEGER:
          case prqlParser.FLOAT:
          case prqlParser.IDENT:
          case prqlParser.STRING:
            this.state = 71;
            this.pipeline();
            break;
          default:
            throw new antlr4.error.NoViableAltException(this);
        }
        this.state = 77;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
        while (_la === prqlParser.NEWLINE || _la === prqlParser.COMMENT) {
          this.state = 74;
          this.nl();
          this.state = 79;
          this._errHandler.sync(this);
          _la = this._input.LA(1);
        }
        this.state = 84;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
      }
      this.state = 85;
      this.match(prqlParser.EOF);
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  programIntro() {
    let localctx = new ProgramIntroContext(this, this._ctx, this.state);
    this.enterRule(localctx, 4, prqlParser.RULE_programIntro);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 87;
      this.match(prqlParser.PRQL);
      this.state = 91;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (_la === prqlParser.IDENT) {
        this.state = 88;
        this.namedArg();
        this.state = 93;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
      }
      this.state = 94;
      this.nl();
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  funcDef() {
    let localctx = new FuncDefContext(this, this._ctx, this.state);
    this.enterRule(localctx, 6, prqlParser.RULE_funcDef);
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 96;
      this.match(prqlParser.FUNC);
      this.state = 97;
      this.funcDefName();
      this.state = 98;
      this.funcDefParams();
      this.state = 99;
      this.match(prqlParser.ARROW);
      this.state = 100;
      this.expr(0);
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  funcDefName() {
    let localctx = new FuncDefNameContext(this, this._ctx, this.state);
    this.enterRule(localctx, 8, prqlParser.RULE_funcDefName);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 102;
      this.match(prqlParser.IDENT);
      this.state = 104;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      if (_la === prqlParser.LANG) {
        this.state = 103;
        this.typeDef();
      }
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  funcDefParams() {
    let localctx = new FuncDefParamsContext(this, this._ctx, this.state);
    this.enterRule(localctx, 10, prqlParser.RULE_funcDefParams);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 109;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (_la === prqlParser.IDENT) {
        this.state = 106;
        this.funcDefParam();
        this.state = 111;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
      }
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  funcDefParam() {
    let localctx = new FuncDefParamContext(this, this._ctx, this.state);
    this.enterRule(localctx, 12, prqlParser.RULE_funcDefParam);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 114;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 9, this._ctx);
      switch (la_) {
        case 1:
          this.state = 112;
          this.namedArg();
          break;

        case 2:
          this.state = 113;
          this.match(prqlParser.IDENT);
          break;
      }
      this.state = 117;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      if (_la === prqlParser.LANG) {
        this.state = 116;
        this.typeDef();
      }
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  typeDef() {
    let localctx = new TypeDefContext(this, this._ctx, this.state);
    this.enterRule(localctx, 14, prqlParser.RULE_typeDef);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 119;
      this.match(prqlParser.LANG);
      this.state = 120;
      this.typeTerm();
      this.state = 121;
      this.match(prqlParser.BAR);
      this.state = 125;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (_la === prqlParser.IDENT) {
        this.state = 122;
        this.typeTerm();
        this.state = 127;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
      }
      this.state = 128;
      this.match(prqlParser.RANG);
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  typeTerm() {
    let localctx = new TypeTermContext(this, this._ctx, this.state);
    this.enterRule(localctx, 16, prqlParser.RULE_typeTerm);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 130;
      this.match(prqlParser.IDENT);
      this.state = 132;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      if (_la === prqlParser.LANG) {
        this.state = 131;
        this.typeDef();
      }
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  stmt() {
    let localctx = new StmtContext(this, this._ctx, this.state);
    this.enterRule(localctx, 18, prqlParser.RULE_stmt);
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 134;
      this.assignStmt();
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  assignStmt() {
    let localctx = new AssignStmtContext(this, this._ctx, this.state);
    this.enterRule(localctx, 20, prqlParser.RULE_assignStmt);
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 136;
      this.match(prqlParser.LET);
      this.state = 137;
      this.match(prqlParser.IDENT);
      this.state = 138;
      this.match(prqlParser.ASSIGN);
      this.state = 139;
      this.expr(0);
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  pipeline() {
    let localctx = new PipelineContext(this, this._ctx, this.state);
    this.enterRule(localctx, 22, prqlParser.RULE_pipeline);
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 141;
      this.exprCall();
      this.state = 147;
      this._errHandler.sync(this);
      var _alt = this._interp.adaptivePredict(this._input, 13, this._ctx);
      while (_alt != 2 && _alt != antlr4.atn.ATN.INVALID_ALT_NUMBER) {
        if (_alt === 1) {
          this.state = 142;
          this.nl();
          this.state = 143;
          this.funcCall();
        }
        this.state = 149;
        this._errHandler.sync(this);
        _alt = this._interp.adaptivePredict(this._input, 13, this._ctx);
      }

      this.state = 152;
      this._errHandler.sync(this);
      switch (this._input.LA(1)) {
        case prqlParser.NEWLINE:
        case prqlParser.COMMENT:
          this.state = 150;
          this.nl();
          break;
        case prqlParser.EOF:
          this.state = 151;
          this.match(prqlParser.EOF);
          break;
        default:
          throw new antlr4.error.NoViableAltException(this);
      }
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  inlinePipeline() {
    let localctx = new InlinePipelineContext(this, this._ctx, this.state);
    this.enterRule(localctx, 24, prqlParser.RULE_inlinePipeline);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 154;
      this.exprCall();
      this.state = 159;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (_la === prqlParser.BAR) {
        this.state = 155;
        this.match(prqlParser.BAR);
        this.state = 156;
        this.funcCall();
        this.state = 161;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
      }
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  identBacktick() {
    let localctx = new IdentBacktickContext(this, this._ctx, this.state);
    this.enterRule(localctx, 26, prqlParser.RULE_identBacktick);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 162;
      this.match(prqlParser.BACKTICK);
      this.state = 166;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (
        ((_la & ~0x1f) == 0 &&
          ((1 << _la) &
            ((1 << prqlParser.FUNC) |
              (1 << prqlParser.PRQL) |
              (1 << prqlParser.LET) |
              (1 << prqlParser.ARROW) |
              (1 << prqlParser.ASSIGN) |
              (1 << prqlParser.PLUS) |
              (1 << prqlParser.MINUS) |
              (1 << prqlParser.STAR) |
              (1 << prqlParser.POW) |
              (1 << prqlParser.DIV) |
              (1 << prqlParser.MOD) |
              (1 << prqlParser.MODEL) |
              (1 << prqlParser.EQ) |
              (1 << prqlParser.NE) |
              (1 << prqlParser.LE) |
              (1 << prqlParser.GE) |
              (1 << prqlParser.BAR) |
              (1 << prqlParser.COLON) |
              (1 << prqlParser.COMMA) |
              (1 << prqlParser.DOT) |
              (1 << prqlParser.DOLLAR) |
              (1 << prqlParser.RANGE) |
              (1 << prqlParser.LANG) |
              (1 << prqlParser.RANG) |
              (1 << prqlParser.LBRACKET) |
              (1 << prqlParser.RBRACKET) |
              (1 << prqlParser.LPAREN) |
              (1 << prqlParser.RPAREN) |
              (1 << prqlParser.UNDERSCORE) |
              (1 << prqlParser.DOUBLE_QUOTE))) !==
            0) ||
        (((_la - 32) & ~0x1f) == 0 &&
          ((1 << (_la - 32)) &
            ((1 << (prqlParser.SINGLE_QUOTE - 32)) |
              (1 << (prqlParser.TRIPLE_DOUBLE_QUOTE - 32)) |
              (1 << (prqlParser.TRIPLE_SINGLE_QUOTE - 32)) |
              (1 << (prqlParser.AND - 32)) |
              (1 << (prqlParser.OR - 32)) |
              (1 << (prqlParser.NOT - 32)) |
              (1 << (prqlParser.COALESCE - 32)) |
              (1 << (prqlParser.NULL_ - 32)) |
              (1 << (prqlParser.BOOLEAN - 32)) |
              (1 << (prqlParser.INTEGER - 32)) |
              (1 << (prqlParser.FLOAT - 32)) |
              (1 << (prqlParser.IDENT - 32)) |
              (1 << (prqlParser.IDENT_START - 32)) |
              (1 << (prqlParser.IDENT_NEXT - 32)) |
              (1 << (prqlParser.WHITESPACE - 32)) |
              (1 << (prqlParser.COMMENT - 32)) |
              (1 << (prqlParser.INTERVAL_KIND - 32)) |
              (1 << (prqlParser.STRING - 32)))) !==
            0)
      ) {
        this.state = 163;
        _la = this._input.LA(1);
        if (
          _la <= 0 ||
          _la === prqlParser.BACKTICK ||
          _la === prqlParser.NEWLINE
        ) {
          this._errHandler.recoverInline(this);
        } else {
          this._errHandler.reportMatch(this);
          this.consume();
        }
        this.state = 168;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
      }
      this.state = 169;
      this.match(prqlParser.BACKTICK);
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  funcCall() {
    let localctx = new FuncCallContext(this, this._ctx, this.state);
    this.enterRule(localctx, 28, prqlParser.RULE_funcCall);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 171;
      this.match(prqlParser.IDENT);
      this.state = 175;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (
        ((_la & ~0x1f) == 0 &&
          ((1 << _la) &
            ((1 << prqlParser.PLUS) |
              (1 << prqlParser.MINUS) |
              (1 << prqlParser.LBRACKET) |
              (1 << prqlParser.LPAREN) |
              (1 << prqlParser.BACKTICK))) !==
            0) ||
        (((_la - 37) & ~0x1f) == 0 &&
          ((1 << (_la - 37)) &
            ((1 << (prqlParser.NOT - 37)) |
              (1 << (prqlParser.NULL_ - 37)) |
              (1 << (prqlParser.BOOLEAN - 37)) |
              (1 << (prqlParser.INTEGER - 37)) |
              (1 << (prqlParser.FLOAT - 37)) |
              (1 << (prqlParser.IDENT - 37)) |
              (1 << (prqlParser.STRING - 37)))) !==
            0)
      ) {
        this.state = 172;
        this.funcCallParam();
        this.state = 177;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
      }
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  funcCallParam() {
    let localctx = new FuncCallParamContext(this, this._ctx, this.state);
    this.enterRule(localctx, 30, prqlParser.RULE_funcCallParam);
    try {
      this.state = 181;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 18, this._ctx);
      switch (la_) {
        case 1:
          this.enterOuterAlt(localctx, 1);
          this.state = 178;
          this.namedArg();
          break;

        case 2:
          this.enterOuterAlt(localctx, 2);
          this.state = 179;
          this.assign();
          break;

        case 3:
          this.enterOuterAlt(localctx, 3);
          this.state = 180;
          this.expr(0);
          break;
      }
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  namedArg() {
    let localctx = new NamedArgContext(this, this._ctx, this.state);
    this.enterRule(localctx, 32, prqlParser.RULE_namedArg);
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 183;
      this.match(prqlParser.IDENT);
      this.state = 184;
      this.match(prqlParser.COLON);
      this.state = 187;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 19, this._ctx);
      switch (la_) {
        case 1:
          this.state = 185;
          this.assign();
          break;

        case 2:
          this.state = 186;
          this.expr(0);
          break;
      }
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  assign() {
    let localctx = new AssignContext(this, this._ctx, this.state);
    this.enterRule(localctx, 34, prqlParser.RULE_assign);
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 189;
      this.match(prqlParser.IDENT);
      this.state = 190;
      this.match(prqlParser.ASSIGN);
      this.state = 191;
      this.expr(0);
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  assignCall() {
    let localctx = new AssignCallContext(this, this._ctx, this.state);
    this.enterRule(localctx, 36, prqlParser.RULE_assignCall);
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 193;
      this.match(prqlParser.IDENT);
      this.state = 194;
      this.match(prqlParser.ASSIGN);
      this.state = 195;
      this.exprCall();
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  exprCall() {
    let localctx = new ExprCallContext(this, this._ctx, this.state);
    this.enterRule(localctx, 38, prqlParser.RULE_exprCall);
    try {
      this.state = 199;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 20, this._ctx);
      switch (la_) {
        case 1:
          this.enterOuterAlt(localctx, 1);
          this.state = 197;
          this.expr(0);
          break;

        case 2:
          this.enterOuterAlt(localctx, 2);
          this.state = 198;
          this.funcCall();
          break;
      }
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  expr(_p) {
    if (_p === undefined) {
      _p = 0;
    }
    const _parentctx = this._ctx;
    const _parentState = this.state;
    let localctx = new ExprContext(this, this._ctx, _parentState);
    let _prevctx = localctx;
    const _startState = 40;
    this.enterRecursionRule(localctx, 40, prqlParser.RULE_expr, _p);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 207;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 21, this._ctx);
      switch (la_) {
        case 1:
          this.state = 202;
          this.match(prqlParser.LPAREN);
          this.state = 203;
          this.expr(0);
          this.state = 204;
          this.match(prqlParser.RPAREN);
          break;

        case 2:
          this.state = 206;
          this.term();
          break;
      }
      this._ctx.stop = this._input.LT(-1);
      this.state = 232;
      this._errHandler.sync(this);
      var _alt = this._interp.adaptivePredict(this._input, 23, this._ctx);
      while (_alt != 2 && _alt != antlr4.atn.ATN.INVALID_ALT_NUMBER) {
        if (_alt === 1) {
          if (this._parseListeners !== null) {
            this.triggerExitRuleEvent();
          }
          _prevctx = localctx;
          this.state = 230;
          this._errHandler.sync(this);
          var la_ = this._interp.adaptivePredict(this._input, 22, this._ctx);
          switch (la_) {
            case 1:
              localctx = new ExprContext(this, _parentctx, _parentState);
              this.pushNewRecursionContext(
                localctx,
                _startState,
                prqlParser.RULE_expr
              );
              this.state = 209;
              if (!this.precpred(this._ctx, 9)) {
                throw new antlr4.error.FailedPredicateException(
                  this,
                  "this.precpred(this._ctx, 9)"
                );
              }
              this.state = 210;
              _la = this._input.LA(1);
              if (
                !(
                  (_la & ~0x1f) == 0 &&
                  ((1 << _la) &
                    ((1 << prqlParser.STAR) |
                      (1 << prqlParser.DIV) |
                      (1 << prqlParser.MOD))) !==
                    0
                )
              ) {
                this._errHandler.recoverInline(this);
              } else {
                this._errHandler.reportMatch(this);
                this.consume();
              }
              this.state = 211;
              this.expr(10);
              break;

            case 2:
              localctx = new ExprContext(this, _parentctx, _parentState);
              this.pushNewRecursionContext(
                localctx,
                _startState,
                prqlParser.RULE_expr
              );
              this.state = 212;
              if (!this.precpred(this._ctx, 8)) {
                throw new antlr4.error.FailedPredicateException(
                  this,
                  "this.precpred(this._ctx, 8)"
                );
              }
              this.state = 213;
              _la = this._input.LA(1);
              if (!(_la === prqlParser.PLUS || _la === prqlParser.MINUS)) {
                this._errHandler.recoverInline(this);
              } else {
                this._errHandler.reportMatch(this);
                this.consume();
              }
              this.state = 214;
              this.expr(9);
              break;

            case 3:
              localctx = new ExprContext(this, _parentctx, _parentState);
              this.pushNewRecursionContext(
                localctx,
                _startState,
                prqlParser.RULE_expr
              );
              this.state = 215;
              if (!this.precpred(this._ctx, 7)) {
                throw new antlr4.error.FailedPredicateException(
                  this,
                  "this.precpred(this._ctx, 7)"
                );
              }
              this.state = 216;
              this.match(prqlParser.POW);
              this.state = 217;
              this.expr(8);
              break;

            case 4:
              localctx = new ExprContext(this, _parentctx, _parentState);
              this.pushNewRecursionContext(
                localctx,
                _startState,
                prqlParser.RULE_expr
              );
              this.state = 218;
              if (!this.precpred(this._ctx, 6)) {
                throw new antlr4.error.FailedPredicateException(
                  this,
                  "this.precpred(this._ctx, 6)"
                );
              }
              this.state = 219;
              this.match(prqlParser.MODEL);
              this.state = 220;
              this.expr(7);
              break;

            case 5:
              localctx = new ExprContext(this, _parentctx, _parentState);
              this.pushNewRecursionContext(
                localctx,
                _startState,
                prqlParser.RULE_expr
              );
              this.state = 221;
              if (!this.precpred(this._ctx, 5)) {
                throw new antlr4.error.FailedPredicateException(
                  this,
                  "this.precpred(this._ctx, 5)"
                );
              }
              this.state = 222;
              _la = this._input.LA(1);
              if (
                !(
                  (_la & ~0x1f) == 0 &&
                  ((1 << _la) &
                    ((1 << prqlParser.EQ) |
                      (1 << prqlParser.NE) |
                      (1 << prqlParser.LE) |
                      (1 << prqlParser.GE) |
                      (1 << prqlParser.LANG) |
                      (1 << prqlParser.RANG))) !==
                    0
                )
              ) {
                this._errHandler.recoverInline(this);
              } else {
                this._errHandler.reportMatch(this);
                this.consume();
              }
              this.state = 223;
              this.expr(6);
              break;

            case 6:
              localctx = new ExprContext(this, _parentctx, _parentState);
              this.pushNewRecursionContext(
                localctx,
                _startState,
                prqlParser.RULE_expr
              );
              this.state = 224;
              if (!this.precpred(this._ctx, 4)) {
                throw new antlr4.error.FailedPredicateException(
                  this,
                  "this.precpred(this._ctx, 4)"
                );
              }
              this.state = 225;
              this.match(prqlParser.COALESCE);
              this.state = 226;
              this.expr(5);
              break;

            case 7:
              localctx = new ExprContext(this, _parentctx, _parentState);
              this.pushNewRecursionContext(
                localctx,
                _startState,
                prqlParser.RULE_expr
              );
              this.state = 227;
              if (!this.precpred(this._ctx, 3)) {
                throw new antlr4.error.FailedPredicateException(
                  this,
                  "this.precpred(this._ctx, 3)"
                );
              }
              this.state = 228;
              _la = this._input.LA(1);
              if (!(_la === prqlParser.AND || _la === prqlParser.OR)) {
                this._errHandler.recoverInline(this);
              } else {
                this._errHandler.reportMatch(this);
                this.consume();
              }
              this.state = 229;
              this.expr(4);
              break;
          }
        }
        this.state = 234;
        this._errHandler.sync(this);
        _alt = this._interp.adaptivePredict(this._input, 23, this._ctx);
      }
    } catch (error) {
      if (error instanceof antlr4.error.RecognitionException) {
        localctx.exception = error;
        this._errHandler.reportError(this, error);
        this._errHandler.recover(this, error);
      } else {
        throw error;
      }
    } finally {
      this.unrollRecursionContexts(_parentctx);
    }
    return localctx;
  }

  term() {
    let localctx = new TermContext(this, this._ctx, this.state);
    this.enterRule(localctx, 42, prqlParser.RULE_term);
    try {
      this.state = 240;
      this._errHandler.sync(this);
      switch (this._input.LA(1)) {
        case prqlParser.NULL_:
        case prqlParser.BOOLEAN:
        case prqlParser.INTEGER:
        case prqlParser.FLOAT:
        case prqlParser.IDENT:
        case prqlParser.STRING:
          this.enterOuterAlt(localctx, 1);
          this.state = 235;
          this.literal();
          break;
        case prqlParser.BACKTICK:
          this.enterOuterAlt(localctx, 2);
          this.state = 236;
          this.identBacktick();
          break;
        case prqlParser.PLUS:
        case prqlParser.MINUS:
        case prqlParser.NOT:
          this.enterOuterAlt(localctx, 3);
          this.state = 237;
          this.exprUnary();
          break;
        case prqlParser.LBRACKET:
          this.enterOuterAlt(localctx, 4);
          this.state = 238;
          this.list();
          break;
        case prqlParser.LPAREN:
          this.enterOuterAlt(localctx, 5);
          this.state = 239;
          this.nestedPipeline();
          break;
        default:
          throw new antlr4.error.NoViableAltException(this);
      }
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  exprUnary() {
    let localctx = new ExprUnaryContext(this, this._ctx, this.state);
    this.enterRule(localctx, 44, prqlParser.RULE_exprUnary);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 242;
      _la = this._input.LA(1);
      if (
        !(
          ((_la - 6) & ~0x1f) == 0 &&
          ((1 << (_la - 6)) &
            ((1 << (prqlParser.PLUS - 6)) |
              (1 << (prqlParser.MINUS - 6)) |
              (1 << (prqlParser.NOT - 6)))) !==
            0
        )
      ) {
        this._errHandler.recoverInline(this);
      } else {
        this._errHandler.reportMatch(this);
        this.consume();
      }
      this.state = 246;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 25, this._ctx);
      switch (la_) {
        case 1:
          this.state = 243;
          this.nestedPipeline();
          break;

        case 2:
          this.state = 244;
          this.literal();
          break;

        case 3:
          this.state = 245;
          this.match(prqlParser.IDENT);
          break;
      }
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  literal() {
    let localctx = new LiteralContext(this, this._ctx, this.state);
    this.enterRule(localctx, 46, prqlParser.RULE_literal);
    var _la = 0; // Token type
    try {
      this.state = 259;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 26, this._ctx);
      switch (la_) {
        case 1:
          this.enterOuterAlt(localctx, 1);
          this.state = 248;
          this.match(prqlParser.IDENT);
          break;

        case 2:
          this.enterOuterAlt(localctx, 2);
          this.state = 249;
          this.match(prqlParser.NULL_);
          break;

        case 3:
          this.enterOuterAlt(localctx, 3);
          this.state = 250;
          this.match(prqlParser.BOOLEAN);
          break;

        case 4:
          this.enterOuterAlt(localctx, 4);
          this.state = 251;
          this.match(prqlParser.STRING);
          break;

        case 5:
          this.enterOuterAlt(localctx, 5);
          this.state = 252;
          this.match(prqlParser.INTEGER);
          break;

        case 6:
          this.enterOuterAlt(localctx, 6);
          this.state = 253;
          this.match(prqlParser.FLOAT);
          break;

        case 7:
          this.enterOuterAlt(localctx, 7);
          this.state = 254;
          this.match(prqlParser.FLOAT);
          this.state = 255;
          this.match(prqlParser.INTERVAL_KIND);
          break;

        case 8:
          this.enterOuterAlt(localctx, 8);
          this.state = 256;
          _la = this._input.LA(1);
          if (!(_la === prqlParser.FLOAT || _la === prqlParser.IDENT)) {
            this._errHandler.recoverInline(this);
          } else {
            this._errHandler.reportMatch(this);
            this.consume();
          }
          this.state = 257;
          this.match(prqlParser.RANGE);
          this.state = 258;
          _la = this._input.LA(1);
          if (!(_la === prqlParser.FLOAT || _la === prqlParser.IDENT)) {
            this._errHandler.recoverInline(this);
          } else {
            this._errHandler.reportMatch(this);
            this.consume();
          }
          break;
      }
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  list() {
    let localctx = new ListContext(this, this._ctx, this.state);
    this.enterRule(localctx, 48, prqlParser.RULE_list);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 261;
      this.match(prqlParser.LBRACKET);
      this.state = 294;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      if (
        ((_la & ~0x1f) == 0 &&
          ((1 << _la) &
            ((1 << prqlParser.PLUS) |
              (1 << prqlParser.MINUS) |
              (1 << prqlParser.LBRACKET) |
              (1 << prqlParser.LPAREN) |
              (1 << prqlParser.BACKTICK))) !==
            0) ||
        (((_la - 37) & ~0x1f) == 0 &&
          ((1 << (_la - 37)) &
            ((1 << (prqlParser.NOT - 37)) |
              (1 << (prqlParser.NULL_ - 37)) |
              (1 << (prqlParser.BOOLEAN - 37)) |
              (1 << (prqlParser.INTEGER - 37)) |
              (1 << (prqlParser.FLOAT - 37)) |
              (1 << (prqlParser.IDENT - 37)) |
              (1 << (prqlParser.NEWLINE - 37)) |
              (1 << (prqlParser.COMMENT - 37)) |
              (1 << (prqlParser.STRING - 37)))) !==
            0)
      ) {
        this.state = 265;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
        while (_la === prqlParser.NEWLINE || _la === prqlParser.COMMENT) {
          this.state = 262;
          this.nl();
          this.state = 267;
          this._errHandler.sync(this);
          _la = this._input.LA(1);
        }
        this.state = 270;
        this._errHandler.sync(this);
        var la_ = this._interp.adaptivePredict(this._input, 28, this._ctx);
        switch (la_) {
          case 1:
            this.state = 268;
            this.assignCall();
            break;

          case 2:
            this.state = 269;
            this.exprCall();
            break;
        }
        this.state = 285;
        this._errHandler.sync(this);
        var _alt = this._interp.adaptivePredict(this._input, 31, this._ctx);
        while (_alt != 2 && _alt != antlr4.atn.ATN.INVALID_ALT_NUMBER) {
          if (_alt === 1) {
            this.state = 272;
            this.match(prqlParser.COMMA);
            this.state = 276;
            this._errHandler.sync(this);
            _la = this._input.LA(1);
            while (_la === prqlParser.NEWLINE || _la === prqlParser.COMMENT) {
              this.state = 273;
              this.nl();
              this.state = 278;
              this._errHandler.sync(this);
              _la = this._input.LA(1);
            }
            this.state = 281;
            this._errHandler.sync(this);
            var la_ = this._interp.adaptivePredict(this._input, 30, this._ctx);
            switch (la_) {
              case 1:
                this.state = 279;
                this.assignCall();
                break;

              case 2:
                this.state = 280;
                this.exprCall();
                break;
            }
          }
          this.state = 287;
          this._errHandler.sync(this);
          _alt = this._interp.adaptivePredict(this._input, 31, this._ctx);
        }

        this.state = 289;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
        if (_la === prqlParser.COMMA) {
          this.state = 288;
          this.match(prqlParser.COMMA);
        }

        this.state = 292;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
        if (_la === prqlParser.NEWLINE || _la === prqlParser.COMMENT) {
          this.state = 291;
          this.nl();
        }
      }

      this.state = 296;
      this.match(prqlParser.RBRACKET);
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }

  nestedPipeline() {
    let localctx = new NestedPipelineContext(this, this._ctx, this.state);
    this.enterRule(localctx, 50, prqlParser.RULE_nestedPipeline);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 298;
      this.match(prqlParser.LPAREN);
      this.state = 302;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (_la === prqlParser.NEWLINE || _la === prqlParser.COMMENT) {
        this.state = 299;
        this.nl();
        this.state = 304;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
      }
      this.state = 307;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 36, this._ctx);
      switch (la_) {
        case 1:
          this.state = 305;
          this.pipeline();
          break;

        case 2:
          this.state = 306;
          this.inlinePipeline();
          break;
      }
      this.state = 312;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (_la === prqlParser.NEWLINE || _la === prqlParser.COMMENT) {
        this.state = 309;
        this.nl();
        this.state = 314;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
      }
      this.state = 315;
      this.match(prqlParser.RPAREN);
    } catch (re) {
      if (re instanceof antlr4.error.RecognitionException) {
        localctx.exception = re;
        this._errHandler.reportError(this, re);
        this._errHandler.recover(this, re);
      } else {
        throw re;
      }
    } finally {
      this.exitRule();
    }
    return localctx;
  }
}

prqlParser.EOF = antlr4.Token.EOF;
prqlParser.FUNC = 1;
prqlParser.PRQL = 2;
prqlParser.LET = 3;
prqlParser.ARROW = 4;
prqlParser.ASSIGN = 5;
prqlParser.PLUS = 6;
prqlParser.MINUS = 7;
prqlParser.STAR = 8;
prqlParser.POW = 9;
prqlParser.DIV = 10;
prqlParser.MOD = 11;
prqlParser.MODEL = 12;
prqlParser.EQ = 13;
prqlParser.NE = 14;
prqlParser.LE = 15;
prqlParser.GE = 16;
prqlParser.BAR = 17;
prqlParser.COLON = 18;
prqlParser.COMMA = 19;
prqlParser.DOT = 20;
prqlParser.DOLLAR = 21;
prqlParser.RANGE = 22;
prqlParser.LANG = 23;
prqlParser.RANG = 24;
prqlParser.LBRACKET = 25;
prqlParser.RBRACKET = 26;
prqlParser.LPAREN = 27;
prqlParser.RPAREN = 28;
prqlParser.UNDERSCORE = 29;
prqlParser.BACKTICK = 30;
prqlParser.DOUBLE_QUOTE = 31;
prqlParser.SINGLE_QUOTE = 32;
prqlParser.TRIPLE_DOUBLE_QUOTE = 33;
prqlParser.TRIPLE_SINGLE_QUOTE = 34;
prqlParser.AND = 35;
prqlParser.OR = 36;
prqlParser.NOT = 37;
prqlParser.COALESCE = 38;
prqlParser.NULL_ = 39;
prqlParser.BOOLEAN = 40;
prqlParser.INTEGER = 41;
prqlParser.FLOAT = 42;
prqlParser.IDENT = 43;
prqlParser.IDENT_START = 44;
prqlParser.IDENT_NEXT = 45;
prqlParser.WHITESPACE = 46;
prqlParser.NEWLINE = 47;
prqlParser.COMMENT = 48;
prqlParser.INTERVAL_KIND = 49;
prqlParser.STRING = 50;

prqlParser.RULE_nl = 0;
prqlParser.RULE_program = 1;
prqlParser.RULE_programIntro = 2;
prqlParser.RULE_funcDef = 3;
prqlParser.RULE_funcDefName = 4;
prqlParser.RULE_funcDefParams = 5;
prqlParser.RULE_funcDefParam = 6;
prqlParser.RULE_typeDef = 7;
prqlParser.RULE_typeTerm = 8;
prqlParser.RULE_stmt = 9;
prqlParser.RULE_assignStmt = 10;
prqlParser.RULE_pipeline = 11;
prqlParser.RULE_inlinePipeline = 12;
prqlParser.RULE_identBacktick = 13;
prqlParser.RULE_funcCall = 14;
prqlParser.RULE_funcCallParam = 15;
prqlParser.RULE_namedArg = 16;
prqlParser.RULE_assign = 17;
prqlParser.RULE_assignCall = 18;
prqlParser.RULE_exprCall = 19;
prqlParser.RULE_expr = 20;
prqlParser.RULE_term = 21;
prqlParser.RULE_exprUnary = 22;
prqlParser.RULE_literal = 23;
prqlParser.RULE_list = 24;
prqlParser.RULE_nestedPipeline = 25;

class NlContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_nl;
  }

  NEWLINE() {
    return this.getToken(prqlParser.NEWLINE, 0);
  }

  COMMENT() {
    return this.getToken(prqlParser.COMMENT, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterNl(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitNl(this);
    }
  }
}

class ProgramContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_program;
  }

  EOF() {
    return this.getToken(prqlParser.EOF, 0);
  }

  nl = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(NlContext);
    } else {
      return this.getTypedRuleContext(NlContext, i);
    }
  };

  programIntro() {
    return this.getTypedRuleContext(ProgramIntroContext, 0);
  }

  funcDef = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(FuncDefContext);
    } else {
      return this.getTypedRuleContext(FuncDefContext, i);
    }
  };

  stmt = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(StmtContext);
    } else {
      return this.getTypedRuleContext(StmtContext, i);
    }
  };

  pipeline = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(PipelineContext);
    } else {
      return this.getTypedRuleContext(PipelineContext, i);
    }
  };

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterProgram(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitProgram(this);
    }
  }
}

class ProgramIntroContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_programIntro;
  }

  PRQL() {
    return this.getToken(prqlParser.PRQL, 0);
  }

  nl() {
    return this.getTypedRuleContext(NlContext, 0);
  }

  namedArg = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(NamedArgContext);
    } else {
      return this.getTypedRuleContext(NamedArgContext, i);
    }
  };

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterProgramIntro(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitProgramIntro(this);
    }
  }
}

class FuncDefContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_funcDef;
  }

  FUNC() {
    return this.getToken(prqlParser.FUNC, 0);
  }

  funcDefName() {
    return this.getTypedRuleContext(FuncDefNameContext, 0);
  }

  funcDefParams() {
    return this.getTypedRuleContext(FuncDefParamsContext, 0);
  }

  ARROW() {
    return this.getToken(prqlParser.ARROW, 0);
  }

  expr() {
    return this.getTypedRuleContext(ExprContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterFuncDef(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitFuncDef(this);
    }
  }
}

class FuncDefNameContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_funcDefName;
  }

  IDENT() {
    return this.getToken(prqlParser.IDENT, 0);
  }

  typeDef() {
    return this.getTypedRuleContext(TypeDefContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterFuncDefName(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitFuncDefName(this);
    }
  }
}

class FuncDefParamsContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_funcDefParams;
  }

  funcDefParam = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(FuncDefParamContext);
    } else {
      return this.getTypedRuleContext(FuncDefParamContext, i);
    }
  };

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterFuncDefParams(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitFuncDefParams(this);
    }
  }
}

class FuncDefParamContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_funcDefParam;
  }

  namedArg() {
    return this.getTypedRuleContext(NamedArgContext, 0);
  }

  IDENT() {
    return this.getToken(prqlParser.IDENT, 0);
  }

  typeDef() {
    return this.getTypedRuleContext(TypeDefContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterFuncDefParam(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitFuncDefParam(this);
    }
  }
}

class TypeDefContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_typeDef;
  }

  LANG() {
    return this.getToken(prqlParser.LANG, 0);
  }

  typeTerm = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(TypeTermContext);
    } else {
      return this.getTypedRuleContext(TypeTermContext, i);
    }
  };

  BAR() {
    return this.getToken(prqlParser.BAR, 0);
  }

  RANG() {
    return this.getToken(prqlParser.RANG, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterTypeDef(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitTypeDef(this);
    }
  }
}

class TypeTermContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_typeTerm;
  }

  IDENT() {
    return this.getToken(prqlParser.IDENT, 0);
  }

  typeDef() {
    return this.getTypedRuleContext(TypeDefContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterTypeTerm(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitTypeTerm(this);
    }
  }
}

class StmtContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_stmt;
  }

  assignStmt() {
    return this.getTypedRuleContext(AssignStmtContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterStmt(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitStmt(this);
    }
  }
}

class AssignStmtContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_assignStmt;
  }

  LET() {
    return this.getToken(prqlParser.LET, 0);
  }

  IDENT() {
    return this.getToken(prqlParser.IDENT, 0);
  }

  ASSIGN() {
    return this.getToken(prqlParser.ASSIGN, 0);
  }

  expr() {
    return this.getTypedRuleContext(ExprContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterAssignStmt(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitAssignStmt(this);
    }
  }
}

class PipelineContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_pipeline;
  }

  exprCall() {
    return this.getTypedRuleContext(ExprCallContext, 0);
  }

  nl = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(NlContext);
    } else {
      return this.getTypedRuleContext(NlContext, i);
    }
  };

  EOF() {
    return this.getToken(prqlParser.EOF, 0);
  }

  funcCall = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(FuncCallContext);
    } else {
      return this.getTypedRuleContext(FuncCallContext, i);
    }
  };

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterPipeline(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitPipeline(this);
    }
  }
}

class InlinePipelineContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_inlinePipeline;
  }

  exprCall() {
    return this.getTypedRuleContext(ExprCallContext, 0);
  }

  BAR = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTokens(prqlParser.BAR);
    } else {
      return this.getToken(prqlParser.BAR, i);
    }
  };

  funcCall = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(FuncCallContext);
    } else {
      return this.getTypedRuleContext(FuncCallContext, i);
    }
  };

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterInlinePipeline(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitInlinePipeline(this);
    }
  }
}

class IdentBacktickContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_identBacktick;
  }

  BACKTICK = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTokens(prqlParser.BACKTICK);
    } else {
      return this.getToken(prqlParser.BACKTICK, i);
    }
  };

  NEWLINE = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTokens(prqlParser.NEWLINE);
    } else {
      return this.getToken(prqlParser.NEWLINE, i);
    }
  };

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterIdentBacktick(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitIdentBacktick(this);
    }
  }
}

class FuncCallContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_funcCall;
  }

  IDENT() {
    return this.getToken(prqlParser.IDENT, 0);
  }

  funcCallParam = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(FuncCallParamContext);
    } else {
      return this.getTypedRuleContext(FuncCallParamContext, i);
    }
  };

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterFuncCall(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitFuncCall(this);
    }
  }
}

class FuncCallParamContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_funcCallParam;
  }

  namedArg() {
    return this.getTypedRuleContext(NamedArgContext, 0);
  }

  assign() {
    return this.getTypedRuleContext(AssignContext, 0);
  }

  expr() {
    return this.getTypedRuleContext(ExprContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterFuncCallParam(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitFuncCallParam(this);
    }
  }
}

class NamedArgContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_namedArg;
  }

  IDENT() {
    return this.getToken(prqlParser.IDENT, 0);
  }

  COLON() {
    return this.getToken(prqlParser.COLON, 0);
  }

  assign() {
    return this.getTypedRuleContext(AssignContext, 0);
  }

  expr() {
    return this.getTypedRuleContext(ExprContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterNamedArg(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitNamedArg(this);
    }
  }
}

class AssignContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_assign;
  }

  IDENT() {
    return this.getToken(prqlParser.IDENT, 0);
  }

  ASSIGN() {
    return this.getToken(prqlParser.ASSIGN, 0);
  }

  expr() {
    return this.getTypedRuleContext(ExprContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterAssign(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitAssign(this);
    }
  }
}

class AssignCallContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_assignCall;
  }

  IDENT() {
    return this.getToken(prqlParser.IDENT, 0);
  }

  ASSIGN() {
    return this.getToken(prqlParser.ASSIGN, 0);
  }

  exprCall() {
    return this.getTypedRuleContext(ExprCallContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterAssignCall(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitAssignCall(this);
    }
  }
}

class ExprCallContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_exprCall;
  }

  expr() {
    return this.getTypedRuleContext(ExprContext, 0);
  }

  funcCall() {
    return this.getTypedRuleContext(FuncCallContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterExprCall(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitExprCall(this);
    }
  }
}

class ExprContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_expr;
  }

  LPAREN() {
    return this.getToken(prqlParser.LPAREN, 0);
  }

  expr = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(ExprContext);
    } else {
      return this.getTypedRuleContext(ExprContext, i);
    }
  };

  RPAREN() {
    return this.getToken(prqlParser.RPAREN, 0);
  }

  term() {
    return this.getTypedRuleContext(TermContext, 0);
  }

  STAR() {
    return this.getToken(prqlParser.STAR, 0);
  }

  DIV() {
    return this.getToken(prqlParser.DIV, 0);
  }

  MOD() {
    return this.getToken(prqlParser.MOD, 0);
  }

  MINUS() {
    return this.getToken(prqlParser.MINUS, 0);
  }

  PLUS() {
    return this.getToken(prqlParser.PLUS, 0);
  }

  POW() {
    return this.getToken(prqlParser.POW, 0);
  }

  MODEL() {
    return this.getToken(prqlParser.MODEL, 0);
  }

  EQ() {
    return this.getToken(prqlParser.EQ, 0);
  }

  NE() {
    return this.getToken(prqlParser.NE, 0);
  }

  GE() {
    return this.getToken(prqlParser.GE, 0);
  }

  LE() {
    return this.getToken(prqlParser.LE, 0);
  }

  LANG() {
    return this.getToken(prqlParser.LANG, 0);
  }

  RANG() {
    return this.getToken(prqlParser.RANG, 0);
  }

  COALESCE() {
    return this.getToken(prqlParser.COALESCE, 0);
  }

  AND() {
    return this.getToken(prqlParser.AND, 0);
  }

  OR() {
    return this.getToken(prqlParser.OR, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterExpr(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitExpr(this);
    }
  }
}

class TermContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_term;
  }

  literal() {
    return this.getTypedRuleContext(LiteralContext, 0);
  }

  identBacktick() {
    return this.getTypedRuleContext(IdentBacktickContext, 0);
  }

  exprUnary() {
    return this.getTypedRuleContext(ExprUnaryContext, 0);
  }

  list() {
    return this.getTypedRuleContext(ListContext, 0);
  }

  nestedPipeline() {
    return this.getTypedRuleContext(NestedPipelineContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterTerm(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitTerm(this);
    }
  }
}

class ExprUnaryContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_exprUnary;
  }

  MINUS() {
    return this.getToken(prqlParser.MINUS, 0);
  }

  PLUS() {
    return this.getToken(prqlParser.PLUS, 0);
  }

  NOT() {
    return this.getToken(prqlParser.NOT, 0);
  }

  nestedPipeline() {
    return this.getTypedRuleContext(NestedPipelineContext, 0);
  }

  literal() {
    return this.getTypedRuleContext(LiteralContext, 0);
  }

  IDENT() {
    return this.getToken(prqlParser.IDENT, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterExprUnary(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitExprUnary(this);
    }
  }
}

class LiteralContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_literal;
  }

  IDENT = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTokens(prqlParser.IDENT);
    } else {
      return this.getToken(prqlParser.IDENT, i);
    }
  };

  NULL_() {
    return this.getToken(prqlParser.NULL_, 0);
  }

  BOOLEAN() {
    return this.getToken(prqlParser.BOOLEAN, 0);
  }

  STRING() {
    return this.getToken(prqlParser.STRING, 0);
  }

  INTEGER() {
    return this.getToken(prqlParser.INTEGER, 0);
  }

  FLOAT = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTokens(prqlParser.FLOAT);
    } else {
      return this.getToken(prqlParser.FLOAT, i);
    }
  };

  INTERVAL_KIND() {
    return this.getToken(prqlParser.INTERVAL_KIND, 0);
  }

  RANGE() {
    return this.getToken(prqlParser.RANGE, 0);
  }

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterLiteral(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitLiteral(this);
    }
  }
}

class ListContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_list;
  }

  LBRACKET() {
    return this.getToken(prqlParser.LBRACKET, 0);
  }

  RBRACKET() {
    return this.getToken(prqlParser.RBRACKET, 0);
  }

  assignCall = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(AssignCallContext);
    } else {
      return this.getTypedRuleContext(AssignCallContext, i);
    }
  };

  exprCall = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(ExprCallContext);
    } else {
      return this.getTypedRuleContext(ExprCallContext, i);
    }
  };

  nl = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(NlContext);
    } else {
      return this.getTypedRuleContext(NlContext, i);
    }
  };

  COMMA = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTokens(prqlParser.COMMA);
    } else {
      return this.getToken(prqlParser.COMMA, i);
    }
  };

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterList(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitList(this);
    }
  }
}

class NestedPipelineContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = prqlParser.RULE_nestedPipeline;
  }

  LPAREN() {
    return this.getToken(prqlParser.LPAREN, 0);
  }

  RPAREN() {
    return this.getToken(prqlParser.RPAREN, 0);
  }

  pipeline() {
    return this.getTypedRuleContext(PipelineContext, 0);
  }

  inlinePipeline() {
    return this.getTypedRuleContext(InlinePipelineContext, 0);
  }

  nl = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(NlContext);
    } else {
      return this.getTypedRuleContext(NlContext, i);
    }
  };

  enterRule(listener) {
    if (listener instanceof prqlListener) {
      listener.enterNestedPipeline(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof prqlListener) {
      listener.exitNestedPipeline(this);
    }
  }
}

prqlParser.NlContext = NlContext;
prqlParser.ProgramContext = ProgramContext;
prqlParser.ProgramIntroContext = ProgramIntroContext;
prqlParser.FuncDefContext = FuncDefContext;
prqlParser.FuncDefNameContext = FuncDefNameContext;
prqlParser.FuncDefParamsContext = FuncDefParamsContext;
prqlParser.FuncDefParamContext = FuncDefParamContext;
prqlParser.TypeDefContext = TypeDefContext;
prqlParser.TypeTermContext = TypeTermContext;
prqlParser.StmtContext = StmtContext;
prqlParser.AssignStmtContext = AssignStmtContext;
prqlParser.PipelineContext = PipelineContext;
prqlParser.InlinePipelineContext = InlinePipelineContext;
prqlParser.IdentBacktickContext = IdentBacktickContext;
prqlParser.FuncCallContext = FuncCallContext;
prqlParser.FuncCallParamContext = FuncCallParamContext;
prqlParser.NamedArgContext = NamedArgContext;
prqlParser.AssignContext = AssignContext;
prqlParser.AssignCallContext = AssignCallContext;
prqlParser.ExprCallContext = ExprCallContext;
prqlParser.ExprContext = ExprContext;
prqlParser.TermContext = TermContext;
prqlParser.ExprUnaryContext = ExprUnaryContext;
prqlParser.LiteralContext = LiteralContext;
prqlParser.ListContext = ListContext;
prqlParser.NestedPipelineContext = NestedPipelineContext;
