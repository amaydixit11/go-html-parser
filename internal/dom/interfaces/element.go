package dom


type Element interface{
	readonly attribute DOMString? namespaceURI;
	readonly attribute DOMString? prefix;
	readonly attribute DOMString localName;
	readonly attribute DOMString tagName;
  
	[CEReactions] attribute DOMString id;
	[CEReactions] attribute DOMString className;
	[SameObject, PutForwards=value] readonly attribute DOMTokenList classList;
	[CEReactions, Unscopable] attribute DOMString slot;
  
	boolean hasAttributes();
	[SameObject] readonly attribute NamedNodeMap attributes;
	sequence<DOMString> getAttributeNames();
	DOMString? getAttribute(DOMString qualifiedName);
	DOMString? getAttributeNS(DOMString? namespace, DOMString localName);
	[CEReactions] undefined setAttribute(DOMString qualifiedName, DOMString value);
	[CEReactions] undefined setAttributeNS(DOMString? namespace, DOMString qualifiedName, DOMString value);
	[CEReactions] undefined removeAttribute(DOMString qualifiedName);
	[CEReactions] undefined removeAttributeNS(DOMString? namespace, DOMString localName);
	[CEReactions] boolean toggleAttribute(DOMString qualifiedName, optional boolean force);
	boolean hasAttribute(DOMString qualifiedName);
	boolean hasAttributeNS(DOMString? namespace, DOMString localName);
  
	Attr? getAttributeNode(DOMString qualifiedName);
	Attr? getAttributeNodeNS(DOMString? namespace, DOMString localName);
	[CEReactions] Attr? setAttributeNode(Attr attr);
	[CEReactions] Attr? setAttributeNodeNS(Attr attr);
	[CEReactions] Attr removeAttributeNode(Attr attr);
  
	ShadowRoot attachShadow(ShadowRootInit init);
	readonly attribute ShadowRoot? shadowRoot;
  
	Element? closest(DOMString selectors);
	boolean matches(DOMString selectors);
	boolean webkitMatchesSelector(DOMString selectors); // legacy alias of .matches
  
	HTMLCollection getElementsByTagName(DOMString qualifiedName);
	HTMLCollection getElementsByTagNameNS(DOMString? namespace, DOMString localName);
	HTMLCollection getElementsByClassName(DOMString classNames);
  
	[CEReactions] Element? insertAdjacentElement(DOMString where, Element element); // legacy
	undefined insertAdjacentText(DOMString where, DOMString data); // legacy
  };
  
  type ShadowRootInit struct {
	mode *ShadowRootMode;
	delegatesFocus bool = false;
	slotAssignment *SlotAssignmentMode = "named";
	clonable bool = false;
	serializable bool = false;
  };