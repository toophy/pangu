module("main", package.seeall)

function OnScreenThreadBegin( t )
	LogDebug("场景线程"..ts:Get_thread_id().." 启动")
	return nil
end

function OnScreenThreadEnd( t )
	LogDebug("场景线程"..ts:Get_thread_id().." 结束")
	return nil
end

function OnInitScreen()
	ts:Add_screen("阿拉斯加2", 1)
	return 1
end
