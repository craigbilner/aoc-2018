package main

import (
	"strings"
	"testing"
)

func TestReadAndTime3(t *testing.T) {
	input := `position=< 9,  1> velocity=< 0,  2>
position=< 7,  0> velocity=<-1,  0>
position=< 3, -2> velocity=<-1,  1>
position=< 6, 10> velocity=<-2, -1>
position=< 2, -4> velocity=< 2,  2>
position=<-6, 10> velocity=< 2, -2>
position=< 1,  8> velocity=< 1, -1>
position=< 1,  7> velocity=< 1,  0>
position=<-3, 11> velocity=< 1, -2>
position=< 7,  6> velocity=<-1, -1>
position=<-2,  3> velocity=< 1,  0>
position=<-4,  3> velocity=< 2,  0>
position=<10, -3> velocity=<-1,  1>
position=< 5, 11> velocity=< 1, -2>
position=< 4,  7> velocity=< 0, -1>
position=< 8, -2> velocity=< 0,  1>
position=<15,  0> velocity=<-2,  0>
position=< 1,  6> velocity=< 1,  0>
position=< 8,  9> velocity=< 0, -1>
position=< 3,  3> velocity=<-1,  1>
position=< 0,  5> velocity=< 0, -1>
position=<-2,  2> velocity=< 2,  0>
position=< 5, -2> velocity=< 1,  2>
position=< 1,  4> velocity=< 2,  1>
position=<-2,  7> velocity=< 2, -2>
position=< 3,  6> velocity=<-1, -1>
position=< 5,  0> velocity=< 1,  0>
position=<-6,  0> velocity=< 2,  0>
position=< 5,  9> velocity=< 1, -2>
position=<14,  7> velocity=<-2,  0>
position=<-3,  6> velocity=< 2, -1>
`

	got := readAndTime(strings.NewReader(input))

	if got != 3 {
		t.Errorf("Wanted 3 got %v\n", got)
	}
}

func TestReadAndTime10159(t *testing.T) {
	input := `position=<-30302,  30614> velocity=< 3, -3>
position=<-20164, -10027> velocity=< 2,  1>
position=<-20135,  50933> velocity=< 2, -5>
position=< 40801, -40503> velocity=<-4,  4>
position=<-50633,  10291> velocity=< 5, -1>
position=< 40811, -50658> velocity=<-4,  5>
position=< 30674,  10292> velocity=<-3, -1>
position=<-40445,  30609> velocity=< 4, -3>
position=<-20155,  20447> velocity=< 2, -2>
position=<-30283, -50665> velocity=< 3,  5>
position=< -9995, -20185> velocity=< 1,  2>
position=< -9978,  20451> velocity=< 1, -2>
position=< 30671, -40507> velocity=<-3,  4>
position=<-20139,  50930> velocity=< 2, -5>
position=<-50596, -30341> velocity=< 5,  3>
position=<-40445, -40499> velocity=< 4,  4>
position=<-30318, -10030> velocity=< 3,  1>
position=<-40458,  30615> velocity=< 4, -3>
position=< 30671, -40500> velocity=<-3,  4>
position=< -9973, -20184> velocity=< 1,  2>
position=< 20513,  20451> velocity=<-2, -2>
position=< 30663, -10024> velocity=<-3,  1>
position=< 10313,  50929> velocity=<-1, -5>
position=< 50960,  20447> velocity=<-5, -2>
position=< 51001,  30609> velocity=<-5, -3>
position=< 30631, -20181> velocity=<-3,  2>
position=< 50986, -50657> velocity=<-5,  5>
position=< 20480,  20451> velocity=<-2, -2>
position=< 40798,  30611> velocity=<-4, -3>
position=< 50965,  40769> velocity=<-5, -4>
position=< 10358,  40766> velocity=<-1, -4>
position=<-20116,  50931> velocity=< 2, -5>
position=< -9960, -50659> velocity=< 1,  5>
position=<-50625,  10289> velocity=< 5, -1>
position=< 20506, -50662> velocity=<-2,  5>
position=< 10332, -50662> velocity=<-1,  5>
position=< 40806,  30612> velocity=<-4, -3>
position=< 50941,  30613> velocity=<-5, -3>
position=<-50589, -10026> velocity=< 5,  1>
position=<-40469, -10029> velocity=< 4,  1>
position=< 10305,  20455> velocity=<-1, -2>
position=<-30283,  30614> velocity=< 3, -3>
position=< 40822, -20180> velocity=<-4,  2>
position=< 50989,  30614> velocity=<-5, -3>
position=< 50983,  20451> velocity=<-5, -2>
position=< -9968,  50930> velocity=< 1, -5>
position=< -9973,  10288> velocity=< 1, -1>
position=< -9973, -40504> velocity=< 1,  4>
position=< 50967,  20451> velocity=<-5, -2>
position=<-30307,  20456> velocity=< 3, -2>
position=< 10334,  10290> velocity=<-1, -1>
position=<-50601,  50925> velocity=< 5, -5>
position=<-30323,  50924> velocity=< 3, -5>
position=< 30639,  50931> velocity=<-3, -5>
position=<-50601,  40774> velocity=< 5, -4>
position=<-30283, -20184> velocity=< 3,  2>
position=< 50941, -20183> velocity=<-5,  2>
position=<-30323, -40504> velocity=< 3,  4>
position=< 10315,  40768> velocity=<-1, -4>
position=<-20156, -20188> velocity=< 2,  2>
position=<-40445,  20452> velocity=< 4, -2>
position=< 50965, -20188> velocity=<-5,  2>
position=< -9973,  10290> velocity=< 1, -1>
position=< 40790,  30607> velocity=<-4, -3>
position=<-40456,  30611> velocity=< 4, -3>
position=< 20509, -40507> velocity=<-2,  4>
position=<-30311, -40503> velocity=< 3,  4>
position=< 10306, -10021> velocity=<-1,  1>
position=< 20504, -30345> velocity=<-2,  3>
position=<-30315,  20448> velocity=< 3, -2>
position=< -9973, -40498> velocity=< 1,  4>
position=<-20162, -30342> velocity=< 2,  3>
position=< 20497, -20180> velocity=<-2,  2>
position=<-30294, -30346> velocity=< 3,  3>
position=<-20139,  20456> velocity=< 2, -2>
position=< 50941, -20188> velocity=<-5,  2>
position=< 30683,  10288> velocity=<-3, -1>
position=< -9969,  40769> velocity=< 1, -4>
position=<-40456,  10288> velocity=< 4, -1>
position=< 30668,  40773> velocity=<-3, -4>
position=<-50633,  50931> velocity=< 5, -5>
position=< 40806, -10028> velocity=<-4,  1>
position=<-50596, -20181> velocity=< 5,  2>
position=<-10002, -30341> velocity=< 1,  3>
position=<-30323,  40774> velocity=< 3, -4>
position=<-10013, -50666> velocity=< 1,  5>
position=< 10358,  20448> velocity=<-1, -2>
position=< -9984, -20189> velocity=< 1,  2>
position=<-20160, -20188> velocity=< 2,  2>
position=< 20497, -10021> velocity=<-2,  1>
position=< 10339, -10021> velocity=<-1,  1>
position=<-40445, -30347> velocity=< 4,  3>
position=< 50994,  50930> velocity=<-5, -5>
position=<-40461,  20448> velocity=< 4, -2>
position=< -9953, -30340> velocity=< 1,  3>
position=< 10341, -30339> velocity=<-1,  3>
position=<-20115, -30339> velocity=< 2,  3>
position=<-20143, -40502> velocity=< 2,  4>
position=<-50589, -20184> velocity=< 5,  2>
position=< 30658, -50662> velocity=<-3,  5>
position=<-30331, -30341> velocity=< 3,  3>
position=< 20472, -30342> velocity=<-2,  3>
position=< 50951, -10024> velocity=<-5,  1>
position=< 20472,  50930> velocity=<-2, -5>
position=<-40438,  20451> velocity=< 4, -2>
position=< 20524,  10292> velocity=<-2, -1>
position=< 20517, -50666> velocity=<-2,  5>
position=<-30291, -10029> velocity=< 3,  1>
position=< 10315,  50927> velocity=<-1, -5>
position=< 50986, -30348> velocity=<-5,  3>
position=< 40782,  20455> velocity=<-4, -2>
position=<-30275,  30614> velocity=< 3, -3>
position=< -9964, -50662> velocity=< 1,  5>
position=< 50993, -20185> velocity=<-5,  2>
position=< 40801, -40503> velocity=<-4,  4>
position=< -9968, -30342> velocity=< 1,  3>
position=< 51001, -20187> velocity=<-5,  2>
position=<-30286, -50662> velocity=< 3,  5>
position=< 40842,  30607> velocity=<-4, -3>
position=<-40461,  40772> velocity=< 4, -4>
position=< 20492, -20185> velocity=<-2,  2>
position=<-40445,  50928> velocity=< 4, -5>
position=< 10305, -50657> velocity=<-1,  5>
position=<-50608,  10292> velocity=< 5, -1>
position=< 10305, -20182> velocity=<-1,  2>
position=<-20151, -10028> velocity=< 2,  1>
position=< 30659,  10291> velocity=<-3, -1>
position=<-50589,  30612> velocity=< 5, -3>
position=<-30291, -50662> velocity=< 3,  5>
position=< 40816, -10021> velocity=<-4,  1>
position=< -9997, -40501> velocity=< 1,  4>
position=< 30644, -30345> velocity=<-3,  3>
position=< 40834,  20451> velocity=<-4, -2>
position=<-30302,  30612> velocity=< 3, -3>
position=< 10350,  10289> velocity=<-1, -1>
position=< 20477,  10288> velocity=<-2, -1>
position=< 30647, -50657> velocity=<-3,  5>
position=<-50641, -20183> velocity=< 5,  2>
position=<-30314,  50928> velocity=< 3, -5>
position=< 20485, -40504> velocity=<-2,  4>
position=< 50981,  20449> velocity=<-5, -2>
position=<-50625, -40500> velocity=< 5,  4>
position=< 30623,  40771> velocity=<-3, -4>
position=<-30274, -40498> velocity=< 3,  4>
position=<-20156, -50662> velocity=< 2,  5>
position=< 30626,  20456> velocity=<-3, -2>
position=<-20127,  10292> velocity=< 2, -1>
position=< 40798, -20188> velocity=<-4,  2>
position=< 10314, -50662> velocity=<-1,  5>
position=< 50981,  30611> velocity=<-5, -3>
position=< 20490,  50928> velocity=<-2, -5>
position=< 30683, -20183> velocity=<-3,  2>
position=< 51001,  20449> velocity=<-5, -2>
position=<-50646, -20180> velocity=< 5,  2>
position=< 40816,  10293> velocity=<-4, -1>
position=< 10323,  40765> velocity=<-1, -4>
position=<-20156, -20187> velocity=< 2,  2>
position=<-50637,  10296> velocity=< 5, -1>
position=<-20111, -20189> velocity=< 2,  2>
position=< 40815,  10288> velocity=<-4, -1>
position=<-20132, -40505> velocity=< 2,  4>
position=< 40784,  10297> velocity=<-4, -1>
position=< 50989, -40503> velocity=<-5,  4>
position=< 30684, -50666> velocity=<-3,  5>
position=< 40816, -50661> velocity=<-4,  5>
position=<-20138,  30615> velocity=< 2, -3>
position=<-40466,  40765> velocity=< 4, -4>
position=<-40445,  20456> velocity=< 4, -2>
position=<-30314,  30606> velocity=< 3, -3>
position=< 40790, -40502> velocity=<-4,  4>
position=<-20170,  30615> velocity=< 2, -3>
position=<-20143, -10030> velocity=< 2,  1>
position=< 10329, -50658> velocity=<-1,  5>
position=<-10004, -40502> velocity=< 1,  4>
position=<-30323, -20189> velocity=< 3,  2>
position=<-50612, -30348> velocity=< 5,  3>
position=<-20143,  40771> velocity=< 2, -4>
position=< 40784,  20456> velocity=<-4, -2>
position=< 20477,  30615> velocity=<-2, -3>
position=< 40814, -10022> velocity=<-4,  1>
position=< 10334,  40767> velocity=<-1, -4>
position=< 10329,  10296> velocity=<-1, -1>
position=< 40811, -20180> velocity=<-4,  2>
position=<-40477, -30339> velocity=< 4,  3>
position=<-20148, -10024> velocity=< 2,  1>
position=<-40445,  20454> velocity=< 4, -2>
position=< 50941,  20449> velocity=<-5, -2>
position=< 10316,  50926> velocity=<-1, -5>
position=<-50609,  20453> velocity=< 5, -2>
position=<-40431,  30615> velocity=< 4, -3>
position=< 50970, -10028> velocity=<-5,  1>
position=<-20112, -50666> velocity=< 2,  5>
position=< 30668, -10022> velocity=<-3,  1>
position=< 20501,  50933> velocity=<-2, -5>
position=<-50604, -10025> velocity=< 5,  1>
position=<-20135,  30606> velocity=< 2, -3>
position=<-20167, -20180> velocity=< 2,  2>
position=<-30323,  30614> velocity=< 3, -3>
position=< 30636, -40498> velocity=<-3,  4>
position=< 40786, -50657> velocity=<-4,  5>
position=<-50609, -30343> velocity=< 5,  3>
position=< 40814,  50924> velocity=<-4, -5>
position=<-20137, -50666> velocity=< 2,  5>
position=<-50609,  50924> velocity=< 5, -5>
position=< 40835, -30344> velocity=<-4,  3>
position=< 10324, -30348> velocity=<-1,  3>
position=<-20172, -10027> velocity=< 2,  1>
position=< 50941,  50929> velocity=<-5, -5>
position=<-50630, -40503> velocity=< 5,  4>
position=< 40842, -10029> velocity=<-4,  1>
position=<-50601,  10297> velocity=< 5, -1>
position=< 51001,  20452> velocity=<-5, -2>
position=<-40461, -50665> velocity=< 4,  5>
position=<-40466,  50932> velocity=< 4, -5>
position=< 30683,  10296> velocity=<-3, -1>
position=< 40782,  40771> velocity=<-4, -4>
position=< -9953,  50926> velocity=< 1, -5>
position=<-40466, -20180> velocity=< 4,  2>
position=< 30671, -40501> velocity=<-3,  4>
position=< 50970,  40773> velocity=<-5, -4>
position=<-30307, -30341> velocity=< 3,  3>
position=< 10364, -50666> velocity=<-1,  5>
position=<-40466, -40505> velocity=< 4,  4>
position=< 30639, -10024> velocity=<-3,  1>
position=< 30623,  50928> velocity=<-3, -5>
position=< 50957,  10296> velocity=<-5, -1>
position=<-20137, -20180> velocity=< 2,  2>
position=< 50968, -50662> velocity=<-5,  5>
position=<-40434,  40773> velocity=< 4, -4>
position=< 50965,  30613> velocity=<-5, -3>
position=< -9965, -10025> velocity=< 1,  1>
position=<-40437,  40770> velocity=< 4, -4>
position=<-40473,  50928> velocity=< 4, -5>
position=< 10313,  10290> velocity=<-1, -1>
position=< 40786,  10297> velocity=<-4, -1>
position=< 50970,  50927> velocity=<-5, -5>
position=< 50970,  20454> velocity=<-5, -2>
position=< -9981, -30339> velocity=< 1,  3>
position=<-40437, -20183> velocity=< 4,  2>
position=< 40811,  40769> velocity=<-4, -4>
position=< 30671, -20182> velocity=<-3,  2>
position=< 40790,  10289> velocity=<-4, -1>
position=< 50965, -50663> velocity=<-5,  5>
position=< 30676, -50662> velocity=<-3,  5>
position=< 30644, -10028> velocity=<-3,  1>
position=<-40474,  40767> velocity=< 4, -4>
position=< 40819, -20188> velocity=<-4,  2>
position=<-20132, -10029> velocity=< 2,  1>
position=<-40437,  20447> velocity=< 4, -2>
position=<-50617,  40772> velocity=< 5, -4>
position=< 40790,  50931> velocity=<-4, -5>
position=< 30623, -40498> velocity=<-3,  4>
position=< 20496,  40765> velocity=<-2, -4>
position=< 40806,  30611> velocity=<-4, -3>
position=< 40802, -30348> velocity=<-4,  3>
position=<-30291, -30347> velocity=< 3,  3>
position=< 20500, -10030> velocity=<-2,  1>
position=< 20517, -10030> velocity=<-2,  1>
position=< 50994, -10027> velocity=<-5,  1>
position=< 50949,  30613> velocity=<-5, -3>
position=< -9965,  10295> velocity=< 1, -1>
position=<-30287, -10026> velocity=< 3,  1>
position=< 40798, -20180> velocity=<-4,  2>
position=< 50978, -40506> velocity=<-5,  4>
position=<-40486,  10297> velocity=< 4, -1>
position=<-20155,  40769> velocity=< 2, -4>
position=< 10326,  40768> velocity=<-1, -4>
position=< 40798, -10022> velocity=<-4,  1>
position=< 50970,  50925> velocity=<-5, -5>
position=< 20509, -50663> velocity=<-2,  5>
position=<-30283,  40768> velocity=< 3, -4>
position=<-40430, -10022> velocity=< 4,  1>
position=<-20137,  40769> velocity=< 2, -4>
position=<-20143, -40504> velocity=< 2,  4>
position=< 30663, -10024> velocity=<-3,  1>
position=< 30655, -40499> velocity=<-3,  4>
position=< -9969,  40769> velocity=< 1, -4>
position=< 50994,  50933> velocity=<-5, -5>
position=< 30676,  20449> velocity=<-3, -2>
position=< -9968, -10024> velocity=< 1,  1>
position=< 50984, -20185> velocity=<-5,  2>
position=< 30631,  30607> velocity=<-3, -3>
position=<-30299, -10022> velocity=< 3,  1>
position=<-50601, -30343> velocity=< 5,  3>
position=<-10004, -50662> velocity=< 1,  5>
position=< 30668, -50665> velocity=<-3,  5>
position=<-20112,  40772> velocity=< 2, -4>
position=<-30311,  30606> velocity=< 3, -3>
position=<-20127, -30346> velocity=< 2,  3>
position=< 30681, -40498> velocity=<-3,  4>
position=<-40450, -10023> velocity=< 4,  1>
position=<-40445,  10295> velocity=< 4, -1>
position=<-30310, -40505> velocity=< 3,  4>
position=< 10331,  40769> velocity=<-1, -4>
position=<-40461,  30612> velocity=< 4, -3>
position=<-20112, -50665> velocity=< 2,  5>
position=< 30639,  30611> velocity=<-3, -3>
position=< 20465,  10297> velocity=<-2, -1>
position=< 40830, -40501> velocity=<-4,  4>
position=< 50998, -50657> velocity=<-5,  5>
position=< 10305,  20447> velocity=<-1, -2>
position=< 10305,  10297> velocity=<-1, -1>
position=<-20132,  30615> velocity=< 2, -3>
position=<-20119, -30341> velocity=< 2,  3>
position=< 40790,  40769> velocity=<-4, -4>
position=< 40830, -10022> velocity=<-4,  1>
position=<-30302,  10293> velocity=< 3, -1>
position=< 50949, -30340> velocity=<-5,  3>
position=< 10345, -50662> velocity=<-1,  5>
position=< 30676, -30346> velocity=<-3,  3>
position=<-30323,  40767> velocity=< 3, -4>
position=<-40471,  50924> velocity=< 4, -5>
position=< 10334, -40507> velocity=<-1,  4>
position=< 30631,  50924> velocity=<-3, -5>
position=<-40461, -30339> velocity=< 4,  3>
position=<-40474,  50924> velocity=< 4, -5>
position=<-20148,  20451> velocity=< 2, -2>
position=< 51002,  30606> velocity=<-5, -3>
position=< 10364,  40765> velocity=<-1, -4>
position=<-50612, -40506> velocity=< 5,  4>
position=<-30321, -20183> velocity=< 3,  2>
position=< -9963, -30344> velocity=< 1,  3>
position=< 50966,  40769> velocity=<-5, -4>
position=< 40799,  40765> velocity=<-4, -4>
position=<-20112, -10030> velocity=< 2,  1>
position=<-40431,  10288> velocity=< 4, -1>
position=<-30283, -50662> velocity=< 3,  5>
position=< 40798,  40770> velocity=<-4, -4>
position=<-50625,  50930> velocity=< 5, -5>
position=< 40810,  50928> velocity=<-4, -5>
position=<-30286,  40774> velocity=< 3, -4>
position=<-30331, -20188> velocity=< 3,  2>
position=<-30283,  40774> velocity=< 3, -4>
position=< -9973,  50924> velocity=< 1, -5>
position=< 30671, -20187> velocity=<-3,  2>
position=< 10329, -30345> velocity=<-1,  3>
position=< 40835,  20455> velocity=<-4, -2>
position=< 10325, -20185> velocity=<-1,  2>
position=< 30639, -20180> velocity=<-3,  2>
position=< 40822,  50932> velocity=<-4, -5>
position=<-50601,  10292> velocity=< 5, -1>
position=<-20127, -30341> velocity=< 2,  3>
`

	got := readAndTime(strings.NewReader(input))

	if got != 10159 {
		t.Errorf("Wanted 10159 got %v\n", got)
	}
}