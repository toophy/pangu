module("woLongShanZhuang", package.seeall)

function OnInit(s)
	print("欢迎来到卧龙山庄.")
	-- s:Get_data()["lolo"] = "lolo"
	-- print(s:Get_data()["lolo"])
	ts:PostEventFromLua("woLongShanZhuang","OnHeartBeat",1000,{})
	ts:PostEventFromLua("woLongShanZhuang","Eon_Qiguan",10000,{["log"]="咕咕鸟在鸣叫!"})
end

function OnHeartBeat(s)
	LogInfo("卧龙山庄心跳 "..os.time())
	ts:PostEventFromLua("woLongShanZhuang","OnHeartBeat",1000,{})
	-- print(s:Get_data()["lolo"])
	-- s:Get_data()["lolo"] = "lolo"..os.time()
end

function Eon_Qiguan(t)
	LogInfo(t["log"])
end
