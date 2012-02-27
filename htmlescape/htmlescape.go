// Author: Jim Teeuwen <jimteeuwen@gmail.com>

package htmlescape
 
import "fmt"
import "utf8"
import "regexp"
import "strconv"
 
var reg_entity = regexp.MustCompile("^&#[0-9]+;$");
 
// Converts a single numerical html entity to a regular Go utf-token.
//    ex: "&#9827;" -> "♣"
func HtmlToUTF8(entity string) string {
	// Make sure we have a valid entity: &#123;
	ok := reg_entity.MatchString(entity);
	if !ok { return "" }
 
	// Convert entity to number
	num, err := strconv.Atoi(entity[2:len(entity)-1]);
	if err != nil { return "" }
 
	var arr [3]byte;
	size := utf8.EncodeRune(num, &arr);
	if size == 0 { return "" }
 
	return string(&arr);
}
 
// Converts a single Go utf-token to it's an Html entity.
//   ex: "♣" -> "&#9827;"
func UTF8ToHtml(token string) string {
	rune, size := utf8.DecodeRuneInString(token);
	if size == 0 { return "" }
	return fmt.Sprintf("&#%d;", rune);
}
 
 
/*
	http://www.w3.org/TR/html4/sgml/entities.html
 
	Portions © International Organization for Standardization 1986
	Permission to copy in any form is granted for use with
	conforming SGML systems and applications as defined in
	ISO 8879, provided this notice is included in all copies.
 
	Fills the supplied map with html entities mapped to their Go utf8
	equivalents. This map can be assigned to xml.Parser.Entity
	It will be used to map non-standard xml entities to a proper value.
	If the parser encounters any unknown entities, it will throw a syntax
	error and abort the parsing. Hence the ability to supply this map.
 */
func loadNonStandardEntities(em *map[string]string) {
	// Generic entities string([]uint8{160});
	(*em)["nbsp"] = " ";
	(*em)["iexcl"] = "¡";
	(*em)["cent"] = "¢";
	(*em)["pound"] = "£";
	(*em)["curren"] = "¤";
	(*em)["yen"] = "¥";
	(*em)["brvbar"] = "¦";
	(*em)["sect"] = "§";
	(*em)["uml"] = "¨";
	(*em)["copy"] = "©";
	(*em)["ordf"] = "ª";
	(*em)["laquo"] = "«";
	(*em)["not"] = "¬";
	(*em)["shy"] = "­";
	(*em)["reg"] = "®";
	(*em)["macr"] = "¯";
	(*em)["deg"] = "°";
	(*em)["plusmn"] = "±";
	(*em)["sup"] = "²";
	(*em)["sup"] = "³";
	(*em)["acute"] = "´";
	(*em)["micro"] = "µ";
	(*em)["para"] = "¶";
	(*em)["middot"] = "·";
	(*em)["cedil"] = "¸";
	(*em)["sup"] = "¹";
	(*em)["ordm"] = "º";
	(*em)["raquo"] = "»";
	(*em)["frac14"] = "¼";
	(*em)["frac12"] = "½";
	(*em)["frac34"] = "¾";
	(*em)["iquest"] = "¿";
	(*em)["Agrave"] = "À";
	(*em)["Aacute"] = "Á";
	(*em)["Acirc"] = "Â";
	(*em)["Atilde"] = "Ã";
	(*em)["Auml"] = "Ä";
	(*em)["Aring"] = "Å";
	(*em)["AElig"] = "Æ";
	(*em)["Ccedil"] = "Ç";
	(*em)["Egrave"] = "È";
	(*em)["Eacute"] = "É";
	(*em)["Ecirc"] = "Ê";
	(*em)["Euml"] = "Ë";
	(*em)["Igrave"] = "Ì";
	(*em)["Iacute"] = "Í";
	(*em)["Icirc"] = "Î";
	(*em)["Iuml"] = "Ï";
	(*em)["ETH"] = "Ð";
	(*em)["Ntilde"] = "Ñ";
	(*em)["Ograve"] = "Ò";
	(*em)["Oacute"] = "Ó";
	(*em)["Ocirc"] = "Ô";
	(*em)["Otilde"] = "Õ";
	(*em)["Ouml"] = "Ö";
	(*em)["times"] = "×";
	(*em)["Oslash"] = "Ø";
	(*em)["Ugrave"] = "Ù";
	(*em)["Uacute"] = "Ú";
	(*em)["Ucirc"] = "Û";
	(*em)["Uuml"] = "Ü";
	(*em)["Yacute"] = "Ý";
	(*em)["THORN"] = "Þ";
	(*em)["szlig"] = "ß";
	(*em)["agrave"] = "à";
	(*em)["aacute"] = "á";
	(*em)["acirc"] = "â";
	(*em)["atilde"] = "ã";
	(*em)["auml"] = "ä";
	(*em)["aring"] = "å";
	(*em)["aelig"] = "æ";
	(*em)["ccedil"] = "ç";
	(*em)["egrave"] = "è";
	(*em)["eacute"] = "é";
	(*em)["ecirc"] = "ê";
	(*em)["euml"] = "ë";
	(*em)["igrave"] = "ì";
	(*em)["iacute"] = "í";
	(*em)["icirc"] = "î";
	(*em)["iuml"] = "ï";
	(*em)["eth"] = "ð";
	(*em)["ntilde"] = "ñ";
	(*em)["ograve"] = "ò";
	(*em)["oacute"] = "ó";
	(*em)["ocirc"] = "ô";
	(*em)["otilde"] = "õ";
	(*em)["ouml"] = "ö";
	(*em)["divide"] = "÷";
	(*em)["oslash"] = "ø";
	(*em)["ugrave"] = "ù";
	(*em)["uacute"] = "ú";
	(*em)["ucirc"] = "û";
	(*em)["uuml"] = "ü";
	(*em)["yacute"] = "ý";
	(*em)["thorn"] = "þ";
	(*em)["yuml"] = "ÿ";
	(*em)["fnof"] = "ƒ";
	(*em)["Alpha"] = "Α";
	(*em)["Beta"] = "Β";
	(*em)["Gamma"] = "Γ";
	(*em)["Delta"] = "Δ";
	(*em)["Epsilon"] = "Ε";
	(*em)["Zeta"] = "Ζ";
	(*em)["Eta"] = "Η";
	(*em)["Theta"] = "Θ";
	(*em)["Iota"] = "Ι";
	(*em)["Kappa"] = "Κ";
	(*em)["Lambda"] = "Λ";
	(*em)["Mu"] = "Μ";
	(*em)["Nu"] = "Ν";
	(*em)["Xi"] = "Ξ";
	(*em)["Omicron"] = "Ο";
	(*em)["Pi"] = "Π";
	(*em)["Rho"] = "Ρ";
	(*em)["Sigma"] = "Σ";
	(*em)["Tau"] = "Τ";
	(*em)["Upsilon"] = "Υ";
	(*em)["Phi"] = "Φ";
	(*em)["Chi"] = "Χ";
	(*em)["Psi"] = "Ψ";
	(*em)["Omega"] = "Ω";
	(*em)["alpha"] = "α";
	(*em)["beta"] = "β";
	(*em)["gamma"] = "γ";
	(*em)["delta"] = "δ";
	(*em)["epsilon"] = "ε";
	(*em)["zeta"] = "ζ";
	(*em)["eta"] = "η";
	(*em)["theta"] = "θ";
	(*em)["iota"] = "ι";
	(*em)["kappa"] = "κ";
	(*em)["lambda"] = "λ";
	(*em)["mu"] = "μ";
	(*em)["nu"] = "ν";
	(*em)["xi"] = "ξ";
	(*em)["omicron"] = "ο";
	(*em)["pi"] = "π";
	(*em)["rho"] = "ρ";
	(*em)["sigmaf"] = "ς";
	(*em)["sigma"] = "σ";
	(*em)["tau"] = "τ";
	(*em)["upsilon"] = "υ";
	(*em)["phi"] = "φ";
	(*em)["chi"] = "χ";
	(*em)["psi"] = "ψ";
	(*em)["omega"] = "ω";
	(*em)["thetasym"] = "ϑ";
	(*em)["upsih"] = "ϒ";
	(*em)["piv"] = "ϖ";
	(*em)["bull"] = "•";
	(*em)["hellip"] = "…";
	(*em)["prime"] = "′";
	(*em)["Prime"] = "″";
	(*em)["oline"] = "‾";
	(*em)["frasl"] = "⁄";
	(*em)["weierp"] = "℘";
	(*em)["image"] = "ℑ";
	(*em)["real"] = "ℜ";
	(*em)["trade"] = "™";
	(*em)["alefsym"] = "ℵ";
	(*em)["larr"] = "←";
	(*em)["uarr"] = "↑";
	(*em)["rarr"] = "→";
	(*em)["darr"] = "↓";
	(*em)["harr"] = "↔";
	(*em)["crarr"] = "↵";
	(*em)["lArr"] = "⇐";
	(*em)["uArr"] = "⇑";
	(*em)["rArr"] = "⇒";
	(*em)["dArr"] = "⇓";
	(*em)["hArr"] = "⇔";
	(*em)["forall"] = "∀";
	(*em)["part"] = "∂";
	(*em)["exist"] = "∃";
	(*em)["empty"] = "∅";
	(*em)["nabla"] = "∇";
	(*em)["isin"] = "∈";
	(*em)["notin"] = "∉";
	(*em)["ni"] = "∋";
	(*em)["prod"] = "∏";
	(*em)["sum"] = "∑";
	(*em)["minus"] = "−";
	(*em)["lowast"] = "∗";
	(*em)["radic"] = "√";
	(*em)["prop"] = "∝";
	(*em)["infin"] = "∞";
	(*em)["ang"] = "∠";
	(*em)["and"] = "∧";
	(*em)["or"] = "∨";
	(*em)["cap"] = "∩";
	(*em)["cup"] = "∪";
	(*em)["int"] = "∫";
	(*em)["there4"] = "∴";
	(*em)["sim"] = "∼";
	(*em)["cong"] = "≅";
	(*em)["asymp"] = "≈";
	(*em)["ne"] = "≠";
	(*em)["equiv"] = "≡";
	(*em)["le"] = "≤";
	(*em)["ge"] = "≥";
	(*em)["sub"] = "⊂";
	(*em)["sup"] = "⊃";
	(*em)["nsub"] = "⊄";
	(*em)["sube"] = "⊆";
	(*em)["supe"] = "⊇";
	(*em)["oplus"] = "⊕";
	(*em)["otimes"] = "⊗";
	(*em)["perp"] = "⊥";
	(*em)["sdot"] = "⋅";
	(*em)["lceil"] = "⌈";
	(*em)["rceil"] = "⌉";
	(*em)["lfloor"] = "⌊";
	(*em)["rfloor"] = "⌋";
	(*em)["lang"] = "〈";
	(*em)["rang"] = "〉";
	(*em)["loz"] = "◊";
	(*em)["spades"] = "♠";
	(*em)["clubs"] = "♣";
	(*em)["hearts"] = "♥";
	(*em)["diams"] = "♦";
	(*em)["quot"] = "\"";
	(*em)["amp"] = "&";
	(*em)["lt"] = "<";
	(*em)["gt"] = ">";
	(*em)["OElig"] = "Œ";
	(*em)["oelig"] = "œ";
	(*em)["Scaron"] = "Š";
	(*em)["scaron"] = "š";
	(*em)["Yuml"] = "Ÿ";
	(*em)["circ"] = "ˆ";
	(*em)["tilde"] = "˜";
	(*em)["ensp"] = " ";
	(*em)["emsp"] = " ";
	(*em)["thinsp"] = " ";
	(*em)["zwnj"] = "‌";
	(*em)["zwj"] = "‍";
	(*em)["lrm"] = "‎";
	(*em)["rlm"] = "‏";
	(*em)["ndash"] = "–";
	(*em)["mdash"] = "—";
	(*em)["lsquo"] = "‘";
	(*em)["rsquo"] = "’";
	(*em)["sbquo"] = "‚";
	(*em)["ldquo"] = "“";
	(*em)["rdquo"] = "”";
	(*em)["bdquo"] = "„";
	(*em)["dagger"] = "†";
	(*em)["Dagger"] = "‡";
	(*em)["permil"] = "‰";
	(*em)["lsaquo"] = "‹";
	(*em)["rsaquo"] = "›";
	(*em)["euro"] = "€";
}
 
