package com.dattus.poc.main;

import java.util.concurrent.ScheduledThreadPoolExecutor;
import java.util.concurrent.TimeUnit;
import java.util.logging.Logger;

import com.dattus.poc.common.Constants;


public class CDTest {
	
	static {
		Logger.getAnonymousLogger().info("Version : "+Constants.version);
	}
	
	public static void main(String[] args) {
		ScheduledThreadPoolExecutor executor=new ScheduledThreadPoolExecutor(1);
		executor.scheduleAtFixedRate(new Runnable() {
			
			@Override
			public void run() {
				Logger.getAnonymousLogger().info("Ping!");
			}
			
		}, 0, 1, TimeUnit.SECONDS);
		Runtime.getRuntime().addShutdownHook(new Thread(new Runnable() {
			
			@Override
			public void run() {
				Logger.getAnonymousLogger().info("Shutting down!");
			}
		}));
	}
}
