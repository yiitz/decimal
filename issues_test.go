package decimal

import (
	"math"
	"math/big"
	"testing"
)

func TestIssue20(t *testing.T) {
	x := New(10240000000000, 0)
	x.Mul(x, New(976563, 9))
	if v, _ := x.Int64(); v != 10000005120 {
		t.Fatal("error int64: ", v, x.Int(nil).Int64())
	}
}

func TestIssue65(t *testing.T) {
	const expected = "999999999000000000000000000000"
	r, _ := new(big.Rat).SetString(expected)
	r2 := new(Big).SetRat(r).Rat(nil)
	if r.Cmp(r2) != 0 {
		t.Fatalf("expected %q, got %q", r, r2)
	}
}

func TestIssue71(t *testing.T) {
	x, _ := new(Big).SetString("-433997231707950814777029946371807573425840064343095193931191306942897586882.200850175108941825587256711340679426793690849230895605323379098449524300541372392806145820741928")
	y := New(5, 0)
	ctx := Context{RoundingMode: ToZero, Precision: 364}

	z := new(Big)
	ctx.Quo(z, x, y)

	r, _ := new(Big).SetString("-86799446341590162955405989274361514685168012868619038786238261388579517376.4401700350217883651174513422681358853587381698461791210646758196899048601082744785612291641483856")
	if z.Cmp(r) != 0 || z.Scale() != r.Scale() {
		t.Fatalf(`Quo(%s, %s)
wanted: %s (%d)
got   : %s (%d)
`, x, y, r, -r.Scale(), z, -z.Scale())
	}
}

func TestIssue72(t *testing.T) {
	x, _ := new(Big).SetString("-8.45632792449080367076920780185655231664817924617196338687858969707575095137356626097186102468204972266270655439471710162223657196797091956190618036249568250856310899052975275153410779062120467574000771085625757386351708361318971283364474972153263288762761014798575650687906566E+474")
	y, _ := new(Big).SetString("4394389707820271499500265597691058417332780189928068605060129835915231607024733174128123086028964659911263805538968425927408117535552905751413991847682423230507052480632597367974353369255973450023914.06480266537851511912348920528179447782332532576762774258658035423323623047681531444628650113938865866058071268742035039370988065347285125745597527162817805470262344343643075954571122548882320506470664701832116848314413975179616459225485097673077072340532232446317251990415268245406080149594165531067657351225251495644780695372152557650401209918010537469259193951365404947434164664325966741900020673085975334136592934327584453217952431999450960191719318690339387778325911")
	z := new(Big)
	ctx := Context{Precision: 276}
	ctx.Rem(z, x, y)
	if !z.IsNaN(+1) {
		t.Fatalf(`Rem(%s, %s)
wanted: NaN
got   : %s
`, x, y, z)
	}
}

func TestIssue99(t *testing.T) {
	bad := &Big{compact: 11234567890123457890, exp: -20, precision: 20}
	ref := &Big{compact: 2, exp: -1, precision: 1}

	cmp := bad.Cmp(ref)
	if cmp != -1 {
		t.Errorf("expected %s smaller than %s\n", bad.String(), ref.String())
	}
	cmp = ref.Cmp(bad)
	if cmp != 1 {
		t.Errorf("expected %s larger than %s\n", ref.String(), bad.String())
	}
}

func TestIssue104(t *testing.T) {
	x := WithContext(Context128).SetUint64(math.MaxUint64)
	y, ok := x.Uint64()
	if !ok {
		t.Error("failed to convert back")
	}
	if y != math.MaxUint64 {
		t.Errorf("conversion 'succeeded' but the value failed to make the round trip - got %d", y)
	}
}

func TestIssue105(t *testing.T) {
	var b Big
	b.SetString("6190.000000000000")
	if _, ok := b.Float64(); !ok {
		t.Fatal("6190 should fit in a float just fine")
	}
}

func TestIssue114(t *testing.T) {
	val := New(-1, 0)
	f, ok := val.Float64()
	if !ok {
		t.Fatal("expected true, got false")
	}
	if f != -1 {
		t.Fatalf("expected -1, got %f", f)
	}
}

func TestIssue129(t *testing.T) {
	const ys = "3032043016321464119267109897502707536081241662295108925759281083" +
		"5908762948729460330525095674778062760636202846843095908778429447" +
		"9796814873070054845295002271736741151110471789535774982123036612" +
		"1550796316003599954950128589658853017487299129885425412133193933" +
		"2498903413531228420735747116119687278947157473738664145261533343" +
		"2722170532493633174112130233429772097184909350130374469788301052" +
		"3482495638079964143016931353366591300743390994201267050772808522" +
		"2187455270795897148327571092934363351270820420433419984534292033" +
		"9675739631958189684458858870083031556825078177416742007427966636" +
		"0514006572844227668041406273671143700886007701183515021593636919" +
		"9207495924756741875118532305644201331724779670102065881743376681" +
		"3203206640845218097463669557522703962743327392152581126298069254" +
		"1371675752848838153664971001230177808500079048087555874537264894" +
		"5157865259620204365998628723043246349215190836616719846461617004" +
		"0624913343900552612138426875638012514387540454105288758707095367" +
		"5948584994999135084215747465116952738745718450623240269941708511" +
		"1366930012340692857027275009242547379294036673865462739988473971" +
		"3767679168406015019734720008609931393762820307855221374905055757" +
		"2746510383957479894295440342407620814450367153478537997173075792" +
		"9072648135598878031235574144702184920386067657206481787776552780" +
		"2239635888373911848341450646956855422463406393911413229129783107" +
		"0522017864168360550852212880497096745423571929600869430698413665" +
		"1045303500743527841031091286949740380378950815031897548721882510" +
		"0087851028898152865730649781472706179514641042315466508133319155" +
		"3448197158089924374869320322729535800201072851864995814974154852" +
		"5092983652452143232513638504502096492324385227003586236182108361" +
		"5438379862109314812939144480678690473819818969197631889660418795" +
		"2142693741203725348608038844901481302569635930101992656179294873" +
		"91111033363129672689273333070"

	var x, y, z Big
	x.SetInf(false)
	y.SetString(ys)
	z.Mul(&x, &y)
	if !z.IsInf(+1) {
		t.Fatalf("expected Inf, got %s", &z)
	}
}

func TestIssue70(t *testing.T) {
	x, _ := new(Big).SetString("1E+41")
	ctx := Context{Precision: 1}
	ctx.Log10(x, x)
	if x.Cmp(New(40, 0)) != 0 {
		t.Fatalf(`Log10(1e+41)
wanted: 40
got   : %s
`, x)
	}
}

func TestIssue69(t *testing.T) {
	maxSqrt := uint64(4294967295)
	if testing.Short() {
		maxSqrt = 1e6
	}
	for i := uint64(1); i <= maxSqrt; i++ {
		var x Big
		x.SetUint64(i * i)
		var ctx Context
		if v, ok := ctx.Sqrt(&x, &x).Uint64(); !ok || v != i {
			t.Fatalf(`Sqrt(%d)
wanted: %d (0)
got   : %d (%d)
`, i*i, i, v, -x.Scale())
		}
	}
}

func TestIssue73(t *testing.T) {
	x := New(16, 2)
	z := new(Big)
	ctx := Context{Precision: 4100}
	ctx.Sqrt(z, x)
	r := New(4, 1)
	if z.Cmp(r) != 0 || z.Scale() != r.Scale() || z.Context.Conditions != r.Context.Conditions {
		t.Fatalf(`Sqrt(%d)
wanted: %s (%d) %s
got   : %s (%d) %s
`, x, r, -r.Scale(), r.Context.Conditions, z, -z.Scale(), z.Context.Conditions)
	}
}

func TestIssue75(t *testing.T) {
	x := New(576, 2)
	z := new(Big)
	ctx := Context{Precision: 2}
	ctx.Sqrt(z, x)
	r := New(24, 1)
	if z.Cmp(r) != 0 || z.Scale() != r.Scale() || z.Context.Conditions != r.Context.Conditions {
		t.Fatalf(`Sqrt(%d)
wanted: %s (%d) %s
got   : %s (%d) %s
`, x, r, -r.Scale(), r.Context.Conditions, z, -z.Scale(), z.Context.Conditions)
	}
}

func TestIssue146(t *testing.T) {
	var ctx Context
	for i := int64(0); i < 10; i++ {
		n := New(i, 1)
		firstPow := ctx.Pow(new(Big), n, one.get())
		if n.Cmp(firstPow) != 0 {
			t.Errorf("%v^1 != %v", n, firstPow)
		} else {
			t.Logf("%v^1 == %v", n, firstPow)
		}
	}
}
