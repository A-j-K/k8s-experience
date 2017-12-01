<?php

if(! class_exists("RdKafka\Producer")) {
	die("RdKafka PECL module not available");
}

$broker = getenv("KAFKA_BROKER");
if($broker === false) {
	die("No broker supplied");
}

$topic = getenv("KAFKA_TOPIC");
if($topic === false) {
	die("No topic supplied");
}

$hostname = gethomename();

$rk = new RdKafka\Producer();
$rk->setLogLevel(LOG_DEBUG);
$rk->addBrokers($broker);
$topic = $rk->newTopic($topic);

while(true) {
	for ($i = 0; $i < 10; $i++) {
		$msg = "$hostname sending message $i";
		echo "Sending $msg...\n";
		$topic->produce(RD_KAFKA_PARTITION_UA, 0, $msg);
		$rk->poll(0);
	}
	while ($rk->getOutQLen() > 0) {
		$rk->poll(50);
	}
	sleep(1);
}


