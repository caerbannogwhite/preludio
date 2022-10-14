// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import * as vscode from "vscode";
import { transpile } from "./prql2sas/prql2sas.js";

// this method is called when your extension is activated
// your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {
  // The command has been defined in the package.json file
  // Now provide the implementation of the command with registerCommand
  // The commandId parameter must match the command field in package.json
  let disposable = vscode.commands.registerCommand("prqlvs.Run", () => {
    const textEditor = vscode.window.activeTextEditor;
    if (textEditor) {
      const endOfLine = "\n";
      const firstLine = textEditor.selection.start.line;
      const lastLine = textEditor.selection.end.line;

      let selectedText =
        textEditor.document
          .lineAt(firstLine)
          .text.slice(textEditor.selection.start.character) + endOfLine;
      for (let i = firstLine + 1; i < lastLine; ++i) {
        selectedText += textEditor.document.lineAt(i).text + endOfLine;
      }
      selectedText += textEditor.document
        .lineAt(lastLine)
        .text.slice(0, textEditor.selection.end.character);

      transpile(selectedText);
    }

    vscode.window.showInformationMessage("PRQL code succesfully executed.");
  });

  context.subscriptions.push(disposable);
}

// this method is called when your extension is deactivated
export function deactivate() {}
