package hint_codes

const IMPORT_SECP256R1_ALPHA = "from starkware.cairo.common.cairo_secp.secp256r1_utils import SECP256R1_ALPHA as ALPHA"
const IMPORT_SECP256R1_N = "from starkware.cairo.common.cairo_secp.secp256r1_utils import SECP256R1_N as N"
const IMPORT_SECP256R1_P = "from starkware.cairo.common.cairo_secp.secp256r1_utils import SECP256R1_P as SECP_P"
const VERIFY_ZERO_EXTERNAL_SECP = "from starkware.cairo.common.cairo_secp.secp_utils import pack\n\nq, r = divmod(pack(ids.val, PRIME), SECP_P)\nassert r == 0, f\"verify_zero: Invalid input {ids.val.d0, ids.val.d1, ids.val.d2}.\"\nids.q = q % PRIME"
