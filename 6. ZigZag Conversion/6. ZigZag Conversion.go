package convert

/*
 * set numRows := 5 then
 * row0: 0,  ,  ,  , 8,  ,  ,  ,16, ... (0,8,16,...)
 * row1: 1,  ,  , 7, 9,  ,  ,15,17, ... (1,7,9,15,17,...)
 * row2: 2,  , 6,  ,10,  ,14,  ,18, ... (2,6,10,14,18,...)
 * row3: 3, 5,  ,  ,11,13,  ,  ,19, ... (3,5,11,13,19,...)
 * row4: 4,  ,  ,  ,12,  ,  ,  ,20, ... (4,12,20,...)
 * we can set a num `p` := numRows * 2 - 2
 * then we can get an equation
 * the first line: 0*p, 1*p, 2*p, ...
 * line r:          0*p+r, 1*p-r, 1*p+r, ...
 * the last line:  (numRow-1)+0*p, (numRow-1)+1*p, (numRow-1)+2*p, ...
 */

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	bytes := []byte(s)
	rows := make([][]byte, numRows)

	// set p
	p := numRows*2 - 2

	//first line
	for i := 0; i < len(s); i += p {
		rows[0] = append(rows[0], bytes[i])
	}

	//line r
	for r := 1; r <= numRows-2; r++ {
		rows[r] = append(rows[r], bytes[r])
		for i := p; i-r < len(s); i += p {
			rows[r] = append(rows[r], bytes[i-r])
			if i+r < len(s) {
				rows[r] = append(rows[r], bytes[i+r])
			}
		}
	}

	//last line
	for i := numRows - 1; i < len(s); i += p {
		rows[numRows-1] = append(rows[numRows-1], bytes[i])
	}

	// return result
	res := []byte{}
	for i := range rows {
		res = append(res, rows[i]...)
	}
	return string(res)
}
