from lexer.lexer import Lexer

#####
# RUN
#####

def run(fn, text):
    lexer = Lexer(fn, text)
    tokens, error = lexer.make_tokens()

    return tokens, error