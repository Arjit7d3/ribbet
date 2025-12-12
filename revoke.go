package main

import (
	"net/http"

	"github.com/Arjit7d3/ribbet/internal/auth"
)

func (cfg *apiConfig) handleRevoke(w http.ResponseWriter, r *http.Request) {
	refreshTokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't find token", err)
		return
	}
	err = cfg.db.RevokeRefreshToken(r.Context(), refreshTokenString)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't revoke session", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
