from lexer.lexer import Lexer
from parser.parser import Parser
from interpreter.interpreter import Interpreter, Context, SymbolTable, Number


#####
# RUN
#####

global_symbol_table = SymbolTable()
global_symbol_table.set("nill", Number(0))
global_symbol_table.set("true", Number(1))
global_symbol_table.set("false", Number(0))


def run(fn, text):
    # Generate tokens
    lexer = Lexer(fn, text)
    tokens, error = lexer.make_tokens()
    if error:
        return None, error

    # Generate AST
    parser = Parser(tokens)
    ast = parser.parse()
    if ast.error:
        return None, ast.error

    # Run program
    interpreter = Interpreter()
    context = Context('<program>')
    context.symbol_table = global_symbol_table
    result = interpreter.visit(ast.node, context)

    # return ast.node, ast.error
    return result.value, result.error
