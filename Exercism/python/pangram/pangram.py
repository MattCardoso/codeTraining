def is_pangram(sentence):
    letters = {}
    # "zaz\sd".
    for c in sentence.lower():
        if letters[c] is None:
            letters[c] = 1
    return True if len(letters.keys()) == 26 else False
