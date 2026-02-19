package ascii

import "malasfetch/tampilan"

type FungsiASCII func() ([]string, string)

func DapatkanLogoArch() ([]string, string) {
	s1 := tampilan.Sian
	s2 := tampilan.SianCerah
	r := tampilan.Reset
	t := tampilan.Tebal

	return []string{
		t + s1 + "                   -`" + r,
		t + s1 + "                 .o+`" + r,
		t + s1 + "                `ooo/" + r,
		t + s1 + "               `+oooo:" + r,
		t + s1 + "              `+oooooo:" + r,
		t + s1 + "              -+oooooo+:" + r,
		t + s1 + "            `/:-:++oooo+:" + r,
		t + s1 + "           `/++++/+++++++:" + r,
		t + s1 + "          `/++++++++++++++:" + r,
		t + s1 + "         `/+++o" + s2 + "oooooooo" + s1 + "oooo/`" + r,
		t + s1 + "        ./" + s2 + "ooosssso++osssssso" + s1 + "+`" + r,
		t + s2 + "       .oossssso-````/ossssss+" + r,
		t + s2 + "      -osssssso.      :ssssssso." + r,
		t + s2 + "     :osssssss/        osssso+++." + r,
		t + s2 + "    /ossssssss/        +ssssooo/-" + r,
		t + s2 + "  `/ossssso+/:-        -:/+osssso+-" + r,
		t + s2 + " `+sso+:-`                 `.-/+oso:" + r,
		t + s2 + "++:.                           `-/+/" + r,
		t + s2 + ".`                                 `/" + r,
	}, ""
}

func DapatkanLogoUbuntu() ([]string, string) {
	m1 := tampilan.Merah
	m2 := tampilan.MerahCerah
	r := tampilan.Reset
	t := tampilan.Tebal

	return []string{
		"                                              ....",
		t + m2 + "              .',:clooo:  " + m1 + ".:looooo:." + r,
		t + m2 + "           .;looooooooc  " + m1 + ".oooooooooo'" + r,
		t + m2 + "        .;looooool:,''.  " + m1 + ":ooooooooooc" + r,
		t + m2 + "       ;looool;.         " + m1 + "'oooooooooo," + r,
		t + m2 + "      ;clool'             " + m1 + ".cooooooc.  " + m2 + ",," + r,
		t + m2 + "         ...                " + m1 + "......  " + m2 + ".:oo," + r,
		t + m2 + "  " + t + m1 + ".;clol:,.                        " + t + m2 + ".loooo'" + r,
		t + m2 + " " + t + m1 + ":ooooooooo,                        " + t + m2 + "'ooool" + r,
		t + m2 + " " + t + m1 + "'ooooooooooo.                        " + t + m2 + "loooo." + r,
		t + m2 + " " + t + m1 + "'ooooooooool                         " + t + m2 + "coooo." + r,
		t + m2 + " " + t + m1 + ",loooooooc.                        " + t + m2 + ".loooo." + r,
		t + m2 + "   " + t + m1 + ".,;;;'.                          " + t + m2 + ";ooooc" + r,
		t + m2 + "       ...                         ,ooool." + r,
		t + m2 + "    .cooooc.              " + m1 + "..',,'.  " + m2 + ".cooo." + r,
		t + m2 + "      ;ooooo:.           " + m1 + ";oooooooc.  " + m2 + ":l." + r,
		t + m2 + "       .coooooc,..      " + m1 + "coooooooooo." + r,
		t + m2 + "         .:ooooooolc:. " + m1 + ".ooooooooooo'" + r,
		t + m2 + "           .':loooooo;  " + m1 + ",oooooooooc" + r,
		t + m2 + "               ..';::c'  " + m1 + ".;loooo:'" + r,
	}, ""
}

func DapatkanLogoFedora() ([]string, string) {
	return []string{
		"             .',;::::;,'.      ",
		"         .';:cccccccccccc:;,.  ",
		"      .;cccccccccccccccccccccc;.",
		"    .:ccccccccccccccccccccccccccc:.",
		"  .;ccccccccccccc;.:dddl:.;ccccccc;.",
		" .:ccccccccccccc;OWMKOOXMWd;ccccccc:.",
		".:ccccccccccccc;KMMc;cc;xMMc;ccccccc:.",
		",cccccccccccccc;MMM.;cc;;WW:;cccccccc,",
		":cccccccccccccc;MMM.;cccccccccccccccc:",
		":ccccccc;oxOOOo;MMM000k.;cccccccccccc:",
		"cccccc;0MMKxdd:;MMMkddc.;cccccccccccc;",
		"ccccc;XMO';cccc;MMM.;cccccccccccccccc'",
		"ccccc;MMo;ccccc;MMW.;ccccccccccccccc;",
		"ccccc;0MNc.ccc.xMMd;ccccccccccccccc;",
		"cccccc;dNMWXXXWM0:;cccccccccccccc:, ",
		"cccccccc;.:odl:.;cccccccccccccc:,.  ",
		"ccccccccccccccccccccccccccccc:'.    ",
		":ccccccccccccccccccc:;,..       ",
		" ':cccccccccccccccc::;,.            ",
	}, tampilan.Biru
}

func DapatkanLogoDefault() ([]string, string) {
	s := tampilan.SianCerah
	b := tampilan.BiruCerah
	r := tampilan.Reset
	t := tampilan.Tebal

	return []string{
		"                  " + t + b + "****************" + r,
		"             " + t + b + "************************" + r,
		"          " + t + b + "******************************" + r,
		"        " + t + b + "**********************************" + r,
		"      " + t + b + "**************************************" + r,
		"    " + t + b + "**********************" + t + s + "+-:...:::-" + t + b + "**********" + r,
		"   " + t + b + "********************" + t + s + "+............+" + t + b + "**********" + r,
		"  " + t + b + "********************" + t + s + "-.............+" + t + b + "***********" + r,
		" " + t + b + "********************" + t + s + "=..............+" + t + b + "************" + r,
		" " + t + b + "********************" + t + s + "........:+" + t + b + "*******************" + r,
		t + b + "*********************" + t + s + "........+" + t + b + "********************" + r,
		t + b + "********************" + t + s + "+........+" + t + b + "********************" + r,
		t + b + "********************" + t + s + "+........+" + t + b + "********************" + r,
		t + b + "**************" + t + s + "+------........------=" + t + b + "**************" + r,
		t + b + "**************" + t + s + "=....................:" + t + b + "**************" + r,
		t + b + "**************" + t + s + "=....................=" + t + b + "**************" + r,
		t + b + "**************" + t + s + "=...................." + t + b + "***************" + r,
		" " + t + b + "*************" + t + s + "+======........======" + t + b + "***************" + r,
		" " + t + b + "*******************" + t + s + "+........+" + t + b + "*******************" + r,
		"  " + t + b + "******************" + t + s + "+........+" + t + b + "******************" + r,
		"   " + t + b + "*****************" + t + s + "+........+" + t + b + "*****************" + r,
		"    " + t + b + "****************" + t + s + "+........+" + t + b + "****************" + r,
		"      " + t + b + "**************" + t + s + "+........+" + t + b + "**************" + r,
		"        " + t + b + "************" + t + s + "+........+" + t + b + "************" + r,
		"          " + t + b + "**********" + t + s + "+........+" + t + b + "**********" + r,
		"             " + t + b + "*******" + t + s + "+........+" + t + b + "*******" + r,
		"                 " + t + b + "***" + t + s + "+........+" + t + b + "***" + r,
	}, ""
}
