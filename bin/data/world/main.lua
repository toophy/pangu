module("main", package.seeall)

function OnWorldBegin( t )
	LogDebug("世界线程启动")
	return nil
end

function OnWorldEnd( t )
	LogDebug("世界线程结束")
	return nil
end
