// Generated from .\prql.g4 by ANTLR 4.9.3
// jshint ignore: start
import antlr4 from 'antlr4';
import prqlListener from './prqlListener.js';

const serializedATN = ["\u0003\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786",
    "\u5964\u00031\u012d\u0004\u0002\t\u0002\u0004\u0003\t\u0003\u0004\u0004",
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
    "\n\u0005\n\u0087\n\n\u0003\u000b\u0003\u000b\u0003\u000b\u0003\u000b",
    "\u0003\u000b\u0003\f\u0003\f\u0005\f\u0090\n\f\u0003\r\u0003\r\u0003",
    "\r\u0003\r\u0007\r\u0096\n\r\f\r\u000e\r\u0099\u000b\r\u0003\u000e\u0003",
    "\u000e\u0007\u000e\u009d\n\u000e\f\u000e\u000e\u000e\u00a0\u000b\u000e",
    "\u0003\u000e\u0003\u000e\u0003\u000f\u0003\u000f\u0003\u000f\u0003\u0010",
    "\u0003\u0010\u0003\u0011\u0003\u0011\u0003\u0011\u0003\u0011\u0006\u0011",
    "\u00ad\n\u0011\r\u0011\u000e\u0011\u00ae\u0003\u0012\u0003\u0012\u0003",
    "\u0012\u0003\u0012\u0005\u0012\u00b5\n\u0012\u0003\u0013\u0003\u0013",
    "\u0003\u0013\u0003\u0013\u0003\u0014\u0003\u0014\u0003\u0014\u0003\u0014",
    "\u0003\u0015\u0003\u0015\u0005\u0015\u00c1\n\u0015\u0003\u0016\u0003",
    "\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0005\u0016\u00c9",
    "\n\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016",
    "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016",
    "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0007\u0016\u00da\n",
    "\u0016\f\u0016\u000e\u0016\u00dd\u000b\u0016\u0003\u0017\u0003\u0017",
    "\u0003\u0017\u0003\u0017\u0003\u0017\u0005\u0017\u00e4\n\u0017\u0003",
    "\u0018\u0003\u0018\u0003\u0018\u0003\u0018\u0005\u0018\u00ea\n\u0018",
    "\u0003\u0019\u0003\u0019\u0003\u0019\u0003\u0019\u0003\u0019\u0003\u0019",
    "\u0003\u0019\u0003\u0019\u0003\u0019\u0003\u0019\u0005\u0019\u00f6\n",
    "\u0019\u0003\u001a\u0003\u001a\u0007\u001a\u00fa\n\u001a\f\u001a\u000e",
    "\u001a\u00fd\u000b\u001a\u0003\u001a\u0003\u001a\u0005\u001a\u0101\n",
    "\u001a\u0003\u001a\u0003\u001a\u0007\u001a\u0105\n\u001a\f\u001a\u000e",
    "\u001a\u0108\u000b\u001a\u0003\u001a\u0003\u001a\u0005\u001a\u010c\n",
    "\u001a\u0007\u001a\u010e\n\u001a\f\u001a\u000e\u001a\u0111\u000b\u001a",
    "\u0003\u001a\u0005\u001a\u0114\n\u001a\u0003\u001a\u0005\u001a\u0117",
    "\n\u001a\u0005\u001a\u0119\n\u001a\u0003\u001a\u0003\u001a\u0003\u001b",
    "\u0003\u001b\u0007\u001b\u011f\n\u001b\f\u001b\u000e\u001b\u0122\u000b",
    "\u001b\u0003\u001b\u0003\u001b\u0007\u001b\u0126\n\u001b\f\u001b\u000e",
    "\u001b\u0129\u000b\u001b\u0003\u001b\u0003\u001b\u0003\u001b\u0002\u0003",
    "*\u001c\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018",
    "\u001a\u001c\u001e \"$&(*,.024\u0002\u000b\u0003\u0002./\u0004\u0002",
    "\u001e\u001e..\u0003\u0002\b\t\u0003\u0002\u0003\u0005\u0003\u0002\n",
    "\f\u0004\u0002\r\u0010\u0017\u0018\u0003\u0002#$\u0004\u0002\b\t%%\u0003",
    "\u0002)*\u0002\u0144\u00026\u0003\u0002\u0002\u0002\u0004;\u0003\u0002",
    "\u0002\u0002\u0006Y\u0003\u0002\u0002\u0002\bb\u0003\u0002\u0002\u0002",
    "\nh\u0003\u0002\u0002\u0002\fo\u0003\u0002\u0002\u0002\u000et\u0003",
    "\u0002\u0002\u0002\u0010y\u0003\u0002\u0002\u0002\u0012\u0084\u0003",
    "\u0002\u0002\u0002\u0014\u0088\u0003\u0002\u0002\u0002\u0016\u008f\u0003",
    "\u0002\u0002\u0002\u0018\u0091\u0003\u0002\u0002\u0002\u001a\u009a\u0003",
    "\u0002\u0002\u0002\u001c\u00a3\u0003\u0002\u0002\u0002\u001e\u00a6\u0003",
    "\u0002\u0002\u0002 \u00a8\u0003\u0002\u0002\u0002\"\u00b0\u0003\u0002",
    "\u0002\u0002$\u00b6\u0003\u0002\u0002\u0002&\u00ba\u0003\u0002\u0002",
    "\u0002(\u00c0\u0003\u0002\u0002\u0002*\u00c8\u0003\u0002\u0002\u0002",
    ",\u00e3\u0003\u0002\u0002\u0002.\u00e5\u0003\u0002\u0002\u00020\u00f5",
    "\u0003\u0002\u0002\u00022\u00f7\u0003\u0002\u0002\u00024\u011c\u0003",
    "\u0002\u0002\u000267\t\u0002\u0002\u00027\u0003\u0003\u0002\u0002\u0002",
    "8:\u0005\u0002\u0002\u000298\u0003\u0002\u0002\u0002:=\u0003\u0002\u0002",
    "\u0002;9\u0003\u0002\u0002\u0002;<\u0003\u0002\u0002\u0002<?\u0003\u0002",
    "\u0002\u0002=;\u0003\u0002\u0002\u0002>@\u0005\u0006\u0004\u0002?>\u0003",
    "\u0002\u0002\u0002?@\u0003\u0002\u0002\u0002@D\u0003\u0002\u0002\u0002",
    "AC\u0005\u0002\u0002\u0002BA\u0003\u0002\u0002\u0002CF\u0003\u0002\u0002",
    "\u0002DB\u0003\u0002\u0002\u0002DE\u0003\u0002\u0002\u0002ET\u0003\u0002",
    "\u0002\u0002FD\u0003\u0002\u0002\u0002GK\u0005\b\u0005\u0002HK\u0005",
    "\u0014\u000b\u0002IK\u0005\u0018\r\u0002JG\u0003\u0002\u0002\u0002J",
    "H\u0003\u0002\u0002\u0002JI\u0003\u0002\u0002\u0002KO\u0003\u0002\u0002",
    "\u0002LN\u0005\u0002\u0002\u0002ML\u0003\u0002\u0002\u0002NQ\u0003\u0002",
    "\u0002\u0002OM\u0003\u0002\u0002\u0002OP\u0003\u0002\u0002\u0002PS\u0003",
    "\u0002\u0002\u0002QO\u0003\u0002\u0002\u0002RJ\u0003\u0002\u0002\u0002",
    "SV\u0003\u0002\u0002\u0002TR\u0003\u0002\u0002\u0002TU\u0003\u0002\u0002",
    "\u0002UW\u0003\u0002\u0002\u0002VT\u0003\u0002\u0002\u0002WX\u0007\u0002",
    "\u0002\u0003X\u0005\u0003\u0002\u0002\u0002Y]\u0007\u0004\u0002\u0002",
    "Z\\\u0005\"\u0012\u0002[Z\u0003\u0002\u0002\u0002\\_\u0003\u0002\u0002",
    "\u0002][\u0003\u0002\u0002\u0002]^\u0003\u0002\u0002\u0002^`\u0003\u0002",
    "\u0002\u0002_]\u0003\u0002\u0002\u0002`a\u0005\u0002\u0002\u0002a\u0007",
    "\u0003\u0002\u0002\u0002bc\u0007\u0003\u0002\u0002cd\u0005\n\u0006\u0002",
    "de\u0005\f\u0007\u0002ef\u0007\u0006\u0002\u0002fg\u0005*\u0016\u0002",
    "g\t\u0003\u0002\u0002\u0002hj\u0007*\u0002\u0002ik\u0005\u0010\t\u0002",
    "ji\u0003\u0002\u0002\u0002jk\u0003\u0002\u0002\u0002k\u000b\u0003\u0002",
    "\u0002\u0002ln\u0005\u000e\b\u0002ml\u0003\u0002\u0002\u0002nq\u0003",
    "\u0002\u0002\u0002om\u0003\u0002\u0002\u0002op\u0003\u0002\u0002\u0002",
    "p\r\u0003\u0002\u0002\u0002qo\u0003\u0002\u0002\u0002ru\u0005\"\u0012",
    "\u0002su\u0007*\u0002\u0002tr\u0003\u0002\u0002\u0002ts\u0003\u0002",
    "\u0002\u0002uw\u0003\u0002\u0002\u0002vx\u0005\u0010\t\u0002wv\u0003",
    "\u0002\u0002\u0002wx\u0003\u0002\u0002\u0002x\u000f\u0003\u0002\u0002",
    "\u0002yz\u0007\u0017\u0002\u0002z{\u0005\u0012\n\u0002{\u007f\u0007",
    "\u0011\u0002\u0002|~\u0005\u0012\n\u0002}|\u0003\u0002\u0002\u0002~",
    "\u0081\u0003\u0002\u0002\u0002\u007f}\u0003\u0002\u0002\u0002\u007f",
    "\u0080\u0003\u0002\u0002\u0002\u0080\u0082\u0003\u0002\u0002\u0002\u0081",
    "\u007f\u0003\u0002\u0002\u0002\u0082\u0083\u0007\u0018\u0002\u0002\u0083",
    "\u0011\u0003\u0002\u0002\u0002\u0084\u0086\u0007*\u0002\u0002\u0085",
    "\u0087\u0005\u0010\t\u0002\u0086\u0085\u0003\u0002\u0002\u0002\u0086",
    "\u0087\u0003\u0002\u0002\u0002\u0087\u0013\u0003\u0002\u0002\u0002\u0088",
    "\u0089\u0007\u0005\u0002\u0002\u0089\u008a\u0007*\u0002\u0002\u008a",
    "\u008b\u0007\u0007\u0002\u0002\u008b\u008c\u00054\u001b\u0002\u008c",
    "\u0015\u0003\u0002\u0002\u0002\u008d\u0090\u0005\u0002\u0002\u0002\u008e",
    "\u0090\u0007\u0011\u0002\u0002\u008f\u008d\u0003\u0002\u0002\u0002\u008f",
    "\u008e\u0003\u0002\u0002\u0002\u0090\u0017\u0003\u0002\u0002\u0002\u0091",
    "\u0097\u0005(\u0015\u0002\u0092\u0093\u0005\u0016\f\u0002\u0093\u0094",
    "\u0005(\u0015\u0002\u0094\u0096\u0003\u0002\u0002\u0002\u0095\u0092",
    "\u0003\u0002\u0002\u0002\u0096\u0099\u0003\u0002\u0002\u0002\u0097\u0095",
    "\u0003\u0002\u0002\u0002\u0097\u0098\u0003\u0002\u0002\u0002\u0098\u0019",
    "\u0003\u0002\u0002\u0002\u0099\u0097\u0003\u0002\u0002\u0002\u009a\u009e",
    "\u0007\u001e\u0002\u0002\u009b\u009d\n\u0003\u0002\u0002\u009c\u009b",
    "\u0003\u0002\u0002\u0002\u009d\u00a0\u0003\u0002\u0002\u0002\u009e\u009c",
    "\u0003\u0002\u0002\u0002\u009e\u009f\u0003\u0002\u0002\u0002\u009f\u00a1",
    "\u0003\u0002\u0002\u0002\u00a0\u009e\u0003\u0002\u0002\u0002\u00a1\u00a2",
    "\u0007\u001e\u0002\u0002\u00a2\u001b\u0003\u0002\u0002\u0002\u00a3\u00a4",
    "\t\u0004\u0002\u0002\u00a4\u00a5\u0007*\u0002\u0002\u00a5\u001d\u0003",
    "\u0002\u0002\u0002\u00a6\u00a7\t\u0005\u0002\u0002\u00a7\u001f\u0003",
    "\u0002\u0002\u0002\u00a8\u00ac\u0007*\u0002\u0002\u00a9\u00ad\u0005",
    "\"\u0012\u0002\u00aa\u00ad\u0005$\u0013\u0002\u00ab\u00ad\u0005*\u0016",
    "\u0002\u00ac\u00a9\u0003\u0002\u0002\u0002\u00ac\u00aa\u0003\u0002\u0002",
    "\u0002\u00ac\u00ab\u0003\u0002\u0002\u0002\u00ad\u00ae\u0003\u0002\u0002",
    "\u0002\u00ae\u00ac\u0003\u0002\u0002\u0002\u00ae\u00af\u0003\u0002\u0002",
    "\u0002\u00af!\u0003\u0002\u0002\u0002\u00b0\u00b1\u0007*\u0002\u0002",
    "\u00b1\u00b4\u0007\u0012\u0002\u0002\u00b2\u00b5\u0005$\u0013\u0002",
    "\u00b3\u00b5\u0005*\u0016\u0002\u00b4\u00b2\u0003\u0002\u0002\u0002",
    "\u00b4\u00b3\u0003\u0002\u0002\u0002\u00b5#\u0003\u0002\u0002\u0002",
    "\u00b6\u00b7\u0007*\u0002\u0002\u00b7\u00b8\u0007\u0007\u0002\u0002",
    "\u00b8\u00b9\u0005*\u0016\u0002\u00b9%\u0003\u0002\u0002\u0002\u00ba",
    "\u00bb\u0007*\u0002\u0002\u00bb\u00bc\u0007\u0007\u0002\u0002\u00bc",
    "\u00bd\u0005(\u0015\u0002\u00bd\'\u0003\u0002\u0002\u0002\u00be\u00c1",
    "\u0005 \u0011\u0002\u00bf\u00c1\u0005*\u0016\u0002\u00c0\u00be\u0003",
    "\u0002\u0002\u0002\u00c0\u00bf\u0003\u0002\u0002\u0002\u00c1)\u0003",
    "\u0002\u0002\u0002\u00c2\u00c3\b\u0016\u0001\u0002\u00c3\u00c4\u0007",
    "\u001b\u0002\u0002\u00c4\u00c5\u0005*\u0016\u0002\u00c5\u00c6\u0007",
    "\u001c\u0002\u0002\u00c6\u00c9\u0003\u0002\u0002\u0002\u00c7\u00c9\u0005",
    ",\u0017\u0002\u00c8\u00c2\u0003\u0002\u0002\u0002\u00c8\u00c7\u0003",
    "\u0002\u0002\u0002\u00c9\u00db\u0003\u0002\u0002\u0002\u00ca\u00cb\f",
    "\t\u0002\u0002\u00cb\u00cc\t\u0006\u0002\u0002\u00cc\u00da\u0005*\u0016",
    "\n\u00cd\u00ce\f\b\u0002\u0002\u00ce\u00cf\t\u0004\u0002\u0002\u00cf",
    "\u00da\u0005*\u0016\t\u00d0\u00d1\f\u0007\u0002\u0002\u00d1\u00d2\t",
    "\u0007\u0002\u0002\u00d2\u00da\u0005*\u0016\b\u00d3\u00d4\f\u0006\u0002",
    "\u0002\u00d4\u00d5\u0007&\u0002\u0002\u00d5\u00da\u0005*\u0016\u0007",
    "\u00d6\u00d7\f\u0005\u0002\u0002\u00d7\u00d8\t\b\u0002\u0002\u00d8\u00da",
    "\u0005*\u0016\u0006\u00d9\u00ca\u0003\u0002\u0002\u0002\u00d9\u00cd",
    "\u0003\u0002\u0002\u0002\u00d9\u00d0\u0003\u0002\u0002\u0002\u00d9\u00d3",
    "\u0003\u0002\u0002\u0002\u00d9\u00d6\u0003\u0002\u0002\u0002\u00da\u00dd",
    "\u0003\u0002\u0002\u0002\u00db\u00d9\u0003\u0002\u0002\u0002\u00db\u00dc",
    "\u0003\u0002\u0002\u0002\u00dc+\u0003\u0002\u0002\u0002\u00dd\u00db",
    "\u0003\u0002\u0002\u0002\u00de\u00e4\u00050\u0019\u0002\u00df\u00e4",
    "\u0005\u001a\u000e\u0002\u00e0\u00e4\u0005.\u0018\u0002\u00e1\u00e4",
    "\u00052\u001a\u0002\u00e2\u00e4\u00054\u001b\u0002\u00e3\u00de\u0003",
    "\u0002\u0002\u0002\u00e3\u00df\u0003\u0002\u0002\u0002\u00e3\u00e0\u0003",
    "\u0002\u0002\u0002\u00e3\u00e1\u0003\u0002\u0002\u0002\u00e3\u00e2\u0003",
    "\u0002\u0002\u0002\u00e4-\u0003\u0002\u0002\u0002\u00e5\u00e9\t\t\u0002",
    "\u0002\u00e6\u00ea\u00054\u001b\u0002\u00e7\u00ea\u00050\u0019\u0002",
    "\u00e8\u00ea\u0007*\u0002\u0002\u00e9\u00e6\u0003\u0002\u0002\u0002",
    "\u00e9\u00e7\u0003\u0002\u0002\u0002\u00e9\u00e8\u0003\u0002\u0002\u0002",
    "\u00ea/\u0003\u0002\u0002\u0002\u00eb\u00f6\u0007\'\u0002\u0002\u00ec",
    "\u00f6\u0007(\u0002\u0002\u00ed\u00f6\u0007)\u0002\u0002\u00ee\u00f6",
    "\u00071\u0002\u0002\u00ef\u00f6\u0007*\u0002\u0002\u00f0\u00f1\u0007",
    ")\u0002\u0002\u00f1\u00f6\u00070\u0002\u0002\u00f2\u00f3\t\n\u0002\u0002",
    "\u00f3\u00f4\u0007\u0016\u0002\u0002\u00f4\u00f6\t\n\u0002\u0002\u00f5",
    "\u00eb\u0003\u0002\u0002\u0002\u00f5\u00ec\u0003\u0002\u0002\u0002\u00f5",
    "\u00ed\u0003\u0002\u0002\u0002\u00f5\u00ee\u0003\u0002\u0002\u0002\u00f5",
    "\u00ef\u0003\u0002\u0002\u0002\u00f5\u00f0\u0003\u0002\u0002\u0002\u00f5",
    "\u00f2\u0003\u0002\u0002\u0002\u00f61\u0003\u0002\u0002\u0002\u00f7",
    "\u0118\u0007\u0019\u0002\u0002\u00f8\u00fa\u0005\u0002\u0002\u0002\u00f9",
    "\u00f8\u0003\u0002\u0002\u0002\u00fa\u00fd\u0003\u0002\u0002\u0002\u00fb",
    "\u00f9\u0003\u0002\u0002\u0002\u00fb\u00fc\u0003\u0002\u0002\u0002\u00fc",
    "\u0100\u0003\u0002\u0002\u0002\u00fd\u00fb\u0003\u0002\u0002\u0002\u00fe",
    "\u0101\u0005&\u0014\u0002\u00ff\u0101\u0005(\u0015\u0002\u0100\u00fe",
    "\u0003\u0002\u0002\u0002\u0100\u00ff\u0003\u0002\u0002\u0002\u0101\u010f",
    "\u0003\u0002\u0002\u0002\u0102\u0106\u0007\u0013\u0002\u0002\u0103\u0105",
    "\u0005\u0002\u0002\u0002\u0104\u0103\u0003\u0002\u0002\u0002\u0105\u0108",
    "\u0003\u0002\u0002\u0002\u0106\u0104\u0003\u0002\u0002\u0002\u0106\u0107",
    "\u0003\u0002\u0002\u0002\u0107\u010b\u0003\u0002\u0002\u0002\u0108\u0106",
    "\u0003\u0002\u0002\u0002\u0109\u010c\u0005&\u0014\u0002\u010a\u010c",
    "\u0005(\u0015\u0002\u010b\u0109\u0003\u0002\u0002\u0002\u010b\u010a",
    "\u0003\u0002\u0002\u0002\u010c\u010e\u0003\u0002\u0002\u0002\u010d\u0102",
    "\u0003\u0002\u0002\u0002\u010e\u0111\u0003\u0002\u0002\u0002\u010f\u010d",
    "\u0003\u0002\u0002\u0002\u010f\u0110\u0003\u0002\u0002\u0002\u0110\u0113",
    "\u0003\u0002\u0002\u0002\u0111\u010f\u0003\u0002\u0002\u0002\u0112\u0114",
    "\u0007\u0013\u0002\u0002\u0113\u0112\u0003\u0002\u0002\u0002\u0113\u0114",
    "\u0003\u0002\u0002\u0002\u0114\u0116\u0003\u0002\u0002\u0002\u0115\u0117",
    "\u0005\u0002\u0002\u0002\u0116\u0115\u0003\u0002\u0002\u0002\u0116\u0117",
    "\u0003\u0002\u0002\u0002\u0117\u0119\u0003\u0002\u0002\u0002\u0118\u00fb",
    "\u0003\u0002\u0002\u0002\u0118\u0119\u0003\u0002\u0002\u0002\u0119\u011a",
    "\u0003\u0002\u0002\u0002\u011a\u011b\u0007\u001a\u0002\u0002\u011b3",
    "\u0003\u0002\u0002\u0002\u011c\u0120\u0007\u001b\u0002\u0002\u011d\u011f",
    "\u0005\u0002\u0002\u0002\u011e\u011d\u0003\u0002\u0002\u0002\u011f\u0122",
    "\u0003\u0002\u0002\u0002\u0120\u011e\u0003\u0002\u0002\u0002\u0120\u0121",
    "\u0003\u0002\u0002\u0002\u0121\u0123\u0003\u0002\u0002\u0002\u0122\u0120",
    "\u0003\u0002\u0002\u0002\u0123\u0127\u0005\u0018\r\u0002\u0124\u0126",
    "\u0005\u0002\u0002\u0002\u0125\u0124\u0003\u0002\u0002\u0002\u0126\u0129",
    "\u0003\u0002\u0002\u0002\u0127\u0125\u0003\u0002\u0002\u0002\u0127\u0128",
    "\u0003\u0002\u0002\u0002\u0128\u012a\u0003\u0002\u0002\u0002\u0129\u0127",
    "\u0003\u0002\u0002\u0002\u012a\u012b\u0007\u001c\u0002\u0002\u012b5",
    "\u0003\u0002\u0002\u0002&;?DJOT]jotw\u007f\u0086\u008f\u0097\u009e\u00ac",
    "\u00ae\u00b4\u00c0\u00c8\u00d9\u00db\u00e3\u00e9\u00f5\u00fb\u0100\u0106",
    "\u010b\u010f\u0113\u0116\u0118\u0120\u0127"].join("");


const atn = new antlr4.atn.ATNDeserializer().deserialize(serializedATN);

const decisionsToDFA = atn.decisionToState.map( (ds, index) => new antlr4.dfa.DFA(ds, index) );

const sharedContextCache = new antlr4.PredictionContextCache();

export default class prqlParser extends antlr4.Parser {

    static grammarFileName = "prql.g4";
    static literalNames = [ null, "'func'", "'prql'", "'table'", "'->'", 
                            "'='", "'+'", "'-'", "'*'", "'/'", "'%'", "'=='", 
                            "'!='", "'<='", "'>='", "'|'", "':'", "','", 
                            "'.'", "'$'", "'..'", "'<'", "'>'", "'['", "']'", 
                            "'('", "')'", "'_'", "'`'", "'\"'", "'''", "'\"\"\"'", 
                            "'''''", "'and'", "'or'", "'!'", "'??'", "'null'" ];
    static symbolicNames = [ null, "FUNC", "PRQL", "TABLE", "ARROW", "ASSIGN", 
                             "PLUS", "MINUS", "STAR", "DIV", "MOD", "EQ", 
                             "NE", "LE", "GE", "BAR", "COLON", "COMMA", 
                             "DOT", "DOLLAR", "RANGE", "LANG", "RANG", "LBRACKET", 
                             "RBRACKET", "LPAREN", "RPAREN", "UNDERSCORE", 
                             "BACKTICK", "DOUBLE_QUOTE", "SINGLE_QUOTE", 
                             "TRIPLE_DOUBLE_QUOTE", "TRIPLE_SINGLE_QUOTE", 
                             "AND", "OR", "NOT", "COALESCE", "NULL_", "BOOLEAN", 
                             "NUMBER", "IDENT", "IDENT_START", "IDENT_NEXT", 
                             "WHITESPACE", "NEWLINE", "COMMENT", "INTERVAL_KIND", 
                             "STRING" ];
    static ruleNames = [ "nl", "query", "queryDef", "funcDef", "funcDefName", 
                         "funcDefParams", "funcDefParam", "typeDef", "typeTerm", 
                         "table", "pipe", "pipeline", "identBacktick", "signedIdent", 
                         "keyword", "funcCall", "namedArg", "assign", "assignCall", 
                         "exprCall", "expr", "term", "exprUnary", "literal", 
                         "list", "nestedPipeline" ];

    constructor(input) {
        super(input);
        this._interp = new antlr4.atn.ParserATNSimulator(this, atn, decisionsToDFA, sharedContextCache);
        this.ruleNames = prqlParser.ruleNames;
        this.literalNames = prqlParser.literalNames;
        this.symbolicNames = prqlParser.symbolicNames;
    }

    get atn() {
        return atn;
    }

    sempred(localctx, ruleIndex, predIndex) {
    	switch(ruleIndex) {
    	case 20:
    	    		return this.expr_sempred(localctx, predIndex);
        default:
            throw "No predicate with index:" + ruleIndex;
       }
    }

    expr_sempred(localctx, predIndex) {
    	switch(predIndex) {
    		case 0:
    			return this.precpred(this._ctx, 7);
    		case 1:
    			return this.precpred(this._ctx, 6);
    		case 2:
    			return this.precpred(this._ctx, 5);
    		case 3:
    			return this.precpred(this._ctx, 4);
    		case 4:
    			return this.precpred(this._ctx, 3);
    		default:
    			throw "No predicate with index:" + predIndex;
    	}
    };




	nl() {
	    let localctx = new NlContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 0, prqlParser.RULE_nl);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 52;
	        _la = this._input.LA(1);
	        if(!(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT)) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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



	query() {
	    let localctx = new QueryContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 2, prqlParser.RULE_query);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 57;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,0,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                this.state = 54;
	                this.nl(); 
	            }
	            this.state = 59;
	            this._errHandler.sync(this);
	            _alt = this._interp.adaptivePredict(this._input,0,this._ctx);
	        }

	        this.state = 61;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.PRQL) {
	            this.state = 60;
	            this.queryDef();
	        }

	        this.state = 66;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	            this.state = 63;
	            this.nl();
	            this.state = 68;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 82;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.FUNC) | (1 << prqlParser.TABLE) | (1 << prqlParser.PLUS) | (1 << prqlParser.MINUS) | (1 << prqlParser.LBRACKET) | (1 << prqlParser.LPAREN) | (1 << prqlParser.BACKTICK))) !== 0) || ((((_la - 35)) & ~0x1f) == 0 && ((1 << (_la - 35)) & ((1 << (prqlParser.NOT - 35)) | (1 << (prqlParser.NULL_ - 35)) | (1 << (prqlParser.BOOLEAN - 35)) | (1 << (prqlParser.NUMBER - 35)) | (1 << (prqlParser.IDENT - 35)) | (1 << (prqlParser.STRING - 35)))) !== 0)) {
	            this.state = 72;
	            this._errHandler.sync(this);
	            switch(this._input.LA(1)) {
	            case prqlParser.FUNC:
	                this.state = 69;
	                this.funcDef();
	                break;
	            case prqlParser.TABLE:
	                this.state = 70;
	                this.table();
	                break;
	            case prqlParser.PLUS:
	            case prqlParser.MINUS:
	            case prqlParser.LBRACKET:
	            case prqlParser.LPAREN:
	            case prqlParser.BACKTICK:
	            case prqlParser.NOT:
	            case prqlParser.NULL_:
	            case prqlParser.BOOLEAN:
	            case prqlParser.NUMBER:
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
	            while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
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
	    	if(re instanceof antlr4.error.RecognitionException) {
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



	queryDef() {
	    let localctx = new QueryDefContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 4, prqlParser.RULE_queryDef);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 87;
	        this.match(prqlParser.PRQL);
	        this.state = 91;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.IDENT) {
	            this.state = 88;
	            this.namedArg();
	            this.state = 93;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 94;
	        this.nl();
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	        if(_la===prqlParser.LANG) {
	            this.state = 103;
	            this.typeDef();
	        }

	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	        while(_la===prqlParser.IDENT) {
	            this.state = 106;
	            this.funcDefParam();
	            this.state = 111;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	        var la_ = this._interp.adaptivePredict(this._input,9,this._ctx);
	        switch(la_) {
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
	        if(_la===prqlParser.LANG) {
	            this.state = 116;
	            this.typeDef();
	        }

	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	        while(_la===prqlParser.IDENT) {
	            this.state = 122;
	            this.typeTerm();
	            this.state = 127;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 128;
	        this.match(prqlParser.RANG);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	        if(_la===prqlParser.LANG) {
	            this.state = 131;
	            this.typeDef();
	        }

	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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



	table() {
	    let localctx = new TableContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 18, prqlParser.RULE_table);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 134;
	        this.match(prqlParser.TABLE);
	        this.state = 135;
	        this.match(prqlParser.IDENT);
	        this.state = 136;
	        this.match(prqlParser.ASSIGN);
	        this.state = 137;
	        this.nestedPipeline();
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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



	pipe() {
	    let localctx = new PipeContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 20, prqlParser.RULE_pipe);
	    try {
	        this.state = 141;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case prqlParser.NEWLINE:
	        case prqlParser.COMMENT:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 139;
	            this.nl();
	            break;
	        case prqlParser.BAR:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 140;
	            this.match(prqlParser.BAR);
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	        this.state = 143;
	        this.exprCall();
	        this.state = 149;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,14,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                this.state = 144;
	                this.pipe();
	                this.state = 145;
	                this.exprCall(); 
	            }
	            this.state = 151;
	            this._errHandler.sync(this);
	            _alt = this._interp.adaptivePredict(this._input,14,this._ctx);
	        }

	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	    this.enterRule(localctx, 24, prqlParser.RULE_identBacktick);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 152;
	        this.match(prqlParser.BACKTICK);
	        this.state = 156;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.FUNC) | (1 << prqlParser.PRQL) | (1 << prqlParser.TABLE) | (1 << prqlParser.ARROW) | (1 << prqlParser.ASSIGN) | (1 << prqlParser.PLUS) | (1 << prqlParser.MINUS) | (1 << prqlParser.STAR) | (1 << prqlParser.DIV) | (1 << prqlParser.MOD) | (1 << prqlParser.EQ) | (1 << prqlParser.NE) | (1 << prqlParser.LE) | (1 << prqlParser.GE) | (1 << prqlParser.BAR) | (1 << prqlParser.COLON) | (1 << prqlParser.COMMA) | (1 << prqlParser.DOT) | (1 << prqlParser.DOLLAR) | (1 << prqlParser.RANGE) | (1 << prqlParser.LANG) | (1 << prqlParser.RANG) | (1 << prqlParser.LBRACKET) | (1 << prqlParser.RBRACKET) | (1 << prqlParser.LPAREN) | (1 << prqlParser.RPAREN) | (1 << prqlParser.UNDERSCORE) | (1 << prqlParser.DOUBLE_QUOTE) | (1 << prqlParser.SINGLE_QUOTE) | (1 << prqlParser.TRIPLE_DOUBLE_QUOTE))) !== 0) || ((((_la - 32)) & ~0x1f) == 0 && ((1 << (_la - 32)) & ((1 << (prqlParser.TRIPLE_SINGLE_QUOTE - 32)) | (1 << (prqlParser.AND - 32)) | (1 << (prqlParser.OR - 32)) | (1 << (prqlParser.NOT - 32)) | (1 << (prqlParser.COALESCE - 32)) | (1 << (prqlParser.NULL_ - 32)) | (1 << (prqlParser.BOOLEAN - 32)) | (1 << (prqlParser.NUMBER - 32)) | (1 << (prqlParser.IDENT - 32)) | (1 << (prqlParser.IDENT_START - 32)) | (1 << (prqlParser.IDENT_NEXT - 32)) | (1 << (prqlParser.WHITESPACE - 32)) | (1 << (prqlParser.COMMENT - 32)) | (1 << (prqlParser.INTERVAL_KIND - 32)) | (1 << (prqlParser.STRING - 32)))) !== 0)) {
	            this.state = 153;
	            _la = this._input.LA(1);
	            if(_la<=0 || _la===prqlParser.BACKTICK || _la===prqlParser.NEWLINE) {
	            this._errHandler.recoverInline(this);
	            }
	            else {
	            	this._errHandler.reportMatch(this);
	                this.consume();
	            }
	            this.state = 158;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 159;
	        this.match(prqlParser.BACKTICK);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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



	signedIdent() {
	    let localctx = new SignedIdentContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 26, prqlParser.RULE_signedIdent);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 161;
	        _la = this._input.LA(1);
	        if(!(_la===prqlParser.PLUS || _la===prqlParser.MINUS)) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 162;
	        this.match(prqlParser.IDENT);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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



	keyword() {
	    let localctx = new KeywordContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 28, prqlParser.RULE_keyword);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 164;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.FUNC) | (1 << prqlParser.PRQL) | (1 << prqlParser.TABLE))) !== 0))) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	    this.enterRule(localctx, 30, prqlParser.RULE_funcCall);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 166;
	        this.match(prqlParser.IDENT);
	        this.state = 170; 
	        this._errHandler.sync(this);
	        var _alt = 1;
	        do {
	        	switch (_alt) {
	        	case 1:
	        		this.state = 170;
	        		this._errHandler.sync(this);
	        		var la_ = this._interp.adaptivePredict(this._input,16,this._ctx);
	        		switch(la_) {
	        		case 1:
	        		    this.state = 167;
	        		    this.namedArg();
	        		    break;

	        		case 2:
	        		    this.state = 168;
	        		    this.assign();
	        		    break;

	        		case 3:
	        		    this.state = 169;
	        		    this.expr(0);
	        		    break;

	        		}
	        		break;
	        	default:
	        		throw new antlr4.error.NoViableAltException(this);
	        	}
	        	this.state = 172; 
	        	this._errHandler.sync(this);
	        	_alt = this._interp.adaptivePredict(this._input,17, this._ctx);
	        } while ( _alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER );
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	        this.state = 174;
	        this.match(prqlParser.IDENT);
	        this.state = 175;
	        this.match(prqlParser.COLON);
	        this.state = 178;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,18,this._ctx);
	        switch(la_) {
	        case 1:
	            this.state = 176;
	            this.assign();
	            break;

	        case 2:
	            this.state = 177;
	            this.expr(0);
	            break;

	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	        this.state = 180;
	        this.match(prqlParser.IDENT);
	        this.state = 181;
	        this.match(prqlParser.ASSIGN);
	        this.state = 182;
	        this.expr(0);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	        this.state = 184;
	        this.match(prqlParser.IDENT);
	        this.state = 185;
	        this.match(prqlParser.ASSIGN);
	        this.state = 186;
	        this.exprCall();
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	        this.state = 190;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,19,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 188;
	            this.funcCall();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 189;
	            this.expr(0);
	            break;

	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
		if(_p===undefined) {
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
	        this.state = 198;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,20,this._ctx);
	        switch(la_) {
	        case 1:
	            this.state = 193;
	            this.match(prqlParser.LPAREN);
	            this.state = 194;
	            this.expr(0);
	            this.state = 195;
	            this.match(prqlParser.RPAREN);
	            break;

	        case 2:
	            this.state = 197;
	            this.term();
	            break;

	        }
	        this._ctx.stop = this._input.LT(-1);
	        this.state = 217;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,22,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                if(this._parseListeners!==null) {
	                    this.triggerExitRuleEvent();
	                }
	                _prevctx = localctx;
	                this.state = 215;
	                this._errHandler.sync(this);
	                var la_ = this._interp.adaptivePredict(this._input,21,this._ctx);
	                switch(la_) {
	                case 1:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 200;
	                    if (!( this.precpred(this._ctx, 7))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 7)");
	                    }
	                    this.state = 201;
	                    _la = this._input.LA(1);
	                    if(!((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.STAR) | (1 << prqlParser.DIV) | (1 << prqlParser.MOD))) !== 0))) {
	                    this._errHandler.recoverInline(this);
	                    }
	                    else {
	                    	this._errHandler.reportMatch(this);
	                        this.consume();
	                    }
	                    this.state = 202;
	                    this.expr(8);
	                    break;

	                case 2:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 203;
	                    if (!( this.precpred(this._ctx, 6))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 6)");
	                    }
	                    this.state = 204;
	                    _la = this._input.LA(1);
	                    if(!(_la===prqlParser.PLUS || _la===prqlParser.MINUS)) {
	                    this._errHandler.recoverInline(this);
	                    }
	                    else {
	                    	this._errHandler.reportMatch(this);
	                        this.consume();
	                    }
	                    this.state = 205;
	                    this.expr(7);
	                    break;

	                case 3:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 206;
	                    if (!( this.precpred(this._ctx, 5))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 5)");
	                    }
	                    this.state = 207;
	                    _la = this._input.LA(1);
	                    if(!((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.EQ) | (1 << prqlParser.NE) | (1 << prqlParser.LE) | (1 << prqlParser.GE) | (1 << prqlParser.LANG) | (1 << prqlParser.RANG))) !== 0))) {
	                    this._errHandler.recoverInline(this);
	                    }
	                    else {
	                    	this._errHandler.reportMatch(this);
	                        this.consume();
	                    }
	                    this.state = 208;
	                    this.expr(6);
	                    break;

	                case 4:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 209;
	                    if (!( this.precpred(this._ctx, 4))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 4)");
	                    }
	                    this.state = 210;
	                    this.match(prqlParser.COALESCE);
	                    this.state = 211;
	                    this.expr(5);
	                    break;

	                case 5:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 212;
	                    if (!( this.precpred(this._ctx, 3))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 3)");
	                    }
	                    this.state = 213;
	                    _la = this._input.LA(1);
	                    if(!(_la===prqlParser.AND || _la===prqlParser.OR)) {
	                    this._errHandler.recoverInline(this);
	                    }
	                    else {
	                    	this._errHandler.reportMatch(this);
	                        this.consume();
	                    }
	                    this.state = 214;
	                    this.expr(4);
	                    break;

	                } 
	            }
	            this.state = 219;
	            this._errHandler.sync(this);
	            _alt = this._interp.adaptivePredict(this._input,22,this._ctx);
	        }

	    } catch( error) {
	        if(error instanceof antlr4.error.RecognitionException) {
		        localctx.exception = error;
		        this._errHandler.reportError(this, error);
		        this._errHandler.recover(this, error);
		    } else {
		    	throw error;
		    }
	    } finally {
	        this.unrollRecursionContexts(_parentctx)
	    }
	    return localctx;
	}



	term() {
	    let localctx = new TermContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 42, prqlParser.RULE_term);
	    try {
	        this.state = 225;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case prqlParser.NULL_:
	        case prqlParser.BOOLEAN:
	        case prqlParser.NUMBER:
	        case prqlParser.IDENT:
	        case prqlParser.STRING:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 220;
	            this.literal();
	            break;
	        case prqlParser.BACKTICK:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 221;
	            this.identBacktick();
	            break;
	        case prqlParser.PLUS:
	        case prqlParser.MINUS:
	        case prqlParser.NOT:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 222;
	            this.exprUnary();
	            break;
	        case prqlParser.LBRACKET:
	            this.enterOuterAlt(localctx, 4);
	            this.state = 223;
	            this.list();
	            break;
	        case prqlParser.LPAREN:
	            this.enterOuterAlt(localctx, 5);
	            this.state = 224;
	            this.nestedPipeline();
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	        this.state = 227;
	        _la = this._input.LA(1);
	        if(!(((((_la - 6)) & ~0x1f) == 0 && ((1 << (_la - 6)) & ((1 << (prqlParser.PLUS - 6)) | (1 << (prqlParser.MINUS - 6)) | (1 << (prqlParser.NOT - 6)))) !== 0))) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 231;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,24,this._ctx);
	        switch(la_) {
	        case 1:
	            this.state = 228;
	            this.nestedPipeline();
	            break;

	        case 2:
	            this.state = 229;
	            this.literal();
	            break;

	        case 3:
	            this.state = 230;
	            this.match(prqlParser.IDENT);
	            break;

	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	        this.state = 243;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,25,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 233;
	            this.match(prqlParser.NULL_);
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 234;
	            this.match(prqlParser.BOOLEAN);
	            break;

	        case 3:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 235;
	            this.match(prqlParser.NUMBER);
	            break;

	        case 4:
	            this.enterOuterAlt(localctx, 4);
	            this.state = 236;
	            this.match(prqlParser.STRING);
	            break;

	        case 5:
	            this.enterOuterAlt(localctx, 5);
	            this.state = 237;
	            this.match(prqlParser.IDENT);
	            break;

	        case 6:
	            this.enterOuterAlt(localctx, 6);
	            this.state = 238;
	            this.match(prqlParser.NUMBER);
	            this.state = 239;
	            this.match(prqlParser.INTERVAL_KIND);
	            break;

	        case 7:
	            this.enterOuterAlt(localctx, 7);
	            this.state = 240;
	            _la = this._input.LA(1);
	            if(!(_la===prqlParser.NUMBER || _la===prqlParser.IDENT)) {
	            this._errHandler.recoverInline(this);
	            }
	            else {
	            	this._errHandler.reportMatch(this);
	                this.consume();
	            }
	            this.state = 241;
	            this.match(prqlParser.RANGE);
	            this.state = 242;
	            _la = this._input.LA(1);
	            if(!(_la===prqlParser.NUMBER || _la===prqlParser.IDENT)) {
	            this._errHandler.recoverInline(this);
	            }
	            else {
	            	this._errHandler.reportMatch(this);
	                this.consume();
	            }
	            break;

	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	        this.state = 245;
	        this.match(prqlParser.LBRACKET);
	        this.state = 278;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.PLUS) | (1 << prqlParser.MINUS) | (1 << prqlParser.LBRACKET) | (1 << prqlParser.LPAREN) | (1 << prqlParser.BACKTICK))) !== 0) || ((((_la - 35)) & ~0x1f) == 0 && ((1 << (_la - 35)) & ((1 << (prqlParser.NOT - 35)) | (1 << (prqlParser.NULL_ - 35)) | (1 << (prqlParser.BOOLEAN - 35)) | (1 << (prqlParser.NUMBER - 35)) | (1 << (prqlParser.IDENT - 35)) | (1 << (prqlParser.NEWLINE - 35)) | (1 << (prqlParser.COMMENT - 35)) | (1 << (prqlParser.STRING - 35)))) !== 0)) {
	            this.state = 249;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	                this.state = 246;
	                this.nl();
	                this.state = 251;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	            }
	            this.state = 254;
	            this._errHandler.sync(this);
	            var la_ = this._interp.adaptivePredict(this._input,27,this._ctx);
	            switch(la_) {
	            case 1:
	                this.state = 252;
	                this.assignCall();
	                break;

	            case 2:
	                this.state = 253;
	                this.exprCall();
	                break;

	            }
	            this.state = 269;
	            this._errHandler.sync(this);
	            var _alt = this._interp.adaptivePredict(this._input,30,this._ctx)
	            while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	                if(_alt===1) {
	                    this.state = 256;
	                    this.match(prqlParser.COMMA);
	                    this.state = 260;
	                    this._errHandler.sync(this);
	                    _la = this._input.LA(1);
	                    while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	                        this.state = 257;
	                        this.nl();
	                        this.state = 262;
	                        this._errHandler.sync(this);
	                        _la = this._input.LA(1);
	                    }
	                    this.state = 265;
	                    this._errHandler.sync(this);
	                    var la_ = this._interp.adaptivePredict(this._input,29,this._ctx);
	                    switch(la_) {
	                    case 1:
	                        this.state = 263;
	                        this.assignCall();
	                        break;

	                    case 2:
	                        this.state = 264;
	                        this.exprCall();
	                        break;

	                    } 
	                }
	                this.state = 271;
	                this._errHandler.sync(this);
	                _alt = this._interp.adaptivePredict(this._input,30,this._ctx);
	            }

	            this.state = 273;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===prqlParser.COMMA) {
	                this.state = 272;
	                this.match(prqlParser.COMMA);
	            }

	            this.state = 276;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	                this.state = 275;
	                this.nl();
	            }

	        }

	        this.state = 280;
	        this.match(prqlParser.RBRACKET);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
	        this.state = 282;
	        this.match(prqlParser.LPAREN);
	        this.state = 286;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	            this.state = 283;
	            this.nl();
	            this.state = 288;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 289;
	        this.pipeline();
	        this.state = 293;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	            this.state = 290;
	            this.nl();
	            this.state = 295;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 296;
	        this.match(prqlParser.RPAREN);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
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
prqlParser.TABLE = 3;
prqlParser.ARROW = 4;
prqlParser.ASSIGN = 5;
prqlParser.PLUS = 6;
prqlParser.MINUS = 7;
prqlParser.STAR = 8;
prqlParser.DIV = 9;
prqlParser.MOD = 10;
prqlParser.EQ = 11;
prqlParser.NE = 12;
prqlParser.LE = 13;
prqlParser.GE = 14;
prqlParser.BAR = 15;
prqlParser.COLON = 16;
prqlParser.COMMA = 17;
prqlParser.DOT = 18;
prqlParser.DOLLAR = 19;
prqlParser.RANGE = 20;
prqlParser.LANG = 21;
prqlParser.RANG = 22;
prqlParser.LBRACKET = 23;
prqlParser.RBRACKET = 24;
prqlParser.LPAREN = 25;
prqlParser.RPAREN = 26;
prqlParser.UNDERSCORE = 27;
prqlParser.BACKTICK = 28;
prqlParser.DOUBLE_QUOTE = 29;
prqlParser.SINGLE_QUOTE = 30;
prqlParser.TRIPLE_DOUBLE_QUOTE = 31;
prqlParser.TRIPLE_SINGLE_QUOTE = 32;
prqlParser.AND = 33;
prqlParser.OR = 34;
prqlParser.NOT = 35;
prqlParser.COALESCE = 36;
prqlParser.NULL_ = 37;
prqlParser.BOOLEAN = 38;
prqlParser.NUMBER = 39;
prqlParser.IDENT = 40;
prqlParser.IDENT_START = 41;
prqlParser.IDENT_NEXT = 42;
prqlParser.WHITESPACE = 43;
prqlParser.NEWLINE = 44;
prqlParser.COMMENT = 45;
prqlParser.INTERVAL_KIND = 46;
prqlParser.STRING = 47;

prqlParser.RULE_nl = 0;
prqlParser.RULE_query = 1;
prqlParser.RULE_queryDef = 2;
prqlParser.RULE_funcDef = 3;
prqlParser.RULE_funcDefName = 4;
prqlParser.RULE_funcDefParams = 5;
prqlParser.RULE_funcDefParam = 6;
prqlParser.RULE_typeDef = 7;
prqlParser.RULE_typeTerm = 8;
prqlParser.RULE_table = 9;
prqlParser.RULE_pipe = 10;
prqlParser.RULE_pipeline = 11;
prqlParser.RULE_identBacktick = 12;
prqlParser.RULE_signedIdent = 13;
prqlParser.RULE_keyword = 14;
prqlParser.RULE_funcCall = 15;
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
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_nl;
    }

	NEWLINE() {
	    return this.getToken(prqlParser.NEWLINE, 0);
	};

	COMMENT() {
	    return this.getToken(prqlParser.COMMENT, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterNl(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitNl(this);
		}
	}


}



class QueryContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_query;
    }

	EOF() {
	    return this.getToken(prqlParser.EOF, 0);
	};

	nl = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(NlContext);
	    } else {
	        return this.getTypedRuleContext(NlContext,i);
	    }
	};

	queryDef() {
	    return this.getTypedRuleContext(QueryDefContext,0);
	};

	funcDef = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(FuncDefContext);
	    } else {
	        return this.getTypedRuleContext(FuncDefContext,i);
	    }
	};

	table = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(TableContext);
	    } else {
	        return this.getTypedRuleContext(TableContext,i);
	    }
	};

	pipeline = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(PipelineContext);
	    } else {
	        return this.getTypedRuleContext(PipelineContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterQuery(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitQuery(this);
		}
	}


}



class QueryDefContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_queryDef;
    }

	PRQL() {
	    return this.getToken(prqlParser.PRQL, 0);
	};

	nl() {
	    return this.getTypedRuleContext(NlContext,0);
	};

	namedArg = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(NamedArgContext);
	    } else {
	        return this.getTypedRuleContext(NamedArgContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterQueryDef(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitQueryDef(this);
		}
	}


}



class FuncDefContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_funcDef;
    }

	FUNC() {
	    return this.getToken(prqlParser.FUNC, 0);
	};

	funcDefName() {
	    return this.getTypedRuleContext(FuncDefNameContext,0);
	};

	funcDefParams() {
	    return this.getTypedRuleContext(FuncDefParamsContext,0);
	};

	ARROW() {
	    return this.getToken(prqlParser.ARROW, 0);
	};

	expr() {
	    return this.getTypedRuleContext(ExprContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterFuncDef(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitFuncDef(this);
		}
	}


}



class FuncDefNameContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_funcDefName;
    }

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	typeDef() {
	    return this.getTypedRuleContext(TypeDefContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterFuncDefName(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitFuncDefName(this);
		}
	}


}



class FuncDefParamsContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_funcDefParams;
    }

	funcDefParam = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(FuncDefParamContext);
	    } else {
	        return this.getTypedRuleContext(FuncDefParamContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterFuncDefParams(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitFuncDefParams(this);
		}
	}


}



class FuncDefParamContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_funcDefParam;
    }

	namedArg() {
	    return this.getTypedRuleContext(NamedArgContext,0);
	};

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	typeDef() {
	    return this.getTypedRuleContext(TypeDefContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterFuncDefParam(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitFuncDefParam(this);
		}
	}


}



class TypeDefContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_typeDef;
    }

	LANG() {
	    return this.getToken(prqlParser.LANG, 0);
	};

	typeTerm = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(TypeTermContext);
	    } else {
	        return this.getTypedRuleContext(TypeTermContext,i);
	    }
	};

	BAR() {
	    return this.getToken(prqlParser.BAR, 0);
	};

	RANG() {
	    return this.getToken(prqlParser.RANG, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterTypeDef(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitTypeDef(this);
		}
	}


}



class TypeTermContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_typeTerm;
    }

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	typeDef() {
	    return this.getTypedRuleContext(TypeDefContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterTypeTerm(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitTypeTerm(this);
		}
	}


}



class TableContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_table;
    }

	TABLE() {
	    return this.getToken(prqlParser.TABLE, 0);
	};

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	ASSIGN() {
	    return this.getToken(prqlParser.ASSIGN, 0);
	};

	nestedPipeline() {
	    return this.getTypedRuleContext(NestedPipelineContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterTable(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitTable(this);
		}
	}


}



class PipeContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_pipe;
    }

	nl() {
	    return this.getTypedRuleContext(NlContext,0);
	};

	BAR() {
	    return this.getToken(prqlParser.BAR, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterPipe(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitPipe(this);
		}
	}


}



class PipelineContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_pipeline;
    }

	exprCall = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(ExprCallContext);
	    } else {
	        return this.getTypedRuleContext(ExprCallContext,i);
	    }
	};

	pipe = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(PipeContext);
	    } else {
	        return this.getTypedRuleContext(PipeContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterPipeline(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitPipeline(this);
		}
	}


}



class IdentBacktickContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_identBacktick;
    }

	BACKTICK = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(prqlParser.BACKTICK);
	    } else {
	        return this.getToken(prqlParser.BACKTICK, i);
	    }
	};


	NEWLINE = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(prqlParser.NEWLINE);
	    } else {
	        return this.getToken(prqlParser.NEWLINE, i);
	    }
	};


	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterIdentBacktick(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitIdentBacktick(this);
		}
	}


}



class SignedIdentContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_signedIdent;
    }

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	PLUS() {
	    return this.getToken(prqlParser.PLUS, 0);
	};

	MINUS() {
	    return this.getToken(prqlParser.MINUS, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterSignedIdent(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitSignedIdent(this);
		}
	}


}



class KeywordContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_keyword;
    }

	PRQL() {
	    return this.getToken(prqlParser.PRQL, 0);
	};

	TABLE() {
	    return this.getToken(prqlParser.TABLE, 0);
	};

	FUNC() {
	    return this.getToken(prqlParser.FUNC, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterKeyword(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitKeyword(this);
		}
	}


}



class FuncCallContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_funcCall;
    }

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	namedArg = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(NamedArgContext);
	    } else {
	        return this.getTypedRuleContext(NamedArgContext,i);
	    }
	};

	assign = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(AssignContext);
	    } else {
	        return this.getTypedRuleContext(AssignContext,i);
	    }
	};

	expr = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(ExprContext);
	    } else {
	        return this.getTypedRuleContext(ExprContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterFuncCall(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitFuncCall(this);
		}
	}


}



class NamedArgContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_namedArg;
    }

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	COLON() {
	    return this.getToken(prqlParser.COLON, 0);
	};

	assign() {
	    return this.getTypedRuleContext(AssignContext,0);
	};

	expr() {
	    return this.getTypedRuleContext(ExprContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterNamedArg(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitNamedArg(this);
		}
	}


}



class AssignContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_assign;
    }

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	ASSIGN() {
	    return this.getToken(prqlParser.ASSIGN, 0);
	};

	expr() {
	    return this.getTypedRuleContext(ExprContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterAssign(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitAssign(this);
		}
	}


}



class AssignCallContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_assignCall;
    }

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	ASSIGN() {
	    return this.getToken(prqlParser.ASSIGN, 0);
	};

	exprCall() {
	    return this.getTypedRuleContext(ExprCallContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterAssignCall(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitAssignCall(this);
		}
	}


}



class ExprCallContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_exprCall;
    }

	funcCall() {
	    return this.getTypedRuleContext(FuncCallContext,0);
	};

	expr() {
	    return this.getTypedRuleContext(ExprContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterExprCall(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitExprCall(this);
		}
	}


}



class ExprContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_expr;
    }

	LPAREN() {
	    return this.getToken(prqlParser.LPAREN, 0);
	};

	expr = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(ExprContext);
	    } else {
	        return this.getTypedRuleContext(ExprContext,i);
	    }
	};

	RPAREN() {
	    return this.getToken(prqlParser.RPAREN, 0);
	};

	term() {
	    return this.getTypedRuleContext(TermContext,0);
	};

	STAR() {
	    return this.getToken(prqlParser.STAR, 0);
	};

	DIV() {
	    return this.getToken(prqlParser.DIV, 0);
	};

	MOD() {
	    return this.getToken(prqlParser.MOD, 0);
	};

	MINUS() {
	    return this.getToken(prqlParser.MINUS, 0);
	};

	PLUS() {
	    return this.getToken(prqlParser.PLUS, 0);
	};

	EQ() {
	    return this.getToken(prqlParser.EQ, 0);
	};

	NE() {
	    return this.getToken(prqlParser.NE, 0);
	};

	GE() {
	    return this.getToken(prqlParser.GE, 0);
	};

	LE() {
	    return this.getToken(prqlParser.LE, 0);
	};

	LANG() {
	    return this.getToken(prqlParser.LANG, 0);
	};

	RANG() {
	    return this.getToken(prqlParser.RANG, 0);
	};

	COALESCE() {
	    return this.getToken(prqlParser.COALESCE, 0);
	};

	AND() {
	    return this.getToken(prqlParser.AND, 0);
	};

	OR() {
	    return this.getToken(prqlParser.OR, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterExpr(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitExpr(this);
		}
	}


}



class TermContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_term;
    }

	literal() {
	    return this.getTypedRuleContext(LiteralContext,0);
	};

	identBacktick() {
	    return this.getTypedRuleContext(IdentBacktickContext,0);
	};

	exprUnary() {
	    return this.getTypedRuleContext(ExprUnaryContext,0);
	};

	list() {
	    return this.getTypedRuleContext(ListContext,0);
	};

	nestedPipeline() {
	    return this.getTypedRuleContext(NestedPipelineContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterTerm(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitTerm(this);
		}
	}


}



class ExprUnaryContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_exprUnary;
    }

	MINUS() {
	    return this.getToken(prqlParser.MINUS, 0);
	};

	PLUS() {
	    return this.getToken(prqlParser.PLUS, 0);
	};

	NOT() {
	    return this.getToken(prqlParser.NOT, 0);
	};

	nestedPipeline() {
	    return this.getTypedRuleContext(NestedPipelineContext,0);
	};

	literal() {
	    return this.getTypedRuleContext(LiteralContext,0);
	};

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterExprUnary(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitExprUnary(this);
		}
	}


}



class LiteralContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_literal;
    }

	NULL_() {
	    return this.getToken(prqlParser.NULL_, 0);
	};

	BOOLEAN() {
	    return this.getToken(prqlParser.BOOLEAN, 0);
	};

	NUMBER = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(prqlParser.NUMBER);
	    } else {
	        return this.getToken(prqlParser.NUMBER, i);
	    }
	};


	STRING() {
	    return this.getToken(prqlParser.STRING, 0);
	};

	IDENT = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(prqlParser.IDENT);
	    } else {
	        return this.getToken(prqlParser.IDENT, i);
	    }
	};


	INTERVAL_KIND() {
	    return this.getToken(prqlParser.INTERVAL_KIND, 0);
	};

	RANGE() {
	    return this.getToken(prqlParser.RANGE, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterLiteral(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitLiteral(this);
		}
	}


}



class ListContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_list;
    }

	LBRACKET() {
	    return this.getToken(prqlParser.LBRACKET, 0);
	};

	RBRACKET() {
	    return this.getToken(prqlParser.RBRACKET, 0);
	};

	assignCall = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(AssignCallContext);
	    } else {
	        return this.getTypedRuleContext(AssignCallContext,i);
	    }
	};

	exprCall = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(ExprCallContext);
	    } else {
	        return this.getTypedRuleContext(ExprCallContext,i);
	    }
	};

	nl = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(NlContext);
	    } else {
	        return this.getTypedRuleContext(NlContext,i);
	    }
	};

	COMMA = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(prqlParser.COMMA);
	    } else {
	        return this.getToken(prqlParser.COMMA, i);
	    }
	};


	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterList(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitList(this);
		}
	}


}



class NestedPipelineContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_nestedPipeline;
    }

	LPAREN() {
	    return this.getToken(prqlParser.LPAREN, 0);
	};

	pipeline() {
	    return this.getTypedRuleContext(PipelineContext,0);
	};

	RPAREN() {
	    return this.getToken(prqlParser.RPAREN, 0);
	};

	nl = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(NlContext);
	    } else {
	        return this.getTypedRuleContext(NlContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterNestedPipeline(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitNestedPipeline(this);
		}
	}


}




prqlParser.NlContext = NlContext; 
prqlParser.QueryContext = QueryContext; 
prqlParser.QueryDefContext = QueryDefContext; 
prqlParser.FuncDefContext = FuncDefContext; 
prqlParser.FuncDefNameContext = FuncDefNameContext; 
prqlParser.FuncDefParamsContext = FuncDefParamsContext; 
prqlParser.FuncDefParamContext = FuncDefParamContext; 
prqlParser.TypeDefContext = TypeDefContext; 
prqlParser.TypeTermContext = TypeTermContext; 
prqlParser.TableContext = TableContext; 
prqlParser.PipeContext = PipeContext; 
prqlParser.PipelineContext = PipelineContext; 
prqlParser.IdentBacktickContext = IdentBacktickContext; 
prqlParser.SignedIdentContext = SignedIdentContext; 
prqlParser.KeywordContext = KeywordContext; 
prqlParser.FuncCallContext = FuncCallContext; 
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
