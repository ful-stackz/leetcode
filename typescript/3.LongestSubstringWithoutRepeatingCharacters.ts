function lengthOfLongestSubstring(s: string): number {
    if (s.length === 0) return 0
    if (s.length === 1) return 1
    if (s.length === 2) return s[0] === s[1] ? 1 : 2

    let substr = s[0]
    let max = 1
    for (let x = 1; x < s.length; x += 1) {
        const charPos = substr.indexOf(s[x])
        if (charPos < 0) {
        substr += s[x]
        } else {
        max = Math.max(max, substr.length)
        substr = substr.slice(charPos + 1) + s[x]
        if (substr.length > s.length - charPos) return max
        }
    }

    return Math.max(substr.length, max)
}
