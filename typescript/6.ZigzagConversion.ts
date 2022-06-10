function convert(s: string, numRows: number): string {
    if (numRows === 1) return s

    const rows = [] as string[][]

	let row = 0
	let goingDown = false
	for (let i = 0; i < s.length; i++) {
		rows[row] ??= []
		rows[row].push(s[i])

		if (row === 0 || row === numRows - 1)
            goingDown = !goingDown

		row += goingDown ? 1 : -1
	}

	return rows.reduce((prev, curr) => [...prev, curr.join('')], []).join('')
}
