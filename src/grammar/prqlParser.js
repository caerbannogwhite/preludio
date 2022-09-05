// Generated from .\prql.g4 by ANTLR 4.9.3
// jshint ignore: start
import antlr4 from 'antlr4';
import prqlListener from './prqlListener.js';

const serializedATN = ["\u0003\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786",
    "\u5964\u00039\u0160\u0004\u0002\t\u0002\u0004\u0003\t\u0003\u0004\u0004",
    "\t\u0004\u0004\u0005\t\u0005\u0004\u0006\t\u0006\u0004\u0007\t\u0007",
    "\u0004\b\t\b\u0004\t\t\t\u0004\n\t\n\u0004\u000b\t\u000b\u0004\f\t\f",
    "\u0004\r\t\r\u0004\u000e\t\u000e\u0004\u000f\t\u000f\u0004\u0010\t\u0010",
    "\u0004\u0011\t\u0011\u0004\u0012\t\u0012\u0004\u0013\t\u0013\u0004\u0014",
    "\t\u0014\u0004\u0015\t\u0015\u0004\u0016\t\u0016\u0004\u0017\t\u0017",
    "\u0004\u0018\t\u0018\u0004\u0019\t\u0019\u0004\u001a\t\u001a\u0004\u001b",
    "\t\u001b\u0004\u001c\t\u001c\u0004\u001d\t\u001d\u0004\u001e\t\u001e",
    "\u0004\u001f\t\u001f\u0004 \t \u0004!\t!\u0004\"\t\"\u0004#\t#\u0004",
    "$\t$\u0004%\t%\u0003\u0002\u0003\u0002\u0003\u0003\u0007\u0003N\n\u0003",
    "\f\u0003\u000e\u0003Q\u000b\u0003\u0003\u0003\u0005\u0003T\n\u0003\u0003",
    "\u0003\u0007\u0003W\n\u0003\f\u0003\u000e\u0003Z\u000b\u0003\u0003\u0003",
    "\u0003\u0003\u0003\u0003\u0005\u0003_\n\u0003\u0003\u0003\u0007\u0003",
    "b\n\u0003\f\u0003\u000e\u0003e\u000b\u0003\u0007\u0003g\n\u0003\f\u0003",
    "\u000e\u0003j\u000b\u0003\u0003\u0003\u0003\u0003\u0003\u0004\u0003",
    "\u0004\u0007\u0004p\n\u0004\f\u0004\u000e\u0004s\u000b\u0004\u0003\u0004",
    "\u0003\u0004\u0003\u0005\u0003\u0005\u0003\u0005\u0003\u0005\u0003\u0005",
    "\u0003\u0005\u0003\u0006\u0003\u0006\u0005\u0006\u007f\n\u0006\u0003",
    "\u0007\u0007\u0007\u0082\n\u0007\f\u0007\u000e\u0007\u0085\u000b\u0007",
    "\u0003\b\u0003\b\u0005\b\u0089\n\b\u0003\b\u0005\b\u008c\n\b\u0003\t",
    "\u0003\t\u0003\t\u0003\t\u0007\t\u0092\n\t\f\t\u000e\t\u0095\u000b\t",
    "\u0003\t\u0003\t\u0003\n\u0003\n\u0005\n\u009b\n\n\u0003\u000b\u0003",
    "\u000b\u0003\u000b\u0003\u000b\u0003\u000b\u0003\f\u0003\f\u0005\f\u00a4",
    "\n\f\u0003\r\u0003\r\u0003\r\u0003\r\u0007\r\u00aa\n\r\f\r\u000e\r\u00ad",
    "\u000b\r\u0003\u000e\u0003\u000e\u0007\u000e\u00b1\n\u000e\f\u000e\u000e",
    "\u000e\u00b4\u000b\u000e\u0003\u000e\u0003\u000e\u0003\u000f\u0003\u000f",
    "\u0003\u000f\u0003\u0010\u0003\u0010\u0003\u0011\u0003\u0011\u0003\u0011",
    "\u0003\u0011\u0006\u0011\u00c1\n\u0011\r\u0011\u000e\u0011\u00c2\u0003",
    "\u0012\u0003\u0012\u0003\u0012\u0003\u0012\u0005\u0012\u00c9\n\u0012",
    "\u0003\u0013\u0003\u0013\u0003\u0013\u0003\u0013\u0003\u0014\u0003\u0014",
    "\u0003\u0014\u0003\u0014\u0003\u0015\u0003\u0015\u0005\u0015\u00d5\n",
    "\u0015\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003",
    "\u0016\u0005\u0016\u00dd\n\u0016\u0003\u0016\u0003\u0016\u0003\u0016",
    "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016",
    "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016",
    "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0007\u0016",
    "\u00f3\n\u0016\f\u0016\u000e\u0016\u00f6\u000b\u0016\u0003\u0017\u0003",
    "\u0017\u0003\u0017\u0003\u0017\u0003\u0017\u0003\u0017\u0003\u0017\u0005",
    "\u0017\u00ff\n\u0017\u0003\u0018\u0003\u0018\u0003\u0018\u0003\u0018",
    "\u0005\u0018\u0105\n\u0018\u0003\u0019\u0003\u0019\u0003\u0019\u0003",
    "\u0019\u0003\u0019\u0005\u0019\u010c\n\u0019\u0003\u001a\u0003\u001a",
    "\u0007\u001a\u0110\n\u001a\f\u001a\u000e\u001a\u0113\u000b\u001a\u0003",
    "\u001a\u0003\u001a\u0005\u001a\u0117\n\u001a\u0003\u001a\u0003\u001a",
    "\u0007\u001a\u011b\n\u001a\f\u001a\u000e\u001a\u011e\u000b\u001a\u0003",
    "\u001a\u0003\u001a\u0005\u001a\u0122\n\u001a\u0007\u001a\u0124\n\u001a",
    "\f\u001a\u000e\u001a\u0127\u000b\u001a\u0003\u001a\u0005\u001a\u012a",
    "\n\u001a\u0003\u001a\u0005\u001a\u012d\n\u001a\u0005\u001a\u012f\n\u001a",
    "\u0003\u001a\u0003\u001a\u0003\u001b\u0003\u001b\u0007\u001b\u0135\n",
    "\u001b\f\u001b\u000e\u001b\u0138\u000b\u001b\u0003\u001b\u0003\u001b",
    "\u0007\u001b\u013c\n\u001b\f\u001b\u000e\u001b\u013f\u000b\u001b\u0003",
    "\u001b\u0003\u001b\u0003\u001c\u0003\u001c\u0003\u001c\u0003\u001c\u0003",
    "\u001d\u0003\u001d\u0003\u001d\u0003\u001d\u0003\u001d\u0003\u001d\u0005",
    "\u001d\u014d\n\u001d\u0003\u001e\u0003\u001e\u0003\u001f\u0003\u001f",
    "\u0003 \u0003 \u0003!\u0003!\u0003\"\u0003\"\u0003#\u0003#\u0003$\u0003",
    "$\u0003%\u0003%\u0003%\u0003%\u0002\u0003*&\u0002\u0004\u0006\b\n\f",
    "\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.0246",
    "8:<>@BDFH\u0002\u000b\u0003\u000278\u0004\u0002\'\'77\u0003\u0002\u0011",
    "\u0012\u0003\u0002\f\u000e\u0004\u0002\u0011\u0012..\u0003\u0002\u0013",
    "\u0015\u0004\u0002\u0016\u0019 !\u0003\u0002,-\u0003\u0002\u0003\u000b",
    "\u0002\u0172\u0002J\u0003\u0002\u0002\u0002\u0004O\u0003\u0002\u0002",
    "\u0002\u0006m\u0003\u0002\u0002\u0002\bv\u0003\u0002\u0002\u0002\n|",
    "\u0003\u0002\u0002\u0002\f\u0083\u0003\u0002\u0002\u0002\u000e\u0088",
    "\u0003\u0002\u0002\u0002\u0010\u008d\u0003\u0002\u0002\u0002\u0012\u0098",
    "\u0003\u0002\u0002\u0002\u0014\u009c\u0003\u0002\u0002\u0002\u0016\u00a3",
    "\u0003\u0002\u0002\u0002\u0018\u00a5\u0003\u0002\u0002\u0002\u001a\u00ae",
    "\u0003\u0002\u0002\u0002\u001c\u00b7\u0003\u0002\u0002\u0002\u001e\u00ba",
    "\u0003\u0002\u0002\u0002 \u00bc\u0003\u0002\u0002\u0002\"\u00c4\u0003",
    "\u0002\u0002\u0002$\u00ca\u0003\u0002\u0002\u0002&\u00ce\u0003\u0002",
    "\u0002\u0002(\u00d4\u0003\u0002\u0002\u0002*\u00dc\u0003\u0002\u0002",
    "\u0002,\u00fe\u0003\u0002\u0002\u0002.\u0100\u0003\u0002\u0002\u0002",
    "0\u010b\u0003\u0002\u0002\u00022\u010d\u0003\u0002\u0002\u00024\u0132",
    "\u0003\u0002\u0002\u00026\u0142\u0003\u0002\u0002\u00028\u014c\u0003",
    "\u0002\u0002\u0002:\u014e\u0003\u0002\u0002\u0002<\u0150\u0003\u0002",
    "\u0002\u0002>\u0152\u0003\u0002\u0002\u0002@\u0154\u0003\u0002\u0002",
    "\u0002B\u0156\u0003\u0002\u0002\u0002D\u0158\u0003\u0002\u0002\u0002",
    "F\u015a\u0003\u0002\u0002\u0002H\u015c\u0003\u0002\u0002\u0002JK\t\u0002",
    "\u0002\u0002K\u0003\u0003\u0002\u0002\u0002LN\u0005\u0002\u0002\u0002",
    "ML\u0003\u0002\u0002\u0002NQ\u0003\u0002\u0002\u0002OM\u0003\u0002\u0002",
    "\u0002OP\u0003\u0002\u0002\u0002PS\u0003\u0002\u0002\u0002QO\u0003\u0002",
    "\u0002\u0002RT\u0005\u0006\u0004\u0002SR\u0003\u0002\u0002\u0002ST\u0003",
    "\u0002\u0002\u0002TX\u0003\u0002\u0002\u0002UW\u0005\u0002\u0002\u0002",
    "VU\u0003\u0002\u0002\u0002WZ\u0003\u0002\u0002\u0002XV\u0003\u0002\u0002",
    "\u0002XY\u0003\u0002\u0002\u0002Yh\u0003\u0002\u0002\u0002ZX\u0003\u0002",
    "\u0002\u0002[_\u0005\b\u0005\u0002\\_\u0005\u0014\u000b\u0002]_\u0005",
    "\u0018\r\u0002^[\u0003\u0002\u0002\u0002^\\\u0003\u0002\u0002\u0002",
    "^]\u0003\u0002\u0002\u0002_c\u0003\u0002\u0002\u0002`b\u0005\u0002\u0002",
    "\u0002a`\u0003\u0002\u0002\u0002be\u0003\u0002\u0002\u0002ca\u0003\u0002",
    "\u0002\u0002cd\u0003\u0002\u0002\u0002dg\u0003\u0002\u0002\u0002ec\u0003",
    "\u0002\u0002\u0002f^\u0003\u0002\u0002\u0002gj\u0003\u0002\u0002\u0002",
    "hf\u0003\u0002\u0002\u0002hi\u0003\u0002\u0002\u0002ik\u0003\u0002\u0002",
    "\u0002jh\u0003\u0002\u0002\u0002kl\u0007\u0002\u0002\u0003l\u0005\u0003",
    "\u0002\u0002\u0002mq\u0007\r\u0002\u0002np\u0005\"\u0012\u0002on\u0003",
    "\u0002\u0002\u0002ps\u0003\u0002\u0002\u0002qo\u0003\u0002\u0002\u0002",
    "qr\u0003\u0002\u0002\u0002rt\u0003\u0002\u0002\u0002sq\u0003\u0002\u0002",
    "\u0002tu\u0005\u0002\u0002\u0002u\u0007\u0003\u0002\u0002\u0002vw\u0007",
    "\f\u0002\u0002wx\u0005\n\u0006\u0002xy\u0005\f\u0007\u0002yz\u0007\u000f",
    "\u0002\u0002z{\u0005*\u0016\u0002{\t\u0003\u0002\u0002\u0002|~\u0007",
    "3\u0002\u0002}\u007f\u0005\u0010\t\u0002~}\u0003\u0002\u0002\u0002~",
    "\u007f\u0003\u0002\u0002\u0002\u007f\u000b\u0003\u0002\u0002\u0002\u0080",
    "\u0082\u0005\u000e\b\u0002\u0081\u0080\u0003\u0002\u0002\u0002\u0082",
    "\u0085\u0003\u0002\u0002\u0002\u0083\u0081\u0003\u0002\u0002\u0002\u0083",
    "\u0084\u0003\u0002\u0002\u0002\u0084\r\u0003\u0002\u0002\u0002\u0085",
    "\u0083\u0003\u0002\u0002\u0002\u0086\u0089\u0005\"\u0012\u0002\u0087",
    "\u0089\u00073\u0002\u0002\u0088\u0086\u0003\u0002\u0002\u0002\u0088",
    "\u0087\u0003\u0002\u0002\u0002\u0089\u008b\u0003\u0002\u0002\u0002\u008a",
    "\u008c\u0005\u0010\t\u0002\u008b\u008a\u0003\u0002\u0002\u0002\u008b",
    "\u008c\u0003\u0002\u0002\u0002\u008c\u000f\u0003\u0002\u0002\u0002\u008d",
    "\u008e\u0007 \u0002\u0002\u008e\u008f\u0005\u0012\n\u0002\u008f\u0093",
    "\u0007\u001a\u0002\u0002\u0090\u0092\u0005\u0012\n\u0002\u0091\u0090",
    "\u0003\u0002\u0002\u0002\u0092\u0095\u0003\u0002\u0002\u0002\u0093\u0091",
    "\u0003\u0002\u0002\u0002\u0093\u0094\u0003\u0002\u0002\u0002\u0094\u0096",
    "\u0003\u0002\u0002\u0002\u0095\u0093\u0003\u0002\u0002\u0002\u0096\u0097",
    "\u0007!\u0002\u0002\u0097\u0011\u0003\u0002\u0002\u0002\u0098\u009a",
    "\u00073\u0002\u0002\u0099\u009b\u0005\u0010\t\u0002\u009a\u0099\u0003",
    "\u0002\u0002\u0002\u009a\u009b\u0003\u0002\u0002\u0002\u009b\u0013\u0003",
    "\u0002\u0002\u0002\u009c\u009d\u0007\u000e\u0002\u0002\u009d\u009e\u0007",
    "3\u0002\u0002\u009e\u009f\u0007\u0010\u0002\u0002\u009f\u00a0\u0005",
    "4\u001b\u0002\u00a0\u0015\u0003\u0002\u0002\u0002\u00a1\u00a4\u0005",
    "\u0002\u0002\u0002\u00a2\u00a4\u0007\u001a\u0002\u0002\u00a3\u00a1\u0003",
    "\u0002\u0002\u0002\u00a3\u00a2\u0003\u0002\u0002\u0002\u00a4\u0017\u0003",
    "\u0002\u0002\u0002\u00a5\u00ab\u0005(\u0015\u0002\u00a6\u00a7\u0005",
    "\u0016\f\u0002\u00a7\u00a8\u0005(\u0015\u0002\u00a8\u00aa\u0003\u0002",
    "\u0002\u0002\u00a9\u00a6\u0003\u0002\u0002\u0002\u00aa\u00ad\u0003\u0002",
    "\u0002\u0002\u00ab\u00a9\u0003\u0002\u0002\u0002\u00ab\u00ac\u0003\u0002",
    "\u0002\u0002\u00ac\u0019\u0003\u0002\u0002\u0002\u00ad\u00ab\u0003\u0002",
    "\u0002\u0002\u00ae\u00b2\u0007\'\u0002\u0002\u00af\u00b1\n\u0003\u0002",
    "\u0002\u00b0\u00af\u0003\u0002\u0002\u0002\u00b1\u00b4\u0003\u0002\u0002",
    "\u0002\u00b2\u00b0\u0003\u0002\u0002\u0002\u00b2\u00b3\u0003\u0002\u0002",
    "\u0002\u00b3\u00b5\u0003\u0002\u0002\u0002\u00b4\u00b2\u0003\u0002\u0002",
    "\u0002\u00b5\u00b6\u0007\'\u0002\u0002\u00b6\u001b\u0003\u0002\u0002",
    "\u0002\u00b7\u00b8\t\u0004\u0002\u0002\u00b8\u00b9\u00073\u0002\u0002",
    "\u00b9\u001d\u0003\u0002\u0002\u0002\u00ba\u00bb\t\u0005\u0002\u0002",
    "\u00bb\u001f\u0003\u0002\u0002\u0002\u00bc\u00c0\u00073\u0002\u0002",
    "\u00bd\u00c1\u0005\"\u0012\u0002\u00be\u00c1\u0005$\u0013\u0002\u00bf",
    "\u00c1\u0005*\u0016\u0002\u00c0\u00bd\u0003\u0002\u0002\u0002\u00c0",
    "\u00be\u0003\u0002\u0002\u0002\u00c0\u00bf\u0003\u0002\u0002\u0002\u00c1",
    "\u00c2\u0003\u0002\u0002\u0002\u00c2\u00c0\u0003\u0002\u0002\u0002\u00c2",
    "\u00c3\u0003\u0002\u0002\u0002\u00c3!\u0003\u0002\u0002\u0002\u00c4",
    "\u00c5\u00073\u0002\u0002\u00c5\u00c8\u0007\u001b\u0002\u0002\u00c6",
    "\u00c9\u0005$\u0013\u0002\u00c7\u00c9\u0005*\u0016\u0002\u00c8\u00c6",
    "\u0003\u0002\u0002\u0002\u00c8\u00c7\u0003\u0002\u0002\u0002\u00c9#",
    "\u0003\u0002\u0002\u0002\u00ca\u00cb\u00073\u0002\u0002\u00cb\u00cc",
    "\u0007\u0010\u0002\u0002\u00cc\u00cd\u0005*\u0016\u0002\u00cd%\u0003",
    "\u0002\u0002\u0002\u00ce\u00cf\u00073\u0002\u0002\u00cf\u00d0\u0007",
    "\u0010\u0002\u0002\u00d0\u00d1\u0005(\u0015\u0002\u00d1\'\u0003\u0002",
    "\u0002\u0002\u00d2\u00d5\u0005 \u0011\u0002\u00d3\u00d5\u0005*\u0016",
    "\u0002\u00d4\u00d2\u0003\u0002\u0002\u0002\u00d4\u00d3\u0003\u0002\u0002",
    "\u0002\u00d5)\u0003\u0002\u0002\u0002\u00d6\u00d7\b\u0016\u0001\u0002",
    "\u00d7\u00d8\u0007$\u0002\u0002\u00d8\u00d9\u0005*\u0016\u0002\u00d9",
    "\u00da\u0007%\u0002\u0002\u00da\u00dd\u0003\u0002\u0002\u0002\u00db",
    "\u00dd\u0005,\u0017\u0002\u00dc\u00d6\u0003\u0002\u0002\u0002\u00dc",
    "\u00db\u0003\u0002\u0002\u0002\u00dd\u00f4\u0003\u0002\u0002\u0002\u00de",
    "\u00df\f\t\u0002\u0002\u00df\u00e0\u0005<\u001f\u0002\u00e0\u00e1\u0005",
    "*\u0016\n\u00e1\u00f3\u0003\u0002\u0002\u0002\u00e2\u00e3\f\b\u0002",
    "\u0002\u00e3\u00e4\u0005> \u0002\u00e4\u00e5\u0005*\u0016\t\u00e5\u00f3",
    "\u0003\u0002\u0002\u0002\u00e6\u00e7\f\u0007\u0002\u0002\u00e7\u00e8",
    "\u0005@!\u0002\u00e8\u00e9\u0005*\u0016\b\u00e9\u00f3\u0003\u0002\u0002",
    "\u0002\u00ea\u00eb\f\u0006\u0002\u0002\u00eb\u00ec\u0005D#\u0002\u00ec",
    "\u00ed\u0005*\u0016\u0007\u00ed\u00f3\u0003\u0002\u0002\u0002\u00ee",
    "\u00ef\f\u0005\u0002\u0002\u00ef\u00f0\u0005B\"\u0002\u00f0\u00f1\u0005",
    "*\u0016\u0006\u00f1\u00f3\u0003\u0002\u0002\u0002\u00f2\u00de\u0003",
    "\u0002\u0002\u0002\u00f2\u00e2\u0003\u0002\u0002\u0002\u00f2\u00e6\u0003",
    "\u0002\u0002\u0002\u00f2\u00ea\u0003\u0002\u0002\u0002\u00f2\u00ee\u0003",
    "\u0002\u0002\u0002\u00f3\u00f6\u0003\u0002\u0002\u0002\u00f4\u00f2\u0003",
    "\u0002\u0002\u0002\u00f4\u00f5\u0003\u0002\u0002\u0002\u00f5+\u0003",
    "\u0002\u0002\u0002\u00f6\u00f4\u0003\u0002\u0002\u0002\u00f7\u00ff\u0005",
    "6\u001c\u0002\u00f8\u00ff\u00050\u0019\u0002\u00f9\u00ff\u00073\u0002",
    "\u0002\u00fa\u00ff\u0005\u001a\u000e\u0002\u00fb\u00ff\u0005.\u0018",
    "\u0002\u00fc\u00ff\u00052\u001a\u0002\u00fd\u00ff\u00054\u001b\u0002",
    "\u00fe\u00f7\u0003\u0002\u0002\u0002\u00fe\u00f8\u0003\u0002\u0002\u0002",
    "\u00fe\u00f9\u0003\u0002\u0002\u0002\u00fe\u00fa\u0003\u0002\u0002\u0002",
    "\u00fe\u00fb\u0003\u0002\u0002\u0002\u00fe\u00fc\u0003\u0002\u0002\u0002",
    "\u00fe\u00fd\u0003\u0002\u0002\u0002\u00ff-\u0003\u0002\u0002\u0002",
    "\u0100\u0104\u0005:\u001e\u0002\u0101\u0105\u00054\u001b\u0002\u0102",
    "\u0105\u00050\u0019\u0002\u0103\u0105\u00073\u0002\u0002\u0104\u0101",
    "\u0003\u0002\u0002\u0002\u0104\u0102\u0003\u0002\u0002\u0002\u0104\u0103",
    "\u0003\u0002\u0002\u0002\u0105/\u0003\u0002\u0002\u0002\u0106\u010c",
    "\u0005H%\u0002\u0107\u010c\u00072\u0002\u0002\u0108\u010c\u00071\u0002",
    "\u0002\u0109\u010c\u00070\u0002\u0002\u010a\u010c\u00079\u0002\u0002",
    "\u010b\u0106\u0003\u0002\u0002\u0002\u010b\u0107\u0003\u0002\u0002\u0002",
    "\u010b\u0108\u0003\u0002\u0002\u0002\u010b\u0109\u0003\u0002\u0002\u0002",
    "\u010b\u010a\u0003\u0002\u0002\u0002\u010c1\u0003\u0002\u0002\u0002",
    "\u010d\u012e\u0007\"\u0002\u0002\u010e\u0110\u0005\u0002\u0002\u0002",
    "\u010f\u010e\u0003\u0002\u0002\u0002\u0110\u0113\u0003\u0002\u0002\u0002",
    "\u0111\u010f\u0003\u0002\u0002\u0002\u0111\u0112\u0003\u0002\u0002\u0002",
    "\u0112\u0116\u0003\u0002\u0002\u0002\u0113\u0111\u0003\u0002\u0002\u0002",
    "\u0114\u0117\u0005&\u0014\u0002\u0115\u0117\u0005(\u0015\u0002\u0116",
    "\u0114\u0003\u0002\u0002\u0002\u0116\u0115\u0003\u0002\u0002\u0002\u0117",
    "\u0125\u0003\u0002\u0002\u0002\u0118\u011c\u0007\u001c\u0002\u0002\u0119",
    "\u011b\u0005\u0002\u0002\u0002\u011a\u0119\u0003\u0002\u0002\u0002\u011b",
    "\u011e\u0003\u0002\u0002\u0002\u011c\u011a\u0003\u0002\u0002\u0002\u011c",
    "\u011d\u0003\u0002\u0002\u0002\u011d\u0121\u0003\u0002\u0002\u0002\u011e",
    "\u011c\u0003\u0002\u0002\u0002\u011f\u0122\u0005&\u0014\u0002\u0120",
    "\u0122\u0005(\u0015\u0002\u0121\u011f\u0003\u0002\u0002\u0002\u0121",
    "\u0120\u0003\u0002\u0002\u0002\u0122\u0124\u0003\u0002\u0002\u0002\u0123",
    "\u0118\u0003\u0002\u0002\u0002\u0124\u0127\u0003\u0002\u0002\u0002\u0125",
    "\u0123\u0003\u0002\u0002\u0002\u0125\u0126\u0003\u0002\u0002\u0002\u0126",
    "\u0129\u0003\u0002\u0002\u0002\u0127\u0125\u0003\u0002\u0002\u0002\u0128",
    "\u012a\u0007\u001c\u0002\u0002\u0129\u0128\u0003\u0002\u0002\u0002\u0129",
    "\u012a\u0003\u0002\u0002\u0002\u012a\u012c\u0003\u0002\u0002\u0002\u012b",
    "\u012d\u0005\u0002\u0002\u0002\u012c\u012b\u0003\u0002\u0002\u0002\u012c",
    "\u012d\u0003\u0002\u0002\u0002\u012d\u012f\u0003\u0002\u0002\u0002\u012e",
    "\u0111\u0003\u0002\u0002\u0002\u012e\u012f\u0003\u0002\u0002\u0002\u012f",
    "\u0130\u0003\u0002\u0002\u0002\u0130\u0131\u0007#\u0002\u0002\u0131",
    "3\u0003\u0002\u0002\u0002\u0132\u0136\u0007$\u0002\u0002\u0133\u0135",
    "\u0005\u0002\u0002\u0002\u0134\u0133\u0003\u0002\u0002\u0002\u0135\u0138",
    "\u0003\u0002\u0002\u0002\u0136\u0134\u0003\u0002\u0002\u0002\u0136\u0137",
    "\u0003\u0002\u0002\u0002\u0137\u0139\u0003\u0002\u0002\u0002\u0138\u0136",
    "\u0003\u0002\u0002\u0002\u0139\u013d\u0005\u0018\r\u0002\u013a\u013c",
    "\u0005\u0002\u0002\u0002\u013b\u013a\u0003\u0002\u0002\u0002\u013c\u013f",
    "\u0003\u0002\u0002\u0002\u013d\u013b\u0003\u0002\u0002\u0002\u013d\u013e",
    "\u0003\u0002\u0002\u0002\u013e\u0140\u0003\u0002\u0002\u0002\u013f\u013d",
    "\u0003\u0002\u0002\u0002\u0140\u0141\u0007%\u0002\u0002\u01415\u0003",
    "\u0002\u0002\u0002\u0142\u0143\u00050\u0019\u0002\u0143\u0144\u0007",
    "\u001f\u0002\u0002\u0144\u0145\u00050\u0019\u0002\u01457\u0003\u0002",
    "\u0002\u0002\u0146\u014d\u0005:\u001e\u0002\u0147\u014d\u0005<\u001f",
    "\u0002\u0148\u014d\u0005> \u0002\u0149\u014d\u0005@!\u0002\u014a\u014d",
    "\u0005B\"\u0002\u014b\u014d\u0005D#\u0002\u014c\u0146\u0003\u0002\u0002",
    "\u0002\u014c\u0147\u0003\u0002\u0002\u0002\u014c\u0148\u0003\u0002\u0002",
    "\u0002\u014c\u0149\u0003\u0002\u0002\u0002\u014c\u014a\u0003\u0002\u0002",
    "\u0002\u014c\u014b\u0003\u0002\u0002\u0002\u014d9\u0003\u0002\u0002",
    "\u0002\u014e\u014f\t\u0006\u0002\u0002\u014f;\u0003\u0002\u0002\u0002",
    "\u0150\u0151\t\u0007\u0002\u0002\u0151=\u0003\u0002\u0002\u0002\u0152",
    "\u0153\t\u0004\u0002\u0002\u0153?\u0003\u0002\u0002\u0002\u0154\u0155",
    "\t\b\u0002\u0002\u0155A\u0003\u0002\u0002\u0002\u0156\u0157\t\t\u0002",
    "\u0002\u0157C\u0003\u0002\u0002\u0002\u0158\u0159\u0007/\u0002\u0002",
    "\u0159E\u0003\u0002\u0002\u0002\u015a\u015b\t\n\u0002\u0002\u015bG\u0003",
    "\u0002\u0002\u0002\u015c\u015d\u00072\u0002\u0002\u015d\u015e\u0005",
    "F$\u0002\u015eI\u0003\u0002\u0002\u0002\'OSX^chq~\u0083\u0088\u008b",
    "\u0093\u009a\u00a3\u00ab\u00b2\u00c0\u00c2\u00c8\u00d4\u00dc\u00f2\u00f4",
    "\u00fe\u0104\u010b\u0111\u0116\u011c\u0121\u0125\u0129\u012c\u012e\u0136",
    "\u013d\u014c"].join("");


const atn = new antlr4.atn.ATNDeserializer().deserialize(serializedATN);

const decisionsToDFA = atn.decisionToState.map( (ds, index) => new antlr4.dfa.DFA(ds, index) );

const sharedContextCache = new antlr4.PredictionContextCache();

export default class prqlParser extends antlr4.Parser {

    static grammarFileName = "prql.g4";
    static literalNames = [ null, "'microseconds'", "'milliseconds'", "'seconds'", 
                            "'minutes'", "'hours'", "'days'", "'weeks'", 
                            "'months'", "'years'", "'func'", "'prql'", "'table'", 
                            "'->'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'", 
                            "'=='", "'!='", "'<='", "'>='", "'|'", "':'", 
                            "','", "'.'", "'$'", "'..'", "'<'", "'>'", "'['", 
                            "']'", "'('", "')'", "'_'", "'`'", "'\"'", "'''", 
                            "'\"\"\"'", "'''''", "'and'", "'or'", "'!'", 
                            "'??'", "'null'" ];
    static symbolicNames = [ null, null, null, null, null, null, null, null, 
                             null, null, "FUNC", "PRQL", "TABLE", "ARROW", 
                             "ASSIGN", "PLUS", "MINUS", "STAR", "DIV", "MOD", 
                             "EQ", "NE", "LE", "GE", "BAR", "COLON", "COMMA", 
                             "DOT", "DOLLAR", "RANGE", "LANG", "RANG", "LBRACKET", 
                             "RBRACKET", "LPAREN", "RPAREN", "UNDERSCORE", 
                             "BACKTICK", "DOUBLE_QUOTE", "SINGLE_QUOTE", 
                             "TRIPLE_DOUBLE_QUOTE", "TRIPLE_SINGLE_QUOTE", 
                             "AND", "OR", "NOT", "COALESCE", "NULL_", "BOOLEAN", 
                             "NUMBER", "IDENT", "IDENT_START", "IDENT_NEXT", 
                             "WHITESPACE", "NEWLINE", "COMMENT", "STRING" ];
    static ruleNames = [ "nl", "query", "queryDef", "funcDef", "funcDefName", 
                         "funcDefParams", "funcDefParam", "typeDef", "typeTerm", 
                         "table", "pipe", "pipeline", "identBackticks", 
                         "signedIdent", "keyword", "funcCall", "namedArg", 
                         "assign", "assignCall", "exprCall", "expr", "term", 
                         "exprUnary", "literal", "list", "nestedPipeline", 
                         "range", "operator", "operatorUnary", "operatorMul", 
                         "operatorAdd", "operatorCompare", "operatorLogical", 
                         "operatorCoalesce", "intervalKind", "interval" ];

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
	        this.state = 72;
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
	        this.state = 77;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,0,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                this.state = 74;
	                this.nl(); 
	            }
	            this.state = 79;
	            this._errHandler.sync(this);
	            _alt = this._interp.adaptivePredict(this._input,0,this._ctx);
	        }

	        this.state = 81;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.PRQL) {
	            this.state = 80;
	            this.queryDef();
	        }

	        this.state = 86;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	            this.state = 83;
	            this.nl();
	            this.state = 88;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 102;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.FUNC) | (1 << prqlParser.TABLE) | (1 << prqlParser.PLUS) | (1 << prqlParser.MINUS))) !== 0) || ((((_la - 32)) & ~0x1f) == 0 && ((1 << (_la - 32)) & ((1 << (prqlParser.LBRACKET - 32)) | (1 << (prqlParser.LPAREN - 32)) | (1 << (prqlParser.BACKTICK - 32)) | (1 << (prqlParser.NOT - 32)) | (1 << (prqlParser.NULL_ - 32)) | (1 << (prqlParser.BOOLEAN - 32)) | (1 << (prqlParser.NUMBER - 32)) | (1 << (prqlParser.IDENT - 32)) | (1 << (prqlParser.STRING - 32)))) !== 0)) {
	            this.state = 92;
	            this._errHandler.sync(this);
	            switch(this._input.LA(1)) {
	            case prqlParser.FUNC:
	                this.state = 89;
	                this.funcDef();
	                break;
	            case prqlParser.TABLE:
	                this.state = 90;
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
	                this.state = 91;
	                this.pipeline();
	                break;
	            default:
	                throw new antlr4.error.NoViableAltException(this);
	            }
	            this.state = 97;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	                this.state = 94;
	                this.nl();
	                this.state = 99;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	            }
	            this.state = 104;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 105;
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
	        this.state = 107;
	        this.match(prqlParser.PRQL);
	        this.state = 111;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.IDENT) {
	            this.state = 108;
	            this.namedArg();
	            this.state = 113;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 114;
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
	        this.state = 116;
	        this.match(prqlParser.FUNC);
	        this.state = 117;
	        this.funcDefName();
	        this.state = 118;
	        this.funcDefParams();
	        this.state = 119;
	        this.match(prqlParser.ARROW);
	        this.state = 120;
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
	        this.state = 122;
	        this.match(prqlParser.IDENT);
	        this.state = 124;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.LANG) {
	            this.state = 123;
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
	        this.state = 129;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.IDENT) {
	            this.state = 126;
	            this.funcDefParam();
	            this.state = 131;
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
	        this.state = 134;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,9,this._ctx);
	        switch(la_) {
	        case 1:
	            this.state = 132;
	            this.namedArg();
	            break;

	        case 2:
	            this.state = 133;
	            this.match(prqlParser.IDENT);
	            break;

	        }
	        this.state = 137;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.LANG) {
	            this.state = 136;
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
	        this.state = 139;
	        this.match(prqlParser.LANG);
	        this.state = 140;
	        this.typeTerm();
	        this.state = 141;
	        this.match(prqlParser.BAR);
	        this.state = 145;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.IDENT) {
	            this.state = 142;
	            this.typeTerm();
	            this.state = 147;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 148;
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
	        this.state = 150;
	        this.match(prqlParser.IDENT);
	        this.state = 152;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.LANG) {
	            this.state = 151;
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
	        this.state = 154;
	        this.match(prqlParser.TABLE);
	        this.state = 155;
	        this.match(prqlParser.IDENT);
	        this.state = 156;
	        this.match(prqlParser.ASSIGN);
	        this.state = 157;
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
	        this.state = 161;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case prqlParser.NEWLINE:
	        case prqlParser.COMMENT:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 159;
	            this.nl();
	            break;
	        case prqlParser.BAR:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 160;
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
	        this.state = 163;
	        this.exprCall();
	        this.state = 169;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,14,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                this.state = 164;
	                this.pipe();
	                this.state = 165;
	                this.exprCall(); 
	            }
	            this.state = 171;
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



	identBackticks() {
	    let localctx = new IdentBackticksContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 24, prqlParser.RULE_identBackticks);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 172;
	        this.match(prqlParser.BACKTICK);
	        this.state = 176;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.T__0) | (1 << prqlParser.T__1) | (1 << prqlParser.T__2) | (1 << prqlParser.T__3) | (1 << prqlParser.T__4) | (1 << prqlParser.T__5) | (1 << prqlParser.T__6) | (1 << prqlParser.T__7) | (1 << prqlParser.T__8) | (1 << prqlParser.FUNC) | (1 << prqlParser.PRQL) | (1 << prqlParser.TABLE) | (1 << prqlParser.ARROW) | (1 << prqlParser.ASSIGN) | (1 << prqlParser.PLUS) | (1 << prqlParser.MINUS) | (1 << prqlParser.STAR) | (1 << prqlParser.DIV) | (1 << prqlParser.MOD) | (1 << prqlParser.EQ) | (1 << prqlParser.NE) | (1 << prqlParser.LE) | (1 << prqlParser.GE) | (1 << prqlParser.BAR) | (1 << prqlParser.COLON) | (1 << prqlParser.COMMA) | (1 << prqlParser.DOT) | (1 << prqlParser.DOLLAR) | (1 << prqlParser.RANGE) | (1 << prqlParser.LANG) | (1 << prqlParser.RANG))) !== 0) || ((((_la - 32)) & ~0x1f) == 0 && ((1 << (_la - 32)) & ((1 << (prqlParser.LBRACKET - 32)) | (1 << (prqlParser.RBRACKET - 32)) | (1 << (prqlParser.LPAREN - 32)) | (1 << (prqlParser.RPAREN - 32)) | (1 << (prqlParser.UNDERSCORE - 32)) | (1 << (prqlParser.DOUBLE_QUOTE - 32)) | (1 << (prqlParser.SINGLE_QUOTE - 32)) | (1 << (prqlParser.TRIPLE_DOUBLE_QUOTE - 32)) | (1 << (prqlParser.TRIPLE_SINGLE_QUOTE - 32)) | (1 << (prqlParser.AND - 32)) | (1 << (prqlParser.OR - 32)) | (1 << (prqlParser.NOT - 32)) | (1 << (prqlParser.COALESCE - 32)) | (1 << (prqlParser.NULL_ - 32)) | (1 << (prqlParser.BOOLEAN - 32)) | (1 << (prqlParser.NUMBER - 32)) | (1 << (prqlParser.IDENT - 32)) | (1 << (prqlParser.IDENT_START - 32)) | (1 << (prqlParser.IDENT_NEXT - 32)) | (1 << (prqlParser.WHITESPACE - 32)) | (1 << (prqlParser.COMMENT - 32)) | (1 << (prqlParser.STRING - 32)))) !== 0)) {
	            this.state = 173;
	            _la = this._input.LA(1);
	            if(_la<=0 || _la===prqlParser.BACKTICK || _la===prqlParser.NEWLINE) {
	            this._errHandler.recoverInline(this);
	            }
	            else {
	            	this._errHandler.reportMatch(this);
	                this.consume();
	            }
	            this.state = 178;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 179;
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
	        this.state = 181;
	        _la = this._input.LA(1);
	        if(!(_la===prqlParser.PLUS || _la===prqlParser.MINUS)) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 182;
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
	        this.state = 184;
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
	        this.state = 186;
	        this.match(prqlParser.IDENT);
	        this.state = 190; 
	        this._errHandler.sync(this);
	        var _alt = 1;
	        do {
	        	switch (_alt) {
	        	case 1:
	        		this.state = 190;
	        		this._errHandler.sync(this);
	        		var la_ = this._interp.adaptivePredict(this._input,16,this._ctx);
	        		switch(la_) {
	        		case 1:
	        		    this.state = 187;
	        		    this.namedArg();
	        		    break;

	        		case 2:
	        		    this.state = 188;
	        		    this.assign();
	        		    break;

	        		case 3:
	        		    this.state = 189;
	        		    this.expr(0);
	        		    break;

	        		}
	        		break;
	        	default:
	        		throw new antlr4.error.NoViableAltException(this);
	        	}
	        	this.state = 192; 
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
	        this.state = 194;
	        this.match(prqlParser.IDENT);
	        this.state = 195;
	        this.match(prqlParser.COLON);
	        this.state = 198;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,18,this._ctx);
	        switch(la_) {
	        case 1:
	            this.state = 196;
	            this.assign();
	            break;

	        case 2:
	            this.state = 197;
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
	        this.state = 200;
	        this.match(prqlParser.IDENT);
	        this.state = 201;
	        this.match(prqlParser.ASSIGN);
	        this.state = 202;
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
	        this.state = 204;
	        this.match(prqlParser.IDENT);
	        this.state = 205;
	        this.match(prqlParser.ASSIGN);
	        this.state = 206;
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
	        this.state = 210;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,19,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 208;
	            this.funcCall();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 209;
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
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 218;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,20,this._ctx);
	        switch(la_) {
	        case 1:
	            this.state = 213;
	            this.match(prqlParser.LPAREN);
	            this.state = 214;
	            this.expr(0);
	            this.state = 215;
	            this.match(prqlParser.RPAREN);
	            break;

	        case 2:
	            this.state = 217;
	            this.term();
	            break;

	        }
	        this._ctx.stop = this._input.LT(-1);
	        this.state = 242;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,22,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                if(this._parseListeners!==null) {
	                    this.triggerExitRuleEvent();
	                }
	                _prevctx = localctx;
	                this.state = 240;
	                this._errHandler.sync(this);
	                var la_ = this._interp.adaptivePredict(this._input,21,this._ctx);
	                switch(la_) {
	                case 1:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 220;
	                    if (!( this.precpred(this._ctx, 7))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 7)");
	                    }
	                    this.state = 221;
	                    this.operatorMul();
	                    this.state = 222;
	                    this.expr(8);
	                    break;

	                case 2:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 224;
	                    if (!( this.precpred(this._ctx, 6))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 6)");
	                    }
	                    this.state = 225;
	                    this.operatorAdd();
	                    this.state = 226;
	                    this.expr(7);
	                    break;

	                case 3:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 228;
	                    if (!( this.precpred(this._ctx, 5))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 5)");
	                    }
	                    this.state = 229;
	                    this.operatorCompare();
	                    this.state = 230;
	                    this.expr(6);
	                    break;

	                case 4:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 232;
	                    if (!( this.precpred(this._ctx, 4))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 4)");
	                    }
	                    this.state = 233;
	                    this.operatorCoalesce();
	                    this.state = 234;
	                    this.expr(5);
	                    break;

	                case 5:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 236;
	                    if (!( this.precpred(this._ctx, 3))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 3)");
	                    }
	                    this.state = 237;
	                    this.operatorLogical();
	                    this.state = 238;
	                    this.expr(4);
	                    break;

	                } 
	            }
	            this.state = 244;
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
	        this.state = 252;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,23,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 245;
	            this.range();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 246;
	            this.literal();
	            break;

	        case 3:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 247;
	            this.match(prqlParser.IDENT);
	            break;

	        case 4:
	            this.enterOuterAlt(localctx, 4);
	            this.state = 248;
	            this.identBackticks();
	            break;

	        case 5:
	            this.enterOuterAlt(localctx, 5);
	            this.state = 249;
	            this.exprUnary();
	            break;

	        case 6:
	            this.enterOuterAlt(localctx, 6);
	            this.state = 250;
	            this.list();
	            break;

	        case 7:
	            this.enterOuterAlt(localctx, 7);
	            this.state = 251;
	            this.nestedPipeline();
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



	exprUnary() {
	    let localctx = new ExprUnaryContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 44, prqlParser.RULE_exprUnary);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 254;
	        this.operatorUnary();
	        this.state = 258;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case prqlParser.LPAREN:
	            this.state = 255;
	            this.nestedPipeline();
	            break;
	        case prqlParser.NULL_:
	        case prqlParser.BOOLEAN:
	        case prqlParser.NUMBER:
	        case prqlParser.STRING:
	            this.state = 256;
	            this.literal();
	            break;
	        case prqlParser.IDENT:
	            this.state = 257;
	            this.match(prqlParser.IDENT);
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



	literal() {
	    let localctx = new LiteralContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 46, prqlParser.RULE_literal);
	    try {
	        this.state = 265;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,25,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 260;
	            this.interval();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 261;
	            this.match(prqlParser.NUMBER);
	            break;

	        case 3:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 262;
	            this.match(prqlParser.BOOLEAN);
	            break;

	        case 4:
	            this.enterOuterAlt(localctx, 4);
	            this.state = 263;
	            this.match(prqlParser.NULL_);
	            break;

	        case 5:
	            this.enterOuterAlt(localctx, 5);
	            this.state = 264;
	            this.match(prqlParser.STRING);
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
	        this.state = 267;
	        this.match(prqlParser.LBRACKET);
	        this.state = 300;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.PLUS || _la===prqlParser.MINUS || ((((_la - 32)) & ~0x1f) == 0 && ((1 << (_la - 32)) & ((1 << (prqlParser.LBRACKET - 32)) | (1 << (prqlParser.LPAREN - 32)) | (1 << (prqlParser.BACKTICK - 32)) | (1 << (prqlParser.NOT - 32)) | (1 << (prqlParser.NULL_ - 32)) | (1 << (prqlParser.BOOLEAN - 32)) | (1 << (prqlParser.NUMBER - 32)) | (1 << (prqlParser.IDENT - 32)) | (1 << (prqlParser.NEWLINE - 32)) | (1 << (prqlParser.COMMENT - 32)) | (1 << (prqlParser.STRING - 32)))) !== 0)) {
	            this.state = 271;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	                this.state = 268;
	                this.nl();
	                this.state = 273;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	            }
	            this.state = 276;
	            this._errHandler.sync(this);
	            var la_ = this._interp.adaptivePredict(this._input,27,this._ctx);
	            switch(la_) {
	            case 1:
	                this.state = 274;
	                this.assignCall();
	                break;

	            case 2:
	                this.state = 275;
	                this.exprCall();
	                break;

	            }
	            this.state = 291;
	            this._errHandler.sync(this);
	            var _alt = this._interp.adaptivePredict(this._input,30,this._ctx)
	            while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	                if(_alt===1) {
	                    this.state = 278;
	                    this.match(prqlParser.COMMA);
	                    this.state = 282;
	                    this._errHandler.sync(this);
	                    _la = this._input.LA(1);
	                    while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	                        this.state = 279;
	                        this.nl();
	                        this.state = 284;
	                        this._errHandler.sync(this);
	                        _la = this._input.LA(1);
	                    }
	                    this.state = 287;
	                    this._errHandler.sync(this);
	                    var la_ = this._interp.adaptivePredict(this._input,29,this._ctx);
	                    switch(la_) {
	                    case 1:
	                        this.state = 285;
	                        this.assignCall();
	                        break;

	                    case 2:
	                        this.state = 286;
	                        this.exprCall();
	                        break;

	                    } 
	                }
	                this.state = 293;
	                this._errHandler.sync(this);
	                _alt = this._interp.adaptivePredict(this._input,30,this._ctx);
	            }

	            this.state = 295;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===prqlParser.COMMA) {
	                this.state = 294;
	                this.match(prqlParser.COMMA);
	            }

	            this.state = 298;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	                this.state = 297;
	                this.nl();
	            }

	        }

	        this.state = 302;
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
	        this.state = 304;
	        this.match(prqlParser.LPAREN);
	        this.state = 308;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	            this.state = 305;
	            this.nl();
	            this.state = 310;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 311;
	        this.pipeline();
	        this.state = 315;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	            this.state = 312;
	            this.nl();
	            this.state = 317;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 318;
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



	range() {
	    let localctx = new RangeContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 52, prqlParser.RULE_range);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 320;
	        this.literal();
	        this.state = 321;
	        this.match(prqlParser.RANGE);
	        this.state = 322;
	        this.literal();
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



	operator() {
	    let localctx = new OperatorContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 54, prqlParser.RULE_operator);
	    try {
	        this.state = 330;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,36,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 324;
	            this.operatorUnary();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 325;
	            this.operatorMul();
	            break;

	        case 3:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 326;
	            this.operatorAdd();
	            break;

	        case 4:
	            this.enterOuterAlt(localctx, 4);
	            this.state = 327;
	            this.operatorCompare();
	            break;

	        case 5:
	            this.enterOuterAlt(localctx, 5);
	            this.state = 328;
	            this.operatorLogical();
	            break;

	        case 6:
	            this.enterOuterAlt(localctx, 6);
	            this.state = 329;
	            this.operatorCoalesce();
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



	operatorUnary() {
	    let localctx = new OperatorUnaryContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 56, prqlParser.RULE_operatorUnary);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 332;
	        _la = this._input.LA(1);
	        if(!(((((_la - 15)) & ~0x1f) == 0 && ((1 << (_la - 15)) & ((1 << (prqlParser.PLUS - 15)) | (1 << (prqlParser.MINUS - 15)) | (1 << (prqlParser.NOT - 15)))) !== 0))) {
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



	operatorMul() {
	    let localctx = new OperatorMulContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 58, prqlParser.RULE_operatorMul);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 334;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.STAR) | (1 << prqlParser.DIV) | (1 << prqlParser.MOD))) !== 0))) {
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



	operatorAdd() {
	    let localctx = new OperatorAddContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 60, prqlParser.RULE_operatorAdd);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 336;
	        _la = this._input.LA(1);
	        if(!(_la===prqlParser.PLUS || _la===prqlParser.MINUS)) {
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



	operatorCompare() {
	    let localctx = new OperatorCompareContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 62, prqlParser.RULE_operatorCompare);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 338;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.EQ) | (1 << prqlParser.NE) | (1 << prqlParser.LE) | (1 << prqlParser.GE) | (1 << prqlParser.LANG) | (1 << prqlParser.RANG))) !== 0))) {
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



	operatorLogical() {
	    let localctx = new OperatorLogicalContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 64, prqlParser.RULE_operatorLogical);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 340;
	        _la = this._input.LA(1);
	        if(!(_la===prqlParser.AND || _la===prqlParser.OR)) {
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



	operatorCoalesce() {
	    let localctx = new OperatorCoalesceContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 66, prqlParser.RULE_operatorCoalesce);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 342;
	        this.match(prqlParser.COALESCE);
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



	intervalKind() {
	    let localctx = new IntervalKindContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 68, prqlParser.RULE_intervalKind);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 344;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.T__0) | (1 << prqlParser.T__1) | (1 << prqlParser.T__2) | (1 << prqlParser.T__3) | (1 << prqlParser.T__4) | (1 << prqlParser.T__5) | (1 << prqlParser.T__6) | (1 << prqlParser.T__7) | (1 << prqlParser.T__8))) !== 0))) {
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



	interval() {
	    let localctx = new IntervalContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 70, prqlParser.RULE_interval);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 346;
	        this.match(prqlParser.NUMBER);
	        this.state = 347;
	        this.intervalKind();
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
prqlParser.T__0 = 1;
prqlParser.T__1 = 2;
prqlParser.T__2 = 3;
prqlParser.T__3 = 4;
prqlParser.T__4 = 5;
prqlParser.T__5 = 6;
prqlParser.T__6 = 7;
prqlParser.T__7 = 8;
prqlParser.T__8 = 9;
prqlParser.FUNC = 10;
prqlParser.PRQL = 11;
prqlParser.TABLE = 12;
prqlParser.ARROW = 13;
prqlParser.ASSIGN = 14;
prqlParser.PLUS = 15;
prqlParser.MINUS = 16;
prqlParser.STAR = 17;
prqlParser.DIV = 18;
prqlParser.MOD = 19;
prqlParser.EQ = 20;
prqlParser.NE = 21;
prqlParser.LE = 22;
prqlParser.GE = 23;
prqlParser.BAR = 24;
prqlParser.COLON = 25;
prqlParser.COMMA = 26;
prqlParser.DOT = 27;
prqlParser.DOLLAR = 28;
prqlParser.RANGE = 29;
prqlParser.LANG = 30;
prqlParser.RANG = 31;
prqlParser.LBRACKET = 32;
prqlParser.RBRACKET = 33;
prqlParser.LPAREN = 34;
prqlParser.RPAREN = 35;
prqlParser.UNDERSCORE = 36;
prqlParser.BACKTICK = 37;
prqlParser.DOUBLE_QUOTE = 38;
prqlParser.SINGLE_QUOTE = 39;
prqlParser.TRIPLE_DOUBLE_QUOTE = 40;
prqlParser.TRIPLE_SINGLE_QUOTE = 41;
prqlParser.AND = 42;
prqlParser.OR = 43;
prqlParser.NOT = 44;
prqlParser.COALESCE = 45;
prqlParser.NULL_ = 46;
prqlParser.BOOLEAN = 47;
prqlParser.NUMBER = 48;
prqlParser.IDENT = 49;
prqlParser.IDENT_START = 50;
prqlParser.IDENT_NEXT = 51;
prqlParser.WHITESPACE = 52;
prqlParser.NEWLINE = 53;
prqlParser.COMMENT = 54;
prqlParser.STRING = 55;

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
prqlParser.RULE_identBackticks = 12;
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
prqlParser.RULE_range = 26;
prqlParser.RULE_operator = 27;
prqlParser.RULE_operatorUnary = 28;
prqlParser.RULE_operatorMul = 29;
prqlParser.RULE_operatorAdd = 30;
prqlParser.RULE_operatorCompare = 31;
prqlParser.RULE_operatorLogical = 32;
prqlParser.RULE_operatorCoalesce = 33;
prqlParser.RULE_intervalKind = 34;
prqlParser.RULE_interval = 35;

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



class IdentBackticksContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_identBackticks;
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
	        listener.enterIdentBackticks(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitIdentBackticks(this);
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

	operatorMul() {
	    return this.getTypedRuleContext(OperatorMulContext,0);
	};

	operatorAdd() {
	    return this.getTypedRuleContext(OperatorAddContext,0);
	};

	operatorCompare() {
	    return this.getTypedRuleContext(OperatorCompareContext,0);
	};

	operatorCoalesce() {
	    return this.getTypedRuleContext(OperatorCoalesceContext,0);
	};

	operatorLogical() {
	    return this.getTypedRuleContext(OperatorLogicalContext,0);
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

	range() {
	    return this.getTypedRuleContext(RangeContext,0);
	};

	literal() {
	    return this.getTypedRuleContext(LiteralContext,0);
	};

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	identBackticks() {
	    return this.getTypedRuleContext(IdentBackticksContext,0);
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

	operatorUnary() {
	    return this.getTypedRuleContext(OperatorUnaryContext,0);
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

	interval() {
	    return this.getTypedRuleContext(IntervalContext,0);
	};

	NUMBER() {
	    return this.getToken(prqlParser.NUMBER, 0);
	};

	BOOLEAN() {
	    return this.getToken(prqlParser.BOOLEAN, 0);
	};

	NULL_() {
	    return this.getToken(prqlParser.NULL_, 0);
	};

	STRING() {
	    return this.getToken(prqlParser.STRING, 0);
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



class RangeContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_range;
    }

	literal = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(LiteralContext);
	    } else {
	        return this.getTypedRuleContext(LiteralContext,i);
	    }
	};

	RANGE() {
	    return this.getToken(prqlParser.RANGE, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterRange(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitRange(this);
		}
	}


}



class OperatorContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_operator;
    }

	operatorUnary() {
	    return this.getTypedRuleContext(OperatorUnaryContext,0);
	};

	operatorMul() {
	    return this.getTypedRuleContext(OperatorMulContext,0);
	};

	operatorAdd() {
	    return this.getTypedRuleContext(OperatorAddContext,0);
	};

	operatorCompare() {
	    return this.getTypedRuleContext(OperatorCompareContext,0);
	};

	operatorLogical() {
	    return this.getTypedRuleContext(OperatorLogicalContext,0);
	};

	operatorCoalesce() {
	    return this.getTypedRuleContext(OperatorCoalesceContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterOperator(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitOperator(this);
		}
	}


}



class OperatorUnaryContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_operatorUnary;
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

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterOperatorUnary(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitOperatorUnary(this);
		}
	}


}



class OperatorMulContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_operatorMul;
    }

	STAR() {
	    return this.getToken(prqlParser.STAR, 0);
	};

	DIV() {
	    return this.getToken(prqlParser.DIV, 0);
	};

	MOD() {
	    return this.getToken(prqlParser.MOD, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterOperatorMul(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitOperatorMul(this);
		}
	}


}



class OperatorAddContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_operatorAdd;
    }

	MINUS() {
	    return this.getToken(prqlParser.MINUS, 0);
	};

	PLUS() {
	    return this.getToken(prqlParser.PLUS, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterOperatorAdd(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitOperatorAdd(this);
		}
	}


}



class OperatorCompareContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_operatorCompare;
    }

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

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterOperatorCompare(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitOperatorCompare(this);
		}
	}


}



class OperatorLogicalContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_operatorLogical;
    }

	AND() {
	    return this.getToken(prqlParser.AND, 0);
	};

	OR() {
	    return this.getToken(prqlParser.OR, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterOperatorLogical(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitOperatorLogical(this);
		}
	}


}



class OperatorCoalesceContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_operatorCoalesce;
    }

	COALESCE() {
	    return this.getToken(prqlParser.COALESCE, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterOperatorCoalesce(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitOperatorCoalesce(this);
		}
	}


}



class IntervalKindContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_intervalKind;
    }


	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterIntervalKind(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitIntervalKind(this);
		}
	}


}



class IntervalContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_interval;
    }

	NUMBER() {
	    return this.getToken(prqlParser.NUMBER, 0);
	};

	intervalKind() {
	    return this.getTypedRuleContext(IntervalKindContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterInterval(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitInterval(this);
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
prqlParser.IdentBackticksContext = IdentBackticksContext; 
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
prqlParser.RangeContext = RangeContext; 
prqlParser.OperatorContext = OperatorContext; 
prqlParser.OperatorUnaryContext = OperatorUnaryContext; 
prqlParser.OperatorMulContext = OperatorMulContext; 
prqlParser.OperatorAddContext = OperatorAddContext; 
prqlParser.OperatorCompareContext = OperatorCompareContext; 
prqlParser.OperatorLogicalContext = OperatorLogicalContext; 
prqlParser.OperatorCoalesceContext = OperatorCoalesceContext; 
prqlParser.IntervalKindContext = IntervalKindContext; 
prqlParser.IntervalContext = IntervalContext; 
