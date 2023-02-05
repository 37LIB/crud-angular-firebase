package normalize

var cmnPhrases = "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"
var cmnCharacters = "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