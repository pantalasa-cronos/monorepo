"""Tiny Flask service for the pantalasa-cronos monorepo api-python component."""

# Touches api-python so the path-filtered code collector fires for this
# sub-component alongside the per-service CI job (LUNAR_COMPONENT attribution test).
# Re-run after reverting the snyk TEMP pin that wedged the cronos hub manifest.
# ci-probe 2026-06-08: switched to LUNAR_COMPONENT_INFER — this touch makes
# api-python appear in the PR diff so changed-paths inference is exercised too.

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
