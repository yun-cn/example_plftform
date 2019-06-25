#!/usr/bin/env bash

# mapping: https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping.html
#
# geo_point: https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-point.html
#
# id_on_platform: Some platforms might use uuids rather than ids. So use keyword instead of long.

index_name="$1"

echo "=== CREATE INDEX === -> ${index_name}_listings"
curl -sS -XPUT "http://localhost:9200/${index_name}_listings" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "_doc": {
      "properties": {
        "data": {
          "type": "object",
          "enabled": false
        },
        "type": {
          "type": "keyword"
        },
        "search": {
          "type": "text",
          "analyzer": "my_analyzer"
        },
        "id": {
          "type": "long"
        },
        "id_on_platform": {
          "type": "keyword"
        },
        "refreshed_at": {
          "type": "date"
        },
        "date": {
          "type": "date"
        },
        "review_count": {
           "type": "integer"
        },
        "event_type": {
           "type": "keyword"
        },
        "amenities": {
           "type": "keyword"
        },
        "location": {
          "type": "geo_point"
        },
        "platform": {
          "type": "keyword"
        }
      }
    }
  },
  "settings": {
    "index": {
      "number_of_shards": 1,
      "max_ngram_diff": 10
    },
    "analysis": {
      "analyzer": {
        "my_analyzer": {
          "tokenizer": "my_tokenizer",
          "filter": [
            "lowercase"
          ]
        }
      },
      "tokenizer": {
        "my_tokenizer": {
          "type": "ngram",
          "min_gram": 1,
          "max_gram": 10,
          "token_chars": [
            "letter",
            "digit",
            "symbol",
            "punctuation"
          ]
        }
      }
    }
  }
}'
echo ''

echo "=== CREATE INDEX === -> ${index_name}_search_requests"
curl -sS -XPUT "http://localhost:9200/${index_name}_search_requests" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "_doc": {
      "properties": {
        "platform": {
          "type": "keyword"
        },
        "results": {
          "type": "object",
          "enabled": false
        },
        "listing_ids": {
          "type": "keyword"
        },
        "refreshed_at": {
          "type": "date"
        }
      }
    }
  },
  "settings": {
    "index": {
      "number_of_shards": 1,
      "max_ngram_diff": 10
    },
    "analysis": {
      "analyzer": {
        "my_analyzer": {
          "tokenizer": "my_tokenizer",
          "filter": [
            "lowercase"
          ]
        }
      },
      "tokenizer": {
        "my_tokenizer": {
          "type": "ngram",
          "min_gram": 1,
          "max_gram": 10,
          "token_chars": [
            "letter",
            "digit",
            "symbol",
            "punctuation"
          ]
        }
      }
    }
  }
}'
echo ''

echo "=== CREATE INDEX === -> ${index_name}_place_suggestions"
curl -sS -XPUT "http://localhost:9200/${index_name}_place_suggestions" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "_doc": {
      "properties": {
        "results": {
          "type": "object",
          "enabled": false
        },
        "name": {
          "type": "text"
        },
        "ward_id": {
          "type": "long"
        },
        "area_id": {
          "type": "long"
        },
        "prefecture_id": {
          "type": "long"
        }
      }
    }
  },
  "settings": {
    "index": {
      "number_of_shards": 1,
      "max_ngram_diff": 10
    },
    "analysis": {
      "analyzer": {
        "my_analyzer": {
          "tokenizer": "my_tokenizer",
          "filter": [
            "lowercase"
          ]
        }
      },
      "tokenizer": {
        "my_tokenizer": {
          "type": "ngram",
          "min_gram": 1,
          "max_gram": 10,
          "token_chars": [
            "letter",
            "digit",
            "symbol",
            "punctuation"
          ]
        }
      }
    }
  }
}'
echo ''

echo "=== CREATE INDEX === -> ${index_name}_users"
curl -sS -XPUT "http://localhost:9200/${index_name}_users" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "_doc": {
      "properties": {
        "data": {
          "type": "object",
          "enabled": false
        },
        "platform": {
          "type": "keyword"
        },
        "id_on_platform": {
          "type": "keyword"
        },
        "name": {
          "type": "text"
        },
        "avatar_url": {
          "type": "keyword"
        }
      }
    }
  },
  "settings": {
    "index": {
      "number_of_shards": 1,
      "max_ngram_diff": 10
    },
    "analysis": {
      "analyzer": {
        "my_analyzer": {
          "tokenizer": "my_tokenizer",
          "filter": [
            "lowercase"
          ]
        }
      },
      "tokenizer": {
        "my_tokenizer": {
          "type": "ngram",
          "min_gram": 1,
          "max_gram": 10,
          "token_chars": [
            "letter",
            "digit",
            "symbol",
            "punctuation"
          ]
        }
      }
    }
  }
}'
echo ''

echo "=== CREATE INDEX === -> ${index_name}_reviews"
curl -sS -XPUT "http://localhost:9200/${index_name}_reviews" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "_doc": {
      "properties": {
        "user_name": {
          "type": "text"
        },
        "description": {
          "type": "text"
        },
        "platform": {
          "type": "keyword"
        },
        "id_on_platform": {
          "type": "keyword"
        },
        "user_id_on_platform": {
          "type": "keyword"
        },
        "user_avatar_url": {
          "type": "keyword"
        },
        "listing_id_on_platform": {
          "type": "keyword"
        }
      }
    }
  },
  "settings": {
    "index": {
      "number_of_shards": 1,
      "max_ngram_diff": 10
    },
    "analysis": {
      "analyzer": {
        "my_analyzer": {
          "tokenizer": "my_tokenizer",
          "filter": [
            "lowercase"
          ]
        }
      },
      "tokenizer": {
        "my_tokenizer": {
          "type": "ngram",
          "min_gram": 1,
          "max_gram": 10,
          "token_chars": [
            "letter",
            "digit",
            "symbol",
            "punctuation"
          ]
        }
      }
    }
  }
}'
echo ''
