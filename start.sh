export STAR_HOSTNAME="star"
export STAR_PORT=9000
export MAGNETAR_HOSTNAME="magnetar"
export MAGNETAR_PORT=5000

export NATS_HOSTNAME="nats"
export NATS_PORT=4222

export ETCD_HOSTNAME="etcd"
export ETCD_PORT=2379

export REGISTRATION_SUBJECT="register"
export REGISTRATION_REQ_TIMEOUT_MILLISECONDS=1000
export MAX_REGISTRATION_RETRIES=5

export NODE_ID_DIR_PATH="/etc/c12s"
export NODE_ID_FILE_NAME="nodeid"

export OORT_HOSTNAME=oort
export OORT_PORT=8000

export NEO4J_HOSTNAME=neo4j
export NEO4J_BOLT_PORT=7687
export NEO4J_HTTP_PORT=7474
export NEO4J_AUTH_ENABLED=false
export NEO4J_DBNAME=neo4j
export NEO4J_apoc_export_file_enabled=true
export NEO4J_apoc_import_file_enabled=true
export NEO4J_apoc_import_file_use__neo4j__config=true
export NEO4J_PLUGINS="[\"apoc\"]"

docker-compose build
#docker-compose up --build