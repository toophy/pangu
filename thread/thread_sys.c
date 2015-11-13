#ifdef WIN32
	#include <windows.h>
#else
	#include <pthread.h>
#endif

#include <thread_sys.h>

unsigned int PthreadSelf()
{
	#ifdef WIN32
		return GetCurrentThreadId();
	#else
		return thread_self();
	#endif
}
