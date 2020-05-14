from lexer.lexer import Lexer
from parser.parser import Parser

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

    return ast.node, ast.error
