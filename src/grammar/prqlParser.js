// Generated from .\prql.g4 by ANTLR 4.9.2
// jshint ignore: start
import antlr4 from 'antlr4';
import prqlListener from './prqlListener.js';

const serializedATN = ["\u0003\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786",
    "\u5964\u00038\u016f\u0004\u0002\t\u0002\u0004\u0003\t\u0003\u0004\u0004",
    "\t\u0004\u0004\u0005\t\u0005\u0004\u0006\t\u0006\u0004\u0007\t\u0007",
    "\u0004\b\t\b\u0004\t\t\t\u0004\n\t\n\u0004\u000b\t\u000b\u0004\f\t\f",
    "\u0004\r\t\r\u0004\u000e\t\u000e\u0004\u000f\t\u000f\u0004\u0010\t\u0010",
    "\u0004\u0011\t\u0011\u0004\u0012\t\u0012\u0004\u0013\t\u0013\u0004\u0014",
    "\t\u0014\u0004\u0015\t\u0015\u0004\u0016\t\u0016\u0004\u0017\t\u0017",
    "\u0004\u0018\t\u0018\u0004\u0019\t\u0019\u0004\u001a\t\u001a\u0004\u001b",
    "\t\u001b\u0004\u001c\t\u001c\u0004\u001d\t\u001d\u0004\u001e\t\u001e",
    "\u0004\u001f\t\u001f\u0004 \t \u0004!\t!\u0004\"\t\"\u0004#\t#\u0004",
    "$\t$\u0004%\t%\u0004&\t&\u0004\'\t\'\u0004(\t(\u0003\u0002\u0003\u0002",
    "\u0003\u0003\u0007\u0003T\n\u0003\f\u0003\u000e\u0003W\u000b\u0003\u0003",
    "\u0003\u0005\u0003Z\n\u0003\u0003\u0003\u0007\u0003]\n\u0003\f\u0003",
    "\u000e\u0003`\u000b\u0003\u0003\u0003\u0003\u0003\u0003\u0003\u0005",
    "\u0003e\n\u0003\u0003\u0003\u0007\u0003h\n\u0003\f\u0003\u000e\u0003",
    "k\u000b\u0003\u0007\u0003m\n\u0003\f\u0003\u000e\u0003p\u000b\u0003",
    "\u0003\u0003\u0003\u0003\u0003\u0004\u0003\u0004\u0007\u0004v\n\u0004",
    "\f\u0004\u000e\u0004y\u000b\u0004\u0003\u0004\u0003\u0004\u0003\u0005",
    "\u0003\u0005\u0003\u0005\u0003\u0005\u0003\u0005\u0003\u0005\u0003\u0006",
    "\u0003\u0006\u0005\u0006\u0085\n\u0006\u0003\u0007\u0007\u0007\u0088",
    "\n\u0007\f\u0007\u000e\u0007\u008b\u000b\u0007\u0003\b\u0003\b\u0005",
    "\b\u008f\n\b\u0003\b\u0005\b\u0092\n\b\u0003\t\u0003\t\u0003\t\u0003",
    "\t\u0007\t\u0098\n\t\f\t\u000e\t\u009b\u000b\t\u0003\t\u0003\t\u0003",
    "\n\u0003\n\u0005\n\u00a1\n\n\u0003\u000b\u0003\u000b\u0003\u000b\u0003",
    "\u000b\u0003\u000b\u0003\f\u0003\f\u0005\f\u00aa\n\f\u0003\r\u0003\r",
    "\u0003\r\u0003\r\u0007\r\u00b0\n\r\f\r\u000e\r\u00b3\u000b\r\u0003\u000e",
    "\u0003\u000e\u0003\u000e\u0003\u000e\u0003\u000e\u0007\u000e\u00ba\n",
    "\u000e\f\u000e\u000e\u000e\u00bd\u000b\u000e\u0003\u000e\u0003\u000e",
    "\u0003\u000f\u0003\u000f\u0003\u000f\u0003\u0010\u0003\u0010\u0003\u0011",
    "\u0003\u0011\u0003\u0011\u0003\u0011\u0006\u0011\u00ca\n\u0011\r\u0011",
    "\u000e\u0011\u00cb\u0003\u0012\u0003\u0012\u0003\u0012\u0003\u0012\u0005",
    "\u0012\u00d2\n\u0012\u0003\u0013\u0003\u0013\u0003\u0013\u0003\u0013",
    "\u0003\u0014\u0003\u0014\u0003\u0014\u0003\u0014\u0003\u0015\u0003\u0015",
    "\u0005\u0015\u00de\n\u0015\u0003\u0016\u0003\u0016\u0003\u0016\u0003",
    "\u0016\u0003\u0016\u0003\u0016\u0005\u0016\u00e6\n\u0016\u0003\u0016",
    "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016",
    "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016",
    "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016",
    "\u0003\u0016\u0007\u0016\u00fc\n\u0016\f\u0016\u000e\u0016\u00ff\u000b",
    "\u0016\u0003\u0017\u0003\u0017\u0003\u0017\u0003\u0017\u0003\u0017\u0003",
    "\u0017\u0005\u0017\u0107\n\u0017\u0003\u0018\u0003\u0018\u0003\u0018",
    "\u0005\u0018\u010c\n\u0018\u0003\u0019\u0003\u0019\u0003\u0019\u0003",
    "\u0019\u0003\u0019\u0005\u0019\u0113\n\u0019\u0003\u001a\u0003\u001a",
    "\u0007\u001a\u0117\n\u001a\f\u001a\u000e\u001a\u011a\u000b\u001a\u0003",
    "\u001a\u0003\u001a\u0005\u001a\u011e\n\u001a\u0003\u001a\u0003\u001a",
    "\u0007\u001a\u0122\n\u001a\f\u001a\u000e\u001a\u0125\u000b\u001a\u0003",
    "\u001a\u0003\u001a\u0005\u001a\u0129\n\u001a\u0007\u001a\u012b\n\u001a",
    "\f\u001a\u000e\u001a\u012e\u000b\u001a\u0003\u001a\u0005\u001a\u0131",
    "\n\u001a\u0003\u001a\u0005\u001a\u0134\n\u001a\u0005\u001a\u0136\n\u001a",
    "\u0003\u001a\u0003\u001a\u0003\u001b\u0003\u001b\u0007\u001b\u013c\n",
    "\u001b\f\u001b\u000e\u001b\u013f\u000b\u001b\u0003\u001b\u0003\u001b",
    "\u0007\u001b\u0143\n\u001b\f\u001b\u000e\u001b\u0146\u000b\u001b\u0003",
    "\u001b\u0003\u001b\u0003\u001c\u0003\u001c\u0003\u001d\u0003\u001d\u0003",
    "\u001e\u0003\u001e\u0003\u001e\u0003\u001e\u0003\u001f\u0003\u001f\u0003",
    "\u001f\u0003\u001f\u0003 \u0003 \u0003 \u0003 \u0003 \u0003 \u0005 ",
    "\u015c\n \u0003!\u0003!\u0003\"\u0003\"\u0003#\u0003#\u0003$\u0003$",
    "\u0003%\u0003%\u0003&\u0003&\u0003\'\u0003\'\u0003(\u0003(\u0003(\u0003",
    "(\u0002\u0003*)\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016",
    "\u0018\u001a\u001c\u001e \"$&(*,.02468:<>@BDFHJLN\u0002\f\u0003\u0002",
    "78\u0003\u0002\"#\u0003\u0002\u001f!\u0003\u0002\u0006\u0007\u0003\u0002",
    "\b\t\u0004\u0002\u000b\u000b\"#\u0003\u0002\f\u000e\u0004\u0002\u000f",
    "\u0012*+\u0003\u0002\u0013\u0014\u0003\u0002\u0016\u001e\u0002\u017c",
    "\u0002P\u0003\u0002\u0002\u0002\u0004U\u0003\u0002\u0002\u0002\u0006",
    "s\u0003\u0002\u0002\u0002\b|\u0003\u0002\u0002\u0002\n\u0082\u0003\u0002",
    "\u0002\u0002\f\u0089\u0003\u0002\u0002\u0002\u000e\u008e\u0003\u0002",
    "\u0002\u0002\u0010\u0093\u0003\u0002\u0002\u0002\u0012\u009e\u0003\u0002",
    "\u0002\u0002\u0014\u00a2\u0003\u0002\u0002\u0002\u0016\u00a9\u0003\u0002",
    "\u0002\u0002\u0018\u00ab\u0003\u0002\u0002\u0002\u001a\u00b4\u0003\u0002",
    "\u0002\u0002\u001c\u00c0\u0003\u0002\u0002\u0002\u001e\u00c3\u0003\u0002",
    "\u0002\u0002 \u00c5\u0003\u0002\u0002\u0002\"\u00cd\u0003\u0002\u0002",
    "\u0002$\u00d3\u0003\u0002\u0002\u0002&\u00d7\u0003\u0002\u0002\u0002",
    "(\u00dd\u0003\u0002\u0002\u0002*\u00e5\u0003\u0002\u0002\u0002,\u0106",
    "\u0003\u0002\u0002\u0002.\u0108\u0003\u0002\u0002\u00020\u0112\u0003",
    "\u0002\u0002\u00022\u0114\u0003\u0002\u0002\u00024\u0139\u0003\u0002",
    "\u0002\u00026\u0149\u0003\u0002\u0002\u00028\u014b\u0003\u0002\u0002",
    "\u0002:\u014d\u0003\u0002\u0002\u0002<\u0151\u0003\u0002\u0002\u0002",
    ">\u015b\u0003\u0002\u0002\u0002@\u015d\u0003\u0002\u0002\u0002B\u015f",
    "\u0003\u0002\u0002\u0002D\u0161\u0003\u0002\u0002\u0002F\u0163\u0003",
    "\u0002\u0002\u0002H\u0165\u0003\u0002\u0002\u0002J\u0167\u0003\u0002",
    "\u0002\u0002L\u0169\u0003\u0002\u0002\u0002N\u016b\u0003\u0002\u0002",
    "\u0002PQ\t\u0002\u0002\u0002Q\u0003\u0003\u0002\u0002\u0002RT\u0005",
    "\u0002\u0002\u0002SR\u0003\u0002\u0002\u0002TW\u0003\u0002\u0002\u0002",
    "US\u0003\u0002\u0002\u0002UV\u0003\u0002\u0002\u0002VY\u0003\u0002\u0002",
    "\u0002WU\u0003\u0002\u0002\u0002XZ\u0005\u0006\u0004\u0002YX\u0003\u0002",
    "\u0002\u0002YZ\u0003\u0002\u0002\u0002Z^\u0003\u0002\u0002\u0002[]\u0005",
    "\u0002\u0002\u0002\\[\u0003\u0002\u0002\u0002]`\u0003\u0002\u0002\u0002",
    "^\\\u0003\u0002\u0002\u0002^_\u0003\u0002\u0002\u0002_n\u0003\u0002",
    "\u0002\u0002`^\u0003\u0002\u0002\u0002ae\u0005\b\u0005\u0002be\u0005",
    "\u0014\u000b\u0002ce\u0005\u0018\r\u0002da\u0003\u0002\u0002\u0002d",
    "b\u0003\u0002\u0002\u0002dc\u0003\u0002\u0002\u0002ei\u0003\u0002\u0002",
    "\u0002fh\u0005\u0002\u0002\u0002gf\u0003\u0002\u0002\u0002hk\u0003\u0002",
    "\u0002\u0002ig\u0003\u0002\u0002\u0002ij\u0003\u0002\u0002\u0002jm\u0003",
    "\u0002\u0002\u0002ki\u0003\u0002\u0002\u0002ld\u0003\u0002\u0002\u0002",
    "mp\u0003\u0002\u0002\u0002nl\u0003\u0002\u0002\u0002no\u0003\u0002\u0002",
    "\u0002oq\u0003\u0002\u0002\u0002pn\u0003\u0002\u0002\u0002qr\u0007\u0002",
    "\u0002\u0003r\u0005\u0003\u0002\u0002\u0002sw\u0007 \u0002\u0002tv\u0005",
    "\"\u0012\u0002ut\u0003\u0002\u0002\u0002vy\u0003\u0002\u0002\u0002w",
    "u\u0003\u0002\u0002\u0002wx\u0003\u0002\u0002\u0002xz\u0003\u0002\u0002",
    "\u0002yw\u0003\u0002\u0002\u0002z{\u0005\u0002\u0002\u0002{\u0007\u0003",
    "\u0002\u0002\u0002|}\u0007\u001f\u0002\u0002}~\u0005\n\u0006\u0002~",
    "\u007f\u0005\f\u0007\u0002\u007f\u0080\u0007\u0003\u0002\u0002\u0080",
    "\u0081\u0005*\u0016\u0002\u0081\t\u0003\u0002\u0002\u0002\u0082\u0084",
    "\u00073\u0002\u0002\u0083\u0085\u0005\u0010\t\u0002\u0084\u0083\u0003",
    "\u0002\u0002\u0002\u0084\u0085\u0003\u0002\u0002\u0002\u0085\u000b\u0003",
    "\u0002\u0002\u0002\u0086\u0088\u0005\u000e\b\u0002\u0087\u0086\u0003",
    "\u0002\u0002\u0002\u0088\u008b\u0003\u0002\u0002\u0002\u0089\u0087\u0003",
    "\u0002\u0002\u0002\u0089\u008a\u0003\u0002\u0002\u0002\u008a\r\u0003",
    "\u0002\u0002\u0002\u008b\u0089\u0003\u0002\u0002\u0002\u008c\u008f\u0005",
    "\"\u0012\u0002\u008d\u008f\u00073\u0002\u0002\u008e\u008c\u0003\u0002",
    "\u0002\u0002\u008e\u008d\u0003\u0002\u0002\u0002\u008f\u0091\u0003\u0002",
    "\u0002\u0002\u0090\u0092\u0005\u0010\t\u0002\u0091\u0090\u0003\u0002",
    "\u0002\u0002\u0091\u0092\u0003\u0002\u0002\u0002\u0092\u000f\u0003\u0002",
    "\u0002\u0002\u0093\u0094\u0007*\u0002\u0002\u0094\u0095\u0005\u0012",
    "\n\u0002\u0095\u0099\u0007%\u0002\u0002\u0096\u0098\u0005\u0012\n\u0002",
    "\u0097\u0096\u0003\u0002\u0002\u0002\u0098\u009b\u0003\u0002\u0002\u0002",
    "\u0099\u0097\u0003\u0002\u0002\u0002\u0099\u009a\u0003\u0002\u0002\u0002",
    "\u009a\u009c\u0003\u0002\u0002\u0002\u009b\u0099\u0003\u0002\u0002\u0002",
    "\u009c\u009d\u0007+\u0002\u0002\u009d\u0011\u0003\u0002\u0002\u0002",
    "\u009e\u00a0\u00073\u0002\u0002\u009f\u00a1\u0005\u0010\t\u0002\u00a0",
    "\u009f\u0003\u0002\u0002\u0002\u00a0\u00a1\u0003\u0002\u0002\u0002\u00a1",
    "\u0013\u0003\u0002\u0002\u0002\u00a2\u00a3\u0007!\u0002\u0002\u00a3",
    "\u00a4\u00073\u0002\u0002\u00a4\u00a5\u0007$\u0002\u0002\u00a5\u00a6",
    "\u00054\u001b\u0002\u00a6\u0015\u0003\u0002\u0002\u0002\u00a7\u00aa",
    "\u0005\u0002\u0002\u0002\u00a8\u00aa\u0007%\u0002\u0002\u00a9\u00a7",
    "\u0003\u0002\u0002\u0002\u00a9\u00a8\u0003\u0002\u0002\u0002\u00aa\u0017",
    "\u0003\u0002\u0002\u0002\u00ab\u00b1\u0005(\u0015\u0002\u00ac\u00ad",
    "\u0005\u0016\f\u0002\u00ad\u00ae\u0005(\u0015\u0002\u00ae\u00b0\u0003",
    "\u0002\u0002\u0002\u00af\u00ac\u0003\u0002\u0002\u0002\u00b0\u00b3\u0003",
    "\u0002\u0002\u0002\u00b1\u00af\u0003\u0002\u0002\u0002\u00b1\u00b2\u0003",
    "\u0002\u0002\u0002\u00b2\u0019\u0003\u0002\u0002\u0002\u00b3\u00b1\u0003",
    "\u0002\u0002\u0002\u00b4\u00bb\u0007\u0004\u0002\u0002\u00b5\u00b6\u0005",
    "\u0002\u0002\u0002\u00b6\u00b7\u0007\u0004\u0002\u0002\u00b7\u00b8\u000b",
    "\u0002\u0002\u0002\u00b8\u00ba\u0003\u0002\u0002\u0002\u00b9\u00b5\u0003",
    "\u0002\u0002\u0002\u00ba\u00bd\u0003\u0002\u0002\u0002\u00bb\u00b9\u0003",
    "\u0002\u0002\u0002\u00bb\u00bc\u0003\u0002\u0002\u0002\u00bc\u00be\u0003",
    "\u0002\u0002\u0002\u00bd\u00bb\u0003\u0002\u0002\u0002\u00be\u00bf\u0007",
    "\u0004\u0002\u0002\u00bf\u001b\u0003\u0002\u0002\u0002\u00c0\u00c1\t",
    "\u0003\u0002\u0002\u00c1\u00c2\u00073\u0002\u0002\u00c2\u001d\u0003",
    "\u0002\u0002\u0002\u00c3\u00c4\t\u0004\u0002\u0002\u00c4\u001f\u0003",
    "\u0002\u0002\u0002\u00c5\u00c9\u00073\u0002\u0002\u00c6\u00ca\u0005",
    "\"\u0012\u0002\u00c7\u00ca\u0005$\u0013\u0002\u00c8\u00ca\u0005*\u0016",
    "\u0002\u00c9\u00c6\u0003\u0002\u0002\u0002\u00c9\u00c7\u0003\u0002\u0002",
    "\u0002\u00c9\u00c8\u0003\u0002\u0002\u0002\u00ca\u00cb\u0003\u0002\u0002",
    "\u0002\u00cb\u00c9\u0003\u0002\u0002\u0002\u00cb\u00cc\u0003\u0002\u0002",
    "\u0002\u00cc!\u0003\u0002\u0002\u0002\u00cd\u00ce\u00073\u0002\u0002",
    "\u00ce\u00d1\u0007\u0005\u0002\u0002\u00cf\u00d2\u0005$\u0013\u0002",
    "\u00d0\u00d2\u0005*\u0016\u0002\u00d1\u00cf\u0003\u0002\u0002\u0002",
    "\u00d1\u00d0\u0003\u0002\u0002\u0002\u00d2#\u0003\u0002\u0002\u0002",
    "\u00d3\u00d4\u00073\u0002\u0002\u00d4\u00d5\u0007$\u0002\u0002\u00d5",
    "\u00d6\u0005*\u0016\u0002\u00d6%\u0003\u0002\u0002\u0002\u00d7\u00d8",
    "\u00073\u0002\u0002\u00d8\u00d9\u0007$\u0002\u0002\u00d9\u00da\u0005",
    "(\u0015\u0002\u00da\'\u0003\u0002\u0002\u0002\u00db\u00de\u0005 \u0011",
    "\u0002\u00dc\u00de\u0005*\u0016\u0002\u00dd\u00db\u0003\u0002\u0002",
    "\u0002\u00dd\u00dc\u0003\u0002\u0002\u0002\u00de)\u0003\u0002\u0002",
    "\u0002\u00df\u00e0\b\u0016\u0001\u0002\u00e0\u00e1\u0007.\u0002\u0002",
    "\u00e1\u00e2\u0005*\u0016\u0002\u00e2\u00e3\u0007/\u0002\u0002\u00e3",
    "\u00e6\u0003\u0002\u0002\u0002\u00e4\u00e6\u0005,\u0017\u0002\u00e5",
    "\u00df\u0003\u0002\u0002\u0002\u00e5\u00e4\u0003\u0002\u0002\u0002\u00e6",
    "\u00fd\u0003\u0002\u0002\u0002\u00e7\u00e8\f\t\u0002\u0002\u00e8\u00e9",
    "\u0005B\"\u0002\u00e9\u00ea\u0005*\u0016\n\u00ea\u00fc\u0003\u0002\u0002",
    "\u0002\u00eb\u00ec\f\b\u0002\u0002\u00ec\u00ed\u0005D#\u0002\u00ed\u00ee",
    "\u0005*\u0016\t\u00ee\u00fc\u0003\u0002\u0002\u0002\u00ef\u00f0\f\u0007",
    "\u0002\u0002\u00f0\u00f1\u0005F$\u0002\u00f1\u00f2\u0005*\u0016\b\u00f2",
    "\u00fc\u0003\u0002\u0002\u0002\u00f3\u00f4\f\u0006\u0002\u0002\u00f4",
    "\u00f5\u0005J&\u0002\u00f5\u00f6\u0005*\u0016\u0007\u00f6\u00fc\u0003",
    "\u0002\u0002\u0002\u00f7\u00f8\f\u0005\u0002\u0002\u00f8\u00f9\u0005",
    "H%\u0002\u00f9\u00fa\u0005*\u0016\u0006\u00fa\u00fc\u0003\u0002\u0002",
    "\u0002\u00fb\u00e7\u0003\u0002\u0002\u0002\u00fb\u00eb\u0003\u0002\u0002",
    "\u0002\u00fb\u00ef\u0003\u0002\u0002\u0002\u00fb\u00f3\u0003\u0002\u0002",
    "\u0002\u00fb\u00f7\u0003\u0002\u0002\u0002\u00fc\u00ff\u0003\u0002\u0002",
    "\u0002\u00fd\u00fb\u0003\u0002\u0002\u0002\u00fd\u00fe\u0003\u0002\u0002",
    "\u0002\u00fe+\u0003\u0002\u0002\u0002\u00ff\u00fd\u0003\u0002\u0002",
    "\u0002\u0100\u0107\u0005<\u001f\u0002\u0101\u0107\u00050\u0019\u0002",
    "\u0102\u0107\u00073\u0002\u0002\u0103\u0107\u0005.\u0018\u0002\u0104",
    "\u0107\u00052\u001a\u0002\u0105\u0107\u00054\u001b\u0002\u0106\u0100",
    "\u0003\u0002\u0002\u0002\u0106\u0101\u0003\u0002\u0002\u0002\u0106\u0102",
    "\u0003\u0002\u0002\u0002\u0106\u0103\u0003\u0002\u0002\u0002\u0106\u0104",
    "\u0003\u0002\u0002\u0002\u0106\u0105\u0003\u0002\u0002\u0002\u0107-",
    "\u0003\u0002\u0002\u0002\u0108\u010b\u0005@!\u0002\u0109\u010c\u0005",
    "4\u001b\u0002\u010a\u010c\u00073\u0002\u0002\u010b\u0109\u0003\u0002",
    "\u0002\u0002\u010b\u010a\u0003\u0002\u0002\u0002\u010c/\u0003\u0002",
    "\u0002\u0002\u010d\u0113\u0005N(\u0002\u010e\u0113\u00072\u0002\u0002",
    "\u010f\u0113\u00071\u0002\u0002\u0110\u0113\u00070\u0002\u0002\u0111",
    "\u0113\u0005:\u001e\u0002\u0112\u010d\u0003\u0002\u0002\u0002\u0112",
    "\u010e\u0003\u0002\u0002\u0002\u0112\u010f\u0003\u0002\u0002\u0002\u0112",
    "\u0110\u0003\u0002\u0002\u0002\u0112\u0111\u0003\u0002\u0002\u0002\u0113",
    "1\u0003\u0002\u0002\u0002\u0114\u0135\u0007,\u0002\u0002\u0115\u0117",
    "\u0005\u0002\u0002\u0002\u0116\u0115\u0003\u0002\u0002\u0002\u0117\u011a",
    "\u0003\u0002\u0002\u0002\u0118\u0116\u0003\u0002\u0002\u0002\u0118\u0119",
    "\u0003\u0002\u0002\u0002\u0119\u011d\u0003\u0002\u0002\u0002\u011a\u0118",
    "\u0003\u0002\u0002\u0002\u011b\u011e\u0005&\u0014\u0002\u011c\u011e",
    "\u0005(\u0015\u0002\u011d\u011b\u0003\u0002\u0002\u0002\u011d\u011c",
    "\u0003\u0002\u0002\u0002\u011e\u012c\u0003\u0002\u0002\u0002\u011f\u0123",
    "\u0007&\u0002\u0002\u0120\u0122\u0005\u0002\u0002\u0002\u0121\u0120",
    "\u0003\u0002\u0002\u0002\u0122\u0125\u0003\u0002\u0002\u0002\u0123\u0121",
    "\u0003\u0002\u0002\u0002\u0123\u0124\u0003\u0002\u0002\u0002\u0124\u0128",
    "\u0003\u0002\u0002\u0002\u0125\u0123\u0003\u0002\u0002\u0002\u0126\u0129",
    "\u0005&\u0014\u0002\u0127\u0129\u0005(\u0015\u0002\u0128\u0126\u0003",
    "\u0002\u0002\u0002\u0128\u0127\u0003\u0002\u0002\u0002\u0129\u012b\u0003",
    "\u0002\u0002\u0002\u012a\u011f\u0003\u0002\u0002\u0002\u012b\u012e\u0003",
    "\u0002\u0002\u0002\u012c\u012a\u0003\u0002\u0002\u0002\u012c\u012d\u0003",
    "\u0002\u0002\u0002\u012d\u0130\u0003\u0002\u0002\u0002\u012e\u012c\u0003",
    "\u0002\u0002\u0002\u012f\u0131\u0007&\u0002\u0002\u0130\u012f\u0003",
    "\u0002\u0002\u0002\u0130\u0131\u0003\u0002\u0002\u0002\u0131\u0133\u0003",
    "\u0002\u0002\u0002\u0132\u0134\u0005\u0002\u0002\u0002\u0133\u0132\u0003",
    "\u0002\u0002\u0002\u0133\u0134\u0003\u0002\u0002\u0002\u0134\u0136\u0003",
    "\u0002\u0002\u0002\u0135\u0118\u0003\u0002\u0002\u0002\u0135\u0136\u0003",
    "\u0002\u0002\u0002\u0136\u0137\u0003\u0002\u0002\u0002\u0137\u0138\u0007",
    "-\u0002\u0002\u01383\u0003\u0002\u0002\u0002\u0139\u013d\u0007.\u0002",
    "\u0002\u013a\u013c\u0005\u0002\u0002\u0002\u013b\u013a\u0003\u0002\u0002",
    "\u0002\u013c\u013f\u0003\u0002\u0002\u0002\u013d\u013b\u0003\u0002\u0002",
    "\u0002\u013d\u013e\u0003\u0002\u0002\u0002\u013e\u0140\u0003\u0002\u0002",
    "\u0002\u013f\u013d\u0003\u0002\u0002\u0002\u0140\u0144\u0005\u0018\r",
    "\u0002\u0141\u0143\u0005\u0002\u0002\u0002\u0142\u0141\u0003\u0002\u0002",
    "\u0002\u0143\u0146\u0003\u0002\u0002\u0002\u0144\u0142\u0003\u0002\u0002",
    "\u0002\u0144\u0145\u0003\u0002\u0002\u0002\u0145\u0147\u0003\u0002\u0002",
    "\u0002\u0146\u0144\u0003\u0002\u0002\u0002\u0147\u0148\u0007/\u0002",
    "\u0002\u01485\u0003\u0002\u0002\u0002\u0149\u014a\t\u0005\u0002\u0002",
    "\u014a7\u0003\u0002\u0002\u0002\u014b\u014c\t\u0006\u0002\u0002\u014c",
    "9\u0003\u0002\u0002\u0002\u014d\u014e\u0007\u0006\u0002\u0002\u014e",
    "\u014f\u000b\u0002\u0002\u0002\u014f\u0150\u0007\u0006\u0002\u0002\u0150",
    ";\u0003\u0002\u0002\u0002\u0151\u0152\u00050\u0019\u0002\u0152\u0153",
    "\u0007\n\u0002\u0002\u0153\u0154\u00050\u0019\u0002\u0154=\u0003\u0002",
    "\u0002\u0002\u0155\u015c\u0005@!\u0002\u0156\u015c\u0005B\"\u0002\u0157",
    "\u015c\u0005D#\u0002\u0158\u015c\u0005F$\u0002\u0159\u015c\u0005H%\u0002",
    "\u015a\u015c\u0005J&\u0002\u015b\u0155\u0003\u0002\u0002\u0002\u015b",
    "\u0156\u0003\u0002\u0002\u0002\u015b\u0157\u0003\u0002\u0002\u0002\u015b",
    "\u0158\u0003\u0002\u0002\u0002\u015b\u0159\u0003\u0002\u0002\u0002\u015b",
    "\u015a\u0003\u0002\u0002\u0002\u015c?\u0003\u0002\u0002\u0002\u015d",
    "\u015e\t\u0007\u0002\u0002\u015eA\u0003\u0002\u0002\u0002\u015f\u0160",
    "\t\b\u0002\u0002\u0160C\u0003\u0002\u0002\u0002\u0161\u0162\t\u0003",
    "\u0002\u0002\u0162E\u0003\u0002\u0002\u0002\u0163\u0164\t\t\u0002\u0002",
    "\u0164G\u0003\u0002\u0002\u0002\u0165\u0166\t\n\u0002\u0002\u0166I\u0003",
    "\u0002\u0002\u0002\u0167\u0168\u0007\u0015\u0002\u0002\u0168K\u0003",
    "\u0002\u0002\u0002\u0169\u016a\t\u000b\u0002\u0002\u016aM\u0003\u0002",
    "\u0002\u0002\u016b\u016c\u00072\u0002\u0002\u016c\u016d\u0005L\'\u0002",
    "\u016dO\u0003\u0002\u0002\u0002\'UY^dinw\u0084\u0089\u008e\u0091\u0099",
    "\u00a0\u00a9\u00b1\u00bb\u00c9\u00cb\u00d1\u00dd\u00e5\u00fb\u00fd\u0106",
    "\u010b\u0112\u0118\u011d\u0123\u0128\u012c\u0130\u0133\u0135\u013d\u0144",
    "\u015b"].join("");


const atn = new antlr4.atn.ATNDeserializer().deserialize(serializedATN);

const decisionsToDFA = atn.decisionToState.map( (ds, index) => new antlr4.dfa.DFA(ds, index) );

const sharedContextCache = new antlr4.PredictionContextCache();

export default class prqlParser extends antlr4.Parser {

    static grammarFileName = "prql.g4";
    static literalNames = [ null, "'->'", "'`'", "':'", "'\"'", "'''", "'\"\"\"'", 
                            "'''''", "'..'", "'!'", "'*'", "'/'", "'%'", 
                            "'=='", "'!='", "'>='", "'<='", "'and'", "'or'", 
                            "'??'", "'microseconds'", "'milliseconds'", 
                            "'seconds'", "'minutes'", "'hours'", "'days'", 
                            "'weeks'", "'months'", "'years'", "'func'", 
                            "'prql'", "'table'", "'+'", "'-'", "'='", "'|'", 
                            "','", "'.'", "'$'", "'_'", "'<'", "'>'", "'['", 
                            "']'", "'('", "')'", "'null'" ];
    static symbolicNames = [ null, null, null, null, null, null, null, null, 
                             null, null, null, null, null, null, null, null, 
                             null, null, null, null, null, null, null, null, 
                             null, null, null, null, null, "FUNC", "PRQL", 
                             "TABLE", "PLUS", "MINUS", "EQUAL", "BAR", "COMMA", 
                             "DOT", "DOLLAR", "UNDERSCORE", "LANG", "RANG", 
                             "LBRACKET", "RBRACKET", "LPAREN", "RPAREN", 
                             "NULL_", "BOOLEAN", "NUMBER", "IDENT", "IDENT_START", 
                             "IDENT_NEXT", "WHITESPACE", "NEWLINE", "COMMENT" ];
    static ruleNames = [ "nl", "query", "query_def", "func_def", "func_def_name", 
                         "func_def_params", "func_def_param", "type_def", 
                         "type_term", "table", "pipe", "pipeline", "ident_backticks", 
                         "signed_ident", "keyword", "func_call", "named_arg", 
                         "assign", "assign_call", "expr_call", "expr", "term", 
                         "expr_unary", "literal", "list", "nested_pipeline", 
                         "single_quote", "multi_quote", "string", "range", 
                         "operator", "operator_unary", "operator_mul", "operator_add", 
                         "operator_compare", "operator_logical", "operator_coalesce", 
                         "interval_kind", "interval" ];

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
	        this.state = 78;
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
	        this.state = 83;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,0,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                this.state = 80;
	                this.nl(); 
	            }
	            this.state = 85;
	            this._errHandler.sync(this);
	            _alt = this._interp.adaptivePredict(this._input,0,this._ctx);
	        }

	        this.state = 87;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.PRQL) {
	            this.state = 86;
	            this.query_def();
	        }

	        this.state = 92;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	            this.state = 89;
	            this.nl();
	            this.state = 94;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 108;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.T__3) | (1 << prqlParser.T__8) | (1 << prqlParser.FUNC) | (1 << prqlParser.TABLE))) !== 0) || ((((_la - 32)) & ~0x1f) == 0 && ((1 << (_la - 32)) & ((1 << (prqlParser.PLUS - 32)) | (1 << (prqlParser.MINUS - 32)) | (1 << (prqlParser.LBRACKET - 32)) | (1 << (prqlParser.LPAREN - 32)) | (1 << (prqlParser.NULL_ - 32)) | (1 << (prqlParser.BOOLEAN - 32)) | (1 << (prqlParser.NUMBER - 32)) | (1 << (prqlParser.IDENT - 32)))) !== 0)) {
	            this.state = 98;
	            this._errHandler.sync(this);
	            switch(this._input.LA(1)) {
	            case prqlParser.FUNC:
	                this.state = 95;
	                this.func_def();
	                break;
	            case prqlParser.TABLE:
	                this.state = 96;
	                this.table();
	                break;
	            case prqlParser.T__3:
	            case prqlParser.T__8:
	            case prqlParser.PLUS:
	            case prqlParser.MINUS:
	            case prqlParser.LBRACKET:
	            case prqlParser.LPAREN:
	            case prqlParser.NULL_:
	            case prqlParser.BOOLEAN:
	            case prqlParser.NUMBER:
	            case prqlParser.IDENT:
	                this.state = 97;
	                this.pipeline();
	                break;
	            default:
	                throw new antlr4.error.NoViableAltException(this);
	            }
	            this.state = 103;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	                this.state = 100;
	                this.nl();
	                this.state = 105;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	            }
	            this.state = 110;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 111;
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



	query_def() {
	    let localctx = new Query_defContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 4, prqlParser.RULE_query_def);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 113;
	        this.match(prqlParser.PRQL);
	        this.state = 117;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.IDENT) {
	            this.state = 114;
	            this.named_arg();
	            this.state = 119;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 120;
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



	func_def() {
	    let localctx = new Func_defContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 6, prqlParser.RULE_func_def);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 122;
	        this.match(prqlParser.FUNC);
	        this.state = 123;
	        this.func_def_name();
	        this.state = 124;
	        this.func_def_params();
	        this.state = 125;
	        this.match(prqlParser.T__0);
	        this.state = 126;
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



	func_def_name() {
	    let localctx = new Func_def_nameContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 8, prqlParser.RULE_func_def_name);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 128;
	        this.match(prqlParser.IDENT);
	        this.state = 130;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.LANG) {
	            this.state = 129;
	            this.type_def();
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



	func_def_params() {
	    let localctx = new Func_def_paramsContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 10, prqlParser.RULE_func_def_params);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 135;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.IDENT) {
	            this.state = 132;
	            this.func_def_param();
	            this.state = 137;
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



	func_def_param() {
	    let localctx = new Func_def_paramContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 12, prqlParser.RULE_func_def_param);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 140;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,9,this._ctx);
	        switch(la_) {
	        case 1:
	            this.state = 138;
	            this.named_arg();
	            break;

	        case 2:
	            this.state = 139;
	            this.match(prqlParser.IDENT);
	            break;

	        }
	        this.state = 143;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.LANG) {
	            this.state = 142;
	            this.type_def();
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



	type_def() {
	    let localctx = new Type_defContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 14, prqlParser.RULE_type_def);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 145;
	        this.match(prqlParser.LANG);
	        this.state = 146;
	        this.type_term();
	        this.state = 147;
	        this.match(prqlParser.BAR);
	        this.state = 151;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.IDENT) {
	            this.state = 148;
	            this.type_term();
	            this.state = 153;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 154;
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



	type_term() {
	    let localctx = new Type_termContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 16, prqlParser.RULE_type_term);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 156;
	        this.match(prqlParser.IDENT);
	        this.state = 158;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.LANG) {
	            this.state = 157;
	            this.type_def();
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
	        this.state = 160;
	        this.match(prqlParser.TABLE);
	        this.state = 161;
	        this.match(prqlParser.IDENT);
	        this.state = 162;
	        this.match(prqlParser.EQUAL);
	        this.state = 163;
	        this.nested_pipeline();
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
	        this.state = 167;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case prqlParser.NEWLINE:
	        case prqlParser.COMMENT:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 165;
	            this.nl();
	            break;
	        case prqlParser.BAR:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 166;
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
	        this.state = 169;
	        this.expr_call();
	        this.state = 175;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,14,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                this.state = 170;
	                this.pipe();
	                this.state = 171;
	                this.expr_call(); 
	            }
	            this.state = 177;
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



	ident_backticks() {
	    let localctx = new Ident_backticksContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 24, prqlParser.RULE_ident_backticks);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 178;
	        this.match(prqlParser.T__1);
	        this.state = 185;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	            this.state = 179;
	            this.nl();
	            this.state = 180;
	            this.match(prqlParser.T__1);
	            this.state = 181;
	            this.matchWildcard();
	            this.state = 187;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 188;
	        this.match(prqlParser.T__1);
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



	signed_ident() {
	    let localctx = new Signed_identContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 26, prqlParser.RULE_signed_ident);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 190;
	        _la = this._input.LA(1);
	        if(!(_la===prqlParser.PLUS || _la===prqlParser.MINUS)) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 191;
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
	        this.state = 193;
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



	func_call() {
	    let localctx = new Func_callContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 30, prqlParser.RULE_func_call);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 195;
	        this.match(prqlParser.IDENT);
	        this.state = 199; 
	        this._errHandler.sync(this);
	        var _alt = 1;
	        do {
	        	switch (_alt) {
	        	case 1:
	        		this.state = 199;
	        		this._errHandler.sync(this);
	        		var la_ = this._interp.adaptivePredict(this._input,16,this._ctx);
	        		switch(la_) {
	        		case 1:
	        		    this.state = 196;
	        		    this.named_arg();
	        		    break;

	        		case 2:
	        		    this.state = 197;
	        		    this.assign();
	        		    break;

	        		case 3:
	        		    this.state = 198;
	        		    this.expr(0);
	        		    break;

	        		}
	        		break;
	        	default:
	        		throw new antlr4.error.NoViableAltException(this);
	        	}
	        	this.state = 201; 
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



	named_arg() {
	    let localctx = new Named_argContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 32, prqlParser.RULE_named_arg);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 203;
	        this.match(prqlParser.IDENT);
	        this.state = 204;
	        this.match(prqlParser.T__2);
	        this.state = 207;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,18,this._ctx);
	        switch(la_) {
	        case 1:
	            this.state = 205;
	            this.assign();
	            break;

	        case 2:
	            this.state = 206;
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
	        this.state = 209;
	        this.match(prqlParser.IDENT);
	        this.state = 210;
	        this.match(prqlParser.EQUAL);
	        this.state = 211;
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



	assign_call() {
	    let localctx = new Assign_callContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 36, prqlParser.RULE_assign_call);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 213;
	        this.match(prqlParser.IDENT);
	        this.state = 214;
	        this.match(prqlParser.EQUAL);
	        this.state = 215;
	        this.expr_call();
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



	expr_call() {
	    let localctx = new Expr_callContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 38, prqlParser.RULE_expr_call);
	    try {
	        this.state = 219;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,19,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 217;
	            this.func_call();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 218;
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
	        this.state = 227;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,20,this._ctx);
	        switch(la_) {
	        case 1:
	            this.state = 222;
	            this.match(prqlParser.LPAREN);
	            this.state = 223;
	            this.expr(0);
	            this.state = 224;
	            this.match(prqlParser.RPAREN);
	            break;

	        case 2:
	            this.state = 226;
	            this.term();
	            break;

	        }
	        this._ctx.stop = this._input.LT(-1);
	        this.state = 251;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,22,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                if(this._parseListeners!==null) {
	                    this.triggerExitRuleEvent();
	                }
	                _prevctx = localctx;
	                this.state = 249;
	                this._errHandler.sync(this);
	                var la_ = this._interp.adaptivePredict(this._input,21,this._ctx);
	                switch(la_) {
	                case 1:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 229;
	                    if (!( this.precpred(this._ctx, 7))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 7)");
	                    }
	                    this.state = 230;
	                    this.operator_mul();
	                    this.state = 231;
	                    this.expr(8);
	                    break;

	                case 2:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 233;
	                    if (!( this.precpred(this._ctx, 6))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 6)");
	                    }
	                    this.state = 234;
	                    this.operator_add();
	                    this.state = 235;
	                    this.expr(7);
	                    break;

	                case 3:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 237;
	                    if (!( this.precpred(this._ctx, 5))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 5)");
	                    }
	                    this.state = 238;
	                    this.operator_compare();
	                    this.state = 239;
	                    this.expr(6);
	                    break;

	                case 4:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 241;
	                    if (!( this.precpred(this._ctx, 4))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 4)");
	                    }
	                    this.state = 242;
	                    this.operator_coalesce();
	                    this.state = 243;
	                    this.expr(5);
	                    break;

	                case 5:
	                    localctx = new ExprContext(this, _parentctx, _parentState);
	                    this.pushNewRecursionContext(localctx, _startState, prqlParser.RULE_expr);
	                    this.state = 245;
	                    if (!( this.precpred(this._ctx, 3))) {
	                        throw new antlr4.error.FailedPredicateException(this, "this.precpred(this._ctx, 3)");
	                    }
	                    this.state = 246;
	                    this.operator_logical();
	                    this.state = 247;
	                    this.expr(4);
	                    break;

	                } 
	            }
	            this.state = 253;
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
	        this.state = 260;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,23,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 254;
	            this.range();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 255;
	            this.literal();
	            break;

	        case 3:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 256;
	            this.match(prqlParser.IDENT);
	            break;

	        case 4:
	            this.enterOuterAlt(localctx, 4);
	            this.state = 257;
	            this.expr_unary();
	            break;

	        case 5:
	            this.enterOuterAlt(localctx, 5);
	            this.state = 258;
	            this.list();
	            break;

	        case 6:
	            this.enterOuterAlt(localctx, 6);
	            this.state = 259;
	            this.nested_pipeline();
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



	expr_unary() {
	    let localctx = new Expr_unaryContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 44, prqlParser.RULE_expr_unary);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 262;
	        this.operator_unary();
	        this.state = 265;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case prqlParser.LPAREN:
	            this.state = 263;
	            this.nested_pipeline();
	            break;
	        case prqlParser.IDENT:
	            this.state = 264;
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
	        this.state = 272;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,25,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 267;
	            this.interval();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 268;
	            this.match(prqlParser.NUMBER);
	            break;

	        case 3:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 269;
	            this.match(prqlParser.BOOLEAN);
	            break;

	        case 4:
	            this.enterOuterAlt(localctx, 4);
	            this.state = 270;
	            this.match(prqlParser.NULL_);
	            break;

	        case 5:
	            this.enterOuterAlt(localctx, 5);
	            this.state = 271;
	            this.string();
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
	        this.state = 274;
	        this.match(prqlParser.LBRACKET);
	        this.state = 307;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===prqlParser.T__3 || _la===prqlParser.T__8 || ((((_la - 32)) & ~0x1f) == 0 && ((1 << (_la - 32)) & ((1 << (prqlParser.PLUS - 32)) | (1 << (prqlParser.MINUS - 32)) | (1 << (prqlParser.LBRACKET - 32)) | (1 << (prqlParser.LPAREN - 32)) | (1 << (prqlParser.NULL_ - 32)) | (1 << (prqlParser.BOOLEAN - 32)) | (1 << (prqlParser.NUMBER - 32)) | (1 << (prqlParser.IDENT - 32)) | (1 << (prqlParser.NEWLINE - 32)) | (1 << (prqlParser.COMMENT - 32)))) !== 0)) {
	            this.state = 278;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	                this.state = 275;
	                this.nl();
	                this.state = 280;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	            }
	            this.state = 283;
	            this._errHandler.sync(this);
	            var la_ = this._interp.adaptivePredict(this._input,27,this._ctx);
	            switch(la_) {
	            case 1:
	                this.state = 281;
	                this.assign_call();
	                break;

	            case 2:
	                this.state = 282;
	                this.expr_call();
	                break;

	            }
	            this.state = 298;
	            this._errHandler.sync(this);
	            var _alt = this._interp.adaptivePredict(this._input,30,this._ctx)
	            while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	                if(_alt===1) {
	                    this.state = 285;
	                    this.match(prqlParser.COMMA);
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
	                    this.state = 294;
	                    this._errHandler.sync(this);
	                    var la_ = this._interp.adaptivePredict(this._input,29,this._ctx);
	                    switch(la_) {
	                    case 1:
	                        this.state = 292;
	                        this.assign_call();
	                        break;

	                    case 2:
	                        this.state = 293;
	                        this.expr_call();
	                        break;

	                    } 
	                }
	                this.state = 300;
	                this._errHandler.sync(this);
	                _alt = this._interp.adaptivePredict(this._input,30,this._ctx);
	            }

	            this.state = 302;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===prqlParser.COMMA) {
	                this.state = 301;
	                this.match(prqlParser.COMMA);
	            }

	            this.state = 305;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	                this.state = 304;
	                this.nl();
	            }

	        }

	        this.state = 309;
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



	nested_pipeline() {
	    let localctx = new Nested_pipelineContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 50, prqlParser.RULE_nested_pipeline);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 311;
	        this.match(prqlParser.LPAREN);
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
	        this.pipeline();
	        this.state = 322;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===prqlParser.NEWLINE || _la===prqlParser.COMMENT) {
	            this.state = 319;
	            this.nl();
	            this.state = 324;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 325;
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



	single_quote() {
	    let localctx = new Single_quoteContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 52, prqlParser.RULE_single_quote);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 327;
	        _la = this._input.LA(1);
	        if(!(_la===prqlParser.T__3 || _la===prqlParser.T__4)) {
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



	multi_quote() {
	    let localctx = new Multi_quoteContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 54, prqlParser.RULE_multi_quote);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 329;
	        _la = this._input.LA(1);
	        if(!(_la===prqlParser.T__5 || _la===prqlParser.T__6)) {
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



	string() {
	    let localctx = new StringContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 56, prqlParser.RULE_string);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 331;
	        this.match(prqlParser.T__3);
	        this.state = 332;
	        this.matchWildcard();
	        this.state = 333;
	        this.match(prqlParser.T__3);
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
	    this.enterRule(localctx, 58, prqlParser.RULE_range);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 335;
	        this.literal();
	        this.state = 336;
	        this.match(prqlParser.T__7);
	        this.state = 337;
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
	    this.enterRule(localctx, 60, prqlParser.RULE_operator);
	    try {
	        this.state = 345;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,36,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 339;
	            this.operator_unary();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 340;
	            this.operator_mul();
	            break;

	        case 3:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 341;
	            this.operator_add();
	            break;

	        case 4:
	            this.enterOuterAlt(localctx, 4);
	            this.state = 342;
	            this.operator_compare();
	            break;

	        case 5:
	            this.enterOuterAlt(localctx, 5);
	            this.state = 343;
	            this.operator_logical();
	            break;

	        case 6:
	            this.enterOuterAlt(localctx, 6);
	            this.state = 344;
	            this.operator_coalesce();
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



	operator_unary() {
	    let localctx = new Operator_unaryContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 62, prqlParser.RULE_operator_unary);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 347;
	        _la = this._input.LA(1);
	        if(!(((((_la - 9)) & ~0x1f) == 0 && ((1 << (_la - 9)) & ((1 << (prqlParser.T__8 - 9)) | (1 << (prqlParser.PLUS - 9)) | (1 << (prqlParser.MINUS - 9)))) !== 0))) {
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



	operator_mul() {
	    let localctx = new Operator_mulContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 64, prqlParser.RULE_operator_mul);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 349;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.T__9) | (1 << prqlParser.T__10) | (1 << prqlParser.T__11))) !== 0))) {
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



	operator_add() {
	    let localctx = new Operator_addContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 66, prqlParser.RULE_operator_add);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 351;
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



	operator_compare() {
	    let localctx = new Operator_compareContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 68, prqlParser.RULE_operator_compare);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 353;
	        _la = this._input.LA(1);
	        if(!(((((_la - 13)) & ~0x1f) == 0 && ((1 << (_la - 13)) & ((1 << (prqlParser.T__12 - 13)) | (1 << (prqlParser.T__13 - 13)) | (1 << (prqlParser.T__14 - 13)) | (1 << (prqlParser.T__15 - 13)) | (1 << (prqlParser.LANG - 13)) | (1 << (prqlParser.RANG - 13)))) !== 0))) {
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



	operator_logical() {
	    let localctx = new Operator_logicalContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 70, prqlParser.RULE_operator_logical);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 355;
	        _la = this._input.LA(1);
	        if(!(_la===prqlParser.T__16 || _la===prqlParser.T__17)) {
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



	operator_coalesce() {
	    let localctx = new Operator_coalesceContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 72, prqlParser.RULE_operator_coalesce);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 357;
	        this.match(prqlParser.T__18);
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



	interval_kind() {
	    let localctx = new Interval_kindContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 74, prqlParser.RULE_interval_kind);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 359;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) == 0 && ((1 << _la) & ((1 << prqlParser.T__19) | (1 << prqlParser.T__20) | (1 << prqlParser.T__21) | (1 << prqlParser.T__22) | (1 << prqlParser.T__23) | (1 << prqlParser.T__24) | (1 << prqlParser.T__25) | (1 << prqlParser.T__26) | (1 << prqlParser.T__27))) !== 0))) {
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
	    this.enterRule(localctx, 76, prqlParser.RULE_interval);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 361;
	        this.match(prqlParser.NUMBER);
	        this.state = 362;
	        this.interval_kind();
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
prqlParser.T__9 = 10;
prqlParser.T__10 = 11;
prqlParser.T__11 = 12;
prqlParser.T__12 = 13;
prqlParser.T__13 = 14;
prqlParser.T__14 = 15;
prqlParser.T__15 = 16;
prqlParser.T__16 = 17;
prqlParser.T__17 = 18;
prqlParser.T__18 = 19;
prqlParser.T__19 = 20;
prqlParser.T__20 = 21;
prqlParser.T__21 = 22;
prqlParser.T__22 = 23;
prqlParser.T__23 = 24;
prqlParser.T__24 = 25;
prqlParser.T__25 = 26;
prqlParser.T__26 = 27;
prqlParser.T__27 = 28;
prqlParser.FUNC = 29;
prqlParser.PRQL = 30;
prqlParser.TABLE = 31;
prqlParser.PLUS = 32;
prqlParser.MINUS = 33;
prqlParser.EQUAL = 34;
prqlParser.BAR = 35;
prqlParser.COMMA = 36;
prqlParser.DOT = 37;
prqlParser.DOLLAR = 38;
prqlParser.UNDERSCORE = 39;
prqlParser.LANG = 40;
prqlParser.RANG = 41;
prqlParser.LBRACKET = 42;
prqlParser.RBRACKET = 43;
prqlParser.LPAREN = 44;
prqlParser.RPAREN = 45;
prqlParser.NULL_ = 46;
prqlParser.BOOLEAN = 47;
prqlParser.NUMBER = 48;
prqlParser.IDENT = 49;
prqlParser.IDENT_START = 50;
prqlParser.IDENT_NEXT = 51;
prqlParser.WHITESPACE = 52;
prqlParser.NEWLINE = 53;
prqlParser.COMMENT = 54;

prqlParser.RULE_nl = 0;
prqlParser.RULE_query = 1;
prqlParser.RULE_query_def = 2;
prqlParser.RULE_func_def = 3;
prqlParser.RULE_func_def_name = 4;
prqlParser.RULE_func_def_params = 5;
prqlParser.RULE_func_def_param = 6;
prqlParser.RULE_type_def = 7;
prqlParser.RULE_type_term = 8;
prqlParser.RULE_table = 9;
prqlParser.RULE_pipe = 10;
prqlParser.RULE_pipeline = 11;
prqlParser.RULE_ident_backticks = 12;
prqlParser.RULE_signed_ident = 13;
prqlParser.RULE_keyword = 14;
prqlParser.RULE_func_call = 15;
prqlParser.RULE_named_arg = 16;
prqlParser.RULE_assign = 17;
prqlParser.RULE_assign_call = 18;
prqlParser.RULE_expr_call = 19;
prqlParser.RULE_expr = 20;
prqlParser.RULE_term = 21;
prqlParser.RULE_expr_unary = 22;
prqlParser.RULE_literal = 23;
prqlParser.RULE_list = 24;
prqlParser.RULE_nested_pipeline = 25;
prqlParser.RULE_single_quote = 26;
prqlParser.RULE_multi_quote = 27;
prqlParser.RULE_string = 28;
prqlParser.RULE_range = 29;
prqlParser.RULE_operator = 30;
prqlParser.RULE_operator_unary = 31;
prqlParser.RULE_operator_mul = 32;
prqlParser.RULE_operator_add = 33;
prqlParser.RULE_operator_compare = 34;
prqlParser.RULE_operator_logical = 35;
prqlParser.RULE_operator_coalesce = 36;
prqlParser.RULE_interval_kind = 37;
prqlParser.RULE_interval = 38;

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

	query_def() {
	    return this.getTypedRuleContext(Query_defContext,0);
	};

	func_def = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(Func_defContext);
	    } else {
	        return this.getTypedRuleContext(Func_defContext,i);
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



class Query_defContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_query_def;
    }

	PRQL() {
	    return this.getToken(prqlParser.PRQL, 0);
	};

	nl() {
	    return this.getTypedRuleContext(NlContext,0);
	};

	named_arg = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(Named_argContext);
	    } else {
	        return this.getTypedRuleContext(Named_argContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterQuery_def(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitQuery_def(this);
		}
	}


}



class Func_defContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_func_def;
    }

	FUNC() {
	    return this.getToken(prqlParser.FUNC, 0);
	};

	func_def_name() {
	    return this.getTypedRuleContext(Func_def_nameContext,0);
	};

	func_def_params() {
	    return this.getTypedRuleContext(Func_def_paramsContext,0);
	};

	expr() {
	    return this.getTypedRuleContext(ExprContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterFunc_def(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitFunc_def(this);
		}
	}


}



class Func_def_nameContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_func_def_name;
    }

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	type_def() {
	    return this.getTypedRuleContext(Type_defContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterFunc_def_name(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitFunc_def_name(this);
		}
	}


}



class Func_def_paramsContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_func_def_params;
    }

	func_def_param = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(Func_def_paramContext);
	    } else {
	        return this.getTypedRuleContext(Func_def_paramContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterFunc_def_params(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitFunc_def_params(this);
		}
	}


}



class Func_def_paramContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_func_def_param;
    }

	named_arg() {
	    return this.getTypedRuleContext(Named_argContext,0);
	};

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	type_def() {
	    return this.getTypedRuleContext(Type_defContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterFunc_def_param(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitFunc_def_param(this);
		}
	}


}



class Type_defContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_type_def;
    }

	LANG() {
	    return this.getToken(prqlParser.LANG, 0);
	};

	type_term = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(Type_termContext);
	    } else {
	        return this.getTypedRuleContext(Type_termContext,i);
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
	        listener.enterType_def(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitType_def(this);
		}
	}


}



class Type_termContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_type_term;
    }

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	type_def() {
	    return this.getTypedRuleContext(Type_defContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterType_term(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitType_term(this);
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

	EQUAL() {
	    return this.getToken(prqlParser.EQUAL, 0);
	};

	nested_pipeline() {
	    return this.getTypedRuleContext(Nested_pipelineContext,0);
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

	expr_call = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(Expr_callContext);
	    } else {
	        return this.getTypedRuleContext(Expr_callContext,i);
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



class Ident_backticksContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_ident_backticks;
    }

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
	        listener.enterIdent_backticks(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitIdent_backticks(this);
		}
	}


}



class Signed_identContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_signed_ident;
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
	        listener.enterSigned_ident(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitSigned_ident(this);
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



class Func_callContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_func_call;
    }

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	named_arg = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(Named_argContext);
	    } else {
	        return this.getTypedRuleContext(Named_argContext,i);
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
	        listener.enterFunc_call(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitFunc_call(this);
		}
	}


}



class Named_argContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_named_arg;
    }

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	assign() {
	    return this.getTypedRuleContext(AssignContext,0);
	};

	expr() {
	    return this.getTypedRuleContext(ExprContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterNamed_arg(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitNamed_arg(this);
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

	EQUAL() {
	    return this.getToken(prqlParser.EQUAL, 0);
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



class Assign_callContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_assign_call;
    }

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	EQUAL() {
	    return this.getToken(prqlParser.EQUAL, 0);
	};

	expr_call() {
	    return this.getTypedRuleContext(Expr_callContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterAssign_call(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitAssign_call(this);
		}
	}


}



class Expr_callContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_expr_call;
    }

	func_call() {
	    return this.getTypedRuleContext(Func_callContext,0);
	};

	expr() {
	    return this.getTypedRuleContext(ExprContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterExpr_call(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitExpr_call(this);
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

	operator_mul() {
	    return this.getTypedRuleContext(Operator_mulContext,0);
	};

	operator_add() {
	    return this.getTypedRuleContext(Operator_addContext,0);
	};

	operator_compare() {
	    return this.getTypedRuleContext(Operator_compareContext,0);
	};

	operator_coalesce() {
	    return this.getTypedRuleContext(Operator_coalesceContext,0);
	};

	operator_logical() {
	    return this.getTypedRuleContext(Operator_logicalContext,0);
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

	expr_unary() {
	    return this.getTypedRuleContext(Expr_unaryContext,0);
	};

	list() {
	    return this.getTypedRuleContext(ListContext,0);
	};

	nested_pipeline() {
	    return this.getTypedRuleContext(Nested_pipelineContext,0);
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



class Expr_unaryContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_expr_unary;
    }

	operator_unary() {
	    return this.getTypedRuleContext(Operator_unaryContext,0);
	};

	nested_pipeline() {
	    return this.getTypedRuleContext(Nested_pipelineContext,0);
	};

	IDENT() {
	    return this.getToken(prqlParser.IDENT, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterExpr_unary(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitExpr_unary(this);
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

	string() {
	    return this.getTypedRuleContext(StringContext,0);
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

	assign_call = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(Assign_callContext);
	    } else {
	        return this.getTypedRuleContext(Assign_callContext,i);
	    }
	};

	expr_call = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(Expr_callContext);
	    } else {
	        return this.getTypedRuleContext(Expr_callContext,i);
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



class Nested_pipelineContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_nested_pipeline;
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
	        listener.enterNested_pipeline(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitNested_pipeline(this);
		}
	}


}



class Single_quoteContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_single_quote;
    }


	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterSingle_quote(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitSingle_quote(this);
		}
	}


}



class Multi_quoteContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_multi_quote;
    }


	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterMulti_quote(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitMulti_quote(this);
		}
	}


}



class StringContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_string;
    }


	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterString(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitString(this);
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

	operator_unary() {
	    return this.getTypedRuleContext(Operator_unaryContext,0);
	};

	operator_mul() {
	    return this.getTypedRuleContext(Operator_mulContext,0);
	};

	operator_add() {
	    return this.getTypedRuleContext(Operator_addContext,0);
	};

	operator_compare() {
	    return this.getTypedRuleContext(Operator_compareContext,0);
	};

	operator_logical() {
	    return this.getTypedRuleContext(Operator_logicalContext,0);
	};

	operator_coalesce() {
	    return this.getTypedRuleContext(Operator_coalesceContext,0);
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



class Operator_unaryContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_operator_unary;
    }

	MINUS() {
	    return this.getToken(prqlParser.MINUS, 0);
	};

	PLUS() {
	    return this.getToken(prqlParser.PLUS, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterOperator_unary(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitOperator_unary(this);
		}
	}


}



class Operator_mulContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_operator_mul;
    }


	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterOperator_mul(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitOperator_mul(this);
		}
	}


}



class Operator_addContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_operator_add;
    }

	PLUS() {
	    return this.getToken(prqlParser.PLUS, 0);
	};

	MINUS() {
	    return this.getToken(prqlParser.MINUS, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterOperator_add(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitOperator_add(this);
		}
	}


}



class Operator_compareContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_operator_compare;
    }

	RANG() {
	    return this.getToken(prqlParser.RANG, 0);
	};

	LANG() {
	    return this.getToken(prqlParser.LANG, 0);
	};

	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterOperator_compare(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitOperator_compare(this);
		}
	}


}



class Operator_logicalContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_operator_logical;
    }


	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterOperator_logical(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitOperator_logical(this);
		}
	}


}



class Operator_coalesceContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_operator_coalesce;
    }


	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterOperator_coalesce(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitOperator_coalesce(this);
		}
	}


}



class Interval_kindContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = prqlParser.RULE_interval_kind;
    }


	enterRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.enterInterval_kind(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof prqlListener ) {
	        listener.exitInterval_kind(this);
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

	interval_kind() {
	    return this.getTypedRuleContext(Interval_kindContext,0);
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
prqlParser.Query_defContext = Query_defContext; 
prqlParser.Func_defContext = Func_defContext; 
prqlParser.Func_def_nameContext = Func_def_nameContext; 
prqlParser.Func_def_paramsContext = Func_def_paramsContext; 
prqlParser.Func_def_paramContext = Func_def_paramContext; 
prqlParser.Type_defContext = Type_defContext; 
prqlParser.Type_termContext = Type_termContext; 
prqlParser.TableContext = TableContext; 
prqlParser.PipeContext = PipeContext; 
prqlParser.PipelineContext = PipelineContext; 
prqlParser.Ident_backticksContext = Ident_backticksContext; 
prqlParser.Signed_identContext = Signed_identContext; 
prqlParser.KeywordContext = KeywordContext; 
prqlParser.Func_callContext = Func_callContext; 
prqlParser.Named_argContext = Named_argContext; 
prqlParser.AssignContext = AssignContext; 
prqlParser.Assign_callContext = Assign_callContext; 
prqlParser.Expr_callContext = Expr_callContext; 
prqlParser.ExprContext = ExprContext; 
prqlParser.TermContext = TermContext; 
prqlParser.Expr_unaryContext = Expr_unaryContext; 
prqlParser.LiteralContext = LiteralContext; 
prqlParser.ListContext = ListContext; 
prqlParser.Nested_pipelineContext = Nested_pipelineContext; 
prqlParser.Single_quoteContext = Single_quoteContext; 
prqlParser.Multi_quoteContext = Multi_quoteContext; 
prqlParser.StringContext = StringContext; 
prqlParser.RangeContext = RangeContext; 
prqlParser.OperatorContext = OperatorContext; 
prqlParser.Operator_unaryContext = Operator_unaryContext; 
prqlParser.Operator_mulContext = Operator_mulContext; 
prqlParser.Operator_addContext = Operator_addContext; 
prqlParser.Operator_compareContext = Operator_compareContext; 
prqlParser.Operator_logicalContext = Operator_logicalContext; 
prqlParser.Operator_coalesceContext = Operator_coalesceContext; 
prqlParser.Interval_kindContext = Interval_kindContext; 
prqlParser.IntervalContext = IntervalContext; 
