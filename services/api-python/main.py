"""Tiny Flask service for the pantalasa-cronos monorepo api-python component."""

import os

from flask import Flask, jsonify

app = Flask(__name__)


@app.route("/")
def index():
    return jsonify(
        {
            "service": "api-python",
            "message": "hello from pantalasa-cronos monorepo api-python",
        }
    )


@app.route("/healthz")
def healthz():
    return "ok", 200


if __name__ == "__main__":
    addr = os.environ.get("ADDR", "0.0.0.0")
    port = int(os.environ.get("PORT", "8081"))
    app.run(host=addr, port=port)

# QA S2 probe — api-python leg of multi-component PR
