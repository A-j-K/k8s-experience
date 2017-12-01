<?php

if(! class_exists("RdKafka\Conf")) {
	die("RdKafka PECL module not installed");
}

$broker = getenv("KAFKA_BROKER");
if($broker === false) {
	die("No broker details supplied");
}

$topic = getenv("KAFKA_TOPIC");
if($topic === false) {
	die("No topic supplied");
}

$conf = new RdKafka\Conf();

$conf->set('group.id', 'myConsumerGroup');

$rk = new RdKafka\Consumer($conf);
$rk->addBrokers($broker);

$topicConf = new RdKafka\TopicConf();
$topicConf->set('auto.commit.interval.ms', 100);
$topicConf->set('offset.store.method', 'file');
$topicConf->set('offset.store.path', sys_get_temp_dir());
$topicConf->set('auto.offset.reset', 'smallest');

$topic = $rk->newTopic($topic, $topicConf);

// Start consuming partition 0
$topic->consumeStart(0, RD_KAFKA_OFFSET_STORED);

while (true) {
    $message = $topic->consume(0, 120*10000);
    switch ($message->err) {
        case RD_KAFKA_RESP_ERR_NO_ERROR:
            var_dump($message);
            break;
        case RD_KAFKA_RESP_ERR__PARTITION_EOF:
            echo "No more messages; will wait for more\n";
            break;
        case RD_KAFKA_RESP_ERR__TIMED_OUT:
            echo "Timed out\n";
            break;
        default:
            throw new \Exception($message->errstr(), $message->err);
            break;
    }
}

