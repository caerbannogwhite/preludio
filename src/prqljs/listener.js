import antlr4 from "antlr4";

import prqlListener from "../grammar/prqlListener.js";
import prqlLexer from "../grammar/prqlLexer.js";
import prqlParser from "../grammar/prqlParser.js";

export class Prqljs extends prqlListener {
  constructor() {
    super();

    this.funcDefParams = {};
  }

  enterQuery(ctx) {
  }

  exitQuery(ctx) {
  }

  // Enter a parse tree produced by prqlParser#query_def.
  enterQuery_def(ctx) { }

  // Exit a parse tree produced by prqlParser#query_def.
  exitQuery_def(ctx) { }

  // Enter a parse tree produced by prqlParser#func_def.
  enterFunc_def(ctx) { }

  // Exit a parse tree produced by prqlParser#func_def.
  exitFunc_def(ctx) { }

  // Enter a parse tree produced by prqlParser#func_def_name.
  enterFunc_def_name(ctx) { }

  // Exit a parse tree produced by prqlParser#func_def_name.
  exitFunc_def_name(ctx) { }

  // Enter a parse tree produced by prqlParser#func_def_params.
  enterFunc_def_params(ctx) {
    this.funcDefParams = {};
  }

  // Exit a parse tree produced by prqlParser#func_def_params.
  exitFunc_def_params(ctx) { }

  // Enter a parse tree produced by prqlParser#func_def_param.
  enterFunc_def_param(ctx) { }

  // Exit a parse tree produced by prqlParser#func_def_param.
  exitFunc_def_param(ctx) {
    if (ctx.IDENT()) {

    }
  }

  // Enter a parse tree produced by prqlParser#type_def.
  enterType_def(ctx) { }

  // Exit a parse tree produced by prqlParser#type_def.
  exitType_def(ctx) { }

  // Enter a parse tree produced by prqlParser#type_term.
  enterType_term(ctx) { }

  // Exit a parse tree produced by prqlParser#type_term.
  exitType_term(ctx) { }

  // Enter a parse tree produced by prqlParser#table.
  enterTable(ctx) { }

  // Exit a parse tree produced by prqlParser#table.
  exitTable(ctx) { }

  // Enter a parse tree produced by prqlParser#pipe.
  enterPipe(ctx) { }

  // Exit a parse tree produced by prqlParser#pipe.
  exitPipe(ctx) { }

  // Enter a parse tree produced by prqlParser#pipeline.
  enterPipeline(ctx) { }

  // Exit a parse tree produced by prqlParser#pipeline.
  exitPipeline(ctx) { }

  // Enter a parse tree produced by prqlParser#ident_backticks.
  enterIdent_backticks(ctx) { }

  // Exit a parse tree produced by prqlParser#ident_backticks.
  exitIdent_backticks(ctx) { }

  // Enter a parse tree produced by prqlParser#signed_ident.
  enterSigned_ident(ctx) { }

  // Exit a parse tree produced by prqlParser#signed_ident.
  exitSigned_ident(ctx) { }

  // Enter a parse tree produced by prqlParser#keyword.
  enterKeyword(ctx) { }

  // Exit a parse tree produced by prqlParser#keyword.
  exitKeyword(ctx) { }

  // Enter a parse tree produced by prqlParser#func_call.
  enterFunc_call(ctx) {

  }

  // Exit a parse tree produced by prqlParser#func_call.
  exitFunc_call(ctx) {

    // console.log(ctx.IDENT());

    const funcName = ctx.IDENT().getText().toLowerCase();
    switch (funcName) {
      case "average":
        break;

      case "aggregate":
        break;

      case "from":
        break;

      case "derive":
        break;

      case "group":
        break;

      case "import":
        break;

      case "sort":
        break;

      default:
        console.log(`function "${funcName}" NON IMPLEMENTED`);
        break;
    }
  }

  // Enter a parse tree produced by prqlParser#named_arg.
  enterNamed_arg(ctx) { }

  // Exit a parse tree produced by prqlParser#named_arg.
  exitNamed_arg(ctx) { }

  // Enter a parse tree produced by prqlParser#assign.
  enterAssign(ctx) { }

  // Exit a parse tree produced by prqlParser#assign.
  exitAssign(ctx) { }

  // Enter a parse tree produced by prqlParser#assign_call.
  enterAssign_call(ctx) { }

  // Exit a parse tree produced by prqlParser#assign_call.
  exitAssign_call(ctx) { }

  // Enter a parse tree produced by prqlParser#expr_call.
  enterExpr_call(ctx) { }

  // Exit a parse tree produced by prqlParser#expr_call.
  exitExpr_call(ctx) { }

  // Enter a parse tree produced by prqlParser#expr.
  enterExpr(ctx) { }

  // Exit a parse tree produced by prqlParser#expr.
  exitExpr(ctx) { }

  // Enter a parse tree produced by prqlParser#expr_coalesce.
  enterExpr_coalesce(ctx) { }

  // Exit a parse tree produced by prqlParser#expr_coalesce.
  exitExpr_coalesce(ctx) { }

  // Enter a parse tree produced by prqlParser#expr_compare.
  enterExpr_compare(ctx) { }

  // Exit a parse tree produced by prqlParser#expr_compare.
  exitExpr_compare(ctx) { }

  // Enter a parse tree produced by prqlParser#expr_add.
  enterExpr_add(ctx) { }

  // Exit a parse tree produced by prqlParser#expr_add.
  exitExpr_add(ctx) { }

  // Enter a parse tree produced by prqlParser#expr_mul.
  enterExpr_mul(ctx) { }

  // Exit a parse tree produced by prqlParser#expr_mul.
  exitExpr_mul(ctx) { }

  // Enter a parse tree produced by prqlParser#term.
  enterTerm(ctx) { }

  // Exit a parse tree produced by prqlParser#term.
  exitTerm(ctx) { }

  // Enter a parse tree produced by prqlParser#expr_unary.
  enterExpr_unary(ctx) { }

  // Exit a parse tree produced by prqlParser#expr_unary.
  exitExpr_unary(ctx) { }

  // Enter a parse tree produced by prqlParser#literal.
  enterLiteral(ctx) { }

  // Exit a parse tree produced by prqlParser#literal.
  exitLiteral(ctx) {
    console.log(ctx);
  }

  // Enter a parse tree produced by prqlParser#list.
  enterList(ctx) { }

  // Exit a parse tree produced by prqlParser#list.
  exitList(ctx) { }

  // Enter a parse tree produced by prqlParser#nested_pipeline.
  enterNested_pipeline(ctx) { }

  // Exit a parse tree produced by prqlParser#nested_pipeline.
  exitNested_pipeline(ctx) { }

  // Enter a parse tree produced by prqlParser#single_quote.
  enterSingle_quote(ctx) { }

  // Exit a parse tree produced by prqlParser#single_quote.
  exitSingle_quote(ctx) { }

  // Enter a parse tree produced by prqlParser#multi_quote.
  enterMulti_quote(ctx) { }

  // Exit a parse tree produced by prqlParser#multi_quote.
  exitMulti_quote(ctx) { }

  // Enter a parse tree produced by prqlParser#string.
  enterString(ctx) { }

  // Exit a parse tree produced by prqlParser#string.
  exitString(ctx) { }

  // Enter a parse tree produced by prqlParser#number.
  enterNumber(ctx) { }

  // Exit a parse tree produced by prqlParser#number.
  exitNumber(ctx) { }

  // Enter a parse tree produced by prqlParser#range.
  enterRange(ctx) { }

  // Exit a parse tree produced by prqlParser#range.
  exitRange(ctx) { }

  // Enter a parse tree produced by prqlParser#range_edge.
  enterRange_edge(ctx) { }

  // Exit a parse tree produced by prqlParser#range_edge.
  exitRange_edge(ctx) { }

  // Enter a parse tree produced by prqlParser#operator.
  enterOperator(ctx) { }

  // Exit a parse tree produced by prqlParser#operator.
  exitOperator(ctx) { }

  // Enter a parse tree produced by prqlParser#operator_unary.
  enterOperator_unary(ctx) { }

  // Exit a parse tree produced by prqlParser#operator_unary.
  exitOperator_unary(ctx) { }

  // Enter a parse tree produced by prqlParser#operator_mul.
  enterOperator_mul(ctx) { }

  // Exit a parse tree produced by prqlParser#operator_mul.
  exitOperator_mul(ctx) { }

  // Enter a parse tree produced by prqlParser#operator_add.
  enterOperator_add(ctx) { }

  // Exit a parse tree produced by prqlParser#operator_add.
  exitOperator_add(ctx) { }

  // Enter a parse tree produced by prqlParser#operator_compare.
  enterOperator_compare(ctx) { }

  // Exit a parse tree produced by prqlParser#operator_compare.
  exitOperator_compare(ctx) { }

  // Enter a parse tree produced by prqlParser#operator_logical.
  enterOperator_logical(ctx) { }

  // Exit a parse tree produced by prqlParser#operator_logical.
  exitOperator_logical(ctx) { }

  // Enter a parse tree produced by prqlParser#operator_coalesce.
  enterOperator_coalesce(ctx) { }

  // Exit a parse tree produced by prqlParser#operator_coalesce.
  exitOperator_coalesce(ctx) { }
}

export default function runPrqljs(source) {
  const { CommonTokenStream, InputStream } = antlr4;

  var chars = new InputStream(source, true);
  var lexer = new prqlLexer(chars);
  var tokens = new CommonTokenStream(lexer);
  var parser = new prqlParser(tokens);

  parser.buildParseTrees = true;
  var tree = parser.query();
  var transpiler = new Prqljs();
  antlr4.tree.ParseTreeWalker.DEFAULT.walk(transpiler, tree);
}
