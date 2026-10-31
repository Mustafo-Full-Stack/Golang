package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/netip"
	"strconv"

	"demo-project/internal/handler/response"
)

const maxBodyBytes = 1 << 20 // 1 МБ: огромное тело не должно съесть память сервера

// decodeJSON читает тело запроса в dst. При ошибке сам отвечает 400 и возвращает false.
func decodeJSON(w http.ResponseWriter, r *http.Request, dst any) bool {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		response.Error(w, r, http.StatusBadRequest, "invalid json body", err)
		return false
	}
	return true
}

// parseID читает {id} из пути. При ошибке сам отвечает 400 и возвращает false.
func parseID(w http.ResponseWriter, r *http.Request) (int64, bool) {
	raw := r.PathValue("id")
	id, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || id <= 0 {
		response.Error(w, r, http.StatusBadRequest, "invalid id", fmt.Errorf("invalid id %q", raw))
		return 0, false
	}
	return id, true
}

// queryInt читает необязательный целочисленный query-параметр.
// При ошибке сам отвечает 400 и возвращает false. Границы проверяет сервис.
func queryInt(w http.ResponseWriter, r *http.Request, name string, def int) (int, bool) {
	raw := r.URL.Query().Get(name)
	if raw == "" {
		return def, true
	}

	value, err := strconv.Atoi(raw)
	if err != nil {
		response.Error(w, r, http.StatusBadRequest, "invalid "+name, err)
		return 0, false
	}
	return value, true
}

// clientIP — адрес, с которого пришло соединение.
// За reverse proxy это будет адрес proxy: см. README, раздел «За reverse proxy».
func clientIP(r *http.Request) netip.Addr {
	addrPort, err := netip.ParseAddrPort(r.RemoteAddr)
	if err != nil {
		return netip.Addr{}
	}
	return addrPort.Addr()
}
