--
-- 本脚本不得修改, 不得删除
--


-- Set up a custom loader
-- package.loaders[2] = function(name) return System:LoadModule(name) end

-- 禁止动态库和all-in-one库加载
package.loaders[3] = nil
package.loaders[4] = nil

-- 禁止加载其他库
package.loadlib = function() end

-- 禁止OS部分函数
local osOriginal = os

os = 
{
	date = os.date,
	time = os.time, 
	setlocale = os.setlocale,
	clock = os.clock, 
	difftime = os.difftime,
}

for k, v in pairs(osOriginal) do
	if not os[k] and type(v) == "function" then
		os[k] = function() end
	end
end

-- 当前目录
local curr_dir = ""
local path_obj = io.popen("cd")  --如果不在交互模式下，前面可以添加local 
curr_dir = path_obj:read("*all"):sub(1,-3)    --path存放当前路径
path_obj:close()   --关掉句柄

-- 自定义print函数, 指向线程普通信息日志函数
function print(...)
   local result = ""
   for i, v in ipairs{...} do
       result = result .. v .. ' '
   end
   ts:LogInfo(result)
end

-- 日志 : 调试信息
function LogDebug(...)
   local result = ""
   for i, v in ipairs{...} do
       result = result .. v .. ' '
   end
   ts:LogDebug(result)
end

-- 日志 : 普通信息
function LogInfo(...)
   local result = ""
   for i, v in ipairs{...} do
       result = result .. v .. ' '
   end
   ts:LogInfo(result)
end

-- 日志 : 警告
function LogWarn(...)
   local result = ""
   for i, v in ipairs{...} do
       result = result .. v .. ' '
   end
   ts:LogWarn(result)
end

-- 日志 : 普通错误
function LogError(...)
   local result = ""
   for i, v in ipairs{...} do
       result = result .. v .. ' '
   end
   ts:LogError(result)
end

-- 日志 : 严重错误
function LogFatal(...)
   local result = ""
   for i, v in ipairs{...} do
       result = result .. v .. ' '
   end
   ts:LogFatal(result)
end


-- 根据模块名获得脚本模块
function GetScriptModule(ModuleName)
  return package.loaded[ModuleName]
end

-- 打印table
function PrintTable(root)
  local cache = {  [root] = "." }
  local function _dump(t,space,name)
    local temp = {}
    for k,v in pairs(t) do
      local key = tostring(k)
      if cache[v] then
        table.insert(temp,"* " .. key .. " {" .. cache[v].."}")
      elseif type(v) == "table" then
        local new_key = name .. "." .. key
        cache[v] = new_key
        table.insert(temp,"+ " .. key .. _dump(v,space .. (next(t,k) and "|" or " " ).. string.rep(" ",#key),new_key))
      else
        table.insert(temp,"- " .. key .. " [" .. tostring(v).."]")
      end
    end
    return table.concat(temp,"\n"..space)
  end
  print(_dump(root, "",""))
end

-- 获取当前目录
function GetCurrDir()
  return curr_dir
end


-- 初始化随机函数
math.randomseed(os.time())
math.random(1, 100)
