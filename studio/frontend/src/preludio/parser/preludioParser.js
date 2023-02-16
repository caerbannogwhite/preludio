// Generated from .\preludio.g4 by ANTLR 4.9.2
// jshint ignore: start
import antlr4 from "antlr4";
import preludioListener from "./preludioListener.js";

const serializedATN = [
  "\u0003\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786",
  "\u5964\u00034\u0143\u0004\u0002\t\u0002\u0004\u0003\t\u0003\u0004\u0004",
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
  "\u000b\u0010\u0003\u0011\u0003\u0011\u0003\u0011\u0003\u0011\u0005\u0011",
  "\u00b9\n\u0011\u0003\u0012\u0003\u0012\u0003\u0012\u0003\u0012\u0005",
  "\u0012\u00bf\n\u0012\u0003\u0013\u0003\u0013\u0003\u0013\u0003\u0013",
  "\u0003\u0014\u0003\u0014\u0003\u0014\u0003\u0014\u0003\u0015\u0003\u0015",
  "\u0005\u0015\u00cb\n\u0015\u0003\u0016\u0003\u0016\u0003\u0016\u0003",
  "\u0016\u0003\u0016\u0003\u0016\u0005\u0016\u00d3\n\u0016\u0003\u0016",
  "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016",
  "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016",
  "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016",
  "\u0003\u0016\u0003\u0016\u0007\u0016\u00ea\n\u0016\f\u0016\u000e\u0016",
  "\u00ed\u000b\u0016\u0003\u0017\u0003\u0017\u0003\u0017\u0003\u0017\u0003",
  "\u0017\u0005\u0017\u00f4\n\u0017\u0003\u0018\u0003\u0018\u0003\u0018",
  "\u0003\u0018\u0005\u0018\u00fa\n\u0018\u0003\u0019\u0003\u0019\u0003",
  "\u0019\u0003\u0019\u0003\u0019\u0003\u0019\u0003\u0019\u0003\u0019\u0003",
  "\u0019\u0003\u0019\u0003\u0019\u0005\u0019\u0107\n\u0019\u0003\u001a",
  "\u0003\u001a\u0007\u001a\u010b\n\u001a\f\u001a\u000e\u001a\u010e\u000b",
  "\u001a\u0003\u001a\u0003\u001a\u0003\u001a\u0005\u001a\u0113\n\u001a",
  "\u0003\u001a\u0003\u001a\u0007\u001a\u0117\n\u001a\f\u001a\u000e\u001a",
  "\u011a\u000b\u001a\u0003\u001a\u0003\u001a\u0003\u001a\u0005\u001a\u011f",
  "\n\u001a\u0007\u001a\u0121\n\u001a\f\u001a\u000e\u001a\u0124\u000b\u001a",
  "\u0003\u001a\u0005\u001a\u0127\n\u001a\u0003\u001a\u0005\u001a\u012a",
  "\n\u001a\u0005\u001a\u012c\n\u001a\u0003\u001a\u0003\u001a\u0003\u001b",
  "\u0003\u001b\u0007\u001b\u0132\n\u001b\f\u001b\u000e\u001b\u0135\u000b",
  "\u001b\u0003\u001b\u0003\u001b\u0005\u001b\u0139\n\u001b\u0003\u001b",
  "\u0007\u001b\u013c\n\u001b\f\u001b\u000e\u001b\u013f\u000b\u001b\u0003",
  "\u001b\u0003\u001b\u0003\u001b\u0002\u0003*\u001c\u0002\u0004\u0006",
  '\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e "$&(*',
  ",.024\u0002\u000b\u0003\u000212\u0004\u0002  11\u0004\u0002\n\n\f\r",
  "\u0003\u0002\b\t\u0004\u0002\u000f\u0012\u0019\u001a\u0003\u0002%&\u0004",
  "\u0002\b\t''\u0003\u0002+,\u0003\u0002+-\u0002\u0162\u00026\u0003",
  "\u0002\u0002\u0002\u0004;\u0003\u0002\u0002\u0002\u0006Y\u0003\u0002",
  "\u0002\u0002\bb\u0003\u0002\u0002\u0002\nh\u0003\u0002\u0002\u0002\f",
  "o\u0003\u0002\u0002\u0002\u000et\u0003\u0002\u0002\u0002\u0010y\u0003",
  "\u0002\u0002\u0002\u0012\u0084\u0003\u0002\u0002\u0002\u0014\u0088\u0003",
  "\u0002\u0002\u0002\u0016\u008a\u0003\u0002\u0002\u0002\u0018\u008f\u0003",
  "\u0002\u0002\u0002\u001a\u009c\u0003\u0002\u0002\u0002\u001c\u00a4\u0003",
  "\u0002\u0002\u0002\u001e\u00ad\u0003\u0002\u0002\u0002 \u00b8\u0003",
  '\u0002\u0002\u0002"\u00ba\u0003\u0002\u0002\u0002$\u00c0\u0003\u0002',
  "\u0002\u0002&\u00c4\u0003\u0002\u0002\u0002(\u00ca\u0003\u0002\u0002",
  "\u0002*\u00d2\u0003\u0002\u0002\u0002,\u00f3\u0003\u0002\u0002\u0002",
  ".\u00f5\u0003\u0002\u0002\u00020\u0106\u0003\u0002\u0002\u00022\u0108",
  "\u0003\u0002\u0002\u00024\u012f\u0003\u0002\u0002\u000267\t\u0002\u0002",
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
  '\u0002\u0002\u00b3\u00b1\u0003\u0002\u0002\u0002\u00b4\u00b9\u0005"',
  "\u0012\u0002\u00b5\u00b9\u0005$\u0013\u0002\u00b6\u00b9\u0005&\u0014",
  "\u0002\u00b7\u00b9\u0005*\u0016\u0002\u00b8\u00b4\u0003\u0002\u0002",
  "\u0002\u00b8\u00b5\u0003\u0002\u0002\u0002\u00b8\u00b6\u0003\u0002\u0002",
  "\u0002\u00b8\u00b7\u0003\u0002\u0002\u0002\u00b9!\u0003\u0002\u0002",
  "\u0002\u00ba\u00bb\u0007-\u0002\u0002\u00bb\u00be\u0007\u0014\u0002",
  "\u0002\u00bc\u00bf\u0005$\u0013\u0002\u00bd\u00bf\u0005*\u0016\u0002",
  "\u00be\u00bc\u0003\u0002\u0002\u0002\u00be\u00bd\u0003\u0002\u0002\u0002",
  "\u00bf#\u0003\u0002\u0002\u0002\u00c0\u00c1\u0007-\u0002\u0002\u00c1",
  "\u00c2\u0007\u0007\u0002\u0002\u00c2\u00c3\u0005(\u0015\u0002\u00c3",
  "%\u0003\u0002\u0002\u0002\u00c4\u00c5\u00052\u001a\u0002\u00c5\u00c6",
  "\u0007\u0007\u0002\u0002\u00c6\u00c7\u0005(\u0015\u0002\u00c7'\u0003",
  "\u0002\u0002\u0002\u00c8\u00cb\u0005*\u0016\u0002\u00c9\u00cb\u0005",
  "\u001e\u0010\u0002\u00ca\u00c8\u0003\u0002\u0002\u0002\u00ca\u00c9\u0003",
  "\u0002\u0002\u0002\u00cb)\u0003\u0002\u0002\u0002\u00cc\u00cd\b\u0016",
  "\u0001\u0002\u00cd\u00ce\u0007\u001d\u0002\u0002\u00ce\u00cf\u0005*",
  "\u0016\u0002\u00cf\u00d0\u0007\u001e\u0002\u0002\u00d0\u00d3\u0003\u0002",
  "\u0002\u0002\u00d1\u00d3\u0005,\u0017\u0002\u00d2\u00cc\u0003\u0002",
  "\u0002\u0002\u00d2\u00d1\u0003\u0002\u0002\u0002\u00d3\u00eb\u0003\u0002",
  "\u0002\u0002\u00d4\u00d5\f\u000b\u0002\u0002\u00d5\u00d6\t\u0004\u0002",
  "\u0002\u00d6\u00ea\u0005*\u0016\f\u00d7\u00d8\f\n\u0002\u0002\u00d8",
  "\u00d9\t\u0005\u0002\u0002\u00d9\u00ea\u0005*\u0016\u000b\u00da\u00db",
  "\f\t\u0002\u0002\u00db\u00dc\u0007\u000b\u0002\u0002\u00dc\u00ea\u0005",
  "*\u0016\n\u00dd\u00de\f\b\u0002\u0002\u00de\u00df\u0007\u000e\u0002",
  "\u0002\u00df\u00ea\u0005*\u0016\t\u00e0\u00e1\f\u0007\u0002\u0002\u00e1",
  "\u00e2\t\u0006\u0002\u0002\u00e2\u00ea\u0005*\u0016\b\u00e3\u00e4\f",
  "\u0006\u0002\u0002\u00e4\u00e5\u0007(\u0002\u0002\u00e5\u00ea\u0005",
  "*\u0016\u0007\u00e6\u00e7\f\u0005\u0002\u0002\u00e7\u00e8\t\u0007\u0002",
  "\u0002\u00e8\u00ea\u0005*\u0016\u0006\u00e9\u00d4\u0003\u0002\u0002",
  "\u0002\u00e9\u00d7\u0003\u0002\u0002\u0002\u00e9\u00da\u0003\u0002\u0002",
  "\u0002\u00e9\u00dd\u0003\u0002\u0002\u0002\u00e9\u00e0\u0003\u0002\u0002",
  "\u0002\u00e9\u00e3\u0003\u0002\u0002\u0002\u00e9\u00e6\u0003\u0002\u0002",
  "\u0002\u00ea\u00ed\u0003\u0002\u0002\u0002\u00eb\u00e9\u0003\u0002\u0002",
  "\u0002\u00eb\u00ec\u0003\u0002\u0002\u0002\u00ec+\u0003\u0002\u0002",
  "\u0002\u00ed\u00eb\u0003\u0002\u0002\u0002\u00ee\u00f4\u00050\u0019",
  "\u0002\u00ef\u00f4\u0005\u001c\u000f\u0002\u00f0\u00f4\u0005.\u0018",
  "\u0002\u00f1\u00f4\u00052\u001a\u0002\u00f2\u00f4\u00054\u001b\u0002",
  "\u00f3\u00ee\u0003\u0002\u0002\u0002\u00f3\u00ef\u0003\u0002\u0002\u0002",
  "\u00f3\u00f0\u0003\u0002\u0002\u0002\u00f3\u00f1\u0003\u0002\u0002\u0002",
  "\u00f3\u00f2\u0003\u0002\u0002\u0002\u00f4-\u0003\u0002\u0002\u0002",
  "\u00f5\u00f9\t\b\u0002\u0002\u00f6\u00fa\u00054\u001b\u0002\u00f7\u00fa",
  "\u00050\u0019\u0002\u00f8\u00fa\u0007-\u0002\u0002\u00f9\u00f6\u0003",
  "\u0002\u0002\u0002\u00f9\u00f7\u0003\u0002\u0002\u0002\u00f9\u00f8\u0003",
  "\u0002\u0002\u0002\u00fa/\u0003\u0002\u0002\u0002\u00fb\u0107\u0007",
  "-\u0002\u0002\u00fc\u0107\u0007)\u0002\u0002\u00fd\u0107\u0007*\u0002",
  "\u0002\u00fe\u0107\u00074\u0002\u0002\u00ff\u0107\u0007+\u0002\u0002",
  "\u0100\u0107\u0007,\u0002\u0002\u0101\u0102\t\t\u0002\u0002\u0102\u0107",
  "\u00073\u0002\u0002\u0103\u0104\t\n\u0002\u0002\u0104\u0105\u0007\u0018",
  "\u0002\u0002\u0105\u0107\t\n\u0002\u0002\u0106\u00fb\u0003\u0002\u0002",
  "\u0002\u0106\u00fc\u0003\u0002\u0002\u0002\u0106\u00fd\u0003\u0002\u0002",
  "\u0002\u0106\u00fe\u0003\u0002\u0002\u0002\u0106\u00ff\u0003\u0002\u0002",
  "\u0002\u0106\u0100\u0003\u0002\u0002\u0002\u0106\u0101\u0003\u0002\u0002",
  "\u0002\u0106\u0103\u0003\u0002\u0002\u0002\u01071\u0003\u0002\u0002",
  "\u0002\u0108\u012b\u0007\u001b\u0002\u0002\u0109\u010b\u0005\u0002\u0002",
  "\u0002\u010a\u0109\u0003\u0002\u0002\u0002\u010b\u010e\u0003\u0002\u0002",
  "\u0002\u010c\u010a\u0003\u0002\u0002\u0002\u010c\u010d\u0003\u0002\u0002",
  "\u0002\u010d\u0112\u0003\u0002\u0002\u0002\u010e\u010c\u0003\u0002\u0002",
  "\u0002\u010f\u0113\u0005$\u0013\u0002\u0110\u0113\u0005&\u0014\u0002",
  "\u0111\u0113\u0005(\u0015\u0002\u0112\u010f\u0003\u0002\u0002\u0002",
  "\u0112\u0110\u0003\u0002\u0002\u0002\u0112\u0111\u0003\u0002\u0002\u0002",
  "\u0113\u0122\u0003\u0002\u0002\u0002\u0114\u0118\u0007\u0015\u0002\u0002",
  "\u0115\u0117\u0005\u0002\u0002\u0002\u0116\u0115\u0003\u0002\u0002\u0002",
  "\u0117\u011a\u0003\u0002\u0002\u0002\u0118\u0116\u0003\u0002\u0002\u0002",
  "\u0118\u0119\u0003\u0002\u0002\u0002\u0119\u011e\u0003\u0002\u0002\u0002",
  "\u011a\u0118\u0003\u0002\u0002\u0002\u011b\u011f\u0005$\u0013\u0002",
  "\u011c\u011f\u0005&\u0014\u0002\u011d\u011f\u0005(\u0015\u0002\u011e",
  "\u011b\u0003\u0002\u0002\u0002\u011e\u011c\u0003\u0002\u0002\u0002\u011e",
  "\u011d\u0003\u0002\u0002\u0002\u011f\u0121\u0003\u0002\u0002\u0002\u0120",
  "\u0114\u0003\u0002\u0002\u0002\u0121\u0124\u0003\u0002\u0002\u0002\u0122",
  "\u0120\u0003\u0002\u0002\u0002\u0122\u0123\u0003\u0002\u0002\u0002\u0123",
  "\u0126\u0003\u0002\u0002\u0002\u0124\u0122\u0003\u0002\u0002\u0002\u0125",
  "\u0127\u0007\u0015\u0002\u0002\u0126\u0125\u0003\u0002\u0002\u0002\u0126",
  "\u0127\u0003\u0002\u0002\u0002\u0127\u0129\u0003\u0002\u0002\u0002\u0128",
  "\u012a\u0005\u0002\u0002\u0002\u0129\u0128\u0003\u0002\u0002\u0002\u0129",
  "\u012a\u0003\u0002\u0002\u0002\u012a\u012c\u0003\u0002\u0002\u0002\u012b",
  "\u010c\u0003\u0002\u0002\u0002\u012b\u012c\u0003\u0002\u0002\u0002\u012c",
  "\u012d\u0003\u0002\u0002\u0002\u012d\u012e\u0007\u001c\u0002\u0002\u012e",
  "3\u0003\u0002\u0002\u0002\u012f\u0133\u0007\u001d\u0002\u0002\u0130",
  "\u0132\u0005\u0002\u0002\u0002\u0131\u0130\u0003\u0002\u0002\u0002\u0132",
  "\u0135\u0003\u0002\u0002\u0002\u0133\u0131\u0003\u0002\u0002\u0002\u0133",
  "\u0134\u0003\u0002\u0002\u0002\u0134\u0138\u0003\u0002\u0002\u0002\u0135",
  "\u0133\u0003\u0002\u0002\u0002\u0136\u0139\u0005\u0018\r\u0002\u0137",
  "\u0139\u0005\u001a\u000e\u0002\u0138\u0136\u0003\u0002\u0002\u0002\u0138",
  "\u0137\u0003\u0002\u0002\u0002\u0139\u013d\u0003\u0002\u0002\u0002\u013a",
  "\u013c\u0005\u0002\u0002\u0002\u013b\u013a\u0003\u0002\u0002\u0002\u013c",
  "\u013f\u0003\u0002\u0002\u0002\u013d\u013b\u0003\u0002\u0002\u0002\u013d",
  "\u013e\u0003\u0002\u0002\u0002\u013e\u0140\u0003\u0002\u0002\u0002\u013f",
  "\u013d\u0003\u0002\u0002\u0002\u0140\u0141\u0007\u001e\u0002\u0002\u0141",
  "5\u0003\u0002\u0002\u0002(;?DJOT]jotw\u007f\u0086\u0095\u009a\u00a1",
  "\u00a8\u00b1\u00b8\u00be\u00ca\u00d2\u00e9\u00eb\u00f3\u00f9\u0106\u010c",
  "\u0112\u0118\u011e\u0122\u0126\u0129\u012b\u0133\u0138\u013d",
].join("");

const atn = new antlr4.atn.ATNDeserializer().deserialize(serializedATN);

const decisionsToDFA = atn.decisionToState.map((ds, index) => new antlr4.dfa.DFA(ds, index));

const sharedContextCache = new antlr4.PredictionContextCache();

export default class preludioParser extends antlr4.Parser {
  static grammarFileName = "preludio.g4";
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
    "multiAssign",
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
    this._interp = new antlr4.atn.ParserATNSimulator(this, atn, decisionsToDFA, sharedContextCache);
    this.ruleNames = preludioParser.ruleNames;
    this.literalNames = preludioParser.literalNames;
    this.symbolicNames = preludioParser.symbolicNames;
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
    this.enterRule(localctx, 0, preludioParser.RULE_nl);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 52;
      _la = this._input.LA(1);
      if (!(_la === preludioParser.NEWLINE || _la === preludioParser.COMMENT)) {
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
    this.enterRule(localctx, 2, preludioParser.RULE_program);
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
      if (_la === preludioParser.PRQL) {
        this.state = 60;
        this.programIntro();
      }

      this.state = 66;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (_la === preludioParser.NEWLINE || _la === preludioParser.COMMENT) {
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
            ((1 << preludioParser.FUNC) |
              (1 << preludioParser.LET) |
              (1 << preludioParser.PLUS) |
              (1 << preludioParser.MINUS) |
              (1 << preludioParser.LBRACKET) |
              (1 << preludioParser.LPAREN) |
              (1 << preludioParser.BACKTICK))) !==
            0) ||
        (((_la - 37) & ~0x1f) == 0 &&
          ((1 << (_la - 37)) &
            ((1 << (preludioParser.NOT - 37)) |
              (1 << (preludioParser.NULL_ - 37)) |
              (1 << (preludioParser.BOOLEAN - 37)) |
              (1 << (preludioParser.INTEGER - 37)) |
              (1 << (preludioParser.FLOAT - 37)) |
              (1 << (preludioParser.IDENT - 37)) |
              (1 << (preludioParser.STRING - 37)))) !==
            0)
      ) {
        this.state = 72;
        this._errHandler.sync(this);
        switch (this._input.LA(1)) {
          case preludioParser.FUNC:
            this.state = 69;
            this.funcDef();
            break;
          case preludioParser.LET:
            this.state = 70;
            this.stmt();
            break;
          case preludioParser.PLUS:
          case preludioParser.MINUS:
          case preludioParser.LBRACKET:
          case preludioParser.LPAREN:
          case preludioParser.BACKTICK:
          case preludioParser.NOT:
          case preludioParser.NULL_:
          case preludioParser.BOOLEAN:
          case preludioParser.INTEGER:
          case preludioParser.FLOAT:
          case preludioParser.IDENT:
          case preludioParser.STRING:
            this.state = 71;
            this.pipeline();
            break;
          default:
            throw new antlr4.error.NoViableAltException(this);
        }
        this.state = 77;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
        while (_la === preludioParser.NEWLINE || _la === preludioParser.COMMENT) {
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
      this.match(preludioParser.EOF);
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
    this.enterRule(localctx, 4, preludioParser.RULE_programIntro);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 87;
      this.match(preludioParser.PRQL);
      this.state = 91;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (_la === preludioParser.IDENT) {
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
    this.enterRule(localctx, 6, preludioParser.RULE_funcDef);
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 96;
      this.match(preludioParser.FUNC);
      this.state = 97;
      this.funcDefName();
      this.state = 98;
      this.funcDefParams();
      this.state = 99;
      this.match(preludioParser.ARROW);
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
    this.enterRule(localctx, 8, preludioParser.RULE_funcDefName);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 102;
      this.match(preludioParser.IDENT);
      this.state = 104;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      if (_la === preludioParser.LANG) {
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
    this.enterRule(localctx, 10, preludioParser.RULE_funcDefParams);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 109;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (_la === preludioParser.IDENT) {
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
    this.enterRule(localctx, 12, preludioParser.RULE_funcDefParam);
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
          this.match(preludioParser.IDENT);
          break;
      }
      this.state = 117;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      if (_la === preludioParser.LANG) {
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
    this.enterRule(localctx, 14, preludioParser.RULE_typeDef);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 119;
      this.match(preludioParser.LANG);
      this.state = 120;
      this.typeTerm();
      this.state = 121;
      this.match(preludioParser.BAR);
      this.state = 125;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (_la === preludioParser.IDENT) {
        this.state = 122;
        this.typeTerm();
        this.state = 127;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
      }
      this.state = 128;
      this.match(preludioParser.RANG);
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
    this.enterRule(localctx, 16, preludioParser.RULE_typeTerm);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 130;
      this.match(preludioParser.IDENT);
      this.state = 132;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      if (_la === preludioParser.LANG) {
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
    this.enterRule(localctx, 18, preludioParser.RULE_stmt);
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
    this.enterRule(localctx, 20, preludioParser.RULE_assignStmt);
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 136;
      this.match(preludioParser.LET);
      this.state = 137;
      this.match(preludioParser.IDENT);
      this.state = 138;
      this.match(preludioParser.ASSIGN);
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
    this.enterRule(localctx, 22, preludioParser.RULE_pipeline);
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
        case preludioParser.NEWLINE:
        case preludioParser.COMMENT:
          this.state = 150;
          this.nl();
          break;
        case preludioParser.EOF:
          this.state = 151;
          this.match(preludioParser.EOF);
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
    this.enterRule(localctx, 24, preludioParser.RULE_inlinePipeline);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 154;
      this.exprCall();
      this.state = 159;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (_la === preludioParser.BAR) {
        this.state = 155;
        this.match(preludioParser.BAR);
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
    this.enterRule(localctx, 26, preludioParser.RULE_identBacktick);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 162;
      this.match(preludioParser.BACKTICK);
      this.state = 166;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (
        ((_la & ~0x1f) == 0 &&
          ((1 << _la) &
            ((1 << preludioParser.FUNC) |
              (1 << preludioParser.PRQL) |
              (1 << preludioParser.LET) |
              (1 << preludioParser.ARROW) |
              (1 << preludioParser.ASSIGN) |
              (1 << preludioParser.PLUS) |
              (1 << preludioParser.MINUS) |
              (1 << preludioParser.STAR) |
              (1 << preludioParser.POW) |
              (1 << preludioParser.DIV) |
              (1 << preludioParser.MOD) |
              (1 << preludioParser.MODEL) |
              (1 << preludioParser.EQ) |
              (1 << preludioParser.NE) |
              (1 << preludioParser.LE) |
              (1 << preludioParser.GE) |
              (1 << preludioParser.BAR) |
              (1 << preludioParser.COLON) |
              (1 << preludioParser.COMMA) |
              (1 << preludioParser.DOT) |
              (1 << preludioParser.DOLLAR) |
              (1 << preludioParser.RANGE) |
              (1 << preludioParser.LANG) |
              (1 << preludioParser.RANG) |
              (1 << preludioParser.LBRACKET) |
              (1 << preludioParser.RBRACKET) |
              (1 << preludioParser.LPAREN) |
              (1 << preludioParser.RPAREN) |
              (1 << preludioParser.UNDERSCORE) |
              (1 << preludioParser.DOUBLE_QUOTE))) !==
            0) ||
        (((_la - 32) & ~0x1f) == 0 &&
          ((1 << (_la - 32)) &
            ((1 << (preludioParser.SINGLE_QUOTE - 32)) |
              (1 << (preludioParser.TRIPLE_DOUBLE_QUOTE - 32)) |
              (1 << (preludioParser.TRIPLE_SINGLE_QUOTE - 32)) |
              (1 << (preludioParser.AND - 32)) |
              (1 << (preludioParser.OR - 32)) |
              (1 << (preludioParser.NOT - 32)) |
              (1 << (preludioParser.COALESCE - 32)) |
              (1 << (preludioParser.NULL_ - 32)) |
              (1 << (preludioParser.BOOLEAN - 32)) |
              (1 << (preludioParser.INTEGER - 32)) |
              (1 << (preludioParser.FLOAT - 32)) |
              (1 << (preludioParser.IDENT - 32)) |
              (1 << (preludioParser.IDENT_START - 32)) |
              (1 << (preludioParser.IDENT_NEXT - 32)) |
              (1 << (preludioParser.WHITESPACE - 32)) |
              (1 << (preludioParser.COMMENT - 32)) |
              (1 << (preludioParser.INTERVAL_KIND - 32)) |
              (1 << (preludioParser.STRING - 32)))) !==
            0)
      ) {
        this.state = 163;
        _la = this._input.LA(1);
        if (_la <= 0 || _la === preludioParser.BACKTICK || _la === preludioParser.NEWLINE) {
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
      this.match(preludioParser.BACKTICK);
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
    this.enterRule(localctx, 28, preludioParser.RULE_funcCall);
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 171;
      this.match(preludioParser.IDENT);
      this.state = 175;
      this._errHandler.sync(this);
      var _alt = this._interp.adaptivePredict(this._input, 17, this._ctx);
      while (_alt != 2 && _alt != antlr4.atn.ATN.INVALID_ALT_NUMBER) {
        if (_alt === 1) {
          this.state = 172;
          this.funcCallParam();
        }
        this.state = 177;
        this._errHandler.sync(this);
        _alt = this._interp.adaptivePredict(this._input, 17, this._ctx);
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
    this.enterRule(localctx, 30, preludioParser.RULE_funcCallParam);
    try {
      this.state = 182;
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
          this.multiAssign();
          break;

        case 4:
          this.enterOuterAlt(localctx, 4);
          this.state = 181;
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
    this.enterRule(localctx, 32, preludioParser.RULE_namedArg);
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 184;
      this.match(preludioParser.IDENT);
      this.state = 185;
      this.match(preludioParser.COLON);
      this.state = 188;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 19, this._ctx);
      switch (la_) {
        case 1:
          this.state = 186;
          this.assign();
          break;

        case 2:
          this.state = 187;
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
    this.enterRule(localctx, 34, preludioParser.RULE_assign);
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 190;
      this.match(preludioParser.IDENT);
      this.state = 191;
      this.match(preludioParser.ASSIGN);
      this.state = 192;
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

  multiAssign() {
    let localctx = new MultiAssignContext(this, this._ctx, this.state);
    this.enterRule(localctx, 36, preludioParser.RULE_multiAssign);
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 194;
      this.list();
      this.state = 195;
      this.match(preludioParser.ASSIGN);
      this.state = 196;
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
    this.enterRule(localctx, 38, preludioParser.RULE_exprCall);
    try {
      this.state = 200;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 20, this._ctx);
      switch (la_) {
        case 1:
          this.enterOuterAlt(localctx, 1);
          this.state = 198;
          this.expr(0);
          break;

        case 2:
          this.enterOuterAlt(localctx, 2);
          this.state = 199;
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
    this.enterRecursionRule(localctx, 40, preludioParser.RULE_expr, _p);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 208;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 21, this._ctx);
      switch (la_) {
        case 1:
          this.state = 203;
          this.match(preludioParser.LPAREN);
          this.state = 204;
          this.expr(0);
          this.state = 205;
          this.match(preludioParser.RPAREN);
          break;

        case 2:
          this.state = 207;
          this.term();
          break;
      }
      this._ctx.stop = this._input.LT(-1);
      this.state = 233;
      this._errHandler.sync(this);
      var _alt = this._interp.adaptivePredict(this._input, 23, this._ctx);
      while (_alt != 2 && _alt != antlr4.atn.ATN.INVALID_ALT_NUMBER) {
        if (_alt === 1) {
          if (this._parseListeners !== null) {
            this.triggerExitRuleEvent();
          }
          _prevctx = localctx;
          this.state = 231;
          this._errHandler.sync(this);
          var la_ = this._interp.adaptivePredict(this._input, 22, this._ctx);
          switch (la_) {
            case 1:
              localctx = new ExprContext(this, _parentctx, _parentState);
              this.pushNewRecursionContext(localctx, _startState, preludioParser.RULE_expr);
              this.state = 210;
              if (!this.precpred(this._ctx, 9)) {
                throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 9)");
              }
              this.state = 211;
              _la = this._input.LA(1);
              if (
                !(
                  (_la & ~0x1f) == 0 &&
                  ((1 << _la) & ((1 << preludioParser.STAR) | (1 << preludioParser.DIV) | (1 << preludioParser.MOD))) !== 0
                )
              ) {
                this._errHandler.recoverInline(this);
              } else {
                this._errHandler.reportMatch(this);
                this.consume();
              }
              this.state = 212;
              this.expr(10);
              break;

            case 2:
              localctx = new ExprContext(this, _parentctx, _parentState);
              this.pushNewRecursionContext(localctx, _startState, preludioParser.RULE_expr);
              this.state = 213;
              if (!this.precpred(this._ctx, 8)) {
                throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 8)");
              }
              this.state = 214;
              _la = this._input.LA(1);
              if (!(_la === preludioParser.PLUS || _la === preludioParser.MINUS)) {
                this._errHandler.recoverInline(this);
              } else {
                this._errHandler.reportMatch(this);
                this.consume();
              }
              this.state = 215;
              this.expr(9);
              break;

            case 3:
              localctx = new ExprContext(this, _parentctx, _parentState);
              this.pushNewRecursionContext(localctx, _startState, preludioParser.RULE_expr);
              this.state = 216;
              if (!this.precpred(this._ctx, 7)) {
                throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 7)");
              }
              this.state = 217;
              this.match(preludioParser.POW);
              this.state = 218;
              this.expr(8);
              break;

            case 4:
              localctx = new ExprContext(this, _parentctx, _parentState);
              this.pushNewRecursionContext(localctx, _startState, preludioParser.RULE_expr);
              this.state = 219;
              if (!this.precpred(this._ctx, 6)) {
                throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 6)");
              }
              this.state = 220;
              this.match(preludioParser.MODEL);
              this.state = 221;
              this.expr(7);
              break;

            case 5:
              localctx = new ExprContext(this, _parentctx, _parentState);
              this.pushNewRecursionContext(localctx, _startState, preludioParser.RULE_expr);
              this.state = 222;
              if (!this.precpred(this._ctx, 5)) {
                throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 5)");
              }
              this.state = 223;
              _la = this._input.LA(1);
              if (
                !(
                  (_la & ~0x1f) == 0 &&
                  ((1 << _la) &
                    ((1 << preludioParser.EQ) |
                      (1 << preludioParser.NE) |
                      (1 << preludioParser.LE) |
                      (1 << preludioParser.GE) |
                      (1 << preludioParser.LANG) |
                      (1 << preludioParser.RANG))) !==
                    0
                )
              ) {
                this._errHandler.recoverInline(this);
              } else {
                this._errHandler.reportMatch(this);
                this.consume();
              }
              this.state = 224;
              this.expr(6);
              break;

            case 6:
              localctx = new ExprContext(this, _parentctx, _parentState);
              this.pushNewRecursionContext(localctx, _startState, preludioParser.RULE_expr);
              this.state = 225;
              if (!this.precpred(this._ctx, 4)) {
                throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 4)");
              }
              this.state = 226;
              this.match(preludioParser.COALESCE);
              this.state = 227;
              this.expr(5);
              break;

            case 7:
              localctx = new ExprContext(this, _parentctx, _parentState);
              this.pushNewRecursionContext(localctx, _startState, preludioParser.RULE_expr);
              this.state = 228;
              if (!this.precpred(this._ctx, 3)) {
                throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 3)");
              }
              this.state = 229;
              _la = this._input.LA(1);
              if (!(_la === preludioParser.AND || _la === preludioParser.OR)) {
                this._errHandler.recoverInline(this);
              } else {
                this._errHandler.reportMatch(this);
                this.consume();
              }
              this.state = 230;
              this.expr(4);
              break;
          }
        }
        this.state = 235;
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
    this.enterRule(localctx, 42, preludioParser.RULE_term);
    try {
      this.state = 241;
      this._errHandler.sync(this);
      switch (this._input.LA(1)) {
        case preludioParser.NULL_:
        case preludioParser.BOOLEAN:
        case preludioParser.INTEGER:
        case preludioParser.FLOAT:
        case preludioParser.IDENT:
        case preludioParser.STRING:
          this.enterOuterAlt(localctx, 1);
          this.state = 236;
          this.literal();
          break;
        case preludioParser.BACKTICK:
          this.enterOuterAlt(localctx, 2);
          this.state = 237;
          this.identBacktick();
          break;
        case preludioParser.PLUS:
        case preludioParser.MINUS:
        case preludioParser.NOT:
          this.enterOuterAlt(localctx, 3);
          this.state = 238;
          this.exprUnary();
          break;
        case preludioParser.LBRACKET:
          this.enterOuterAlt(localctx, 4);
          this.state = 239;
          this.list();
          break;
        case preludioParser.LPAREN:
          this.enterOuterAlt(localctx, 5);
          this.state = 240;
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
    this.enterRule(localctx, 44, preludioParser.RULE_exprUnary);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 243;
      _la = this._input.LA(1);
      if (
        !(
          ((_la - 6) & ~0x1f) == 0 &&
          ((1 << (_la - 6)) & ((1 << (preludioParser.PLUS - 6)) | (1 << (preludioParser.MINUS - 6)) | (1 << (preludioParser.NOT - 6)))) !==
            0
        )
      ) {
        this._errHandler.recoverInline(this);
      } else {
        this._errHandler.reportMatch(this);
        this.consume();
      }
      this.state = 247;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 25, this._ctx);
      switch (la_) {
        case 1:
          this.state = 244;
          this.nestedPipeline();
          break;

        case 2:
          this.state = 245;
          this.literal();
          break;

        case 3:
          this.state = 246;
          this.match(preludioParser.IDENT);
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
    this.enterRule(localctx, 46, preludioParser.RULE_literal);
    var _la = 0; // Token type
    try {
      this.state = 260;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 26, this._ctx);
      switch (la_) {
        case 1:
          this.enterOuterAlt(localctx, 1);
          this.state = 249;
          this.match(preludioParser.IDENT);
          break;

        case 2:
          this.enterOuterAlt(localctx, 2);
          this.state = 250;
          this.match(preludioParser.NULL_);
          break;

        case 3:
          this.enterOuterAlt(localctx, 3);
          this.state = 251;
          this.match(preludioParser.BOOLEAN);
          break;

        case 4:
          this.enterOuterAlt(localctx, 4);
          this.state = 252;
          this.match(preludioParser.STRING);
          break;

        case 5:
          this.enterOuterAlt(localctx, 5);
          this.state = 253;
          this.match(preludioParser.INTEGER);
          break;

        case 6:
          this.enterOuterAlt(localctx, 6);
          this.state = 254;
          this.match(preludioParser.FLOAT);
          break;

        case 7:
          this.enterOuterAlt(localctx, 7);
          this.state = 255;
          _la = this._input.LA(1);
          if (!(_la === preludioParser.INTEGER || _la === preludioParser.FLOAT)) {
            this._errHandler.recoverInline(this);
          } else {
            this._errHandler.reportMatch(this);
            this.consume();
          }
          this.state = 256;
          this.match(preludioParser.INTERVAL_KIND);
          break;

        case 8:
          this.enterOuterAlt(localctx, 8);
          this.state = 257;
          _la = this._input.LA(1);
          if (
            !(
              ((_la - 41) & ~0x1f) == 0 &&
              ((1 << (_la - 41)) &
                ((1 << (preludioParser.INTEGER - 41)) | (1 << (preludioParser.FLOAT - 41)) | (1 << (preludioParser.IDENT - 41)))) !==
                0
            )
          ) {
            this._errHandler.recoverInline(this);
          } else {
            this._errHandler.reportMatch(this);
            this.consume();
          }
          this.state = 258;
          this.match(preludioParser.RANGE);
          this.state = 259;
          _la = this._input.LA(1);
          if (
            !(
              ((_la - 41) & ~0x1f) == 0 &&
              ((1 << (_la - 41)) &
                ((1 << (preludioParser.INTEGER - 41)) | (1 << (preludioParser.FLOAT - 41)) | (1 << (preludioParser.IDENT - 41)))) !==
                0
            )
          ) {
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
    this.enterRule(localctx, 48, preludioParser.RULE_list);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 262;
      this.match(preludioParser.LBRACKET);
      this.state = 297;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      if (
        ((_la & ~0x1f) == 0 &&
          ((1 << _la) &
            ((1 << preludioParser.PLUS) |
              (1 << preludioParser.MINUS) |
              (1 << preludioParser.LBRACKET) |
              (1 << preludioParser.LPAREN) |
              (1 << preludioParser.BACKTICK))) !==
            0) ||
        (((_la - 37) & ~0x1f) == 0 &&
          ((1 << (_la - 37)) &
            ((1 << (preludioParser.NOT - 37)) |
              (1 << (preludioParser.NULL_ - 37)) |
              (1 << (preludioParser.BOOLEAN - 37)) |
              (1 << (preludioParser.INTEGER - 37)) |
              (1 << (preludioParser.FLOAT - 37)) |
              (1 << (preludioParser.IDENT - 37)) |
              (1 << (preludioParser.NEWLINE - 37)) |
              (1 << (preludioParser.COMMENT - 37)) |
              (1 << (preludioParser.STRING - 37)))) !==
            0)
      ) {
        this.state = 266;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
        while (_la === preludioParser.NEWLINE || _la === preludioParser.COMMENT) {
          this.state = 263;
          this.nl();
          this.state = 268;
          this._errHandler.sync(this);
          _la = this._input.LA(1);
        }
        this.state = 272;
        this._errHandler.sync(this);
        var la_ = this._interp.adaptivePredict(this._input, 28, this._ctx);
        switch (la_) {
          case 1:
            this.state = 269;
            this.assign();
            break;

          case 2:
            this.state = 270;
            this.multiAssign();
            break;

          case 3:
            this.state = 271;
            this.exprCall();
            break;
        }
        this.state = 288;
        this._errHandler.sync(this);
        var _alt = this._interp.adaptivePredict(this._input, 31, this._ctx);
        while (_alt != 2 && _alt != antlr4.atn.ATN.INVALID_ALT_NUMBER) {
          if (_alt === 1) {
            this.state = 274;
            this.match(preludioParser.COMMA);
            this.state = 278;
            this._errHandler.sync(this);
            _la = this._input.LA(1);
            while (_la === preludioParser.NEWLINE || _la === preludioParser.COMMENT) {
              this.state = 275;
              this.nl();
              this.state = 280;
              this._errHandler.sync(this);
              _la = this._input.LA(1);
            }
            this.state = 284;
            this._errHandler.sync(this);
            var la_ = this._interp.adaptivePredict(this._input, 30, this._ctx);
            switch (la_) {
              case 1:
                this.state = 281;
                this.assign();
                break;

              case 2:
                this.state = 282;
                this.multiAssign();
                break;

              case 3:
                this.state = 283;
                this.exprCall();
                break;
            }
          }
          this.state = 290;
          this._errHandler.sync(this);
          _alt = this._interp.adaptivePredict(this._input, 31, this._ctx);
        }

        this.state = 292;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
        if (_la === preludioParser.COMMA) {
          this.state = 291;
          this.match(preludioParser.COMMA);
        }

        this.state = 295;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
        if (_la === preludioParser.NEWLINE || _la === preludioParser.COMMENT) {
          this.state = 294;
          this.nl();
        }
      }

      this.state = 299;
      this.match(preludioParser.RBRACKET);
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
    this.enterRule(localctx, 50, preludioParser.RULE_nestedPipeline);
    var _la = 0; // Token type
    try {
      this.enterOuterAlt(localctx, 1);
      this.state = 301;
      this.match(preludioParser.LPAREN);
      this.state = 305;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (_la === preludioParser.NEWLINE || _la === preludioParser.COMMENT) {
        this.state = 302;
        this.nl();
        this.state = 307;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
      }
      this.state = 310;
      this._errHandler.sync(this);
      var la_ = this._interp.adaptivePredict(this._input, 36, this._ctx);
      switch (la_) {
        case 1:
          this.state = 308;
          this.pipeline();
          break;

        case 2:
          this.state = 309;
          this.inlinePipeline();
          break;
      }
      this.state = 315;
      this._errHandler.sync(this);
      _la = this._input.LA(1);
      while (_la === preludioParser.NEWLINE || _la === preludioParser.COMMENT) {
        this.state = 312;
        this.nl();
        this.state = 317;
        this._errHandler.sync(this);
        _la = this._input.LA(1);
      }
      this.state = 318;
      this.match(preludioParser.RPAREN);
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

preludioParser.EOF = antlr4.Token.EOF;
preludioParser.FUNC = 1;
preludioParser.PRQL = 2;
preludioParser.LET = 3;
preludioParser.ARROW = 4;
preludioParser.ASSIGN = 5;
preludioParser.PLUS = 6;
preludioParser.MINUS = 7;
preludioParser.STAR = 8;
preludioParser.POW = 9;
preludioParser.DIV = 10;
preludioParser.MOD = 11;
preludioParser.MODEL = 12;
preludioParser.EQ = 13;
preludioParser.NE = 14;
preludioParser.LE = 15;
preludioParser.GE = 16;
preludioParser.BAR = 17;
preludioParser.COLON = 18;
preludioParser.COMMA = 19;
preludioParser.DOT = 20;
preludioParser.DOLLAR = 21;
preludioParser.RANGE = 22;
preludioParser.LANG = 23;
preludioParser.RANG = 24;
preludioParser.LBRACKET = 25;
preludioParser.RBRACKET = 26;
preludioParser.LPAREN = 27;
preludioParser.RPAREN = 28;
preludioParser.UNDERSCORE = 29;
preludioParser.BACKTICK = 30;
preludioParser.DOUBLE_QUOTE = 31;
preludioParser.SINGLE_QUOTE = 32;
preludioParser.TRIPLE_DOUBLE_QUOTE = 33;
preludioParser.TRIPLE_SINGLE_QUOTE = 34;
preludioParser.AND = 35;
preludioParser.OR = 36;
preludioParser.NOT = 37;
preludioParser.COALESCE = 38;
preludioParser.NULL_ = 39;
preludioParser.BOOLEAN = 40;
preludioParser.INTEGER = 41;
preludioParser.FLOAT = 42;
preludioParser.IDENT = 43;
preludioParser.IDENT_START = 44;
preludioParser.IDENT_NEXT = 45;
preludioParser.WHITESPACE = 46;
preludioParser.NEWLINE = 47;
preludioParser.COMMENT = 48;
preludioParser.INTERVAL_KIND = 49;
preludioParser.STRING = 50;

preludioParser.RULE_nl = 0;
preludioParser.RULE_program = 1;
preludioParser.RULE_programIntro = 2;
preludioParser.RULE_funcDef = 3;
preludioParser.RULE_funcDefName = 4;
preludioParser.RULE_funcDefParams = 5;
preludioParser.RULE_funcDefParam = 6;
preludioParser.RULE_typeDef = 7;
preludioParser.RULE_typeTerm = 8;
preludioParser.RULE_stmt = 9;
preludioParser.RULE_assignStmt = 10;
preludioParser.RULE_pipeline = 11;
preludioParser.RULE_inlinePipeline = 12;
preludioParser.RULE_identBacktick = 13;
preludioParser.RULE_funcCall = 14;
preludioParser.RULE_funcCallParam = 15;
preludioParser.RULE_namedArg = 16;
preludioParser.RULE_assign = 17;
preludioParser.RULE_multiAssign = 18;
preludioParser.RULE_exprCall = 19;
preludioParser.RULE_expr = 20;
preludioParser.RULE_term = 21;
preludioParser.RULE_exprUnary = 22;
preludioParser.RULE_literal = 23;
preludioParser.RULE_list = 24;
preludioParser.RULE_nestedPipeline = 25;

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
    this.ruleIndex = preludioParser.RULE_nl;
  }

  NEWLINE() {
    return this.getToken(preludioParser.NEWLINE, 0);
  }

  COMMENT() {
    return this.getToken(preludioParser.COMMENT, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterNl(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_program;
  }

  EOF() {
    return this.getToken(preludioParser.EOF, 0);
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
    if (listener instanceof preludioListener) {
      listener.enterProgram(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_programIntro;
  }

  PRQL() {
    return this.getToken(preludioParser.PRQL, 0);
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
    if (listener instanceof preludioListener) {
      listener.enterProgramIntro(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_funcDef;
  }

  FUNC() {
    return this.getToken(preludioParser.FUNC, 0);
  }

  funcDefName() {
    return this.getTypedRuleContext(FuncDefNameContext, 0);
  }

  funcDefParams() {
    return this.getTypedRuleContext(FuncDefParamsContext, 0);
  }

  ARROW() {
    return this.getToken(preludioParser.ARROW, 0);
  }

  expr() {
    return this.getTypedRuleContext(ExprContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterFuncDef(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_funcDefName;
  }

  IDENT() {
    return this.getToken(preludioParser.IDENT, 0);
  }

  typeDef() {
    return this.getTypedRuleContext(TypeDefContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterFuncDefName(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_funcDefParams;
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
    if (listener instanceof preludioListener) {
      listener.enterFuncDefParams(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_funcDefParam;
  }

  namedArg() {
    return this.getTypedRuleContext(NamedArgContext, 0);
  }

  IDENT() {
    return this.getToken(preludioParser.IDENT, 0);
  }

  typeDef() {
    return this.getTypedRuleContext(TypeDefContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterFuncDefParam(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_typeDef;
  }

  LANG() {
    return this.getToken(preludioParser.LANG, 0);
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
    return this.getToken(preludioParser.BAR, 0);
  }

  RANG() {
    return this.getToken(preludioParser.RANG, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterTypeDef(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_typeTerm;
  }

  IDENT() {
    return this.getToken(preludioParser.IDENT, 0);
  }

  typeDef() {
    return this.getTypedRuleContext(TypeDefContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterTypeTerm(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_stmt;
  }

  assignStmt() {
    return this.getTypedRuleContext(AssignStmtContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterStmt(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_assignStmt;
  }

  LET() {
    return this.getToken(preludioParser.LET, 0);
  }

  IDENT() {
    return this.getToken(preludioParser.IDENT, 0);
  }

  ASSIGN() {
    return this.getToken(preludioParser.ASSIGN, 0);
  }

  expr() {
    return this.getTypedRuleContext(ExprContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterAssignStmt(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_pipeline;
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
    return this.getToken(preludioParser.EOF, 0);
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
    if (listener instanceof preludioListener) {
      listener.enterPipeline(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_inlinePipeline;
  }

  exprCall() {
    return this.getTypedRuleContext(ExprCallContext, 0);
  }

  BAR = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTokens(preludioParser.BAR);
    } else {
      return this.getToken(preludioParser.BAR, i);
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
    if (listener instanceof preludioListener) {
      listener.enterInlinePipeline(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_identBacktick;
  }

  BACKTICK = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTokens(preludioParser.BACKTICK);
    } else {
      return this.getToken(preludioParser.BACKTICK, i);
    }
  };

  NEWLINE = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTokens(preludioParser.NEWLINE);
    } else {
      return this.getToken(preludioParser.NEWLINE, i);
    }
  };

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterIdentBacktick(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_funcCall;
  }

  IDENT() {
    return this.getToken(preludioParser.IDENT, 0);
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
    if (listener instanceof preludioListener) {
      listener.enterFuncCall(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_funcCallParam;
  }

  namedArg() {
    return this.getTypedRuleContext(NamedArgContext, 0);
  }

  assign() {
    return this.getTypedRuleContext(AssignContext, 0);
  }

  multiAssign() {
    return this.getTypedRuleContext(MultiAssignContext, 0);
  }

  expr() {
    return this.getTypedRuleContext(ExprContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterFuncCallParam(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_namedArg;
  }

  IDENT() {
    return this.getToken(preludioParser.IDENT, 0);
  }

  COLON() {
    return this.getToken(preludioParser.COLON, 0);
  }

  assign() {
    return this.getTypedRuleContext(AssignContext, 0);
  }

  expr() {
    return this.getTypedRuleContext(ExprContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterNamedArg(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_assign;
  }

  IDENT() {
    return this.getToken(preludioParser.IDENT, 0);
  }

  ASSIGN() {
    return this.getToken(preludioParser.ASSIGN, 0);
  }

  exprCall() {
    return this.getTypedRuleContext(ExprCallContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterAssign(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
      listener.exitAssign(this);
    }
  }
}

class MultiAssignContext extends antlr4.ParserRuleContext {
  constructor(parser, parent, invokingState) {
    if (parent === undefined) {
      parent = null;
    }
    if (invokingState === undefined || invokingState === null) {
      invokingState = -1;
    }
    super(parent, invokingState);
    this.parser = parser;
    this.ruleIndex = preludioParser.RULE_multiAssign;
  }

  list() {
    return this.getTypedRuleContext(ListContext, 0);
  }

  ASSIGN() {
    return this.getToken(preludioParser.ASSIGN, 0);
  }

  exprCall() {
    return this.getTypedRuleContext(ExprCallContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterMultiAssign(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
      listener.exitMultiAssign(this);
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
    this.ruleIndex = preludioParser.RULE_exprCall;
  }

  expr() {
    return this.getTypedRuleContext(ExprContext, 0);
  }

  funcCall() {
    return this.getTypedRuleContext(FuncCallContext, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterExprCall(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_expr;
  }

  LPAREN() {
    return this.getToken(preludioParser.LPAREN, 0);
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
    return this.getToken(preludioParser.RPAREN, 0);
  }

  term() {
    return this.getTypedRuleContext(TermContext, 0);
  }

  STAR() {
    return this.getToken(preludioParser.STAR, 0);
  }

  DIV() {
    return this.getToken(preludioParser.DIV, 0);
  }

  MOD() {
    return this.getToken(preludioParser.MOD, 0);
  }

  MINUS() {
    return this.getToken(preludioParser.MINUS, 0);
  }

  PLUS() {
    return this.getToken(preludioParser.PLUS, 0);
  }

  POW() {
    return this.getToken(preludioParser.POW, 0);
  }

  MODEL() {
    return this.getToken(preludioParser.MODEL, 0);
  }

  EQ() {
    return this.getToken(preludioParser.EQ, 0);
  }

  NE() {
    return this.getToken(preludioParser.NE, 0);
  }

  GE() {
    return this.getToken(preludioParser.GE, 0);
  }

  LE() {
    return this.getToken(preludioParser.LE, 0);
  }

  LANG() {
    return this.getToken(preludioParser.LANG, 0);
  }

  RANG() {
    return this.getToken(preludioParser.RANG, 0);
  }

  COALESCE() {
    return this.getToken(preludioParser.COALESCE, 0);
  }

  AND() {
    return this.getToken(preludioParser.AND, 0);
  }

  OR() {
    return this.getToken(preludioParser.OR, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterExpr(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_term;
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
    if (listener instanceof preludioListener) {
      listener.enterTerm(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_exprUnary;
  }

  MINUS() {
    return this.getToken(preludioParser.MINUS, 0);
  }

  PLUS() {
    return this.getToken(preludioParser.PLUS, 0);
  }

  NOT() {
    return this.getToken(preludioParser.NOT, 0);
  }

  nestedPipeline() {
    return this.getTypedRuleContext(NestedPipelineContext, 0);
  }

  literal() {
    return this.getTypedRuleContext(LiteralContext, 0);
  }

  IDENT() {
    return this.getToken(preludioParser.IDENT, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterExprUnary(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_literal;
  }

  IDENT = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTokens(preludioParser.IDENT);
    } else {
      return this.getToken(preludioParser.IDENT, i);
    }
  };

  NULL_() {
    return this.getToken(preludioParser.NULL_, 0);
  }

  BOOLEAN() {
    return this.getToken(preludioParser.BOOLEAN, 0);
  }

  STRING() {
    return this.getToken(preludioParser.STRING, 0);
  }

  INTEGER = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTokens(preludioParser.INTEGER);
    } else {
      return this.getToken(preludioParser.INTEGER, i);
    }
  };

  FLOAT = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTokens(preludioParser.FLOAT);
    } else {
      return this.getToken(preludioParser.FLOAT, i);
    }
  };

  INTERVAL_KIND() {
    return this.getToken(preludioParser.INTERVAL_KIND, 0);
  }

  RANGE() {
    return this.getToken(preludioParser.RANGE, 0);
  }

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterLiteral(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_list;
  }

  LBRACKET() {
    return this.getToken(preludioParser.LBRACKET, 0);
  }

  RBRACKET() {
    return this.getToken(preludioParser.RBRACKET, 0);
  }

  assign = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(AssignContext);
    } else {
      return this.getTypedRuleContext(AssignContext, i);
    }
  };

  multiAssign = function (i) {
    if (i === undefined) {
      i = null;
    }
    if (i === null) {
      return this.getTypedRuleContexts(MultiAssignContext);
    } else {
      return this.getTypedRuleContext(MultiAssignContext, i);
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
      return this.getTokens(preludioParser.COMMA);
    } else {
      return this.getToken(preludioParser.COMMA, i);
    }
  };

  enterRule(listener) {
    if (listener instanceof preludioListener) {
      listener.enterList(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
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
    this.ruleIndex = preludioParser.RULE_nestedPipeline;
  }

  LPAREN() {
    return this.getToken(preludioParser.LPAREN, 0);
  }

  RPAREN() {
    return this.getToken(preludioParser.RPAREN, 0);
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
    if (listener instanceof preludioListener) {
      listener.enterNestedPipeline(this);
    }
  }

  exitRule(listener) {
    if (listener instanceof preludioListener) {
      listener.exitNestedPipeline(this);
    }
  }
}

preludioParser.NlContext = NlContext;
preludioParser.ProgramContext = ProgramContext;
preludioParser.ProgramIntroContext = ProgramIntroContext;
preludioParser.FuncDefContext = FuncDefContext;
preludioParser.FuncDefNameContext = FuncDefNameContext;
preludioParser.FuncDefParamsContext = FuncDefParamsContext;
preludioParser.FuncDefParamContext = FuncDefParamContext;
preludioParser.TypeDefContext = TypeDefContext;
preludioParser.TypeTermContext = TypeTermContext;
preludioParser.StmtContext = StmtContext;
preludioParser.AssignStmtContext = AssignStmtContext;
preludioParser.PipelineContext = PipelineContext;
preludioParser.InlinePipelineContext = InlinePipelineContext;
preludioParser.IdentBacktickContext = IdentBacktickContext;
preludioParser.FuncCallContext = FuncCallContext;
preludioParser.FuncCallParamContext = FuncCallParamContext;
preludioParser.NamedArgContext = NamedArgContext;
preludioParser.AssignContext = AssignContext;
preludioParser.MultiAssignContext = MultiAssignContext;
preludioParser.ExprCallContext = ExprCallContext;
preludioParser.ExprContext = ExprContext;
preludioParser.TermContext = TermContext;
preludioParser.ExprUnaryContext = ExprUnaryContext;
preludioParser.LiteralContext = LiteralContext;
preludioParser.ListContext = ListContext;
preludioParser.NestedPipelineContext = NestedPipelineContext;
