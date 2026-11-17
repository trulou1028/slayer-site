-- links.lua: makes the links in a copied slayer document work on the site.
--
-- A slayer document links to its neighbours by relative path. The ones the
-- site carries are pointed at their pages here; pictures and films are
-- pointed at the folder they were copied into (the page's media-prefix);
-- every other relative link keeps its words and loses its target, because
-- the document it names is not on the site and a dead link is worse than
-- none. Links that leave for the web are left alone.

local pages = {
  ["native/README.md"] = "ferruccio-native.html",
  ["../ferruccio/README.md"] = "ferruccio.html",
  ["../../ferruccio/README.md"] = "ferruccio.html",
  ["../native/README.md"] = "ferruccio-native.html",
  ["../azazel.md"] = "azazel.html",
  ["../doc/azazel.md"] = "azazel.html",
  ["azazel.md"] = "azazel.html",
  ["visualguide/README.md"] = "visual-guide.html",
}

local prefix = ""

local function relative(target)
  return not target:match("^%a[%w+.-]*:") and not target:match("^#")
    and not target:match("^/")
end

local function media(path)
  return path:match("^images/") or path:match("%.mp4$")
end

return {
  {
    Meta = function(meta)
      if meta["media-prefix"] then
        prefix = pandoc.utils.stringify(meta["media-prefix"])
      end
    end,
  },
  {
    Link = function(el)
      if not relative(el.target) then return nil end
      local path, fragment = el.target:match("^([^#]*)(.*)$")
      if pages[path] then
        el.target = pages[path] .. fragment
        return el
      end
      if media(path) then
        el.target = prefix .. path
        return el
      end
      return el.content
    end,
    Image = function(el)
      if relative(el.src) then el.src = prefix .. el.src end
      return el
    end,
  },
}
