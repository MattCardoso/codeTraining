def is_pangram(sentence):
    letters = dict()
    for c in sentence.lower():
        if c not in letters and c is not None and c.isalpha():
            letters[c] = 1
    return True if len(letters) == 26 else False
