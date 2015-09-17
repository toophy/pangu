module("main", package.seeall)

function OnWorldBegin( t )
	LogDebug("世界线程启动")
	ts:CreateScreenThread(Tid_screen_1, "场景线程1", 100, Evt_lay1_time, 10000)
	ts:CreateScreenThread(Tid_screen_2, "场景线程2", 100, Evt_lay1_time, 10000)
	return nil
end

function OnWorldEnd( t )
	LogDebug("世界线程结束")
	return nil
end
