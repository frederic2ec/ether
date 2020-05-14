from lexer.lexer import Lexer
from parser.parser import Parser
from interpreter.interpreter import Interpreter, Context

#####
# RUN
#####


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
    result = interpreter.visit(ast.node, context)

    # return ast.node, ast.error
    return result.value, result.error
