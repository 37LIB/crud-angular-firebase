package lemmatize

var ENG = "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