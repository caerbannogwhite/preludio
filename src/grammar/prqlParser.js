// Generated from .\prql.g4 by ANTLR 4.9.3
// jshint ignore: start
import antlr4 from 'antlr4';
import prqlListener from './prqlListener.js';

const serializedATN = ["\u0003\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786",
    "\u5964\u00039\u0139\u0004\u0002\t\u0002\u0004\u0003\t\u0003\u0004\u0004",
    "\t\u0004\u0004\u0005\t\u0005\u0004\u0006\t\u0006\u0004\u0007\t\u0007",
    "\u0004\b\t\b\u0004\t\t\t\u0004\n\t\n\u0004\u000b\t\u000b\u0004\f\t\f",
    "\u0004\r\t\r\u0004\u000e\t\u000e\u0004\u000f\t\u000f\u0004\u0010\t\u0010",
    "\u0004\u0011\t\u0011\u0004\u0012\t\u0012\u0004\u0013\t\u0013\u0004\u0014",
    "\t\u0014\u0004\u0015\t\u0015\u0004\u0016\t\u0016\u0004\u0017\t\u0017",
    "\u0004\u0018\t\u0018\u0004\u0019\t\u0019\u0004\u001a\t\u001a\u0004\u001b",
    "\t\u001b\u0004\u001c\t\u001c\u0004\u001d\t\u001d\u0004\u001e\t\u001e",
    "\u0003\u0002\u0003\u0002\u0003\u0003\u0007\u0003@\n\u0003\f\u0003\u000e",
    "\u0003C\u000b\u0003\u0003\u0003\u0005\u0003F\n\u0003\u0003\u0003\u0007",
    "\u0003I\n\u0003\f\u0003\u000e\u0003L\u000b\u0003\u0003\u0003\u0003\u0003",
    "\u0003\u0003\u0005\u0003Q\n\u0003\u0003\u0003\u0007\u0003T\n\u0003\f",
    "\u0003\u000e\u0003W\u000b\u0003\u0007\u0003Y\n\u0003\f\u0003\u000e\u0003",
    "\\\u000b\u0003\u0003\u0003\u0003\u0003\u0003\u0004\u0003\u0004\u0007",
    "\u0004b\n\u0004\f\u0004\u000e\u0004e\u000b\u0004\u0003\u0004\u0003\u0004",
    "\u0003\u0005\u0003\u0005\u0003\u0005\u0003\u0005\u0003\u0005\u0003\u0005",
    "\u0003\u0006\u0003\u0006\u0005\u0006q\n\u0006\u0003\u0007\u0007\u0007",
    "t\n\u0007\f\u0007\u000e\u0007w\u000b\u0007\u0003\b\u0003\b\u0005\b{",
    "\n\b\u0003\b\u0005\b~\n\b\u0003\t\u0003\t\u0003\t\u0003\t\u0007\t\u0084",
    "\n\t\f\t\u000e\t\u0087\u000b\t\u0003\t\u0003\t\u0003\n\u0003\n\u0005",
    "\n\u008d\n\n\u0003\u000b\u0003\u000b\u0003\u000b\u0003\u000b\u0003\u000b",
    "\u0003\f\u0003\f\u0005\f\u0096\n\f\u0003\r\u0003\r\u0003\r\u0003\r\u0007",
    "\r\u009c\n\r\f\r\u000e\r\u009f\u000b\r\u0003\u000e\u0003\u000e\u0007",
    "\u000e\u00a3\n\u000e\f\u000e\u000e\u000e\u00a6\u000b\u000e\u0003\u000e",
    "\u0003\u000e\u0003\u000f\u0003\u000f\u0003\u000f\u0003\u0010\u0003\u0010",
    "\u0003\u0011\u0003\u0011\u0003\u0011\u0003\u0011\u0006\u0011\u00b3\n",
    "\u0011\r\u0011\u000e\u0011\u00b4\u0003\u0012\u0003\u0012\u0003\u0012",
    "\u0003\u0012\u0005\u0012\u00bb\n\u0012\u0003\u0013\u0003\u0013\u0003",
    "\u0013\u0003\u0013\u0003\u0014\u0003\u0014\u0003\u0014\u0003\u0014\u0003",
    "\u0015\u0003\u0015\u0005\u0015\u00c7\n\u0015\u0003\u0016\u0003\u0016",
    "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0005\u0016\u00cf\n",
    "\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003",
    "\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003",
    "\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0007\u0016\u00e0\n\u0016",
    "\f\u0016\u000e\u0016\u00e3\u000b\u0016\u0003\u0017\u0003\u0017\u0003",
    "\u0017\u0003\u0017\u0003\u0017\u0003\u0017\u0003\u0017\u0005\u0017\u00ec",
    "\n\u0017\u0003\u0018\u0003\u0018\u0003\u0018\u0003\u0018\u0005\u0018",
    "\u00f2\n\u0018\u0003\u0019\u0003\u0019\u0003\u0019\u0003\u0019\u0003",
    "\u0019\u0005\u0019\u00f9\n\u0019\u0003\u001a\u0003\u001a\u0007\u001a",
    "\u00fd\n\u001a\f\u001a\u000e\u001a\u0100\u000b\u001a\u0003\u001a\u0003",
    "\u001a\u0005\u001a\u0104\n\u001a\u0003\u001a\u0003\u001a\u0007\u001a",
    "\u0108\n\u001a\f\u001a\u000e\u001a\u010b\u000b\u001a\u0003\u001a\u0003",
    "\u001a\u0005\u001a\u010f\n\u001a\u0007\u001a\u0111\n\u001a\f\u001a\u000e",
    "\u001a\u0114\u000b\u001a\u0003\u001a\u0005\u001a\u0117\n\u001a\u0003",
    "\u001a\u0005\u001a\u011a\n\u001a\u0005\u001a\u011c\n\u001a\u0003\u001a",
    "\u0003\u001a\u0003\u001b\u0003\u001b\u0007\u001b\u0122\n\u001b\f\u001b",
    "\u000e\u001b\u0125\u000b\u001b\u0003\u001b\u0003\u001b\u0007\u001b\u0129",
    "\n\u001b\f\u001b\u000e\u001b\u012c\u000b\u001b\u0003\u001b\u0003\u001b",
    "\u0003\u001c\u0003\u001c\u0003\u001c\u0003\u001c\u0003\u001d\u0003\u001d",
    "\u0003\u001e\u0003\u001e\u0003\u001e\u0003\u001e\u0002\u0003*\u001f",
    "\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c",
    "\u001e \"$&(*,.02468:\u0002\u000b\u0003\u000278\u0004\u0002\'\'77\u0003",
    "\u0002\u0011\u0012\u0003\u0002\f\u000e\u0003\u0002\u0013\u0015\u0004",
    "\u0002\u0016\u0019 !\u0003\u0002,-\u0004\u0002\u0011\u0012..\u0003\u0002",
    "\u0003\u000b\u0002\u014d\u0002<\u0003\u0002\u0002\u0002\u0004A\u0003",
    "\u0002\u0002\u0002\u0006_\u0003\u0002\u0002\u0002\bh\u0003\u0002\u0002",
    "\u0002\nn\u0003\u0002\u0002\u0002\fu\u0003\u0002\u0002\u0002\u000ez",
    "\u0003\u0002\u0002\u0002\u0010\u007f\u0003\u0002\u0002\u0002\u0012\u008a",
    "\u0003\u0002\u0002\u0002\u0014\u008e\u0003\u0002\u0002\u0002\u0016\u0095",
    "\u0003\u0002\u0002\u0002\u0018\u0097\u0003\u0002\u0002\u0002\u001a\u00a0",
    "\u0003\u0002\u0002\u0002\u001c\u00a9\u0003\u0002\u0002\u0002\u001e\u00ac",
    "\u0003\u0002\u0002\u0002 \u00ae\u0003\u0002\u0002\u0002\"\u00b6\u0003",
    "\u0002\u0002\u0002$\u00bc\u0003\u0002\u0002\u0002&\u00c0\u0003\u0002",
    "\u0002\u0002(\u00c6\u0003\u0002\u0002\u0002*\u00ce\u0003\u0002\u0002",
    "\u0002,\u00eb\u0003\u0002\u0002\u0002.\u00ed\u0003\u0002\u0002\u0002",
    "0\u00f8\u0003\u0002\u0002\u00022\u00fa\u0003\u0002\u0002\u00024\u011f",
    "\u0003\u0002\u0002\u00026\u012f\u0003\u0002\u0002\u00028\u0133\u0003",
    "\u0002\u0002\u0002:\u0135\u0003\u0002\u0002\u0002<=\t\u0002\u0002\u0002",
    "=\u0003\u0003\u0002\u0002\u0002>@\u0005\u0002\u0002\u0002?>\u0003\u0002",
    "\u0002\u0002@C\u0003\u0002\u0002\u0002A?\u0003\u0002\u0002\u0002AB\u0003",
    "\u0002\u0002\u0002BE\u0003\u0002\u0002\u0002CA\u0003\u0002\u0002\u0002",
    "DF\u0005\u0006\u0004\u0002ED\u0003\u0002\u0002\u0002EF\u0003\u0002\u0002",
    "\u0002FJ\u0003\u0002\u0002\u0002GI\u0005\u0002\u0002\u0002HG\u0003\u0002",
    "\u0002\u0002IL\u0003\u0002\u0002\u0002JH\u0003\u0002\u0002\u0002JK\u0003",
    "\u0002\u0002\u0002KZ\u0003\u0002\u0002\u0002LJ\u0003\u0002\u0002\u0002",
    "MQ\u0005\b\u0005\u0002NQ\u0005\u0014\u000b\u0002OQ\u0005\u0018\r\u0002",
    "PM\u0003\u0002\u0002\u0002PN\u0003\u0002\u0002\u0002PO\u0003\u0002\u0002",
    "\u0002QU\u0003\u0002\u0002\u0002RT\u0005\u0002\u0002\u0002SR\u0003\u0002",
    "\u0002\u0002TW\u0003\u0002\u0002\u0002US\u0003\u0002\u0002\u0002UV\u0003",
    "\u0002\u0002\u0002VY\u0003\u0002\u0002\u0002WU\u0003\u0002\u0002\u0002",
    "XP\u0003\u0002\u0002\u0002Y\\\u0003\u0002\u0002\u0002ZX\u0003\u0002",
    "\u0002\u0002Z[\u0003\u0002\u0002\u0002[]\u0003\u0002\u0002\u0002\\Z",
    "\u0003\u0002\u0002\u0002]^\u0007\u0002\u0002\u0003^\u0005\u0003\u0002",
    "\u0002\u0002_c\u0007\r\u0002\u0002`b\u0005\"\u0012\u0002a`\u0003\u0002",
    "\u0002\u0002be\u0003\u0002\u0002\u0002ca\u0003\u0002\u0002\u0002cd\u0003",
    "\u0002\u0002\u0002df\u0003\u0002\u0002\u0002ec\u0003\u0002\u0002\u0002",
    "fg\u0005\u0002\u0002\u0002g\u0007\u0003\u0002\u0002\u0002hi\u0007\f",
    "\u0002\u0002ij\u0005\n\u0006\u0002jk\u0005\f\u0007\u0002kl\u0007\u000f",
    "\u0002\u0002lm\u0005*\u0016\u0002m\t\u0003\u0002\u0002\u0002np\u0007",
    "3\u0002\u0002oq\u0005\u0010\t\u0002po\u0003\u0002\u0002\u0002pq\u0003",
    "\u0002\u0002\u0002q\u000b\u0003\u0002\u0002\u0002rt\u0005\u000e\b\u0002",
    "sr\u0003\u0002\u0002\u0002tw\u0003\u0002\u0002\u0002us\u0003\u0002\u0002",
    "\u0002uv\u0003\u0002\u0002\u0002v\r\u0003\u0002\u0002\u0002wu\u0003",
    "\u0002\u0002\u0002x{\u0005\"\u0012\u0002y{\u00073\u0002\u0002zx\u0003",
    "\u0002\u0002\u0002zy\u0003\u0002\u0002\u0002{}\u0003\u0002\u0002\u0002",
    "|~\u0005\u0010\t\u0002}|\u0003\u0002\u0002\u0002}~\u0003\u0002\u0002",
    "\u0002~\u000f\u0003\u0002\u0002\u0002\u007f\u0080\u0007 \u0002\u0002",
    "\u0080\u0081\u0005\u0012\n\u0002\u0081\u0085\u0007\u001a\u0002\u0002",
    "\u0082\u0084\u0005\u0012\n\u0002\u0083\u0082\u0003\u0002\u0002\u0002",
    "\u0084\u0087\u0003\u0002\u0002\u0002\u0085\u0083\u0003\u0002\u0002\u0002",
    "\u0085\u0086\u0003\u0002\u0002\u0002\u0086\u0088\u0003\u0002\u0002\u0002",
    "\u0087\u0085\u0003\u0002\u0002\u0002\u0088\u0089\u0007!\u0002\u0002",
    "\u0089\u0011\u0003\u0002\u0002\u0002\u008a\u008c\u00073\u0002\u0002",
    "\u008b\u008d\u0005\u0010\t\u0002\u008c\u008b\u0003\u0002\u0002\u0002",
    "\u008c\u008d\u0003\u0002\u0002\u0002\u008d\u0013\u0003\u0002\u0002\u0002",
    "\u008e\u008f\u0007\u000e\u0002\u0002\u008f\u0090\u00073\u0002\u0002",
    "\u0090\u0091\u0007\u0010\u0002\u0002\u0091\u0092\u00054\u001b\u0002",
    "\u0092\u0015\u0003\u0002\u0002\u0002\u0093\u0096\u0005\u0002\u0002\u0002",
    "\u0094\u0096\u0007\u001a\u0002\u0002\u0095\u0093\u0003\u0002\u0002\u0002",
    "\u0095\u0094\u0003\u0002\u0002\u0002\u0096\u0017\u0003\u0002\u0002\u0002",
    "\u0097\u009d\u0005(\u0015\u0002\u0098\u0099\u0005\u0016\f\u0002\u0099",
    "\u009a\u0005(\u0015\u0002\u009a\u009c\u0003\u0002\u0002\u0002\u009b",
    "\u0098\u0003\u0002\u0002\u0002\u009c\u009f\u0003\u0002\u0002\u0002\u009d",
    "\u009b\u0003\u0002\u0002\u0002\u009d\u009e\u0003\u0002\u0002\u0002\u009e",
    "\u0019\u0003\u0002\u0002\u0002\u009f\u009d\u0003\u0002\u0002\u0002\u00a0",
    "\u00a4\u0007\'\u0002\u0002\u00a1\u00a3\n\u0003\u0002\u0002\u00a2\u00a1",
    "\u0003\u0002\u0002\u0002\u00a3\u00a6\u0003\u0002\u0002\u0002\u00a4\u00a2",
    "\u0003\u0002\u0002\u0002\u00a4\u00a5\u0003\u0002\u0002\u0002\u00a5\u00a7",
    "\u0003\u0002\u0002\u0002\u00a6\u00a4\u0003\u0002\u0002\u0002\u00a7\u00a8",
    "\u0007\'\u0002\u0002\u00a8\u001b\u0003\u0002\u0002\u0002\u00a9\u00aa",
    "\t\u0004\u0002\u0002\u00aa\u00ab\u00073\u0002\u0002\u00ab\u001d\u0003",
    "\u0002\u0002\u0002\u00ac\u00ad\t\u0005\u0002\u0002\u00ad\u001f\u0003",
    "\u0002\u0002\u0002\u00ae\u00b2\u00073\u0002\u0002\u00af\u00b3\u0005",
    "\"\u0012\u0002\u00b0\u00b3\u0005$\u0013\u0002\u00b1\u00b3\u0005*\u0016",
    "\u0002\u00b2\u00af\u0003\u0002\u0002\u0002\u00b2\u00b0\u0003\u0002\u0002",
    "\u0002\u00b2\u00b1\u0003\u0002\u0002\u0002\u00b3\u00b4\u0003\u0002\u0002",
    "\u0002\u00b4\u00b2\u0003\u0002\u0002\u0002\u00b4\u00b5\u0003\u0002\u0002",
    "\u0002\u00b5!\u0003\u0002\u0002\u0002\u00b6\u00b7\u00073\u0002\u0002",
    "\u00b7\u00ba\u0007\u001b\u0002\u0002\u00b8\u00bb\u0005$\u0013\u0002",
    "\u00b9\u00bb\u0005*\u0016\u0002\u00ba\u00b8\u0003\u0002\u0002\u0002",
    "\u00ba\u00b9\u0003\u0002\u0002\u0002\u00bb#\u0003\u0002\u0002\u0002",
    "\u00bc\u00bd\u00073\u0002\u0002\u00bd\u00be\u0007\u0010\u0002\u0002",
    "\u00be\u00bf\u0005*\u0016\u0002\u00bf%\u0003\u0002\u0002\u0002\u00c0",
    "\u00c1\u00073\u0002\u0002\u00c1\u00c2\u0007\u0010\u0002\u0002\u00c2",
    "\u00c3\u0005(\u0015\u0002\u00c3\'\u0003\u0002\u0002\u0002\u00c4\u00c7",
    "\u0005 \u0011\u0002\u00c5\u00c7\u0005*\u0016\u0002\u00c6\u00c4\u0003",
    "\u0002\u0002\u0002\u00c6\u00c5\u0003\u0002\u0002\u0002\u00c7)\u0003",
    "\u0002\u0002\u0002\u00c8\u00c9\b\u0016\u0001\u0002\u00c9\u00ca\u0007",
    "$\u0002\u0002\u00ca\u00cb\u0005*\u0016\u0002\u00cb\u00cc\u0007%\u0002",
    "\u0002\u00cc\u00cf\u0003\u0002\u0002\u0002\u00cd\u00cf\u0005,\u0017",
    "\u0002\u00ce\u00c8\u0003\u0002\u0002\u0002\u00ce\u00cd\u0003\u0002\u0002",
    "\u0002\u00cf\u00e1\u0003\u0002\u0002\u0002\u00d0\u00d1\f\t\u0002\u0002",
    "\u00d1\u00d2\t\u0006\u0002\u0002\u00d2\u00e0\u0005*\u0016\n\u00d3\u00d4",
    "\f\b\u0002\u0002\u00d4\u00d5\t\u0004\u0002\u0002\u00d5\u00e0\u0005*",
    "\u0016\t\u00d6\u00d7\f\u0007\u0002\u0002\u00d7\u00d8\t\u0007\u0002\u0002",
    "\u00d8\u00e0\u0005*\u0016\b\u00d9\u00da\f\u0006\u0002\u0002\u00da\u00db",
    "\u0007/\u0002\u0002\u00db\u00e0\u0005*\u0016\u0007\u00dc\u00dd\f\u0005",
    "\u0002\u0002\u00dd\u00de\t\b\u0002\u0002\u00de\u00e0\u0005*\u0016\u0006",
    "\u00df\u00d0\u0003\u0002\u0002\u0002\u00df\u00d3\u0003\u0002\u0002\u0002",
    "\u00df\u00d6\u0003\u0002\u0002\u0002\u00df\u00d9\u0003\u0002\u0002\u0002",
    "\u00df\u00dc\u0003\u0002\u0002\u0002\u00e0\u00e3\u0003\u0002\u0002\u0002",
    "\u00e1\u00df\u0003\u0002\u0002\u0002\u00e1\u00e2\u0003\u0002\u0002\u0002",
    "\u00e2+\u0003\u0002\u0002\u0002\u00e3\u00e1\u0003\u0002\u0002\u0002",
    "\u00e4\u00ec\u00056\u001c\u0002\u00e5\u00ec\u00050\u0019\u0002\u00e6",
    "\u00ec\u00073\u0002\u0002\u00e7\u00ec\u0005\u001a\u000e\u0002\u00e8",
    "\u00ec\u0005.\u0018\u0002\u00e9\u00ec\u00052\u001a\u0002\u00ea\u00ec",
    "\u00054\u001b\u0002\u00eb\u00e4\u0003\u0002\u0002\u0002\u00eb\u00e5",
    "\u0003\u0002\u0002\u0002\u00eb\u00e6\u0003\u0002\u0002\u0002\u00eb\u00e7",
    "\u0003\u0002\u0002\u0002\u00eb\u00e8\u0003\u0002\u0002\u0002\u00eb\u00e9",
    "\u0003\u0002\u0002\u0002\u00eb\u00ea\u0003\u0002\u0002\u0002\u00ec-",
    "\u0003\u0002\u0002\u0002\u00ed\u00f1\t\t\u0002\u0002\u00ee\u00f2\u0005",
    "4\u001b\u0002\u00ef\u00f2\u00050\u0019\u0002\u00f0\u00f2\u00073\u0002",
    "\u0002\u00f1\u00ee\u0003\u0002\u0002\u0002\u00f1\u00ef\u0003\u0002\u0002",
    "\u0002\u00f1\u00f0\u0003\u0002\u0002\u0002\u00f2/\u0003\u0002\u0002",
    "\u0002\u00f3\u00f9\u0005:\u001e\u0002\u00f4\u00f9\u00072\u0002\u0002",
    "\u00f5\u00f9\u00071\u0002\u0002\u00f6\u00f9\u00070\u0002\u0002\u00f7",
    "\u00f9\u00079\u0002\u0002\u00f8\u00f3\u0003\u0002\u0002\u0002\u00f8",
    "\u00f4\u0003\u0002\u0002\u0002\u00f8\u00f5\u0003\u0002\u0002\u0002\u00f8",
    "\u00f6\u0003\u0002\u0002\u0002\u00f8\u00f7\u0003\u0002\u0002\u0002\u00f9",
    "1\u0003\u0002\u0002\u0002\u00fa\u011b\u0007\"\u0002\u0002\u00fb\u00fd",
    "\u0005\u0002\u0002\u0002\u00fc\u00fb\u0003\u0002\u0002\u0002\u00fd\u0100",
    "\u0003\u0002\u0002\u0002\u00fe\u00fc\u0003\u0002\u0002\u0002\u00fe\u00ff",
    "\u0003\u0002\u0002\u0002\u00ff\u0103\u0003\u0002\u0002\u0002\u0100\u00fe",
    "\u0003\u0002\u0002\u0002\u0101\u0104\u0005&\u0014\u0002\u0102\u0104",
    "\u0005(\u0015\u0002\u0103\u0101\u0003\u0002\u0002\u0002\u0103\u0102",
    "\u0003\u0002\u0002\u0002\u0104\u0112\u0003\u0002\u0002\u0002\u0105\u0109",
    "\u0007\u001c\u0002\u0002\u0106\u0108\u0005\u0002\u0002\u0002\u0107\u0106",
    "\u0003\u0002\u0002\u0002\u0108\u010b\u0003\u0002\u0002\u0002\u0109\u0107",
    "\u0003\u0002\u0002\u0002\u0109\u010a\u0003\u0002\u0002\u0002\u010a\u010e",
    "\u0003\u0002\u0002\u0002\u010b\u0109\u0003\u0002\u0002\u0002\u010c\u010f",
    "\u0005&\u0014\u0002\u010d\u010f\u0005(\u0015\u0002\u010e\u010c\u0003",
    "\u0002\u0002\u0002\u010e\u010d\u0003\u0002\u0002\u0002\u010f\u0111\u0003",
    "\u0002\u0002\u0002\u0110\u0105\u0003\u0002\u0002\u0002\u0111\u0114\u0003",
    "\u0002\u0002\u0002\u0112\u0110\u0003\u0002\u0002\u0002\u0112\u0113\u0003",
    "\u0002\u0002\u0002\u0113\u0116\u0003\u0002\u0002\u0002\u0114\u0112\u0003",
    "\u0002\u0002\u0002\u0115\u0117\u0007\u001c\u0002\u0002\u0116\u0115\u0003",
    "\u0002\u0002\u0002\u0116\u0117\u0003\u0002\u0002\u0002\u0117\u0119\u0003",
    "\u0002\u0002\u0002\u0118\u011a\u0005\u0002\u0002\u0002\u0119\u0118\u0003",
    "\u0002\u0002\u0002\u0119\u011a\u0003\u0002\u0002\u0002\u011a\u011c\u0003",
    "\u0002\u0002\u0002\u011b\u00fe\u0003\u0002\u0002\u0002\u011b\u011c\u0003",
    "\u0002\u0002\u0002\u011c\u011d\u0003\u0002\u0002\u0002\u011d\u011e\u0007",
    "#\u0002\u0002\u011e3\u0003\u0002\u0002\u0002\u011f\u0123\u0007$\u0002",
    "\u0002\u0120\u0122\u0005\u0002\u0002\u0002\u0121\u0120\u0003\u0002\u0002",
    "\u0002\u0122\u0125\u0003\u0002\u0002\u0002\u0123\u0121\u0003\u0002\u0002",
    "\u0002\u0123\u0124\u0003\u0002\u0002\u0002\u0124\u0126\u0003\u0002\u0002",
    "\u0002\u0125\u0123\u0003\u0002\u0002\u0002\u0126\u012a\u0005\u0018\r",
    "\u0002\u0127\u0129\u0005\u0002\u0002\u0002\u0128\u0127\u0003\u0002\u0002",
    "\u0002\u0129\u012c\u0003\u0002\u0002\u0002\u012a\u0128\u0003\u0002\u0002",
    "\u0002\u012a\u012b\u0003\u0002\u0002\u0002\u012b\u012d\u0003\u0002\u0002",
    "\u0002\u012c\u012a\u0003\u0002\u0002\u0002\u012d\u012e\u0007%\u0002",
    "\u0002\u012e5\u0003\u0002\u0002\u0002\u012f\u0130\u00050\u0019\u0002",
    "\u0130\u0131\u0007\u001f\u0002\u0002\u0131\u0132\u00050\u0019\u0002",
    "\u01327\u0003\u0002\u0002\u0002\u0133\u0134\t\n\u0002\u0002\u01349\u0003",
    "\u0002\u0002\u0002\u0135\u0136\u00072\u0002\u0002\u0136\u0137\u0005",
    "8\u001d\u0002\u0137;\u0003\u0002\u0002\u0002&AEJPUZcpuz}\u0085\u008c",
    "\u0095\u009d\u00a4\u00b2\u00b4\u00ba\u00c6\u00ce\u00df\u00e1\u00eb\u00f1",
    "\u00f8\u00fe\u0103\u0109\u010e\u0112\u0116\u0119\u011b\u0123\u012a"].join("");


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
                         "range", "intervalKind", "interval" ];

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
	        this.state = 58;
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
	        this.state = 63;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,0,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                this.state = 60;
	                this.nl(); 
	            }
	            this.state = 65;
	            this._errHandler.sync(this);
	            _alt = this._interp.adaptivePredict(this._input,0,this._ctx);
	        }

	        this.state = 67;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.PRQL) {
	            this.state = 66;
	            this.queryDef();
	        }

	        this.state = 72;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	            this.state = 69;
	            this.nl();
	            this.state = 74;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 88;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.FUNC) | (1 << prqlParser.TABLE) | (1 << prqlParser.PLUS) | (1 << prqlParser.MINUS))) !== 0) || ((((_la - 32)) & ~0x1f) == 0 && ((1 << (_la - 32)) & ((1 << (prqlParser.LBRACKET - 32)) | (1 << (prqlParser.LPAREN - 32)) | (1 << (prqlParser.BACKTICK - 32)) | (1 << (prqlParser.NOT - 32)) | (1 << (prqlParser.NULL_ - 32)) | (1 << (prqlParser.BOOLEAN - 32)) | (1 << (prqlParser.NUMBER - 32)) | (1 << (prqlParser.IDENT - 32)) | (1 << (prqlParser.STRING - 32)))) !== 0)) {
	            this.state = 78;
	            this._errHandler.sync(this);
	            switch(this._input.LA(1)) {
	            case prqlParser.FUNC:
	                this.state = 75;
	                this.funcDef();
	                break;
	            case prqlParser.TABLE:
	                this.state = 76;
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
	                this.state = 77;
	                this.pipeline();
	                break;
	            default:
	                throw new antlr4.error.NoViableAltException(this);
	            }
	            this.state = 83;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	                this.state = 80;
	                this.nl();
	                this.state = 85;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	            }
	            this.state = 90;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 91;
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
	        this.state = 93;
	        this.match(prqlParser.PRQL);
	        this.state = 97;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.IDENT) {
	            this.state = 94;
	            this.namedArg();
	            this.state = 99;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 100;
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
	        this.state = 102;
	        this.match(prqlParser.FUNC);
	        this.state = 103;
	        this.funcDefName();
	        this.state = 104;
	        this.funcDefParams();
	        this.state = 105;
	        this.match(prqlParser.ARROW);
	        this.state = 106;
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
	        this.state = 108;
	        this.match(prqlParser.IDENT);
	        this.state = 110;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.LANG) {
	            this.state = 109;
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
	        this.state = 115;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.IDENT) {
	            this.state = 112;
	            this.funcDefParam();
	            this.state = 117;
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
	        this.state = 120;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,9,this._ctx);
	        switch(la_) {
	        case 1:
	            this.state = 118;
	            this.namedArg();
	            break;

	        case 2:
	            this.state = 119;
	            this.match(prqlParser.IDENT);
	            break;

	        }
	        this.state = 123;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.LANG) {
	            this.state = 122;
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
	        this.state = 125;
	        this.match(prqlParser.LANG);
	        this.state = 126;
	        this.typeTerm();
	        this.state = 127;
	        this.match(prqlParser.BAR);
	        this.state = 131;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.IDENT) {
	            this.state = 128;
	            this.typeTerm();
	            this.state = 133;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 134;
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
	        this.state = 136;
	        this.match(prqlParser.IDENT);
	        this.state = 138;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.LANG) {
	            this.state = 137;
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
	        this.state = 140;
	        this.match(prqlParser.TABLE);
	        this.state = 141;
	        this.match(prqlParser.IDENT);
	        this.state = 142;
	        this.match(prqlParser.ASSIGN);
	        this.state = 143;
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
	        this.state = 147;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case prqlParser.NEWLINE:
	        case prqlParser.COMMENT:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 145;
	            this.nl();
	            break;
	        case prqlParser.BAR:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 146;
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
	        this.state = 149;
	        this.exprCall();
	        this.state = 155;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,14,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                this.state = 150;
	                this.pipe();
	                this.state = 151;
	                this.exprCall(); 
	            }
	            this.state = 157;
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
	        this.state = 158;
	        this.match(prqlParser.BACKTICK);
	        this.state = 162;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.T__0) | (1 << prqlParser.T__1) | (1 << prqlParser.T__2) | (1 << prqlParser.T__3) | (1 << prqlParser.T__4) | (1 << prqlParser.T__5) | (1 << prqlParser.T__6) | (1 << prqlParser.T__7) | (1 << prqlParser.T__8) | (1 << prqlParser.FUNC) | (1 << prqlParser.PRQL) | (1 << prqlParser.TABLE) | (1 << prqlParser.ARROW) | (1 << prqlParser.ASSIGN) | (1 << prqlParser.PLUS) | (1 << prqlParser.MINUS) | (1 << prqlParser.STAR) | (1 << prqlParser.DIV) | (1 << prqlParser.MOD) | (1 << prqlParser.EQ) | (1 << prqlParser.NE) | (1 << prqlParser.LE) | (1 << prqlParser.GE) | (1 << prqlParser.BAR) | (1 << prqlParser.COLON) | (1 << prqlParser.COMMA) | (1 << prqlParser.DOT) | (1 << prqlParser.DOLLAR) | (1 << prqlParser.RANGE) | (1 << prqlParser.LANG) | (1 << prqlParser.RANG))) !== 0) || ((((_la - 32)) & ~0x1f) == 0 && ((1 << (_la - 32)) & ((1 << (prqlParser.LBRACKET - 32)) | (1 << (prqlParser.RBRACKET - 32)) | (1 << (prqlParser.LPAREN - 32)) | (1 << (prqlParser.RPAREN - 32)) | (1 << (prqlParser.UNDERSCORE - 32)) | (1 << (prqlParser.DOUBLE_QUOTE - 32)) | (1 << (prqlParser.SINGLE_QUOTE - 32)) | (1 << (prqlParser.TRIPLE_DOUBLE_QUOTE - 32)) | (1 << (prqlParser.TRIPLE_SINGLE_QUOTE - 32)) | (1 << (prqlParser.AND - 32)) | (1 << (prqlParser.OR - 32)) | (1 << (prqlParser.NOT - 32)) | (1 << (prqlParser.COALESCE - 32)) | (1 << (prqlParser.NULL_ - 32)) | (1 << (prqlParser.BOOLEAN - 32)) | (1 << (prqlParser.NUMBER - 32)) | (1 << (prqlParser.IDENT - 32)) | (1 << (prqlParser.IDENT_START - 32)) | (1 << (prqlParser.IDENT_NEXT - 32)) | (1 << (prqlParser.WHITESPACE - 32)) | (1 << (prqlParser.COMMENT - 32)) | (1 << (prqlParser.STRING - 32)))) !== 0)) {
	            this.state = 159;
	            _la = this._input.LA(1);
	            if(_la<=0 || _la===prqlParser.BACKTICK || _la===prqlParser.NEWLINE) {
	            this._errHandler.recoverInline(this);
	            }
	            else {
	            	this._errHandler.reportMatch(this);
	                this.consume();
	            }
	            this.state = 164;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 165;
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
	        this.state = 167;
	        _la = this._input.LA(1);
	        if(!(_la===prqlParser.PLUS || _la===prqlParser.MINUS)) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 168;
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
	        this.state = 170;
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
	        this.state = 172;
	        this.match(prqlParser.IDENT);
	        this.state = 176; 
	        this._errHandler.sync(this);
	        var _alt = 1;
	        do {
	        	switch (_alt) {
	        	case 1:
	        		this.state = 176;
	        		this._errHandler.sync(this);
	        		var la_ = this._interp.adaptivePredict(this._input,16,this._ctx);
	        		switch(la_) {
	        		case 1:
	        		    this.state = 173;
	        		    this.namedArg();
	        		    break;

	        		case 2:
	        		    this.state = 174;
	        		    this.assign();
	        		    break;

	        		case 3:
	        		    this.state = 175;
	        		    this.expr(0);
	        		    break;

	        		}
	        		break;
	        	default:
	        		throw new antlr4.error.NoViableAltException(this);
	        	}
	        	this.state = 178; 
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
	        this.state = 180;
	        this.match(prqlParser.IDENT);
	        this.state = 181;
	        this.match(prqlParser.COLON);
	        this.state = 184;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,18,this._ctx);
	        switch(la_) {
	        case 1:
	            this.state = 182;
	            this.assign();
	            break;

	        case 2:
	            this.state = 183;
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
	        this.state = 186;
	        this.match(prqlParser.IDENT);
	        this.state = 187;
	        this.match(prqlParser.ASSIGN);
	        this.state = 188;
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
	        this.state = 190;
	        this.match(prqlParser.IDENT);
	        this.state = 191;
	        this.match(prqlParser.ASSIGN);
	        this.state = 192;
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
	        this.state = 196;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,19,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 194;
	            this.funcCall();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 195;
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
	        this.state = 204;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,20,this._ctx);
	        switch(la_) {
	        case 1:
	            this.state = 199;
	            this.match(prqlParser.LPAREN);
	            this.state = 200;
	            this.expr(0);
	            this.state = 201;
	            this.match(prqlParser.RPAREN);
	            break;

	        case 2:
	            this.state = 203;
	            this.term();
	            break;

	        }
	        this._ctx.stop = this._input.LT(-1);
	        this.state = 223;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,22,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                if(this._parseListeners!==null) {
	                    this.triggerExitRuleEvent();
	                }
	                _prevctx = localctx;
	                this.state = 221;
	                this._errHandler.sync(this);
	                var la_ = this._interp.adaptivePredict(this._input,21,this._ctx);
	                switch(la_) {
	                case 1:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 206;
	                    if (!( this.precpred(this._ctx, 7))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 7)");
	                    }
	                    this.state = 207;
	                    _la = this._input.LA(1);
	                    if(!((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.STAR) | (1 << prqlParser.DIV) | (1 << prqlParser.MOD))) !== 0))) {
	                    this._errHandler.recoverInline(this);
	                    }
	                    else {
	                    	this._errHandler.reportMatch(this);
	                        this.consume();
	                    }
	                    this.state = 208;
	                    this.expr(8);
	                    break;

	                case 2:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 209;
	                    if (!( this.precpred(this._ctx, 6))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 6)");
	                    }
	                    this.state = 210;
	                    _la = this._input.LA(1);
	                    if(!(_la===prqlParser.PLUS || _la===prqlParser.MINUS)) {
	                    this._errHandler.recoverInline(this);
	                    }
	                    else {
	                    	this._errHandler.reportMatch(this);
	                        this.consume();
	                    }
	                    this.state = 211;
	                    this.expr(7);
	                    break;

	                case 3:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 212;
	                    if (!( this.precpred(this._ctx, 5))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 5)");
	                    }
	                    this.state = 213;
	                    _la = this._input.LA(1);
	                    if(!((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.EQ) | (1 << prqlParser.NE) | (1 << prqlParser.LE) | (1 << prqlParser.GE) | (1 << prqlParser.LANG) | (1 << prqlParser.RANG))) !== 0))) {
	                    this._errHandler.recoverInline(this);
	                    }
	                    else {
	                    	this._errHandler.reportMatch(this);
	                        this.consume();
	                    }
	                    this.state = 214;
	                    this.expr(6);
	                    break;

	                case 4:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 215;
	                    if (!( this.precpred(this._ctx, 4))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 4)");
	                    }
	                    this.state = 216;
	                    this.match(prqlParser.COALESCE);
	                    this.state = 217;
	                    this.expr(5);
	                    break;

	                case 5:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 218;
	                    if (!( this.precpred(this._ctx, 3))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 3)");
	                    }
	                    this.state = 219;
	                    _la = this._input.LA(1);
	                    if(!(_la===prqlParser.AND || _la===prqlParser.OR)) {
	                    this._errHandler.recoverInline(this);
	                    }
	                    else {
	                    	this._errHandler.reportMatch(this);
	                        this.consume();
	                    }
	                    this.state = 220;
	                    this.expr(4);
	                    break;

	                } 
	            }
	            this.state = 225;
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
	        this.state = 233;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,23,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 226;
	            this.range();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 227;
	            this.literal();
	            break;

	        case 3:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 228;
	            this.match(prqlParser.IDENT);
	            break;

	        case 4:
	            this.enterOuterAlt(localctx, 4);
	            this.state = 229;
	            this.identBackticks();
	            break;

	        case 5:
	            this.enterOuterAlt(localctx, 5);
	            this.state = 230;
	            this.exprUnary();
	            break;

	        case 6:
	            this.enterOuterAlt(localctx, 6);
	            this.state = 231;
	            this.list();
	            break;

	        case 7:
	            this.enterOuterAlt(localctx, 7);
	            this.state = 232;
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
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 235;
	        _la = this._input.LA(1);
	        if(!(((((_la - 15)) & ~0x1f) == 0 && ((1 << (_la - 15)) & ((1 << (prqlParser.PLUS - 15)) | (1 << (prqlParser.MINUS - 15)) | (1 << (prqlParser.NOT - 15)))) !== 0))) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 239;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case prqlParser.LPAREN:
	            this.state = 236;
	            this.nestedPipeline();
	            break;
	        case prqlParser.NULL_:
	        case prqlParser.BOOLEAN:
	        case prqlParser.NUMBER:
	        case prqlParser.STRING:
	            this.state = 237;
	            this.literal();
	            break;
	        case prqlParser.IDENT:
	            this.state = 238;
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
	        this.state = 246;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,25,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 241;
	            this.interval();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 242;
	            this.match(prqlParser.NUMBER);
	            break;

	        case 3:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 243;
	            this.match(prqlParser.BOOLEAN);
	            break;

	        case 4:
	            this.enterOuterAlt(localctx, 4);
	            this.state = 244;
	            this.match(prqlParser.NULL_);
	            break;

	        case 5:
	            this.enterOuterAlt(localctx, 5);
	            this.state = 245;
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
	        this.state = 248;
	        this.match(prqlParser.LBRACKET);
	        this.state = 281;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.PLUS || _la===prqlParser.MINUS || ((((_la - 32)) & ~0x1f) == 0 && ((1 << (_la - 32)) & ((1 << (prqlParser.LBRACKET - 32)) | (1 << (prqlParser.LPAREN - 32)) | (1 << (prqlParser.BACKTICK - 32)) | (1 << (prqlParser.NOT - 32)) | (1 << (prqlParser.NULL_ - 32)) | (1 << (prqlParser.BOOLEAN - 32)) | (1 << (prqlParser.NUMBER - 32)) | (1 << (prqlParser.IDENT - 32)) | (1 << (prqlParser.NEWLINE - 32)) | (1 << (prqlParser.COMMENT - 32)) | (1 << (prqlParser.STRING - 32)))) !== 0)) {
	            this.state = 252;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	                this.state = 249;
	                this.nl();
	                this.state = 254;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	            }
	            this.state = 257;
	            this._errHandler.sync(this);
	            var la_ = this._interp.adaptivePredict(this._input,27,this._ctx);
	            switch(la_) {
	            case 1:
	                this.state = 255;
	                this.assignCall();
	                break;

	            case 2:
	                this.state = 256;
	                this.exprCall();
	                break;

	            }
	            this.state = 272;
	            this._errHandler.sync(this);
	            var _alt = this._interp.adaptivePredict(this._input,30,this._ctx)
	            while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	                if(_alt===1) {
	                    this.state = 259;
	                    this.match(prqlParser.COMMA);
	                    this.state = 263;
	                    this._errHandler.sync(this);
	                    _la = this._input.LA(1);
	                    while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	                        this.state = 260;
	                        this.nl();
	                        this.state = 265;
	                        this._errHandler.sync(this);
	                        _la = this._input.LA(1);
	                    }
	                    this.state = 268;
	                    this._errHandler.sync(this);
	                    var la_ = this._interp.adaptivePredict(this._input,29,this._ctx);
	                    switch(la_) {
	                    case 1:
	                        this.state = 266;
	                        this.assignCall();
	                        break;

	                    case 2:
	                        this.state = 267;
	                        this.exprCall();
	                        break;

	                    } 
	                }
	                this.state = 274;
	                this._errHandler.sync(this);
	                _alt = this._interp.adaptivePredict(this._input,30,this._ctx);
	            }

	            this.state = 276;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===prqlParser.COMMA) {
	                this.state = 275;
	                this.match(prqlParser.COMMA);
	            }

	            this.state = 279;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	                this.state = 278;
	                this.nl();
	            }

	        }

	        this.state = 283;
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
	        this.state = 285;
	        this.match(prqlParser.LPAREN);
	        this.state = 289;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	            this.state = 286;
	            this.nl();
	            this.state = 291;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 292;
	        this.pipeline();
	        this.state = 296;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	            this.state = 293;
	            this.nl();
	            this.state = 298;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 299;
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
	        this.state = 301;
	        this.literal();
	        this.state = 302;
	        this.match(prqlParser.RANGE);
	        this.state = 303;
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



	intervalKind() {
	    let localctx = new IntervalKindContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 54, prqlParser.RULE_intervalKind);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 305;
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
	    this.enterRule(localctx, 56, prqlParser.RULE_interval);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 307;
	        this.match(prqlParser.NUMBER);
	        this.state = 308;
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
prqlParser.RULE_intervalKind = 27;
prqlParser.RULE_interval = 28;

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
prqlParser.IntervalKindContext = IntervalKindContext; 
prqlParser.IntervalContext = IntervalContext; 
